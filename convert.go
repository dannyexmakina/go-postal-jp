package main

import (
	"archive/zip"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

const (
	Tablename     = "postals"
	TempDirName   = "tml"
	AvoidAddress1 = "イカニケイサイガナイバアイ"
	AvoidAddress2 = "以下に掲載がない場合"
)

/*
Parse CSV from JapanPost

( See http://www.post.japanpost.jp/zipcode/dl/kogaki-zip.html )
*/

func Convert() (e error) {
	csvpath, err := Download(config.Url)
	if err != nil {
		return err
	}

	if len(csvpath) < 1 {
		e = errors.New("Download fail")
		return
	}

	if err := ParseAndInsert(csvpath); err != nil {
		e = err
		return
	}
	temp := filepath.Dir(csvpath)
	os.RemoveAll(temp)
	return
}

func Download(url string) (path string, e error) {

	temp, err := os.MkdirTemp(config.Storagepath, TempDirName)
	if err != nil {
		e = err
		return
	}

	zp := filepath.Join(temp, "zip")
	dst, err := os.Create(zp)
	if err != nil {
		e = err
		return
	}
	defer dst.Close()

	r, err := http.Get(url)
	if err != nil {
		e = err
		return
	}
	defer r.Body.Close()
	if _, err := io.Copy(dst, r.Body); err != nil {
		e = err
		return
	}

	baseDir := filepath.Dir(zp)
	z, err := zip.OpenReader(zp)
	if err != nil {
		e = err
		return
	}
	defer z.Close()

	for _, zf := range z.File {
		zff, err := zf.Open()
		if err != nil {
			continue
		}
		defer zff.Close()

		if zf.Name != config.CsvName {
			return
		}
		path = filepath.Join(baseDir, zf.Name)

		if zf.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
			continue
		}
		csvf, err := os.OpenFile(
			path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zf.Mode())
		if err != nil {
			continue
		}
		defer csvf.Close()
		_, err = io.Copy(csvf, zff)
		if err != nil {
			continue
		}
	}
	return
}

func ParseAndInsert(filename string) (e error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		e = err
		return
	}
	defer f.Close()
	db, err := NewSqlite(config.Sqlite.Path, config.Sqlite.Save)
	if err != nil {
		e = err
		return
	}
	defer db.Close()

	db.Memory.AutoMigrate(&Postal{})

	reader := CsvShiftJISReader(f)

	var counter int
	var r1 = regexp.MustCompile("（(.*?)）")
	var r2 = regexp.MustCompile("\\((.*?)\\)")

	var ps Postals
	for {
		counter++

		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		p := ParseLine(line, r1, r2)
		ps = append(ps, p)
	}

	db.Memory.Table(Tablename).Create(&ps)

	var dups Postals
	e = db.Memory.Table(Tablename).
		Select("code").
		Group("code").
		Having("COUNT(`code`) > (?)", 1).
		Find(&dups).Error
	if e != nil {
		return
	}

	var dupCodes []string
	var forDelete Postals

	for _, pd1 := range dups {
		dupCodes = append(dupCodes, pd1.Code)

		var dupsByCode Postals
		if err := db.Memory.Table(Tablename).
			Select("*").
			Where("code=?", pd1.Code).
			Find(&dupsByCode).Error; err != nil {
			return err
		}
		var added bool
		var counter int
		for _, pd2 := range dupsByCode {
			counter++
			for _, dp := range forDelete {
				if dp.Code == pd2.Code {
					added = true
				}
			}
			if added {
				if err := db.Memory.Table(Tablename).Where("id = ?", pd2.Id).Delete(pd2).Error; err != nil {
					// log error
					continue
				}
			} else {
				forDelete = append(forDelete, pd2)
				pd2.EraseCity()
				if err := db.Memory.Table(Tablename).Save(pd2).Error; err != nil {
					// log error
					continue
				}
			}
			if counter > config.Chunk {
				//        time.Sleep(1000 * time.Millisecond)
				counter = 0
			}
		}
	}

	var results Postals
	db.Memory.Table(Tablename).Find(&results)
	results.SaveSQL(config.Sql.Path)

	db.Save(Tablename)
	return
}

func ParseLine(line []string, r1 *regexp.Regexp, r2 *regexp.Regexp) (p Postal) {
	var newline []string
	for _, v := range line {
		v = KanaHalfToFull(v)
		if v == AvoidAddress1 || v == AvoidAddress2 {
			v = ""
		}
		newline = append(newline, v)
	}
	code := newline[2]
	jiscode := newline[0]

	street := r1.ReplaceAllString(newline[8], "")
	street_kana := r2.ReplaceAllString(newline[5], "")

	p = Postal{
		First:       code[0:3],
		Last:        code[3:],
		Code:        code,
		Prefcode:    jiscode[0:2],
		Pref:        newline[6],
		Pref_kana:   newline[3],
		City:        newline[7],
		City_kana:   newline[4],
		Street:      street,
		Street_kana: street_kana,
		Jiscode:     jiscode,
	}
	return p
}
