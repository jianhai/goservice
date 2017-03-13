package main

import (
    "fmt"
    "flag"
    "encoding/json"

    "./lib"
    m "../src/models"
)

type requestModel struct {
    Client m.Client `json:"client"`
    Email string `json:"email"`
    Password string `json:"password"`
};

var (
    email = flag.String("e", "luanjianhai@13.com", "Email Address")
    password = flag.String("p", "123456", "Default Password")
)

func main() {
    flag.Parse()

    request := requestModel{}
    request.Email = *email
    request.Password = *password

    str,err := json.Marshal(request)
    if err == nil {
        fmt.Println(string(str))
    }

    url := Lib.URL + "/api/user/register"
    Lib.Post(url, str)
}
