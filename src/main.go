package main

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "./controllers/live"
    "./controllers/user"
)

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:Wangyumei521@tcp(127.0.0.1:3306)/goserver?charset=utf8")
}

func main() {
    o := orm.NewOrm()
    o.Using("default")

    // 
    beego.Router("/api/user/register", &UserController.RegisterController{})
    beego.Router("/api/user/signin", &UserController.SigninController{})
    beego.Router("/api/user/signout", &UserController.SignoutController{})
    beego.Router("/api/user/modify", &UserController.ModifyController{})
    beego.Router("/api/user/info", &UserController.InfoController{})

    beego.Router("/api/live/topics", &LiveController.TopicsController{})
    beego.Router("/api/live/givings", &LiveController.GivingController{})
    beego.Router("/api/live/livings", &LiveController.LivingsController{})

    beego.Run()
}
