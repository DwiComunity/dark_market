package controllers

import (
	"net/http"

	"github.com/crownss/dark_market/config"
	"github.com/crownss/dark_market/models"
	"github.com/gin-gonic/gin"
)

func GetAllStuff(c *gin.Context){
	var allstuf []models.Stuff
	getall := config.DB.Find(&allstuf).RowsAffected
	if getall == 0{
		c.JSON(http.StatusBadRequest, models.Response{
			Code: http.StatusBadRequest,
			Message: "Data not found !",
			Status: "error",
		})
		return
	}else if(getall != 0){
		c.JSON(http.StatusOK, models.StuffResponseMany{
			Code: http.StatusOK,
			Message: "Data Found !",
			Status: "error",
			Data: allstuf,
		})
		return
	}
}

func GetStuffTitle(c *gin.Context){
	var getstufftitle models.Stuff
	if err := config.DB.Where("title = ?", c.Param("title")).First(&getstufftitle).Error; err != nil{
		c.JSON(http.StatusBadRequest, models.Response{
			Code: http.StatusBadRequest,
			Message: "Data not found !",
			Status: "error",
		})
		return
	}
	c.JSON(http.StatusOK, models.StuffResponseAny{
		Code: http.StatusOK,
		Message: "Data Found !",
		Status: "error",
		Data: getstufftitle,
	})
}

func PostStuff(c *gin.Context){}

func UpdateStuff(c *gin.Context){}

func DeleteStuff(c *gin.Context){}