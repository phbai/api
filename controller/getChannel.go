package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/phbai/api/logic"
)

func ChannelController(c *gin.Context) {
    
    c.JSON(200, gin.H{
        "message": "pong",
    })
}