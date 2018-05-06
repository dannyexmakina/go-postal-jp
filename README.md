# postal

- 日本郵便が提供する郵便番号CSVを定型のSQL/SQLiteDBに変換するコマンドライン
  - http://www.post.japanpost.jp/zipcode/dl/kogaki-zip.html
- フォームにて郵便番号を入力し、住所を自動入力する用途などを想定している

## 用語

- 町域
  - 大口事業所名・大字・丁目など
  - see. http://www.post.japanpost.jp/zipcode/dl/readme.html
- 全国地方公共団体コード
  - JIS X0401・X0402に基づくコード
- 小書き
  - 促音・拗音を通常表記通りに記載すること（例：ホッカイドウ・シンジュク）
  - 小書きしない場合は促音・拗音を大文字にて記載することを意味する（例：ホツカイドウ・シンジユク）


## 日本郵便のデータ提供運用に基づく前提

- 日本郵便提供CSVは毎月末付近にのみ更新されるため、１ヶ月に１度の頻度で実行することを前提とする


#### 郵便番号仕様に基づく前提・挙動仕様

- 住所が存在しない郵便番号が存在する
  - 例：「以下に掲載が無い場合」など
  - 都道府県のみ有効とし、市区町村郡・町域は空欄とする
- 同一郵便番号に複数の市町村郡・町域が存在する
  - 都道府県のみ有効とし、市区町村郡・町域は空欄とする
- 都道府県をまたがる郵便番号が存在する
  - ごく小数（2018/1現在3件）である
  - CSV内に初出の都道府県のみ有効とし、市区町村郡・町域は空欄とする


## 文字列処理仕様

- 文字コードをCP392 -> UTF-8に変換
- 半角カナ -> 全角カナに変換
- 町域に「（）」「()」が含まれる場合は該当部分を削除
  - 町域範囲についての但し書きであるため

## 専用DBスキーマおよび使用上の注意事項

- `env/postal/docker/mysql/docker-entrypoint-initdb.d/create.sql`
- [docs/db.md](docs/db.md)

## 実行時設定

- [docs/env.md](docs/env.md)
- [docs/config.md](docs/config.md)

## 実行コマンド

- `cd env/postal`
- `docker-compose up`

## 生成されるファイル

- `db/postal.sql`
  - SQLファイル（INSERT文・他DBで使用する用）
- `db/postal.db`
  - SQLite DB（そのまま使用可能）


## API生成用 Swaggerファイル

- [swagger.yml](swagger.yml)

