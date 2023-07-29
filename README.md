# simple-admin-example-api tool v1.1.0 gen by tools v1.5.16
Ent api 生成例子

# 生成命令

```shell
goctls api new example --i18n=true --casbin=true --trans_err=true --module_name=github.com/suyuan32/simple-admin-example-api --port=8081

# or 
# goctls api new example -i -c -a -m github.com/suyuan32/simple-admin-example-api -p 8081 -g


cd example

go mod tidy

# linux
goctls api proto --proto=/home/ryan/GolandProjects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=Example --rpc_service_name=Example --output=./ --model=Student --rpc_name=Example --grpc_package=github.com/suyuan32/simple-admin-example-rpc/types/example
goctls api proto --proto=/home/ryan/GolandProjects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=Example --rpc_service_name=Example --output=./ --model=Teacher --rpc_name=Example --grpc_package=github.com/suyuan32/simple-admin-example-rpc/types/example

# windows
# goctls api proto --proto=D:/projects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=Example --rpc_service_name=Example --output=./ --model=Teacher --rpc_name=Example --grpc_package=github.com/suyuan32/simple-admin-example-rpc/types/example --output=./  overwrite=true
# goctls api proto --proto=D:/projects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=Example --rpc_service_name=Example --output=./ --model=Student --rpc_name=Example --grpc_package=github.com/suyuan32/simple-admin-example-rpc/types/example --output=./  overwrite=true

# 简写 | short
# goctls api proto -p D:/projects/simple-admin-example-rpc/example.proto -s go_zero -a Example -r Example -o ./ -m Student -n Example -g github.com/suyuan32/simple-admin-example-rpc/types/example -w

# goctls api go --api ./desc/all.api --dir ./ --trans_err=true
make gen-api

go mod tidy
```

> 还需在 service context, config, etc 添加 ExampleRpc

> 新增命令 `make help` ，可以查看所有命令，旧RPC项目只需要复制本项目的 `Makefile`, 将内部的第一行的 `PROJECT` 设置为自己的项目名称的小写即可。
> `Dockerfile` 同样复制到自己的项目中修改下 `ARG` 参数

> Add the new command `make help`, you can view all the commands, the old RPC project only needs to copy the `Makefile` of this project, and set `PROJECT` in the first line inside to the lowercase of your own project name.
> `Dockerfile` is also copied to your own project to modify the `ARG` parameter