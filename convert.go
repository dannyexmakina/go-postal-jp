package main

import (
  "path/filepath"
  "os"

  "errors"

)

/*
Parse CSV from JapanPost

( See http://www.post.japanpost.jp/zipcode/dl/kogaki-zip.html )
*/
func Convert(arg []string) (e error) {
  if len(arg) < 2 {
    e = errors.New("Invalid args. input file empty. see help with -h")
    return
  }
  infile := arg[0]
  outfile := arg[1]

// args
  if len(infile) < 1 {
    e = errors.New("Invalid args. input file empty. see help with -h")
    return
  }
  if len(outfile) < 1 {
    e = errors.New("Invalid args. output directory empty. see help with -h")
    return
  }
//
  ipath,iperr := filepath.Abs(infile)
  if iperr != nil {
      return iperr
  }

  opath,operr := filepath.Abs(outfile)
  if operr != nil {
      return operr
  }

  pe := ParseFile(ipath)
  if pe != nil {
    return pe
  }

	stserr := SqliteToSql(opath)
  if stserr != nil {
      return stserr
  }

  if err := os.Remove(config.Sqlite); err != nil {
    Log("Info",err)
  }

  return nil

}
