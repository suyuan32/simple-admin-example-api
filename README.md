# simple-admin-example-api tool v1.10.2 gen by tools v1.10.2
Ent api 生成例子

# 生成命令

```shell
goctls api new example --i18n=true --casbin=true --trans_err=true --module_name=github.com/suyuan32/simple-admin-example-api --port=8081 --use_core_rpc=true

# or 
# goctls api new example -i -c -a -m github.com/suyuan32/simple-admin-example-api -p 8081 -r


cd example

go mod tidy

# linux
goctls api proto --proto=/home/ryan/GolandProjects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=Example --rpc_service_name=Example --output=./ --model=Student --rpc_name=Example --grpc_package=github.com/suyuan32/simple-admin-example-rpc/types/example --i18n=true --api_data=true
goctls api proto --proto=/home/ryan/GolandProjects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=Example --rpc_service_name=Example --output=./ --model=Teacher --rpc_name=Example --grpc_package=github.com/suyuan32/simple-admin-example-rpc/types/example --i18n=true --api_data=true

# windows
# goctls api proto --proto=D:/projects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=Example --rpc_service_name=Example --output=./ --model=Teacher --rpc_name=Example --grpc_package=github.com/suyuan32/simple-admin-example-rpc/types/example --output=./  overwrite=true --i18n=true --api_data=true
# goctls api proto --proto=D:/projects/simple-admin-example-rpc/example.proto --style=go_zero --api_service_name=Example --rpc_service_name=Example --output=./ --model=Student --rpc_name=Example --grpc_package=github.com/suyuan32/simple-admin-example-rpc/types/example --output=./  overwrite=true --i18n=true --api_data=true

# 简写 | short
# goctls api proto -p D:/projects/simple-admin-example-rpc/example.proto -s go_zero -a Example -r Example -o ./ -m Student -n Example -g github.com/suyuan32/simple-admin-example-rpc/types/example -w -i -d

# goctls api go --api ./desc/all.api --dir ./ --trans_err=true
make gen-api

go mod tidy
```

> 还需在 service context, config, etc 添加 ExampleRpc

## 添加初始化代码

在 logic/base 文件夹内手动添加 API 和 menu 初始化代码