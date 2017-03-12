package models

import (
    _ "fmt"
)

type Client struct {
    AppId  string  `json:"appId"`
    AppVersion  string  `json:"appVersion"`
    UserId  int64  `json:"userId"`
    Os  string  `json:"os"`
    DeviceId  string  `json:"deviceId"`
    Longitude  float64  `json:"longitude"`
    Latitude   float64 `json:"latitude"`
}

func GetClient(m map[string] interface{}) (Client, error) {
    client := Client{}

    if m["appId"] != nil {
        client.AppId = m["appId"].(string)
    }

    if m["appVersion"] != nil {
        client.AppVersion = m["appVersion"].(string)
    }

    if m["userId"] != nil {
        client.UserId = int64(m["userId"].(float64))
    }

    if m["os"] != nil {
        client.Os = m["os"].(string)
    }

    if m["deviceId"] != nil {
        client.DeviceId = m["deviceId"].(string)
    }

    if m["longitude"] != nil {
        client.Longitude = m["longitude"].(float64)
    }

    if m["latitude"] != nil {
        client.Latitude = m["latitude"].(float64)
    }

    return client, nil
}
