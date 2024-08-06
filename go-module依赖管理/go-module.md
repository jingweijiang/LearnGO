Go1.11 开始，Go 官方加入Go Module 支持
在Go1.12 成为默认支持

# 使用

## 1. 配置代理
```azure
go env -w GOPROXY=https://goproxy.io,direct
go env -w GONOSUMDB="*"
go env -w GONOSUMDB="github.com/user/project"
```

1. 初始化模块

```azure
go mod init github.com/user/project
```
2. 检测和清理依赖

```azure
go mod tidy
```
3. 安装指定包

```azure
go get -v github.com/user/project/package@v1.0.0 # 拉取指定版本的依赖模块
```
4. 升级依赖

```azure
go get -u   #自动更新项目依赖版本
go get -U github.com/user/project/package@v1.0.0 # 更新指定依赖模块到最新版本

```
5. replace 替换依赖

```azure
go mod edit -replace=github.com/user/project/package=github.com/user/project/package@v1.0.0
```
replace 也可以在 go.mod 文件中配置，如：

```azure    
replace github.com/user/project/package => github.com/user/project/package v1.0.0
```

6. 更多
```azure
go help mod

```

# go.mod 文件

go.mod 文件记录了项目的依赖信息，包括模块路径、版本号、依赖的其他模块等。

```azure
// 定义当前模块的路径S
module github.com/user/project

// 当前模块的依赖及版本
require (
    github.com/user/project/package v1.0.0
)

replace (
    github.com/user/project/package => github.com/user/project/package v1.0.0
)

``` 

# GOPATH

1.src
一般存放项目源码
2.bin
使用go install命令编译后的可执行文件S
3.pkg
存放项目依赖的模块缓存S