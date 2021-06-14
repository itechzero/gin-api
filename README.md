## 关于Gin
Gin 是一个用 Go 编写的 HTTP web 框架，一个拥有更好性能的 API 框架

### Use
```
cp .env.example .env
go mod tidy
go mod vendor
go run main.go
```

### Supervisor
~~~
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/gin-api main.go
~~~
#### Alpine
~~~
apk update
apk add supervisor
~~~
~~~
[program:gin-api-worker]
process_name=gin-api
command=/go/src/gin-api/bin/gin-api
autostart=true
autorestart=true
user=root
numprocs=1
redirect_stderr=true
stdout_logfile=/go/src/gin-api/bin/supervisorctl.log
~~~

### TODO
- [x] 钉钉消息
- [x] 配置文件