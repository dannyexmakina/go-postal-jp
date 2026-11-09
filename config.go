package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type PostalConfig struct {
	TempDir string
	Input   string
	Output  string

	Storagepath string
	Chunk       int
	Url         string
	CsvName     string
	Sql         PostalConfigSql
	Json        PostalConfigJson
	Sqlite      PostalConfigSqlite
	Mysql       PostalConfigMySql
}
type PostalConfigSqlite struct {
	Path string
}
type PostalConfigSql struct {
	Path string
}
type PostalConfigJson struct {
	Path string
}
type PostalConfigMySql struct {
	Dsn string
}

var config PostalConfig

func init() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Printf("Fatal: %+v", err)
	}

	config = PostalConfig{
		Storagepath: getEnv("STORAGE_PATH", "./storage"),
		Url:         getEnv("URL", "https://www.post.japanpost.jp/service/search/zipcode/download/utf/zip/utf_ken_all.zip"),
		Chunk:       getEnvInt("CHUNK", 1000),
		CsvName:     getEnv("CSV_NAME", "utf_ken_all.csv"),
		Sql:         PostalConfigSql{Path: getEnv("SQL_PATH", "./storage/postal.sql")},
		Sqlite:      PostalConfigSqlite{Path: getEnv("SQLITE_PATH", "./storage/postal.sqlite")},
		Json:        PostalConfigJson{Path: getEnv("JSON_PATH", "./storage/postal.json")},
		Mysql:       PostalConfigMySql{Dsn: getEnv("MYSQL_DSN", "")},
	}
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}

func getEnvInt(key string, def int) int {
	v, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		log.Printf("Fatal: invalid %s=%q: %+v", key, v, err)
		return def
	}
	return n
}
