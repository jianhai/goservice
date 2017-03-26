package  LiveController

import (
    "encoding/json"

    "github.com/astaxie/beego"
    m "../../models"
)

type LivingsController struct {
    beego.Controller
}

type requestLivings struct {
    Client m.Client `json:"client"`
    PageNo float64 `json:pageNo`
    PageSize float64 `json:pageSize`
}

func (this *LivingsController) Post() {
    var request requestLivings

    if err := json.Unmarshal(this.Ctx.Input.RequestBody, &request); err != nil {
        this.Data["json"] = m.Respond {
            ResultCode:	1,
	    Message:	"Request is NULL",
        }
        this.ServeJSON()
    }

    data := m.GetAllUser()

    this.Data["json"] = m.Respond {
        ResultCode:	0,
	Message:	"Success",
        Data:		data,
    }
    this.ServeJSON()
}
