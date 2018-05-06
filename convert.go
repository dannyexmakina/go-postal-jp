package main

import (
  "path/filepath"
  "os"
//  "errors"
  "io"
  "net/http"
  "io/ioutil"
  "archive/zip"
  "golang.org/x/text/encoding/japanese"
  "golang.org/x/text/transform"
  "encoding/csv"

	"github.com/gocraft/dbr"
  _ "github.com/mattn/go-sqlite3"

  "bytes"
  "regexp"

)

/*
Parse CSV from JapanPost

( See http://www.post.japanpost.jp/zipcode/dl/kogaki-zip.html )
*/
func Convert() (e error) {

  p,pe := NewPostal()
  if pe != nil {
    return pe
  }
  defer p.Close()

  n,ne := p.Download()
  if ne != nil {
    Log("Warn",ne)
    return ne
  }

  c,ce := p.Unzip(n)
  if ce != nil {
    Log("Warn",ce)
    return ce
  }

  if len(c) < 1 {
    Log("Warn", "unzip fail")
    return
  }

  se := p.ParseAndInsert(c)
  if se != nil {
    Log("Warn",se)
    return se
  }

  Log("Info", "end")

  return nil
}

func NewPostal() (p Postal, e error) {
  
  Log("Info", "start")
  
	temp,te := ioutil.TempDir("","postal")
	if te != nil {
		return p,te
	}
  p.Dir = temp
  return p,nil
}

func (p Postal) Close() error {
  return os.RemoveAll(p.Dir)
}

func (p Postal) Download() (n string, e error) {

  Log("Info","start download")

	temp,te := ioutil.TempDir("","postal")
	if te != nil {
    Log("Warn", te)
		return n,te
	}

  n = filepath.Join(temp, "zip")
  dst, err := os.Create(n)
  if err != nil {
    Log("Warn", err)
      return n,err
  }
  defer dst.Close()

  r,err := http.Get(config.Url)
  if err != nil {
    Log("Warn", err, config.Url)
    return n,err
  }
  defer r.Body.Close()
  _,e = io.Copy(dst, r.Body) 
  if e != nil {
    Log("Warn", e)
    return n,e
  }
  return n,nil

}


func (p Postal) Unzip(zipfile string) (path string, e error) {

  Log("Info","start unzip")

  z, err := zip.OpenReader(zipfile)
  if err != nil {
    return path,err
  }
  defer func(){
    z.Close()
    os.Remove(zipfile)
  }()

  for _,f := range z.File {
    if f.Name != config.CsvName {
      return
    }

    path, e = func(fp *zip.File) (pa string, e error) {
      zf, err := fp.Open()
      defer zf.Close()
      if err != nil {
        return path,err
      }
      pa = filepath.Join(p.Dir, f.Name)
      if fp.FileInfo().IsDir() {
        os.MkdirAll(pa, os.ModePerm)
        return
      }
      buf := make([]byte, fp.UncompressedSize)
      _, e = io.ReadFull(zf, buf)
      if e != nil {
        return
      }
      e = ioutil.WriteFile(pa, buf, fp.Mode())
      return
    }(f)
    if e != nil {
      return
    }
  }
  return
}


func (p Postal) ParseAndInsert(filename string) (e error) {

  Log("Info","start parse")

  f, fe := os.OpenFile(filename, os.O_RDONLY, 0644)
  if fe != nil {
    Log("Warn",fe,filename)
      return fe
  }
  defer f.Close()

  reader := csv.NewReader(transform.NewReader(f,japanese.ShiftJIS.NewDecoder()))
  reader.Comma = ','
  reader.LazyQuotes = true

  connection, err := dbr.Open("sqlite3", config.Sqlite.Path, nil)
  if err != nil {
    return err
  }
  defer connection.Close()
  session := connection.NewSession(nil)

  Log("Info", "start initialize")

  session.Exec("CREATE TABLE postal (id integer primary key autoincrement, first, last, code, prefcode, pref, pref_kana, city, city_kana, street, street_kana, jiscode);")

  Log("Info","start insert")

  var counter int
  var r1 = regexp.MustCompile("（(.*?)）")
  var r2 = regexp.MustCompile("\\((.*?)\\)")

  for {
    counter++

    line, err := reader.Read()
    if err == io.EOF {
      break
    } else if err != nil {
      Log("Warn",err)
      return err
    }

    p := ParseLine(line, r1, r2)
  
    _,re := session.InsertInto("postal").
                Columns("first", "last", "code", "prefcode", "pref", "pref_kana", "city", "city_kana", "street", "street_kana", "jiscode").
                Record(p).
                Exec()
   if re != nil {
     return re
   }
  }

  Log("Info","start screening")

	var dups []PostalData
	e = session.Select("code").
				From("postal").
				GroupBy("code").
				Having(dbr.Expr("COUNT(`code`) > 1")).
				LoadStruct(&dups)
	if e != nil {
    Log("Warn", e)
		return e
	}

	var dupCodes []string
	var forDelete []PostalData
	for _,pd1 := range dups {
		dupCodes = append(dupCodes, pd1.Code)

		var dupsByCode []PostalData
		dupe := session.Select("*").
					From("postal").
					Where("code=?", pd1.Code).
					LoadStruct(&dupsByCode)
		if dupe != nil {
			return dupe
		}
		var added bool
    var counter int
		for _,pd2 := range dupsByCode{
      counter++
			for _,dp := range forDelete {
				if dp.Code == pd2.Code {
					added = true
				}
			}
			if added {
				_,de := session.DeleteFrom("postal").
							Where("id = ?", pd2.Id).
							Exec()
				if de != nil {
					Log("Warn", de)
				}
			} else {
				forDelete = append(forDelete, pd2)
				_,ue := session.Update("postal").
							Set("City","").
							Set("City_kana","").
							Set("Street","").
							Set("Street_kana","").
							Where("id = ?", pd2.Id).
							Exec()
				if ue != nil {
					Log("Warn", ue)
				}
			}
      if counter > config.Chunk {
//        time.Sleep(1000 * time.Millisecond)
        counter = 0
      }
		}
	}

  Log("Info","start dump")

	var results []PostalData
	e = session.Select("*").
				From("postal").
				LoadStruct(&results)
	if e != nil {
    Log("Warn", e)
		return e
	}
  
  buf := bytes.NewBufferString("")
  for _,v := range results {
    sqlValue := "'" + v.First + "','"+
              v.Last + "','" + 
              v.Code + "','" +
              v.Prefcode + "','" +
              v.Pref + "','" +
              v.Pref_kana + "','" +
              v.City + "','" +
              v.City_kana + "','" +
              v.Street + "','" +
              v.Street_kana + "','" +
              v.Jiscode + "'"
    sql := "INSERT INTO postal (first, last, code, prefcode, pref, pref_kana, city, city_kana, street, street_kana, jiscode) VALUES (" + sqlValue + ");\n"
    buf.WriteString(sql)
  }
  file,fe := os.Create(config.Sql.Path)
  defer file.Close()
  buf.WriteTo(file)

  return nil
}

func ParseLine(line []string, r1 *regexp.Regexp, r2 *regexp.Regexp) (p PostalData) {
  var newline []string
  for _,v := range line {
    v = KanaHalfToFull(v)
    if(v == "イカニケイサイガナイバアイ" || v == "以下に掲載がない場合"){
      v = ""
    }
    newline = append(newline,v);
  }
  code := newline[2]
  jiscode := newline[0]

  street := r1.ReplaceAllString(newline[8], "")
  street_kana := r2.ReplaceAllString(newline[5], "")

	p = PostalData{
		First: code[0:3],
		Last: code[3:],
		Code: code,
		Prefcode: jiscode[0:2],
		Pref: newline[6],
		Pref_kana: newline[3],
    City: newline[7],
    City_kana: newline[4],
    Street: street,
    Street_kana: street_kana,
		Jiscode: jiscode,
	}
	return p
}


