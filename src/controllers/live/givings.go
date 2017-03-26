package  LiveController

import (
     "encoding/json"

    "github.com/astaxie/beego"
    m "../../models"
)

type GivingController struct {
    beego.Controller
}

type requestGiving struct {
    Client m.Client `json:"client"`
}

func (this *GivingController) Post() {
    var request requestGiving

    if err := json.Unmarshal(this.Ctx.Input.RequestBody, &request); err != nil {
        this.Data["json"] = m.Respond {
            ResultCode:	1,
	    Message:	"Request is NULL",
        }
        this.ServeJSON()
    }

    data := m.GetAllGiving()

    this.Data["json"] = m.Respond {
        ResultCode:	0,
	Message:	"Success",
        Data:		data,
    }
    this.ServeJSON()
}
