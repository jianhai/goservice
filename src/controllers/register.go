package controllers

import (
    "encoding/json"

    "github.com/astaxie/beego"
    m "../models"
)

type RegisterController struct {
    beego.Controller
}

type requestRegister struct {
    Client m.Client `json:"client"`
    Email string `json:"email"`
    Password string `json:"password"`
};

func registerUser(email string, password string) (m.User, error){
    user := m.User { EMail: email, Password:password, }

    id, err := m.AddUser(&user)
    if err != nil {
        return m.User{}, err
    }

    user = m.GetUserById(id)
    return user, nil
}

func (this *RegisterController) Post() {
    var request requestRegister

    if err := json.Unmarshal(this.Ctx.Input.RequestBody, &request); err != nil {
        this.Data["json"] = m.Respond {
            ResultCode:	1,
	    Message:	"Request is NULL",
        }
        this.ServeJSON()
    }

    user := m.GetUserByEmail(request.Email)
    if user.Id != 0 {
        this.Data["json"] = m.Respond {
            ResultCode:	1,
	    Message:	"EMail have exist",
        }
        this.ServeJSON()
    }

    data, err := registerUser(request.Email, request.Password)
    if err != nil {
        this.Data["json"] = m.Respond {
            ResultCode:	1,
	    Message:	"Register Error",
        }
        this.ServeJSON()
    }

    this.Data["json"] = m.Respond {
        ResultCode:	0,
	Message:	"Success",
        Data:		data,
    }
    this.ServeJSON()
}
