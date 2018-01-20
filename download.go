package main

import (
  "os"
  "io"
  "path/filepath"
  "net/http"

)

func Download() (n string, e error) {
  url := os.Getenv("url")
  temp := os.Getenv("temp")
  if _,e = os.Stat(temp); e != nil {
    return
  }
  n = filepath.Join(temp, "zip")
  t, err := os.Create(n)
  if err != nil {
      return n,err
  }
  defer t.Close()

  r,err := http.Get(url)
  if err != nil {
    return n,err
  }
  defer r.Body.Close()
  _,e = io.Copy(t, r.Body) 
  if e != nil {
    return n,e
  }
  return n,nil
}
