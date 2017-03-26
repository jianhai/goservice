package UserController

import (
    "encoding/json"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    m "../../models"
)

type ModifyController struct {
    beego.Controller
}

type requestModify struct {
    Client m.Client `json:"client"`
    Sex int `json:"sex"`
    NickName string `json:"nickName"`
    WantSex int `json:"wantSex"`
    HeadURL string `json:"headURL"`
    Birthday float64 `json:"birthday"`
    IsOnline int64  `json:"isOnline"`
    Topic int64 `json:"topicId"`
};

func (this *ModifyController) Post() {
    var request requestModify

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

    if request.Sex != 0 {
        user.Sex = request.Sex
    }

    if len(request.NickName) > 0 {
        user.NickName = request.NickName
    }

    if request.WantSex != 0 {
        user.WantSex = request.WantSex
    }

    if len(request.HeadURL) > 0 {
        user.HeadURL = request.HeadURL
    }

    if request.Birthday != 0 {
        user.Birthday = request.Birthday
    }

    if request.Birthday != 0 {
        user.Birthday = request.Birthday
    }

    if request.Client.Longitude != 0 {
        user.Longitude = request.Client.Longitude
    }

    if request.Client.Latitude != 0 {
        user.Latitude = request.Client.Latitude
    }

    if request.Topic != 0 {
        user.Topic = request.Topic
    }

    if request.IsOnline !=0 {
        user.IsOnline = request.IsOnline
    }

    if num, err := o.Update(&user); err != nil {
        fmt.Println("%d:%s", num, err)
    }

    this.Data["json"] = m.Respond {
        ResultCode:	0,
	Message:	"Success",
        Data:		m.GetRespUser(user),
    }
    this.ServeJSON()
}
