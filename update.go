package main

import (
  "path/filepath"
  "os"
  "errors"

)


func UpdateAuto() (e error) {
  name, e := Download()
  if e != nil {
    return e
  }
  files,ze := Unzip(name)
  if ze != nil {
    return ze
  }
  var infile string
  for _,file := range files {
    if filepath.Base(file) == postal_csv_name {
      infile = file
      break
    }
  }
  e = ParseFile(infile)
  if e != nil {
    Log("Warn",e)
    return
  }

	e = UpdateDB()
  if e != nil {
    Log("Warn",e)
      return
  }

  if err := os.Remove(config.Sqlite); err != nil {
    Log("Info",err)
  }

  e = os.Remove(infile)
  if e != nil {
    Log("Warn",e)
  }

  return nil
}


func Update() (e error) {
  if len(config.Input) < 1 {
    e = errors.New("Invalid args. input file empty. see help with -h")
    return
  }
  infile := config.Input
//
  ipath,iperr := filepath.Abs(infile)
  if iperr != nil {
      return iperr
  }

  e = ParseFile(ipath)
  if e != nil {
    return
  }

	e = UpdateDB()
  if e != nil {
      return
  }

  if e = os.Remove(config.Sqlite); e != nil {
    Log("Info",e)
  }

  if e = os.Remove(infile); e != nil {
    Log("Warn",e)
  }

  return nil
}
