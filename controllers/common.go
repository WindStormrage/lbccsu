package controllers

import (
	"github.com/astaxie/beego"
	"lbmodel"
)

type Common struct {
	beego.Controller
	User lbmodel.User
	IsLogin bool
}


type CtrPrepare interface {
	CtrPrepare()
}
















func (this *Common) ReturnJson(status int,message string,args ...interface{})  {
	result:=make(map[string]interface{})
	result["status"]=status
	result["message"]=message

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