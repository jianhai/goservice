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

func getTopic() [] models.Topic {
    var data = [] models.Topic {{1, "Item1"}, {2, "Item2"}}

    return data
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

    fmt.Println(request)
    data := getTopic()
    this.Data["json"] = models.Respond {
        ResultCode:	0,
	Message:	"Success",
        Data:		data,
    }
    this.ServeJSON()
}
