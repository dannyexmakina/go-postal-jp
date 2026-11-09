package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// Download は日本郵便の郵便番号データ（UTF-8版ZIP）を取得し、
// 中のCSVを Storagepath に展開してそのパスを返す。
func Download() (dst string, e error) {
	if err := os.MkdirAll(config.Storagepath, 0755); err != nil {
		return "", err
	}
	temp, err := os.MkdirTemp(config.Storagepath, TempDirName)
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(temp)

	zippath := filepath.Join(temp, "zip")
	if err := downloadFile(config.Url, zippath); err != nil {
		return "", err
	}

	z, err := zip.OpenReader(zippath)
	if err != nil {
		return "", err
	}
	defer z.Close()

	for _, zf := range z.File {
		if zf.FileInfo().IsDir() || zf.Name != config.CsvName {
			continue
		}
		dst = filepath.Join(config.Storagepath, zf.Name)
		if err := extractFile(zf, dst); err != nil {
			return "", err
		}
		return dst, nil
	}

	return "", fmt.Errorf("%s not found in %s", config.CsvName, config.Url)
}

func downloadFile(url, path string) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: %s: %s", url, r.Status)
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, r.Body)
	return err
}

func extractFile(zf *zip.File, dst string) error {
	src, err := zf.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	f, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, src)
	return err
}
