# simple-admin-example-api tool v0.3.2
Ent api 生成例子

# 生成命令

```shell
goctls api new example --i18n=true --casbin=true --go_zero_version=v1.5.1 --tool_version=v0.3.2 --trans_err=true --module_name=github.com/suyuan32/simple-admin-example-api --port=8081

cd example

go mod tidy

# linux
goctls api proto --proto=/home/ryan/GolandProjects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=example --rpc_service_name=example --output=./ --model=Student --rpc_name=Example --grpc_package=github.com/suyuan32/simple-admin-example-rpc/types/example
goctls api proto --proto=/home/ryan/GolandProjects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=example --rpc_service_name=example --output=./ --model=Teacher --rpc_name=Example --grpc_package=github.com/suyuan32/simple-admin-example-rpc/types/example

# windows
# goctls api proto --proto=D:/projects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=example --rpc_service_name=example --output=./ --model=Teacher --rpc_name=Example --grpc_package=github.com/suyuan32/simple-admin-example-rpc/types/example --outputverwrite=true
# goctls api proto --proto=D:/projects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=example --rpc_service_name=example --output=./ --model=Student --rpc_name=Example --grpc_package=github.com/suyuan32/simple-admin-example-rpc/types/example --outputverwrite=true

# 简写 | short
# goctls api proto -p D:/projects/simple-admin-example-rpc/example.proto -s go_zero -a example -r example -o ./ -m Student -n Example -g github.com/suyuan32/simple-admin-example-rpc/types/example -w

# goctls api go --api ./desc/all.api --dir ./ --trans_err=true
make gen-api

go mod tidy
```

> 还需在 service context, config, etc 添加 ExampleRpc

> 新增命令 `make help` ，可以查看所有命令，旧RPC项目只需要复制本项目的 `Makefile`, 将内部的第一行的 `PROJECT` 设置为自己的项目名称的小写即可。
> `Dockerfile` 同样复制到自己的项目中修改下 `ARG` 参数

> Add the new command `make help`, you can view all the commands, the old RPC project only needs to copy the `Makefile` of this project, and set `PROJECT` in the first line inside to the lowercase of your own project name.
> `Dockerfile` is also copied to your own project to modify the `ARG` parameter