package main 

import (
    "github.com/gin-gonic/gin"
    "github.com/phbai/api/controller"
)

func main() {
    r := gin.Default()
    
    r.GET("/getChannel"  ,  controller.ChannelController)
    r.GET("/getActor"    ,  controller.ActorController)
    r.GET("/getClass"    ,  controller.ClassController)
    r.GET("/getMovieInfo",  controller.MovieInfoController)
    r.GET("/getMovies"   ,  controller.MoviesController)
    r.GET("/play"        ,  controller.PlayController)
    
    r.Run() // listen and serve on 0.0.0.0:8080
}