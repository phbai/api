package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/phbai/api/logic"
    "github.com/phbai/api/util"
    "github.com/bitly/go-simplejson"
)

func ClassController(c *gin.Context) {
    result := make(chan *simplejson.Json)
    var params util.Class

    params.PageIndex = c.Query("index");
    params.PageSize = c.Query("size");
    
    go logic.GetClass(params, result)
    // res := <- result

    c.JSON(200, gin.H{
        "result": <- result,
    })
}