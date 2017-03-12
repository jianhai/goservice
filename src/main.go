package main

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "./controllers"
)

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:Wangyumei521@tcp(127.0.0.1:3306)/goserver?charset=utf8")
}

func main() {
    o := orm.NewOrm()
    o.Using("default")

    // 
    beego.Router("/api/user/login", &controllers.LoginController{})
    beego.Router("/api/user/register", &controllers.RegisterController{})
    beego.Router("/api/user/info", &controllers.InfoController{})

    beego.Router("/api/live/topics", &controllers.TopicsController{})

    //controllers.Test()

    beego.Run()
}
