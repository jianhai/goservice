package main

import (
    "fmt"
    "encoding/json"
    m "../src/models"
    "./lib"
)

type LiveTopics struct {
    Client m.Client `json:"client"`
};

func main() {
    object := LiveTopics{}
    str, err := json.Marshal(object)
    if err == nil {
        fmt.Println(string(str))
    }

    url := Lib.URL + "/api/live/topics"
    Lib.Post(url, str)
}
