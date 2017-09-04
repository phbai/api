package main

import (
    "fmt"
)

func Play() {
    url := "Movie/Play"
    data := make(map[string]string)
    data["MovieID"] = "5040679"

    SendRequest(url, data, func(j *simplejson.Json) {
        fmt.Println(j);
    })
}