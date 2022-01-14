# golang-study


## GO modules
```
PS D:\go\src\golang-study> go env
set GO111MODULE=on
set GOARCH=amd64
```
goproxy 设置包下载代理
GOPRIVATE 设置私有项目，不需要通过代理下载，也不校验完整性
` go env -w GOPRIVATE="*.example.com"`

