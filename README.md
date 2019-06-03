# xiao-gogo

### golang snip

1. 重试：https://github.com/kamilsk/retry
2. web服务器：https://github.com/gin-gonic/gin

### go mod
0. 创建环境变量
- export GO111MODULE=auto
- export GOPROXY=https://goproxy.io

1. 初始化 
- go mod init github.com/githubao/xiao-gogo

1. 导入所有依赖的包 // 使用哪个包就导入哪个更新go.mod
- go build main.go

1. 下载最近的包到mod文件夹
- go mod download 

1. 替换到本地的包
replace old => new

### 其他备忘
1. unset GOPATH
1. use vgo in goland

