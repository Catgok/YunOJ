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
     etcdctl get --prefix ""
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
    ```
   - 4 统计代码行数
    ```shell
   find ./ -name "*.go" -exec cat {} + | wc -l
    ```



