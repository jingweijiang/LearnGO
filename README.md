Golang 学习笔记

# 环境搭建
## GOPROXY
```
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOPROXY=https://goproxy.io,direct
```
## GONOPROXY
direct 表示不要走代理，直接源码库下载
如果是公司内，可以使用 GONOPROXY 来屏蔽某些包的拉取，比如：
```azure
go env -w GOPROXY=https://gitlab.mycorp.com,direct
```
## GONOSUMDB
如果是公司内，可以使用 GONOSUMDB 来屏蔽某些包的校验，比如：
```azure
go env -w GONOSUMDB=gitlab.mycorp.com
```

## GOPRIVATE
如果设置了GOPRIVATE，它是GONOPROXY和GONOSUMDB的默认值，可以不用设置GONOPROXY和GONOSUMDB。
如果是公司内，可以使用 GOPRIVATE 来指定某些包的私有化拉取，比如：
```azure
go env -w GOPRIVATE=gitlab.mycorp.com/myproject
```