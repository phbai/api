package logic

import (
    "github.com/bitly/go-simplejson"
    "github.com/phbai/api/util"
)

func GetToken(ch chan <- *simplejson.Json) {
    url := "http://todayapi.com/vodapi.html"
    data := make(map[string]string)
    data["data"] = "{\"Action\": \"CreateToken\", \"Message\": {\"UID\": \"" + util.RandomUserId() + "\"}}"

    util.SendNormalRequest(url, data, func(j *simplejson.Json) {
        ch <- j
    })
}