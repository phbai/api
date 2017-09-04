package main

import (
    "fmt"
)

func GetMovieInfo() {
    url := "Movie/GetMovieInfo"
    data := make(map[string]string)
    data["MovieID"] = "5040679"

    SendRequest(url, data, func(j *simplejson.Json) {
        fmt.Println(j);
    })
}