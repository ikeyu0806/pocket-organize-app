webpackのbuild

```
webpack --config webpack.dev.js
```

Dockerでのサーバ立ち上げ

```
docker build ./ -t pocket_app
docker run -it --env-file .env -v $(pwd):/workspace/pocket_app -p 5000:80 pocket_app
```
