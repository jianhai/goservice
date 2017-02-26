package  controllers

import (
    "fmt"
     "encoding/json"

    "github.com/astaxie/beego"
    "../models"
)

type TopicsController struct {
    beego.Controller
}

func (this *TopicsController) Post() {
    var request map[string] interface{}

    fmt.Println(this.Ctx.Input.RequestBody)

    if err := json.Unmarshal(this.Ctx.Input.RequestBody, &request); err != nil {
        fmt.Println("Bad")
    }

    resp := models.Resp {
        ResultCode:	0,
	Message:	"Success",
        Data:		[] string{ "Crimson" , "Red" , "Ruby" , "Maroon" },
    }

    json,_ :=json.Marshal(&resp)
    this.Data["json"]=json
    this.ServeJSON()
}
