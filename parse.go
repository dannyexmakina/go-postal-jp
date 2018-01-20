package main

import (
  "golang.org/x/text/encoding/japanese"
  "golang.org/x/text/transform"
  "encoding/csv"
  "os"
  "io"
  "strings"
  "time"

)

func ParseFile(filename string) (e error) {
  f, fe := os.OpenFile(filename, os.O_RDONLY, 0644)
  if fe != nil {
      return fe
  }
  defer f.Close()

  reader := csv.NewReader(transform.NewReader(f,japanese.ShiftJIS.NewDecoder()))
  reader.Comma = ','
  reader.LazyQuotes = true

  var ps Postals
  var counter int
  for {
    counter++

    line, err := reader.Read()
    if err == io.EOF {
      break
    } else if err != nil {
      return err
    }
    p := ParseLine(line)
    ps = append(ps, p)

    if counter > config.Chunk {
      ser := ps.SaveToSqlite() 
      if ser != nil {
        return ser
      }
      counter = 0
      ps = Postals{}
      time.Sleep(1000 * time.Millisecond)
    }

  }

  err := ScreeningSqlite()
  if err != nil {
    Log("Info",err)
  }

  return nil
}



func ParseLine(line []string) (p Postal) {
  var newline []string
  for _,v := range line {
    v = KanaHalfToFull(v)
    if(v == "イカニケイサイガナイバアイ" || v == "以下に掲載がない場合"){
      v = ""
    }
    newline = append(newline,v);
  }
  code := newline[2]
  jiscode := newline[0]
  ps := strings.Split(code,"")
	p = Postal{
		First: code[0:3],
		Last: code[3:],
		P1: ps[0],
		P2: ps[1],
		P3: ps[2],
		P4: ps[3],
		P5: ps[4],
		P6: ps[5],
		P7: ps[6],
		Code: code,
		Prefcode: jiscode[0:2],
		Pref: newline[6],
		Pref_kana: newline[3],
		Jiscode: jiscode,
    Street: newline[8],
    Street_kana: newline[5],
    City: newline[7],
    City_kana: newline[4],
	}
	return p
}
