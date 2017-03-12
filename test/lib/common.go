package Lib

import (
    "fmt"
    "net/http"
     "io/ioutil"
     "log"
     "bytes"
)

var URL = "http://127.0.0.1:8080/"

func Post(url string, b []byte) {
    body := bytes.NewBuffer(b)
    res, err := http.Post(url, "application/json;charset=utf-8", body)
    if err != nil {
        log.Fatal(err)
        return
    }

    result, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if  err != nil {
        log.Fatal(err)
        return
    }
    fmt.Printf("%s", result)
}
