package main

import (
    "fmt"
)

func GetActor() {
    url := "Movie/GetActor"
    data := make(map[string]string)
    data["PageIndex"] = "1"
	data["PageSize"] = "50"

    SendRequest(url, data, func(j *simplejson.Json) {
        fmt.Println(j);
    })
}