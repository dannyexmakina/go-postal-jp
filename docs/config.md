# 実行時設定

- `.env`（`.env.example` をコピーして作成。未設定の項目は既定値を使用）

#### `URL`

- ZIPファイル取得先URL
  - 通常は変更する必要が無い

#### `CHUNK`

- レコードを分割処理する件数
  - 通常は 1000 のまま変更する必要が無い

#### `STORAGE_PATH`

- 生成物の保存ディレクトリ

#### `CSV_NAME`

- ZIP内のCSVファイル名

#### `SQL_PATH`

- SQLファイル保存先

#### `JSON_PATH`

- JSONファイル保存先

#### `SQLITE_PATH`

- SQLiteファイルの生成先

#### `MYSQL_DSN`

- DB接続先
  - 例）`MYSQL_DSN=root:password@tcp(127.0.0.1:3306)/test?charset=utf8`
