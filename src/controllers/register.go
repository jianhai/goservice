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

func registerUser() models.User {
   user := models.User {
   }

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
    fmt.Println(request)

    data := registerUser()
    this.Data["json"] = models.Respond {
        ResultCode:	0,
	Message:	"Success",
        Data:		data,
    }
    this.ServeJSON()
}
