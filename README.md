Pocketの記事を表示、タグ検索、まとめてタグ付ができるサービスです。

フロントエンドはVue、バックエンドはGo言語を使用しています。

## 開発環境立ち上げ
webpackのbuild

```
# 開発環境
yarn build:dev
```

Dockerでのバックエンドサーバ立ち上げ

```
docker build ./ -f Dockerfile.local -t pocket_app
docker run -it --env-file .env -v $(pwd)/go/src/pocket-organize-app:/usr/local/go/src/pocket-organize-app -v $(pwd)/dist/js:/usr/local/go/src/pocket-organize-app/dist/js -p 5000:80 pocket_app
```

## 画面イメージ
![image](https://user-images.githubusercontent.com/30525452/119951507-fe94c000-bfd6-11eb-8441-417e7a4d0101.png)

## Author
Yuki Ikegaya
