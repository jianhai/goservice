package controllers

import (
    "fmt"
    "encoding/json"

    "github.com/astaxie/beego"
    "../models"
)

type LoginController struct {
    beego.Controller
}

func init() {
}

func (this *LoginController) Post() {
    var request map[string] interface{}

    if err := json.Unmarshal(this.Ctx.Input.RequestBody, &request); err != nil {
        this.Data["json"] = models.Respond {
            ResultCode:	1,
	    Message:	"Request is NULL",
        }
        this.ServeJSON()
    }

    var email = request["email"]
    if email == nil {
        this.Data["json"] = models.Respond {
            ResultCode:	1,
	    Message:	"Email is NULL",
        }
        this.ServeJSON()
    }

    var password = request["password"]
    if password == nil {
        this.Data["json"] = models.Respond {
            ResultCode:	1,
	    Message:	"Password is NULL",
        }
        this.ServeJSON()
    }

    fmt.Println(request["client"])
    user := models.GetUserByEmailandPassword(fmt.Sprintf("%s", email), fmt.Sprintf("%s", password))
    if user.Id == 0 {
        this.Data["json"] = models.Respond {
            ResultCode:	1,
	    Message:	"EMail OR Password incorrect",
        }
        this.ServeJSON()
    }

    this.Data["json"] = models.Respond {
        ResultCode:	0,
	Message:	"Success",
        Data:		user,
    }
    this.ServeJSON()
}
