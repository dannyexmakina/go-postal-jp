package main

import (
  "path/filepath"
  "os"
  "io"
  "io/ioutil"
  "archive/zip"
)


func Unzip(zipfile string) (paths []string, e error) {
  temp := filepath.Dir(zipfile)
  z, err := zip.OpenReader(zipfile)
  if err != nil {
    return paths,err
  }
  defer func(){
    z.Close()
    os.Remove(zipfile)
  }()

  for _,f := range z.File {
    path, err := func(fp *zip.File) (p string, e error) {
      zf, err := fp.Open()
      defer zf.Close()
      if err != nil {
        return p,err
      }
      p = filepath.Join(temp, f.Name)
      if fp.FileInfo().IsDir() {
        os.MkdirAll(p, os.ModePerm)
        return
      }
      buf := make([]byte, fp.UncompressedSize)
      _, e = io.ReadFull(zf, buf)
      if e != nil {
        return
      }
      e = ioutil.WriteFile(p, buf, fp.Mode())
      return
    }(f)
    if err != nil {
      return paths,err
    }
    if len(path) > 0 {
      paths = append(paths, path)
    }
  }
  return
}

