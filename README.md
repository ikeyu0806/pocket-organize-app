webpackのbuild

```
# 開発環境
yarn build:dev
```

Dockerでのバックエンドサーバ立ち上げ

```
docker build ./ -t pocket_app
docker run -it --env-file .env -v $(pwd)/go/src/pocket-organize-app:/usr/local/go/src/pocket-organize-app -v $(pwd)/dist/js:/usr/local/go/src/pocket-organize-app/dist/js -p 5000:80 pocket_app
```
