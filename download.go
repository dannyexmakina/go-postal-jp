package main

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Download() (dst string, e error) {

	temp, err := os.MkdirTemp(config.Storagepath, TempDirName)
	if err != nil {
		e = err
		return
	}
	defer os.RemoveAll(temp)

	zippath := filepath.Join(temp, "zip")
	zp, err := os.Create(zippath)
	if err != nil {
		e = err
		return
	}
	defer zp.Close()

	r, err := http.Get(config.Url)
	if err != nil {
		e = err
		return
	}
	defer r.Body.Close()
	if _, err := io.Copy(zp, r.Body); err != nil {
		e = err
		return
	}

	z, err := zip.OpenReader(zippath)
	if err != nil {
		e = err
		return
	}
	defer z.Close()

	var csvName string

	for _, zf := range z.File {
		zff, err := zf.Open()
		if err != nil {
			continue
		}
		defer zff.Close()

		if zf.Name != config.CsvName {
			return
		}

		if zf.FileInfo().IsDir() {
			os.MkdirAll(zf.Name, os.ModePerm)
			continue
		}
		csvName = zf.Name
		csvf, err := os.OpenFile(
			filepath.Join(temp, csvName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zf.Mode())
		if err != nil {
			continue
		}
		defer csvf.Close()
		_, err = io.Copy(csvf, zff)
		if err != nil {
			continue
		}
	}

	sjisByte, err := os.ReadFile(filepath.Join(temp, csvName))
	if err != nil {
		e = err
		return
	}

	utf8byte, err := SjisByteToUtf8(sjisByte)
	if err != nil {
		e = err
		return
	}

	dst = filepath.Join(config.Storagepath, csvName)
	os.WriteFile(dst, utf8byte, 0777)

	return
}
