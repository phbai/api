package main 

import (
    "github.com/gin-gonic/gin"
    "github.com/phbai/api/controller"
)

func main() {
    r := gin.Default()
    
    r.GET("/getChannel", controller.ChannelController)
    
    r.Run() // listen and serve on 0.0.0.0:8080
}