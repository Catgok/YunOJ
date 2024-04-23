# YunOJ

- code 存放代码
- target 存放编译后的代码和docker-compose

## 服务划分

- 1
-
    - 1
-
    - 2

## 常用命令

1. etcd 相关
    - 启动etcd
    ```shell
    etcd
    ```
    - 查看注册到etcd中的服务
    ```shell
    etcdctl --endpoints=serverhost:12379 get --prefix ""
    ```
2. goctl 相关
    - 1 mysql model
    ```shell
    goctl model mysql ddl -c --src *.sql  --dir .
    ```
    - 2 rpc
    ```shell
    goctl rpc protoc  *.proto  --go_out=. --go-grpc_out=. --zrpc_out=.
    ```
    - 3 api
    ```shell
    goctl api go -api ./gateway.api  -dir .
    ```
    - 4 统计代码行数
    ```shell
    find ./ -name "*.go" -exec cat {} + | wc -l
    ```

3. 运行
    - 1 启动服务
    ```shell
    go run ./services/user/rpc/user.go -f ./services/user/rpc/etc/user.yaml
    go run ./services/problem/rpc/problem.go -f ./services/problem/rpc/etc/problem.yaml
    go run ./services/judge/rpc/judge.go -f ./services/judge/rpc/etc/judge.yaml
    ```
    - 2 启动网关
    ```shell
    go run ./services/gateway/api/gateway.go -f ./services/gateway/api/etc/gateway.yaml
    ```
    - 3 启动web
    ```shell
    vue-cli-service serve
    ```

