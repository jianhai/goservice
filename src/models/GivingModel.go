package models

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/astaxie/beego/validation"
    _ "../lib"
)

/*
 */
type Giving struct {
    Id  int64  `orm:"column(Id)" json:"giveId"`
    Type string `orm:"column(type)" json:"type,omitempty"`
    Sex int `orm:"column(sex)" json:"sex,omitempty"`
    Title  string  `orm:"column(title)" json:"title,omitempty"`
    ImageURL  string  `orm:"column(imageURL)" json:"imageURL,omitempty"`
    GifURL  string  `orm:"column(gifURL)" json:"giftURL,omitempty"`
    Price  float64  `orm:"column(price)" json:"price,omitempty"`
}

func (giving *Giving) TableName() string {
	return "t_Giving"
}

func init() {
    orm.RegisterModel(new(Giving))
}

func GetAllGiving() []Giving {
    var givings []Giving
    orm.NewOrm().Raw("select * from t_Giving").QueryRows(&givings)

    return givings
}
