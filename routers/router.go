package routers

import (
	"github.com/astaxie/beego"
	"github.com/hota1024/beego-test/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/users", &controllers.FirstController{}, "get:GetUsers")
}
