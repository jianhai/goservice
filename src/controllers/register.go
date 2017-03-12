package controllers

import (
    "fmt"
    "encoding/json"

    "github.com/astaxie/beego"
    "../models"
)

type RegisterController struct {
    beego.Controller
}

func registerUser(email string, password string) models.User {
    user := models.User { EMail: email, Password:password, }

    id, _ := models.AddUser(&user)
    if id == 0 {
        return user
    }

    user = models.GetUserById(id)
    return user
}

func (this *RegisterController) Post() {
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
    user := models.GetUserByEmail(fmt.Sprintf("%s", email))
    if user.Id != 0 {
        this.Data["json"] = models.Respond {
            ResultCode:	1,
	    Message:	"EMail have exist",
        }
        this.ServeJSON()
    }

    data := registerUser(fmt.Sprintf("%s", email), fmt.Sprintf("%s", password))
    this.Data["json"] = models.Respond {
        ResultCode:	0,
	Message:	"Success",
        Data:		data,
    }
    this.ServeJSON()
}
