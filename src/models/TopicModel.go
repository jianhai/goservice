package models

import (
    "github.com/astaxie/beego/orm"
)

type Topic struct {
    Id  int64 `orm:"column(Id)" json:"topicId"`
    Title string `orm:"colum(title)" json:"title"`
}

func (topic *Topic) TableName() string {
     return "t_Topic"
}

func init() {
    orm.RegisterModel(new(Topic))
}

func GetAllTopic() []Topic {
    var topics []Topic
    orm.NewOrm().Raw("select * from t_Topic").QueryRows(&topics)

    return topics
}

func GetTopicById(id int64) Topic {
   topic  := Topic {Id: id}
   o := orm.NewOrm()
   o.Read(&topic, "Id")

   return topic
}
