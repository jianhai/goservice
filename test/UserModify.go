package main

import (
    "fmt"
    "flag"
    "encoding/json"
    m "../src/models"
    "./lib"
)

type UserModify struct {
    Client m.Client `json:"client"`
    Sex int `json:"sex"`
    NickName string `json:"nickName"`
    WantSex int `json:"wantSex"`
    HeadURL string `json:"headURL"`
    Birthday float64 `json:"birthday"`
    Topic int `json:"topic"`
};

var (
    id = flag.Int64("id", 10000, "User ID")
)

func main() {
    flag.Parse()

    object := UserModify{}
    object.Client.UserId = *id
    object.Client.Longitude = 116.300680217574
    object.Client.Latitude = 39.98267552027887
    object.NickName="栾建海"
    object.HeadURL="https://ss2.baidu.com/6ONYsjip0QIZ8tyhnq/it/u=310311839,4078970003&fm=58"
    object.Birthday=2.90082566e+08
    object.Topic=1
    str, err := json.Marshal(object)
    if err == nil {
        fmt.Println(string(str))
    }

    url := Lib.URL + "/api/user/modify"
    Lib.Post(url, str)
}
