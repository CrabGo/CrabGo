# CrabGo



## 常用命令之 go mod

| 命令       | 说明              |
|----------|-----------------|
| download | 下载依赖包           |
| init     | 初始化模块           |
| tidy     | 拉取缺少的模块,移除不用的模块 |
| vendor   | 将依赖复制到vendor下   |

## tablewriter编译错误

使用gfv2.9.5 最新版本的tablewriter出现编译错误
需将tablewriter将至1.1.0版本

`
go get github.com/olekukonko/tablewriter@v1.1.0
`
对应版本在github上查看:https://github.com/gogf/gf/blob/master/go.mod