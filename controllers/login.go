package controllers

import (
	"github.com/astaxie/beego"
	"lbccsu/service"
	"strconv"
	"lbmodel"
	"crypto/md5"
	"encoding/hex"
	"time"
)

type LoginController struct {
	beego.Controller
	code string
}

//@router /login [*]
func (this *LoginController) Login() {
	this.TplName = "login.tpl"
}


//@router /login/submit [post]
func (this *LoginController) LoginSubmit()  {
	this.ReturnSuccess()
}




//@router /register/sendEmail [post]
func (this *LoginController) SendEmail(){
	email:=this.GetString("email")
	code:=strconv.FormatInt(service.RandInt64(1000,9999),10)
	if err:=service.SendEmail(email,code);err!=nil{
		this.ReturnJson(10001,"send email error")
		return
	}
	this.code=code
	this.ReturnJson(10000,code)
}


//@router /register/submit [post]
func (this *LoginController) Register(){
	username:=this.GetString("username")
	password:=this.GetString("password")


	user := lbmodel.User{UserName:username}

	if err:=user.FindByUserName();err!=nil{
		beego.Debug(err)
		user.Password=StrMd5(password)
		user.UserName=username
		user.CreateTime=time.Now()
		user.LastIp=this.Ctx.Input.IP()
		if _,err:=user.Insert();err!=nil{
			beego.Error("insert error",err)
			this.ReturnJson(10002,"insert error")
			return
		}
	}else{
		this.ReturnJson(10001,"username is register")
		return
	}


	beego.Debug(username,password)
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

func (this *LoginController) ReturnJson(status int,message string,args ...interface{})  {
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

func StrMd5(strings string) string {
	data:=[]byte(strings)
	h:=md5.New()
	h.Write(data)
	s:=hex.EncodeToString(h.Sum(nil))
	return  s
}


