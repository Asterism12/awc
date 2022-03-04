:: windows下交叉编译Linux下的可执行文件
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o gateway
