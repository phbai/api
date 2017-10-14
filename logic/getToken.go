package logic

import (
    "github.com/bitly/go-simplejson"
    "github.com/phbai/api/util"
)

func GetToken(params util.Actor, ch chan <- *simplejson.Json) {
    url := "Movie/GetActor"
    data := make(map[string]string)
    data["PageIndex"] = params.PageIndex
    data["PageSize"] = params.PageSize

    util.SendRequest(url, data, func(j *simplejson.Json) {
        ch <- j
    })
}