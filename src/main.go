package main

import (
    "github.com/astaxie/beego"
    "./controllers"
)

func main() {
    // 
    beego.Router("/api/user/login", &controllers.LoginController{})
    beego.Router("/api/user/register", &controllers.RegisterController{})
    beego.Router("/api/user/info", &controllers.InfoController{})

    beego.Router("/api/live/topic", &controllers.TopicController{})

    beego.Run()
}
