package controllers

import (
	"net/http"

	"github.com/crownss/dark_market/config"
	"github.com/crownss/dark_market/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserBuy(c *gin.Context) {
	var UserInDB models.Users
	var input models.Tx
	session := sessions.Default(c)
	user := session.Get(userkey)
	check_user := config.DB.Where("username = ?", user).Where("is_active = ?", true).Where("is_admin = ?", false).Find(&UserInDB).RowsAffected
	if check_user == 0 {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: "You are not logged in",
			Status:  "error",
		})
		return
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Status:  "error",
		})
		return
	}

	result := models.Tx{
		UserID: input.UserID,
		TxBTC: input.TxBTC,
		Is_Valid: false,
		WhichStuffID: input.WhichStuffID,
	}
	if e := config.DB.Create(&result).Error; e != nil{
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    http.StatusInternalServerError,
			Message: e.Error(),
			Status:  "error",
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Code: http.StatusOK,
		Message: "Waiting admin for accept your tx",
		Status: "success",
	})
}

func Tx_Is_Valid(c *gin.Context) {
	// var UserInDB models.Users
	// session := sessions.Default(c)
	// user := session.Get(userkey)
}
