package controllers

import (
	"github.com/astaxie/beego"
    "calibre-web/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
    this.Data["Website"] = "beego.me"
    this.Data["Email"]   = "astaxie@gmail.com"
    this.TplNames        = "index.tpl"
}
