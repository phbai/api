package logic

import (
    "github.com/bitly/go-simplejson"
    "github.com/phbai/api/util"
)

func Play(params util.MovieInfo, ch chan <- *simplejson.Json) {
    url := "http://todayapi.com/vodapi.html"
    // data := make(map[string]string)
    data := make(map[string]string)
    userID := util.RandomUserId()
    data["data"] = "{\"Action\": \"CreateToken\", \"Message\": {\"UID\": \"" + userID + "\"}}"

    util.SendNormalRequest(url, data, func(j *simplejson.Json) {
        token, _ := j.Get("Message").Get("Token").String()

        data["data"] = "{\"Action\": \"PlayFreeMovie\", \"Message\": {\"UID\": \"" + userID + "\", \"Token\": \"" + token + "\", \"MovieID\": \"" + params.MovieID + "\"}}"
        util.SendNormalRequest(url, data, func(k *simplejson.Json) {
            ch <- k
        })
    })
    // data["MovieID"] = params.MovieID
    
}