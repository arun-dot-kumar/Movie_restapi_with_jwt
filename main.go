package main

import (
	middlewares "movie/MiddleWare"
	"movie/controller"
	"movie/database"
	flag "movie/flags"
	//"movie/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	var connstr string
	if *flag.Instance {
		connstr = "root:arunsham@tcp(localhost:3306)/movie?parseTime=true"
	}
	if *flag.Container {
		connstr = "arun:arunsham@tcp(db:3306)/movie?parseTime=True"
	}
	database.Connect(connstr)
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(":4567")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/")
	{
		api.POST("/user/register", controllers.RegisterUser)
		api.POST("/user/token", controllers.GenerateToken)
		api.POST("/admin/register", controllers.RegisterAdminUser)
		api.POST("/admin/token", controllers.GenerateTokenForAdmin)
		secured := api.Group("/movie").Use(middlewares.Auth())
		{
			secured.POST("/admin/addShow", controllers.AddMovie)
			secured.POST("/getshows", controllers.GetShows)
			secured.POST("/user/bookmovie", controllers.BookMovie)
			secured.GET("/admin/getbookings", controllers.Getbookings)
		}
	}
	return router
}
