package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["lbccsu/controllers/index:IndexController"] = append(beego.GlobalControllerRouter["lbccsu/controllers/index:IndexController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/main`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

}
