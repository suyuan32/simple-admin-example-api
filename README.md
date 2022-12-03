# simple-admin-example-api
Ent api 生成例子

# 生成命令

```shell
goctls api new example --i18n=true --casbin=true --goZeroVersion=v1.4.2 --toolVersion=v0.0.9 --transErr=true --moduleName=github.com/suyuan32/simple-admin-example-api --port=8081

goctls api proto --proto=/home/ryan/GolandProjects/simple-admin-example-rpc/example.proto --style=go_zero --serviceName=example --o=./ --model=Student --rpcName=Example --grpcPackage=github.com/suyuan32/simple-admin-example-rpc/example
```