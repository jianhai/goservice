package  controllers

import (
    "fmt"

    "github.com/astaxie/beego"
)

type TopicController struct {
    beego.Controller
}

func (this *TopicController) Post() {
    fmt.Println("Good Luck")
}
