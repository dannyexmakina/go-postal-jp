# go-postal-jp

[日本語](README.md) | English

- A command-line tool that converts the postal code CSV published by Japan Post into a fixed-format SQL / SQLite DB
  - <http://www.post.japanpost.jp/zipcode/dl/kogaki-zip.html>
- Intended for use cases such as auto-filling an address after a postal code is entered in a form

## Terminology

- Town area (*chōiki*)
  - Large-volume business office names, *ōaza*, *chōme*, etc.
  - See <http://www.post.japanpost.jp/zipcode/dl/readme.html>
- Local government code
  - Code based on JIS X0401 / X0402
- *Kogaki* (small kana)
  - Writing geminate and contracted sounds with small kana as normally written (e.g. ホッカイドウ, シンジュク)
  - Without *kogaki*, those sounds are written with full-size kana (e.g. ホツカイドウ, シンジユク)

## Assumptions based on Japan Post's data publishing

- Japan Post updates the CSV only around the end of each month, so the tool is expected to be run about once a month

### Assumptions and behavior based on postal code specifications

- Some postal codes have no address
  - e.g. "以下に掲載がない場合" (not listed below)
  - Only the prefecture is kept; city and town area are left empty
- A single postal code can map to multiple cities / town areas
  - Only the prefecture is kept; city and town area are left empty
- Some postal codes span multiple prefectures
  - Very few (3 as of January 2018)
  - Only the prefecture that appears first in the CSV is kept; city and town area are left empty

## String processing

- The UTF-8 data (`utf_ken_all.csv`) is used as is (no character encoding conversion)
- Half-width kana -> full-width kana (normally a no-op, since the UTF-8 data already uses full-width kana)
- If the town area contains full-width `（）` or half-width `()` parentheses, the enclosed part is removed
  - It is a note about the extent of the town area
- "以下に掲載がない場合" and "イカニケイサイガナイバアイ" are removed

## Build

- `task build`

## Commands

- Download
  - `./build/go-postal-jp download`

- Parse and generate SQL / SQLite / JSON
  - `./build/go-postal-jp parse {CSV file}`

- Print SQL
  - `./build/go-postal-jp get sql`

- Print JSON
  - `./build/go-postal-jp get json`

- Download, parse, and generate SQL / SQLite / JSON
  - `./build/go-postal-jp run`

## Generated files

- `storage/postal.sql`
  - SQL file (INSERT statements)
- `storage/postal.sqlite`
  - SQLite DB
- `storage/postal.json`
  - JSON

## Notes for server-side use

#### DB schema and usage notes

- [docs/create.sql](docs/create.sql)
- [docs/db.md](docs/db.md)

#### Runtime configuration

- [docs/config.md](docs/config.md)

#### Swagger file for API generation

- [docs/swagger.yml](docs/swagger.yml)

## License

[MIT](LICENSE)
