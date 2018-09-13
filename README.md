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

## ビルド方法
importに追記したら以下を実行する。  
勝手にgo getが走ります。  
```
$ vgo build -o chat
```
