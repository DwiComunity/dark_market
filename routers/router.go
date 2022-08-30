package routers

import (
	"os"

	"github.com/crownss/dark_market/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte(os.Getenv("SECRET_KEY")))
	// r.Use(sessions.Sessions("session", store))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24}) //expired in a day
	r.Use(cors.Default())
	stuff := r.Group("/marketplace")
	stuff.Use(sessions.Sessions("Token", store))
	{
		stuff.GET("/", controllers.GetAllStuff)
		stuff.GET("/:code", controllers.GetStuffCode)
		stuff.POST("/:code", controllers.UserBuy)
		stuff.POST("/posts/", controllers.PostStuff)
		stuff.PUT("/update/:code", controllers.UpdateStuff)
		stuff.DELETE("/delete/:code", controllers.DeleteStuff)
	}
	accounts := r.Group("/accounts")
	accounts.Use(sessions.Sessions("Token", store))
	{
		accounts.GET("/get/all", controllers.GetAllAccount)
		accounts.GET("/get/:username", controllers.GetAccountUsername)
		accounts.GET("/tes", controllers.TesUser)
		accounts.POST("/register", controllers.RegisterAccount)
		accounts.POST("/login", controllers.LoginAccount)
		accounts.POST("/logout", controllers.LogoutAccount)
		accounts.PUT("/change-password", controllers.UpdatePassword)
		accounts.PUT("/inactive/:username", controllers.InactiveAccount)
		accounts.DELETE("/delete/:username", controllers.DeleteAccount)
	}
	return r
}
