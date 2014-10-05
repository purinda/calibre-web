package main

import (
	_ "calibre-web/routers"
    "calibre-web/models"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
)

func init() {
    orm.RegisterDriver("sqlite", orm.DR_Sqlite)
    orm.RegisterDataBase("default", "sqlite3", "db/metadata.db")
    orm.RegisterModel(new(models.Book))
}

func main() {
    orm.Debug = true
	beego.Run()
}
