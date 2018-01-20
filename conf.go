package main

import (
	"github.com/joho/godotenv"

)

type PostalConfig struct {
  Input string
  Output string
  Sqlite string
  Chunk int
}
var config PostalConfig

var postal_default_output = "KEN_ALL.sql"
var postal_csv_name = "KEN_ALL.CSV"

type Postals []Postal

type Postal struct {
  Id int
  First string
  Last string
  P1 string
  P2 string
  P3 string
  P4 string
  P5 string
  P6 string
  P7 string
  Code string
  Prefcode string
  Pref string
  Pref_kana string
  City string
  City_kana string
  Street string
  Street_kana string
  Jiscode string
}

func init(){
  config.Chunk = 1000

  if err := godotenv.Load(); err != nil {
    Log("Fatal",err)
  }
  if err := InitializeSqlite(); err != nil {
    Log("Fatal",err)
  }

}