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

func Test() models.User {
   user := models.User {
   }

   models.AddUser(&user)
   return user
}

func registerUser(email interface{}, password interface{}) models.User {
   user := models.User {
   }

   models.AddUser(&user)
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
    data := registerUser(email, password)
    this.Data["json"] = models.Respond {
        ResultCode:	0,
	Message:	"Success",
        Data:		data,
    }
    this.ServeJSON()
}
