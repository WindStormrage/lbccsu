package controllers

import (
	"github.com/astaxie/beego"
	"lbccsu/service"
	"strconv"
)

type LoginController struct {
	beego.Controller
}

//@router /login [*]
func (this *LoginController) Login() {
	this.TplName = "login.tpl"
}


//@router /register/submit [post]
func (this *LoginController) Register(){
	email:=this.GetString("email")
	code:=strconv.FormatInt(service.RandInt64(1000,9999),10)
	go service.SendEmail(email,code)
	this.ReturnSuccess()
}




func (this *LoginController) ReturnSuccess(args ...interface{})  {
	result:=make(map[string]interface{})
	result["status"]=10000
	result["message"]="success"
	key:=""
	for _,arg:=range args{
		switch arg.(type) {
		case string:
			key=arg.(string)
		default:
			result[key]=arg
		}
	}
	this.Data["json"]=result
	this.ServeJSON()
}