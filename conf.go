package main

import (
  "github.com/joho/godotenv"
	"github.com/BurntSushi/toml"

  "os"
  "path/filepath"
)


type Postal struct {
  Dir string
  Conf PostalConfig

}

type PostalConfig struct {
  Input string
  Output string

  AssetPath string
  Chunk int
  Url string
  CsvName string
  Sql PostalConfigSql
  Sqlite PostalConfigSqlite
  Mysql PostalConfigMySql
}
type PostalConfigSqlite struct {
  Path string
}
type PostalConfigSql struct {
  Path string
}
type PostalConfigMySql struct {
  Dsn string
}

var config PostalConfig

type Postals []PostalData

type PostalData struct {
  Id int
  First string
  Last string
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

type CountRows struct {
	Count int
}

func init(){
  config.Chunk = 1000

  if err := godotenv.Load(); err != nil {
    Log("Fatal",err)
  }

  assetPath := os.Getenv("asset_path")
	_, cerr := toml.DecodeFile(filepath.Join(assetPath, "config.toml"), &config)
	if cerr != nil {
    Log("Fatal",cerr)
	}
  config.AssetPath = assetPath

}