package interfaces

import "github.com/gin-gonic/gin"

type Account interface {
	GetAccountUsername(c *gin.Context)
	GetAllAccount(c *gin.Context)
	RegisterAccount(c *gin.Context)
	LoginAccount(c *gin.Context)
	LogoutAccount(c *gin.Context)
	UpdatePassword(c *gin.Context)
	UpdateAdmin(c *gin.Context)
	InactiveAccount(c *gin.Context)
	ActivateAccount(c *gin.Context)
	DeleteAccount(c *gin.Context)
}

type Stuff interface {
	GetAllStuff(c *gin.Context)
	GetStuffCode(c *gin.Context)
	_RandomString(n int) string
	PostStuff(c *gin.Context)
	UpdateStuff(c *gin.Context)
	DeleteStuff(c *gin.Context)
}
