# web_app
## docker 起動
```
$ docker build -t chat_image .
$ docker run -d -p 8000:8000 chat_image
```


## cloneしたときにすること
```bash
$ go get -u golang.org/x/vgo
```

.envに必要な環境変数を記述
```
GOOGLE_API_KEY=<GOOGLE_CLIENT_KEY>
GOOGLE_API_SECRET=<GOOGLE_CLIENT_SECRET_KEY>
```

## ビルド方法
importに追記したら以下を実行する。  
勝手にgo getが走ります。  
```
$ vgo build -o chat
```


## docker-composeを使う場合

```
$ docker-compose up -d
```
