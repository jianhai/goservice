package  controllers

import (
    _ "fmt"
     "encoding/json"

    "github.com/astaxie/beego"
    "../models"
)

type TopicsController struct {
    beego.Controller
}

func (this *TopicsController) Post() {
    var request map[string] interface{}

    if err := json.Unmarshal(this.Ctx.Input.RequestBody, &request); err != nil {
        this.Data["json"] = models.Respond {
            ResultCode:	1,
	    Message:	"Request is NULL",
        }
        this.ServeJSON()
    }

    this.Data["json"] = models.Respond {
        ResultCode:	0,
	Message:	"Success",
        Data:		[] string{ "Crimson" , "Red" , "Ruby" , "Maroon" },
    }
    this.ServeJSON()
}
