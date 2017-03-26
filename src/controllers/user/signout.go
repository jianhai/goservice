package UserController

import (
    "fmt"
    "encoding/json"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    m "../../models"
)

type SignoutController struct {
    beego.Controller
}

type requestSignout struct {
    Client m.Client `json:"client"`
};

func (this *SignoutController) Post() {
    var request requestSignout

    if err := json.Unmarshal(this.Ctx.Input.RequestBody, &request); err != nil {
        this.Data["json"] = m.Respond {
            ResultCode:	1,
	    Message:	"Request is NULL",
        }
        this.ServeJSON()
    }

    o := orm.NewOrm()
    user := m.User{Id: request.Client.UserId}
    o.Read(&user)
    user.IsOnline = 1

    if num, err := o.Update(&user); err != nil {
        fmt.Println("%d:%s", num, err)
    }

    this.Data["json"] = m.Respond {
        ResultCode:	0,
	Message:	"Success",
    }
    this.ServeJSON()
}
