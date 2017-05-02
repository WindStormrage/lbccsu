package routers

import (
	"lbccsu/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Include(&controllers.LoginController{})
}
