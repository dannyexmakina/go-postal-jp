package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type PostalConfig struct {
	TempDir string `toml:"temp_dir"`
	Input   string
	Output  string

	AssetPath   string             `toml:"asset_path"`
	Storagepath string             `toml:"storage_path"`
	Chunk       int                `toml:"chunk"`
	Url         string             `toml:"url"`
	CsvName     string             `toml:"csv_name"`
	Sql         PostalConfigSql    `toml:"sql"`
	Sqlite      PostalConfigSqlite `toml:"sqlite"`
	Mysql       PostalConfigMySql  `toml:"mysql"`
}
type PostalConfigSqlite struct {
	Path string `toml:"path"`
	Save string `toml:"save"`
}
type PostalConfigSql struct {
	Path string `toml:"path"`
}
type PostalConfigMySql struct {
	Dsn string `toml:"dsn"`
}

var config PostalConfig

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Fatal: %+v", err)
	}

	configPath := os.Getenv("config_path")
	p, err := filepath.Abs(configPath)
	if err != nil {
		log.Printf("Fatal: %+v", err)
	}
	if _, err := toml.DecodeFile(p, &config); err != nil {
		log.Printf("Fatal: %+v", err)
	}

}
