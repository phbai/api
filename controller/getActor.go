package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/phbai/api/logic"
    "github.com/bitly/go-simplejson"
)

func ActorController(c *gin.Context) {
    result := make(chan *simplejson.Json)

    go logic.GetActor(result)
    // res := <- result

    c.JSON(200, gin.H{
        "result": <- result,
    })
}