package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/phbai/api/logic"
    "github.com/phbai/api/util"
    "github.com/bitly/go-simplejson"
)

func MovieInfoController(c *gin.Context) {
    result := make(chan *simplejson.Json)
    params := util.MovieInfo{
        MovieID: c.Param("MovieID")
    }

    go logic.GetMovieInfo(params, result)
    // res := <- result

    c.JSON(200, gin.H{
        "result": <- result,
    })
}