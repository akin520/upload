package controllers

import (
	"os"
	"path/filepath"

	// "fmt"
	"strings"
	"upload/utils"

	"github.com/astaxie/beego"
)

var Cfg = beego.AppConfig

type FileController struct {
	beego.Controller
}

func (c *FileController) Get() {
	paths := strings.Split(c.Ctx.Request.RequestURI, "/")
	beego.Debug(paths)
	beego.Debug(len(paths))
	pathstr := Cfg.String("path")
	add := ""
	if len(paths) <= 3 && paths[2] == "index" {
		beego.Debug("-----------------")
	} else {
		add = strings.Join(paths[2:], "/")
	}
	pathstr = filepath.Join(pathstr, add)
	// beego.Debug(pathstr)
	dirs, files, err := utils.ReadDF(pathstr)
	if err != nil {
		beego.Debug("get file error")
	}
	// beego.Debug(dirs)
	// beego.Debug(files)

	c.Data["Addpath"] = add
	c.Data["error_info"] = ""
	c.Data["Dirs"] = dirs
	c.Data["Files"] = files
	c.TplName = "index.html"
	return
}

func (c *FileController) Post() {
	dirs := c.Input().Get("dirs")
	pathstr := Cfg.String("path")
	pathstr = filepath.Join(pathstr, dirs)
	err := os.MkdirAll(pathstr, os.ModePerm)
	if err != nil {
		beego.Debug("mkdir error")
	}
	c.Redirect("/file/index", 301)
	return
}
