package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/phbai/api/logic"
    "github.com/phbai/api/util"
    "github.com/bitly/go-simplejson"
)

func MoviesController(c *gin.Context) {
    result := make(chan *simplejson.Json)
    var params util.Movies

    params.Type = c.Query("type");
    params.PageIndex = c.Query("index");
    params.PageSize = c.Query("size");
    params.ID = c.Query("id");
    params.Data = c.Query("keyword");

    go logic.GetMovies(params, result)
    // res := <- result

    c.JSON(200, gin.H{
        "result": <- result,
    })
}