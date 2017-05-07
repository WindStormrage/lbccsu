package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["lbccsu/controllers:LoginController"] = append(beego.GlobalControllerRouter["lbccsu/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["lbccsu/controllers:LoginController"] = append(beego.GlobalControllerRouter["lbccsu/controllers:LoginController"],
		beego.ControllerComments{
			Method: "LoginSubmit",
			Router: `/login/submit`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["lbccsu/controllers:LoginController"] = append(beego.GlobalControllerRouter["lbccsu/controllers:LoginController"],
		beego.ControllerComments{
			Method: "SendEmail",
			Router: `/register/sendEmail`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["lbccsu/controllers:LoginController"] = append(beego.GlobalControllerRouter["lbccsu/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register/submit`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
