package models

import (
    "log"
    "errors"

    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/validation"
    . "../lib"
)

/*
 */
type User struct {
    Id  int64  `orm:"column(userID)" json:"userId"`
    EMail string `orm:"column(email)" json:"email"`
    Password string `orm:"column(password)" json:"password"`
    NickName  string  `orm:"column(nickName)" json:"nickName"`
    HeadURL  string  `orm:"column(headURL)" json:"headURL"`
    Sex  int64  `orm:"column(sex)" json:"sex"`
    Birthday  string  `orm:"column(birthday)" json:"birthday"`
    IsOnline  int64  `orm:"column(isOnline)" json:"isOnline"`
    Longitude  int64  `orm:"column(longitude)" json:"longitude"`
    Latitude   int64  `orm:"column(latitude)" json:"latitude"`
    LivingId   string  `orm:"column(livingId)" json:"livingId"`
    LivingState  int64  `orm:"column(livingState)" json:"livingState"`
    LivingDuration  int64  `orm:"column(livingDuration)" json:"livingDuration"`
    Golden  int64  `orm:"column(golden)" json:"golden"`
    Gifts  int64  `orm:"column(gifts)" json:"gifts"`
    WantSex  int  `orm:"column(wantSex)" json:"wantSex"`
    Topic   int  `orm:"column(topicID)" json:"topic"`
}

func (user *User) TableName() string {
	return "t_User"
}

func init() {
    orm.RegisterModel(new(User))
}

func checkUser(u *User) (err error) {
     valid := validation.Validation{}
     b, _ := valid.Valid(&u)
     if !b {
         for _, err := range valid.Errors {
             log.Println(err.Key, err.Message)
             return errors.New(err.Message)
         }
     }
     return nil
}

func AddUser(u *User) (int64, error) {
    if err := checkUser(u); err != nil {
        return 0, err
    }

    o := orm.NewOrm()
    user := new(User)
    user.EMail  = u.EMail
    user.Password = Strtomd5(u.Password)
    id, err := o.Insert(user)
    return  id, err
}

func DelUserById(Id int64) (int64, error) {
    o := orm.NewOrm()
    status, err := o.Delete(&User{Id: Id})
    return status, err
}

func UpdateUser(u *User) (int64, error) {
    return 0, nil
}
