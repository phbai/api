package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/phbai/api/logic"
    "github.com/bitly/go-simplejson"
    "github.com/phbai/api/util"
)

func PlayController(c *gin.Context) {
    result := make(chan *simplejson.Json)
    var params util.MovieInfo

    params.MovieID = c.Query("id");
    go logic.Play(params, result)
    // res := <- result

    c.JSON(200, gin.H{
        "result": <- result,
    })
}