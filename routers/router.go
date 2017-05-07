package routers

import (
	"lbccsu/controllers"
	"github.com/astaxie/beego"
	"lbccsu/controllers/index"
)

func init() {
    	beego.Include(&controllers.LoginController{})
	beego.Include(&index.IndexController{})
}
