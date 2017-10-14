package logic

import (
    "github.com/bitly/go-simplejson"
    "github.com/phbai/api/util"
)

func Login(params util.Login, ch chan <- *simplejson.Json) {
    url := "User/Login"
    data := make(map[string]string)
    data["Token"] = params.Token

    util.SendRequest(url, data, func(j *simplejson.Json) {
        ch <- j
    })
}