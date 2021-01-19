package controllers

import (
	"io"
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
)

// var Cfg = beego.AppConfig

type UploadController struct {
	beego.Controller
}

func (c *UploadController) Get() {
	//vue模板冲突https://blog.csdn.net/qq_34067821/article/details/86572502
	c.TplName = "upload.html"
}

func (c *UploadController) Post() {
	//GetFiles return multi-upload files
	// https://www.cnblogs.com/prince5460/archive/2004/01/13/12000046.html
	files, err := c.GetFiles("myFile")
	add := c.GetString("upload_path")
	pathstr := Cfg.String("path")
	pathstr = filepath.Join(pathstr, add)
	// beego.Debug(files)
	// beego.Debug(pathstr)
	if err != nil {
		beego.Error(err)
		return
	}
	for i, _ := range files {
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			beego.Error(err)
			return
		}
		//create destination file making sure the path is writeable.
		okpath := filepath.Join(pathstr, files[i].Filename)
		beego.Debug(okpath)
		dst, err := os.Create(okpath)
		defer dst.Close()
		if err != nil {
			beego.Error(err)
			return
		}
		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			beego.Error(err)
			return
		}
	}
	c.Redirect("/upload", 301)
	return
}
