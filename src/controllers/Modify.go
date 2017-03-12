package controllers

import (
    "fmt"
    "encoding/json"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "../models"
)

type ModifyController struct {
    beego.Controller
}

func (this *ModifyController) Post() {
    var request map[string] interface{}

    if err := json.Unmarshal(this.Ctx.Input.RequestBody, &request); err != nil {
        this.Data["json"] = models.Respond {
            ResultCode:	1,
	    Message:	"Request is NULL",
        }
        this.ServeJSON()
    }

    fmt.Println(request)
    client := request["client"].(map[string] interface{})
    fmt.Println(client)
    c, _ := models.GetClient(client);
    fmt.Println(c.Latitude)
    o := orm.NewOrm()
    user := models.User{Id: 10000}
    if o.Read(&user) == nil {
        user.NickName = request["nickName"].(string)
        user.HeadURL = request["headURL"].(string)
        user.Sex = 1
        user.WantSex=2
        user.Birthday = request["birthday"].(float64)
        if num, err := o.Update(&user, "nickName", "headURL", "sex", "wantSex", "isOnline"); err == nil {
            fmt.Println(num)
        }
    }

    this.Data["json"] = models.Respond {
        ResultCode:	0,
	Message:	"Success",
        Data:		user,
    }
    this.ServeJSON()
}
