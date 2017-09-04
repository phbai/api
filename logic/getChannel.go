package logic

import (
    "github.com/bitly/go-simplejson"
)

func GetChannel() *simplejson.Json {
    url := "Movie/GetChannel"
    data := make(map[string]string)
    result := make(chan *simplejson.Json)

    SendRequest(url, data, func(j *simplejson.Json) {
    	// finished <- true
        result <- j
    })
    
    return <- result
}