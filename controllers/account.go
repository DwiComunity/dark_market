package controllers

import (
	"net/http"

	"github.com/crownss/dark_market/config"
	"github.com/crownss/dark_market/models"
	"github.com/gin-gonic/gin"
)


func GetAllAccount(c *gin.Context)  {
	var allstuf []models.Users
	getall := config.DB.Find(&allstuf).RowsAffected
	if getall == 0{
		c.JSON(http.StatusNotFound, models.Response{
			Code: http.StatusNotFound,
			Message: "Data not found !",
			Status: "error",
		})
		return
	}else if(getall != 0){
		c.JSON(http.StatusOK, models.UsersResponseMany{
			Code: http.StatusOK,
			Message: "Data Found !",
			Status: "success",
			Data: allstuf,
		})
		return
	}
}

func GetAccountUsername(c *gin.Context){
	var getaccountusername models.Users
	if err := config.DB.Where("username = ?", c.Param("username")).First(&getaccountusername).Error; err != nil{
		c.JSON(http.StatusNotFound, models.Response{
			Code: http.StatusNotFound,
			Message: "Data not found !",
			Status: "error",
		})
		return
	}
	c.JSON(http.StatusOK, models.UsersResponseAny{
		Code: http.StatusOK,
		Message: "Data Found !",
		Status: "success",
		Data: getaccountusername,
	})
}

func RegisterAccount(c *gin.Context){}

func UpdatePassword(c *gin.Context){}

func DeleteAccount(c *gin.Context){}