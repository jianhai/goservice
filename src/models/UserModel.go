package models

import (
    "log"
    "fmt"
    "errors"

    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/validation"
    _ "../lib"
)

/*
 */
type User struct {
    Id  int64  `orm:"column(userId)" form:"userID" json:"userId"`
    EMail string `orm:"column(email)" json:"_"`
    Password string `orm:"column(password)" json:"_"`
    NickName  string  `orm:"column(nickName)" form:"nickName" json:"nickName"`
    HeadURL  string  `orm:"column(headURL)" form:"headURL" json:"headURL"`
    Sex  int  `orm:"column(sex)" form:"sex" json:"sex"`
    Birthday  float64  `orm:"column(birthday)" json:"birthday"`
    IsOnline  int64  `orm:"column(isOnline)" json:"isOnline"`
    Longitude  float64  `orm:"column(longitude)" json:"longitude"`
    Latitude   float64  `orm:"column(latitude)" json:"latitude"`
    LivingId   string  `orm:"column(livingId)" json:"livingId"`
    LivingState  int64  `orm:"column(livingState)" json:"livingState"`
    LivingDuration  int64  `orm:"column(livingDuration)" json:"livingDuration"`
    Golden  int64  `orm:"column(golden)" json:"golden"`
    Gifts  int64  `orm:"column(gifts)" json:"gifts"`
    WantSex  int  `orm:"column(wantSex)" form:"wantSex" json:"wantSex"`
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
    user.Password = u.Password
    /* below commit is only placeholder */
    user.LivingId = "fOKVFiseW"
    id, err := o.Insert(user)

    return id, err
}

func DelUserById(Id int64) (int64, error) {
    o := orm.NewOrm()
    status, err := o.Delete(&User{Id: Id})
    return status, err
}

func UpdateUser(u *User) (int64, error) {
    if err:= checkUser(u); err != nil {
        return 0, err
    }

    o := orm.NewOrm()
    user := make(orm.Params)
    if len(u.NickName) > 0 {
        user["nickName"] = u.NickName
    }
    var table User
    num, err := o.QueryTable(table).Filter("Id", u.Id).Update(user)
    return num, err
}

func GetUserByEmail(email string) (user User) {
    user = User{EMail: email}
    o := orm.NewOrm()
    o.Read(&user, "email")
    return user
}

func GetUserByEmailandPassword(email string, password string) (user User) {
    user = User{EMail: email, Password: password}
    o := orm.NewOrm()
    err := o.Read(&user, "email", "password")
    if err == orm.ErrNoRows {
        fmt.Println("No result found.")
    } else if err == orm.ErrMissPK {
        fmt.Println("No primary key found.")
    } else {
        fmt.Println(user.Id, user.EMail, user.Password)
    }

    return user
}

func GetUserById(id int64) (user User) {
    user = User{Id: id}
    o := orm.NewOrm()
    o.Read(&user, "Id")
    return user
}

func GetAllUser() []User {
    var users []User
    orm.NewOrm().Raw("select * from t_User").QueryRows(&users)

    return users
}
