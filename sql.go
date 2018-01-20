package main

import (
	"path/filepath"
	"os"

)

func SqliteToSql(outpath string) error {

	if(IsDir(outpath)){
	  outpath = filepath.Join(outpath, postal_default_output)
	}
	f, fe := os.OpenFile(outpath, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0777)
	if fe != nil {
		return fe
	}
	defer f.Close()

	ps,pe := GetSqlite()
	if pe != nil {
		return pe
	}

	for _,p := range ps {
		sql := "INSERT INTO postal (first, last, p1, p2, p3, p4, p5, p6, p7, code, prefcode, pref, pref_kana, city, city_kana, street, street_kana, jiscode) VALUES ('"
		sql += p.First + "','"
		sql += p.Last + "','"
		sql += p.P1 + "','"
		sql += p.P2 + "','"
		sql += p.P3 + "','"
		sql += p.P4 + "','"
		sql += p.P5 + "','"
		sql += p.P6 + "','"
		sql += p.P7 + "','"
		sql += p.Code + "','"
		sql += p.Prefcode + "','"
		sql += p.Pref + "','"
		sql += p.Pref_kana + "','"
		sql += p.City + "','"
		sql += p.City_kana + "','"
		sql += p.Street + "','"
		sql += p.Street_kana + "','"
		sql += p.Jiscode + "');\n"
		if _, err := f.WriteString(sql); err != nil {
			return err
		}
	}
	return nil
}
