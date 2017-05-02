package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
	"net/url"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
	"lbmodel"
)

var o orm.Ormer

func init(){
	orm.RegisterModel(
		new(lbmodel.User),
	)
}


func Syncdb(force bool)  {
	CreateDb()
	Connect()
	o=orm.NewOrm()

	name:="default"

	verbose:=true

	err:=orm.RunSyncdb(name,force,verbose)

	if err!=nil{
		fmt.Println(err)
	}
	if force{
		fmt.Println("database init is complete.\nPlease restart the application")
	}
}


func CreateDb() {
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	var dns string
	var sqlstring string
	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=%s", db_user, db_pass, db_host, db_port, db_name, url.QueryEscape("Asia/Shanghai"))
	db,err:=sql.Open("mysql",dns)

	if err!=nil{
		panic(err.Error())
	}
	r,err:=db.Exec(sqlstring)
	if err!=nil{
		log.Println(err)
		log.Println(r)
	}else{
		log.Println("Database",db_name,"created")
	}
	defer db.Close()

}



func Connect(){
	var dns string
	db_host:=beego.AppConfig.String("db_host")
	db_port:=beego.AppConfig.String("db_port")
	db_user:=beego.AppConfig.String("db_user")
	db_pass:=beego.AppConfig.String("db_pass")
	db_name:=beego.AppConfig.String("db_name")

	orm.DefaultTimeLoc,_=time.LoadLocation("Asia/Shanghai")
	orm.RegisterDriver("mysql",orm.DRMySQL)
	dns=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=%s",db_user,db_pass,db_host,db_port,db_name,url.QueryEscape("Asia/Shanghai"))
	orm.RegisterDataBase("default","mysql",dns)
}


