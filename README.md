# simple-admin-example-api tool v0.2.9-beta
Ent api 生成例子

# 生成命令

```shell
goctls api new example --i18n=true --casbin=true --go_zero_version=v1.5.0 --tool_version=v0.2.8 --trans_err=true --module_name=github.com/suyuan32/simple-admin-example-api --port=8081

cd example

go mod tidy

# linux
goctls api proto --proto=/home/ryan/GolandProjects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=example --rpc_service_name=example --o=./ --model=Student --rpc_name=Example --grpc_package=github.com/suyuan32/simple-admin-example-rpc/types/example
goctls api proto --proto=/home/ryan/GolandProjects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=example --rpc_service_name=example --o=./ --model=Teacher --rpc_name=Example --grpc_package=github.com/suyuan32/simple-admin-example-rpc/types/example

# windows
# goctls api proto --proto=D:/projects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=example --rpc_service_name=example --o=./ --model=Teacher --rpc_name=Example --grpc_package=github.com/suyuan32/simple-admin-example-rpc/types/example --overwrite=true
# goctls api proto --proto=D:/projects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=example --rpc_service_name=example --o=./ --model=Student --rpc_name=Example --grpc_package=github.com/suyuan32/simple-admin-example-rpc/types/example --overwrite=true

# goctls api go --api ./desc/all.api --dir ./ --trans_err=true
make gen-api

go mod tidy
```

> 还需在 service context, config, etc 添加 ExampleRpc