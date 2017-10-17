package logic

import (
    "github.com/bitly/go-simplejson"
    "github.com/phbai/api/util"
)

func GetChannel(params util.Channel, ch chan <- *simplejson.Json) {
    url := "Movie/GetChannel"
    data := make(map[string]string)

    util.SendRequest(url, data, func(j *simplejson.Json) {
        // finished <- true
        ch <- j
    })
}