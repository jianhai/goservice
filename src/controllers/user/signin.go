package UserController

import (
    "encoding/json"

    "github.com/astaxie/beego"
    m "../../models"
)

type SigninController struct {
    beego.Controller
}

type requestSignin struct {
    Client m.Client `json:"client"`
    Email string `json:"email"`
    Password string `json:"password"`
};

func (this *SigninController) Post() {
    var request requestSignin

    if err := json.Unmarshal(this.Ctx.Input.RequestBody, &request); err != nil {
        this.Data["json"] = m.Respond {
            ResultCode:	1,
	    Message:	"Request is NULL",
        }
        this.ServeJSON()
    }

    user := m.GetUserByEmailandPassword(request.Email, request.Password)
    if user.Id == 0 {
        this.Data["json"] = m.Respond {
            ResultCode:	1,
	    Message:	"EMail OR Password incorrect",
        }
        this.ServeJSON()
    }

    this.Data["json"] = m.Respond {
        ResultCode:	0,
	Message:	"Success",
        Data:		m.GetRespUser(user),
    }
    this.ServeJSON()
}
