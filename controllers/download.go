package controllers

import (
	"net/url"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
)

type DownloadController struct {
	beego.Controller
}

func (c *DownloadController) Get() {

	//第一个参数是文件的地址，第二个参数是下载显示的文件的名称
	paths := strings.Split(c.Ctx.Request.RequestURI, "/")
	beego.Debug(paths)
	pathstr := Cfg.String("path")
	add := strings.Join(paths[2:], "/")
	pathstr = filepath.Join(pathstr, add)
	beego.Debug(pathstr)
	enEscapeUrl, _ := url.QueryUnescape(pathstr)
	beego.Debug(enEscapeUrl)
	c.Ctx.Output.Download(enEscapeUrl)
}
