package main

/**
e:\code\go\src\upload>bee version
______
| ___ \
| |_/ /  ___   ___
| ___ \ / _ \ / _ \
| |_/ /|  __/|  __/
\____/  \___| \___| v1.6.2

├── Beego     : 1.7.2
├── GoVersion : go1.8
├── GOOS      : windows
├── GOARCH    : amd64
├── NumCPU    : 8
├── GOPATH    : E:\code\go
├── GOROOT    : C:\Go\
├── Compiler  : gc
└── Date      : Tuesday, 19 Jan 2021
**/

import (
	_ "upload/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/auth"
)

func main() {
	authPlugin := auth.NewBasicAuthenticator(SecretAuth, "Authorization Required")
	beego.InsertFilter("*", beego.BeforeRouter, authPlugin)
	beego.Run()
}

//简单的http验证
func SecretAuth(username, password string) bool {
	return username == "admin" && password == "admin.."
}
