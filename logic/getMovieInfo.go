package logic

import (
    "github.com/bitly/go-simplejson"
    "github.com/phbai/api/util"
)

func GetMovieInfo(params util.MovieInfo, ch chan <- *simplejson.Json) {
    url := "Movie/GetMovieInfo"
    data := make(map[string]string)
    data["MovieID"] = "5040679"

    util.SendRequest(url, data, func(j *simplejson.Json) {
        ch <- j
    })
}