package routers

import (
	"upload/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/file/?:path/?:path", &controllers.FileController{})
	// beego.Router("/file/", &controllers.FileController{})
	beego.Router("/file/*.*", &controllers.FileController{})
	beego.Router("/download/*.*", &controllers.DownloadController{})
	beego.Router("/upload", &controllers.UploadController{})
}
