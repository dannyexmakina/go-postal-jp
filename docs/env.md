# 実行時設定

- `.env` によって実行時設定を記述・設置する

## `.env` ファイル

### `url`

- ZIPファイル取得先URL
  - 通常は変更する必要が無い

### `temp`

- テンポラリディレクトリ
  - ファイルDL・ZIPファイル一時保存に使用するディレクトリを指定
  - 相対PATH指定の場合はバイナリ実行時のカレントPATHを基点とする
  - 例）temp=temp

### `sqlite`

- SQLiteファイル保存先
  - SQLiteの仕様上、R/W可能なディレクトリ以下であることが必須
  - 例）sqlite=temp/db

### `dsn`

- DB接続先
  - 例）dsn=root:password@tcp(127.0.0.1:3306)/test?charset=utf8

