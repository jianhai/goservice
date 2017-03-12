package models

type Giving struct {
    GiveId  int64  `json:"giveId"`
    Type  int64  `json:"type"`
    Sex  int64  `json:"sex"`
    Title  string  `json:"title"`
    ImageURL  string  `json:"imageURL"`
    GifURL  string  `json:"gifURL"`
    Price  int `json:"price"`
}
