package main

import (
    "fmt"
    "encoding/json"
    _ "io/ioutil"
    _ "net/http"
    _ "strings"
    m "../src/models"
)

func main() {
    //URL := "http://127.0.0.1:8080/api/user/login"

    //resp, err := http.Post(URL, "application/x-www-form-urlencoded", strings.NewReader("name=cjb"))
    //if err != nil {
    //   fmt.Println(err)
    //}
    //defer resp.Body.Close()

    //body, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println(string(body))

    client := m.Client{}
    if str, err := json.Marshal(client); err == nil {
        fmt.Println(string(str))
    }
}
