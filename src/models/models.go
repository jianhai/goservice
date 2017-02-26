package models

/*
 */
type Client struct {
    AppId  string  `json:"appId"`
    AppVersion  string  `json:"appVersion"`
    UserId  int64  `json:"userId"`
    Os  string  `json:"os"`
    DeviceId  string  `json:"deviceId"`
    Longitude  int64  `json:"longitude"`
    Latitude   int64  `json:"latitude"`
}

/*
 */
type Topic struct {
    TopicId  int64 `json:"topicId"`
    Title string `json:"title"`
}

/*
 */
type User struct {
    UserId  int64  `json:"userId"`
    Sex  int64  `json:"sex"`
    NickName  string  `json:"nickName"`
    HeadURL  string  `json:"headURL"`
    Birthday  string  `json:"birthday"`
    IsOnline  int64  `json:"isOnline"`
    Longitude  int64  `json:"longitude"`
    Latitude   int64  `json:"latitude"`
    LivingId   string  `json:"livingId"`
    LivingState  int64  `json:"livingState"`
    LivingDuration  int64  `json:"livingDuration"`
    Golden  int64  `json:"golden"`
    Gifts  int64  `json:"gifts"`
    WantSex  int  `json:"wantSex"`
    Topic   int  `json:"topic"`
}

/*
 */
type Giving struct {
    GiveId  int64  `json:"giveId"`
    Type  int64  `json:"type"`
    Sex  int64  `json:"sex"`
    Title  string  `json:"title"`
    ImageURL  string  `json:"imageURL"`
    GifURL  string  `json:"gifURL"`
    Price  int `json:"price"`
}
