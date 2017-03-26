package  LiveController

import (
     "encoding/json"

    "github.com/astaxie/beego"
    m "../../models"
)

type TopicsController struct {
    beego.Controller
}

func getTopic() []m.Topic {
    data := m.GetAllTopic()

    return data
}

func (this *TopicsController) Post() {
    var request map[string] interface{}

    if err := json.Unmarshal(this.Ctx.Input.RequestBody, &request); err != nil {
        this.Data["json"] = m.Respond {
            ResultCode:	1,
	    Message:	"Request is NULL",
        }
        this.ServeJSON()
    }

    data := getTopic()
    this.Data["json"] = m.Respond {
        ResultCode:	0,
	Message:	"Success",
        Data:		data,
    }
    this.ServeJSON()
}
