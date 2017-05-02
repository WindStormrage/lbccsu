package main

import (
	_ "lbccsu/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	"os"
	"lbccsu/models"
	"lbmodel"
	"encoding/gob"
)
func init(){
	gob.Register(lbmodel.User{})
	gob.Register(lbmodel.AdminUser{})
	models.Connect()
	initArgs()
}

func main() {
	beego.Run()
}

func initArgs(){
	args:=os.Args

	for _,v:=range args{
		if v== "-syncdb"{
			models.Syncdb(false)
			os.Exit(0)
		}
		if v== "-syncdb-force"{
			models.Syncdb(true)
			os.Exit(0)
		}
	}
}