package  controllers

import (
    "fmt"
     "encoding/json"

    "github.com/astaxie/beego"
)

type TopicsController struct {
    beego.Controller
}

func (this *TopicsController) Post() {
    var request map[string] interface{}

    fmt.Println(this.Ctx.Input.RequestBody)
    fmt.Println(this.Ctx.Input)

    json.Unmarshal(this.Ctx.Input.RequestBody, &request)
    //if err := json.Unmarshal(this.Ctx.Input.RequestBody, &request); err != nil {
    //    fmt.Println("Bad")
    //}

    this.Data["resultCode"] = 0
    this.Data["message"] = "Success"
    this.Data["data"] = ""
    this.ServeJSON()
}
