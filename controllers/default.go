package controllers

import (
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    models "calibre-web/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
    o := orm.NewOrm()
    o.Using("default")

    var books [] *models.Book
    bookInstance := models.Book{}
    num, err := o.QueryTable(bookInstance.TableName()).All(&books)

    if err != orm.ErrNoRows && num > 0 {
        this.Data["books"] = books
    }

    this.TplNames        = "index.tpl"
}
