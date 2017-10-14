package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/phbai/api/logic"
    "github.com/phbai/api/util"
    "github.com/bitly/go-simplejson"
)

func LoginController(c *gin.Context) {
    result := make(chan *simplejson.Json)
    var params util.Login

    params.Token = c.Query("token");

    go logic.Login(params, result)
    // res := <- result

    c.JSON(200, gin.H{
        "result": <- result,
    })
}