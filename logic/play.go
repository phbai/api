package logic

import (
    "github.com/bitly/go-simplejson"
)

func Play(ch chan <- *simplejson.Json) {
    url := "Movie/Play"
    data := make(map[string]string)
    data["MovieID"] = "5040679"

    SendRequest(url, data, func(j *simplejson.Json) {
        ch <- j
    })
}