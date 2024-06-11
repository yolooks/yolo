# yolo
------------------------------
Yolook - Look Good, Feel Good!

## install

```
go install github.com/yolooks/yolo@latest
```
说明: yolo默认安装在$GOPATH/bin目录下

## usage

```
$GOBIN/yolo init -name <项目名> -port 端口号
```

## run

```
cd <项目名>/cmd
go run server.go
```

## test

```
curl http://127.0.0.1:<端口号>/v1/liveness
```
