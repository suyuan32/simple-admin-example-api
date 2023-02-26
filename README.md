# simple-admin-example-api
Ent api 生成例子

# 生成命令

```shell
goctls api new example --i18n=true --casbin=true --go_zero_version=v1.4.4 --tool_version=v0.2.1 --trans_err=true --module_name=github.com/suyuan32/simple-admin-example-api --port=8081

cd example

go mod tidy

goctls api proto --proto=/home/ryan/GolandProjects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=example --rpc_service_name=example --o=./ --model=Student --rpc_name=Example --grpc_package=github.com/suyuan32/simple-admin-example-rpc/example --multiple=true

goctls api proto --proto=/home/ryan/GolandProjects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=example --rpc_service_name=school --o=./ --model=Teacher --rpc_name=School --grpc_package=github.com/suyuan32/simple-admin-example-rpc/example --multiple=true

# goctls api go --api ./desc/all.api --dir ./ --trans_err=true
make gen-api
```

> 还需在 service context, config, etc 添加 ExampleRpc, SchoolRpc