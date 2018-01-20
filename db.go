package main

import (
	"github.com/gocraft/dbr"
	"os"
  _ "github.com/go-sql-driver/mysql"

)

type CountRows struct {
	Count int
}

func UpdateDB() (e error) {

	dsnMy := os.Getenv("dsn")
  cRemote, merr := dbr.Open("mysql", dsnMy, nil)
  if merr != nil {
    return merr
  }
  defer cRemote.Close()

  cLocal, lerr := dbr.Open("sqlite3", config.Sqlite, nil)
  if lerr != nil {
    return lerr
  }
  defer cLocal.Close()

  sRemote := cRemote.NewSession(nil)
  sLocal := cLocal.NewSession(nil)

  var ps []Postal

  e = sLocal.Select("*").
        From("postal").
        LoadStruct(&ps)
  if e != nil {
		return e
  }
	for _,p := range ps {
		var c CountRows
		sRemote.Select("COUNT(*) AS count").
							From("postal").
							Where("code = ?", p.Code).
							LoadStruct(&c)
		if c.Count > 0 {
			_,ue := sRemote.Update("postal").
						Set("first",p.First).
						Set("last",p.Last).
						Set("p1",p.P1).
						Set("p2",p.P2).
						Set("p3",p.P3).
						Set("p4",p.P4).
						Set("p5",p.P5).
						Set("p6",p.P6).
						Set("p7",p.P7).
						Set("prefcode",p.Prefcode).
						Set("pref",p.Pref).
						Set("pref_kana",p.Pref_kana).
						Set("city",p.City).
						Set("city_kana",p.City_kana).
						Set("street",p.Street).
						Set("street_kana",p.Street_kana).
						Set("jiscode",p.Jiscode).
						Where("`code`=?",p.Code).
						Exec()
			if ue != nil {
				return ue
			}
		}else{
			_,re := sRemote.InsertInto("postal").
												Columns("first", "last", "p1", "p2", "p3", "p4", "p5", "p6", "p7", "code", "prefcode", "pref", "pref_kana", "city", "city_kana", "street", "street_kana", "jiscode").
												Values(p.First, p.Last, p.P1, p.P2, p.P3, p.P4, p.P5, p.P6, p.P7, p.Code, p.Prefcode, p.Pref, p.Pref_kana, p.City, p.City_kana, p.Street, p.Street_kana, p.Jiscode).
												Exec()
			if re != nil {
				return re
			}
		}
  }
  return nil
}
