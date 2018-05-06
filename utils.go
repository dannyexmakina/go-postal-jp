package main

import (
  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
  "fmt"
  "os"
  "strings"

)

func Log(level string, messages ...interface{}) {
  log,_ := zap.NewDevelopment()
  log = log.WithOptions(zap.AddCallerSkip(1))
  defer log.Sync()
  
  var list []zapcore.Field
  for i,v := range messages {
    index := fmt.Sprintf("%d",(i + 1))
    list = append(list, zap.Any(index, v))
  }
  switch strings.ToUpper(level) {
    case "INFO":
      if !Quiet {
        log.Info("-", list...)
      }
    case "DEBUG":
      if !Quiet {
        log.Debug("-", list...)
      }
    case "ERROR":
      log.Error("-", list...)
    case "ERR":
      log.Error("-", list...)
    case "WARN":
      log.Warn("-", list...)
    case "WARNING":
      log.Warn("-", list...)
    case "FATAL":
      log.Fatal("-", list...)
    default:
      log.Fatal("-", list...)
  }
}



// Convert Hankaku-Kana to Zenkaku-kana
func KanaHalfToFull(input string) string{
  replacer := strings.NewReplacer(
    "ｶﾞ", "ガ",
    "ｷﾞ", "ギ",
    "ｸﾞ", "グ",
    "ｹﾞ", "ゲ",
    "ｺﾞ", "ゴ",
    "ｻﾞ", "ザ",
    "ｼﾞ", "ジ",
    "ｽﾞ", "ズ",
    "ｾﾞ", "ゼ",
    "ｿﾞ", "ゾ",
    "ﾀﾞ", "ダ",
    "ﾁﾞ", "ヂ",
    "ﾂﾞ", "ヅ",
    "ﾃﾞ", "デ",
    "ﾄﾞ", "ド",
    "ﾊﾞ", "バ",
    "ﾊﾟ", "パ",
    "ﾋﾞ", "ビ",
    "ﾋﾟ", "ピ",
    "ﾌﾞ", "ブ",
    "ﾌﾟ", "プ",
    "ﾍﾞ", "ベ",
    "ﾍﾟ", "ペ",
    "ﾎﾞ", "ボ",
    "ﾎﾟ", "ポ",
    "ｳﾞ", "ヴ",
    "ｱ", "ア",
    "ｲ", "イ",
    "ｳ", "ウ",
    "ｴ", "エ",
    "ｵ", "オ",
    "ｶ", "カ",
    "ｷ", "キ",
    "ｸ", "ク",
    "ｹ", "ケ",
    "ｺ", "コ",
    "ｻ", "サ",
    "ｼ", "シ",
    "ｽ", "ス",
    "ｾ", "セ",
    "ｿ", "ソ",
    "ﾀ", "タ",
    "ﾁ", "チ",
    "ﾂ", "ツ",
    "ﾃ", "テ",
    "ﾄ", "ト",
    "ﾅ", "ナ",
    "ﾆ", "ニ",
    "ﾇ", "ヌ",
    "ﾈ", "ネ",
    "ﾉ", "ノ",
    "ﾊ", "ハ",
    "ﾋ", "ヒ",
    "ﾌ", "フ",
    "ﾍ", "ヘ",
    "ﾎ", "ホ",
    "ﾏ", "マ",
    "ﾐ", "ミ",
    "ﾑ", "ム",
    "ﾒ", "メ",
    "ﾓ", "モ",
    "ﾗ", "ラ",
    "ﾘ", "リ",
    "ﾙ", "ル",
    "ﾚ", "レ",
    "ﾛ", "ロ",
    "ﾔ", "ヤ",
    "ﾕ", "ユ",
    "ﾖ", "ヨ",
    "ｧ", "ァ",
    "ｨ", "ィ",
    "ｩ", "ゥ",
    "ｪ", "ェ",
    "ｫ", "ォ",
    "ｯ", "ッ",
    "ｬ", "ャ",
    "ｭ", "ュ",
    "ｮ", "ョ",
    "ﾜ", "ワ",
    "ﾝ", "ン",
  )
  return replacer.Replace(input);
}

func IsDir(path string) bool {
  fileinfo,err := os.Stat(path)
  if err != nil {
    return false
  }
  return fileinfo.Mode().IsDir()
}
