package index

import "github.com/astaxie/beego"


type IndexController  struct {
	beego.Controller
}

//@router /main [*]
func (this *IndexController) Index(){
	this.TplName="main/main.tpl"
}
