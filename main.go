package main

import(
	"github.com/Dev79844/goCRUD/models"
	"github.com/gin-gonic/gin"
	"github.com/Dev79844/goCRUD/controllers"
)

func main(){
	router := gin.Default()

	models.ConnectDB()

	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.GetPosts)
	router.GET("/post/:id", controllers.GetPostsByID)
	router.PUT("/post/:id", controllers.UpdatePost)
	router.DELETE("/post/:id", controllers.DeletePost)

	router.Run("localhost:5000")
}
