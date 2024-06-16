# echoserver

echo server 是一個 go 專案的基礎框架, 可以用來快速組建一個 RESTful API server

## Run the service as a docker container

```
# build docker image first
$ make docker-image

# run service
$ docker run -itd --name echoserver -p 8087:8080 echoserver
```

## Try it out

```
$ curl http://localhost:8087/api/v1/echo
```

## 透過此專案(template)快速建立其他專案

1. 將此專案 git clone 到本地
2. 將專案內含有字串 "github.com/Krados/echoserver" 的地方全部 replace 成你要的專案名稱

## Test Cloud Build
