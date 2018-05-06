# 実行時設定

- `assets/config.toml`

#### `url`

- ZIPファイル取得先URL
  - 通常は変更する必要が無い

#### `chunk`

- レコードを分割処理する件数
  - 通常は 1000 のまま変更する必要が無い

## `sql`

#### `path`

- SQLファイル保存先


## `sqlite`

#### `path`

- SQLiteファイルの生成先
  - 通常は `":memory:"` のまま変更する必要が無い
  

## `mysql`

#### `dsn`

- DB接続先
  - 例）dsn=root:password@tcp(127.0.0.1:3306)/test?charset=utf8

