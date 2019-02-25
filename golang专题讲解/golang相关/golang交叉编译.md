### 1、Mac 下编译 Linux 和 Windows 64位可执行程序
* CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
* CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

### 2、Linux 下编译 Mac 和 Windows 64位可执行程序
* CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
* CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

### 3、Windows 下编译 Mac 和 Linux 64位可执行程序
* SET CGO_ENABLED=0
  SET GOOS=darwin
  SET GOARCH=amd64
  go build main.go
* SET CGO_ENABLED=0
  SET GOOS=linux
  SET GOARCH=amd64
  go build main.go

