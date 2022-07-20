package controllers

import (
	"net/http"

	"github.com/crownss/dark_market/config"
	"github.com/crownss/dark_market/helpers"
	"github.com/crownss/dark_market/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func RegisterAccount(c *gin.Context){
	validate := helpers.InitValidator()
	var request models.RequestUsersRegister
	if eer := c.BindJSON(&request);eer != nil{
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: eer.Error(),
			Status:  "error",
		})
		return
	}
	if err := validate.Struct(&request);err != nil {
		errTranslated := helpers.TranslateError(err, validate)
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: errTranslated.Error(),
			Status:  "error",
		})
		return
	}
	var UserInDB models.Users
	UserInDB.Username = request.Username
	query := config.DB.Where("username = ?", request.Username).Find(&UserInDB)
	if query.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: "Username is already registered",
			Status:  "error",
		})
		return
	}
	encrypt, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	UserInDB.Password = string(encrypt)
	if e := config.DB.Save(&UserInDB).Error;e != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    http.StatusInternalServerError,
			Message: e.Error(),
			Status:  "error",
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "Succesfuly register with username: "+request.Username,
		Status:  "Success",
	})
}

func LoginAccount(c *gin.Context){}

func UpdatePassword(c *gin.Context){}

func DeleteAccount(c *gin.Context){}
