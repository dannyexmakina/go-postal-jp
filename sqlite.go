package main

import (
	"github.com/gocraft/dbr"
  "time"
  "os"
  "path/filepath"

  "database/sql"
  _ "github.com/mattn/go-sqlite3"

)

func InitializeSqlite() error {
  temp := os.Getenv("sqlite")
  file, e := filepath.Abs(temp)
  if e != nil {
    return e
  }
  _,se := os.Stat(file)
  if se == nil {
    os.Remove(file)
  }
  config.Sqlite = file
  db, _ := sql.Open("sqlite3", config.Sqlite)
  defer db.Close()
  _,err := db.Exec("CREATE TABLE postal (id integer primary key autoincrement, first, last, p1, p2, p3, p4, p5, p6, p7, code, prefcode, pref, pref_kana, city, city_kana, street, street_kana, jiscode);")
  if err != nil {
    return err
  }
  return nil
}

func GetSqlite() (ps []Postal, e error) {
  connection, err := dbr.Open("sqlite3", config.Sqlite, nil)
  if err != nil {
    return ps,err
  }
  defer connection.Close()
  session := connection.NewSession(nil)
  e = session.Select("*").
        From("postal").
        LoadStruct(&ps)
	return
}


func ScreeningSqlite() error {
  connection, err := dbr.Open("sqlite3", config.Sqlite, nil)
  if err != nil {
    return err
  }
  defer connection.Close()
  session := connection.NewSession(nil)

	var dups []Postal
	e := session.Select("code").
				From("postal").
				GroupBy("code").
				Having(dbr.Expr("COUNT(`code`) > 1")).
				LoadStruct(&dups)
	if e != nil {
		return e
	}

	var dupCodes []string
	var forDelete []Postal
	for _,p := range dups {
		dupCodes = append(dupCodes, p.Code)

		var dupsByCode []Postal
		dupe := session.Select("*").
					From("postal").
					Where("code=?", p.Code).
					LoadStruct(&dupsByCode)
		if dupe != nil {
			return dupe
		}
		var added bool
    var counter int
		for _,p := range dupsByCode{
      counter++
			for _,dp := range forDelete {
				if dp.Code == p.Code {
					added = true
				}
			}
			if added {
				_,de := session.DeleteFrom("postal").
							Where("id = ?", p.Id).
							Exec()
				if de != nil {
					Log("Warn", de)
				}
			} else {
				forDelete = append(forDelete, p)
				_,ue := session.Update("postal").
							Set("City","").
							Set("City_kana","").
							Set("Street","").
							Set("Street_kana","").
							Where("id = ?", p.Id).
							Exec()
				if ue != nil {
					Log("Warn", ue)
				}
			}
      if counter > config.Chunk {
        time.Sleep(1000 * time.Millisecond)
        counter = 0
      }
		}
	}

		return nil
}



func (ps Postals) SaveToSqlite() error {
  connection, err := dbr.Open("sqlite3", config.Sqlite, nil)
  if err != nil {
    return err
  }
  defer connection.Close()
  session := connection.NewSession(nil)
  q := session.InsertInto("postal").
                Columns("first", "last", "p1", "p2", "p3", "p4", "p5", "p6", "p7", "code", "prefcode", "pref", "pref_kana", "city", "city_kana", "street", "street_kana", "jiscode")
	for _,p := range ps {
		var check Postal
		e := session.Select("*").
					From("postal").
					Where("code = ?", p.Code).
					LoadStruct(&check)
		if e == nil {
//			Log("Info","dup code: ",p.Code)
			continue
		}
		q.Record(p)
	}
   _,re := q.Exec()
  if re != nil {
    return re
  }
	time.Sleep(100 * time.Millisecond)
	return nil
}

