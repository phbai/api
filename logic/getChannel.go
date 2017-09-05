package logic

import (
    "github.com/bitly/go-simplejson"
)

func GetChannel(ch chan <- *simplejson.Json) {
    url := "Movie/GetChannel"
    data := make(map[string]string)

    SendRequest(url, data, func(j *simplejson.Json) {
    	// finished <- true
        ch <- j
    })
}