# echoserver

echo server 是一個 go 專案的基礎框架, 可以用來快速組建一個 RESTful API server

## Run the service as a docker container

```
# build docker image first
$ make docker-image

# run service
$ make docker-run
```

## Try it out

```
$ curl http://localhost:8087/api/v1/echo
```

## 透過此專案(template)快速建立其他專案

1. 將此專案 git clone 到本地

```bash
git clone https://github.com/Krados/echoserver.git
```

2. 將專案內含有字串 "github.com/Krados" 的地方全部 replace 成你要的專案名稱

```bash
find . -type f \( -name "*.go" -o -name "*.mod" \) -exec sed -i 's|github\.com/Krados|github.com/YourName|g' {} +
```

3. 將專案內含有字串 "echoserver" 的地方 replace 成你想要的

```bash
find . -type f \( -name "*.go" -o -name "*.mod" -o -name "Makefile" \) -exec sed -i 's|echoserver|yourserver|g' {} +
```

4. 重整 go mod

```bash
go mod tidy
```

5. 移除 .git

```bash
rm -rf .git
```

6. 重新命名專案資料夾

```bash
cd ..
mv echoserver yourserver
```

7. 初始化 git

```bash
git init
```

8. 大功告成