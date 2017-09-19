package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/phbai/api/logic"
    "github.com/phbai/api/util"
    "github.com/bitly/go-simplejson"
)

func ChannelController(c *gin.Context) {
    result := make(chan *simplejson.Json)
    params := util.Channel{}

    go logic.GetChannel(params, result)
    // res := <- result

    c.JSON(200, gin.H{
        "result": <- result,
    })
}