package main

import (
    "fmt"
)

func GetChannel() {
    url := "Movie/GetChannel"
    data := make(map[string]string)

    sendRequest(url, data, func(j *simplejson.Json) {
        fmt.Println(j);
    })
}