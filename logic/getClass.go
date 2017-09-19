package logic

import (
    "github.com/bitly/go-simplejson"
    "github.com/phbai/api/util"
)

func GetClass(ch chan <- *simplejson.Json) {
    url := "Movie/GetClass"
    data := make(map[string]string)
    data["PageIndex"] = "1"
    data["PageSize"] = "50"

    util.SendRequest(url, data, func(j *simplejson.Json) {
        ch <- j
    })
}