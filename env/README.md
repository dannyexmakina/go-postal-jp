# inquiry / env

- 開発環境

## 必須事項

- dockerおよびdocker-composeの動作

## 初期設定

- docker-compose build

## コンテナ起動

- docker-compose up -d
  - frontコンテナの起動後、HTTP通信が有効になるまで多少時間がかかることに注意


## Goソースを書き換え後、ビルドしてサーバ再起動

- docker-compose exec front /root/cli/restart

### dep

- docker-compose exec front /root/cli/dep ensure


