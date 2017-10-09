package logic

import (
    "github.com/bitly/go-simplejson"
    "github.com/phbai/api/util"
)

func GetMovies(params util.Movies, ch chan <- *simplejson.Json) {
    url := "Movie/GetMovies"
    data := make(map[string]string)
    data["PageIndex"] = params.PageIndex
    data["PageSize"] = params.PageSize
    data["Type"] = params.Type
    data["ID"] = params.ID
    data["Data"] = params.Data

    util.SendRequest(url, data, func(j *simplejson.Json) {
        ch <- j
    })
}