package logic

import (
    "github.com/bitly/go-simplejson"
)

func GetMovieInfo(ch chan <- *simplejson.Json) {
    url := "Movie/GetMovieInfo"
    data := make(map[string]string)
    data["MovieID"] = "5040679"

    SendRequest(url, data, func(j *simplejson.Json) {
        ch <- j
    })
}