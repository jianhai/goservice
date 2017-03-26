package main

import (
    "fmt"
    "encoding/json"
    m "../src/models"
    "./lib"
)

type LiveGivings struct {
    Client m.Client `json:"client"`
};

func main() {
    object := LiveGivings{}
    str, err := json.Marshal(object)
    if err == nil {
        fmt.Println(string(str))
    }

    url := Lib.URL + "/api/live/givings"
    Lib.Post(url, str)
}
