package logic

import (
    "github.com/bitly/go-simplejson"
    "github.com/phbai/api/util"
)

func GetToken(ch chan <- *simplejson.Json) {
    url := "http://todayapi.com/vodapi.html"
    data := make(map[string]string)
    userID := util.RandomUserId()
    data["data"] = "{\"Action\": \"CreateToken\", \"Message\": {\"UID\": \"" + userID + "\"}}"

    util.SendNormalRequest(url, data, func(j *simplejson.Json) {
        j.Get("Message").Set("UID", userID)
        ch <- j
    })
}