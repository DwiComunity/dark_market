package controllers

import (
	"net/http"

	"github.com/crownss/dark_market/config"
	"github.com/crownss/dark_market/helpers"
	"github.com/crownss/dark_market/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const userkey = "username"

func GetAllAccount(c *gin.Context) {
	var allstuf []models.Users
	getall := config.DB.Find(&allstuf).RowsAffected
	if getall == 0 {
		c.JSON(http.StatusNotFound, models.Response{
			Code:    http.StatusNotFound,
			Message: "Data not found !",
			Status:  "error",
		})
		return
	} else if getall != 0 {
		c.JSON(http.StatusOK, models.UsersResponseMany{
			Code:    http.StatusOK,
			Message: "Data Found !",
			Status:  "success",
			Data:    allstuf,
		})
		return
	}
}

func GetAccountUsername(c *gin.Context) {
	var getaccountusername models.Users
	if err := config.DB.Where("username = ?", c.Param("username")).First(&getaccountusername).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Code:    http.StatusNotFound,
			Message: "Data not found !",
			Status:  "error",
		})
		return
	}
	c.JSON(http.StatusOK, models.UsersResponseAny{
		Code:    http.StatusOK,
		Message: "Data Found !",
		Status:  "success",
		Data:    getaccountusername,
	})
}

func RegisterAccount(c *gin.Context) {
	validate := helpers.InitValidator()
	var request models.RequestUsersRegister
	if eer := c.BindJSON(&request); eer != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: eer.Error(),
			Status:  "error",
		})
		return
	}
	if err := validate.Struct(&request); err != nil {
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
	encrypt, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	UserInDB.Password = string(encrypt)
	if e := config.DB.Save(&UserInDB).Error; e != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    http.StatusInternalServerError,
			Message: e.Error(),
			Status:  "error",
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "Succesfuly register with username: " + request.Username,
		Status:  "Success",
	})
}

func LoginAccount(c *gin.Context) {
	var request models.RequestUsersLogin
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Status:  "error",
		})
		return
	}
	var UserInDB models.Users
	query := config.DB.Where("username = ?", request.Username).Find(&UserInDB)
	if query.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    http.StatusInternalServerError,
			Message: query.Error.Error(),
			Status:  "error",
		})
		return
	}
	if query.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: "Username is not registered",
			Status:  "error",
		})
		return
	}
	check := bcrypt.CompareHashAndPassword([]byte(UserInDB.Password), []byte(request.Password))
	if check != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: "Incorrect password",
			Status:  "error",
		})
		return
	}
	if !UserInDB.Is_Active {
		c.JSON(http.StatusForbidden, models.Response{
			Code:    http.StatusForbidden,
			Message: "Your account has been deactivate",
			Status:  "error",
		})
		return
	}
	session := sessions.Default(c)
	session.Set(userkey, request.Username)
	if e := session.Save(); e != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed to save session",
			Status:  "error",
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "Succesfuly logged in",
		Status:  "success",
	})

}

func TesUser(c *gin.Context) {
	var UserInDB models.Users
	session := sessions.Default(c)
	user := session.Get(userkey)
	check_admin := config.DB.Where("username = ?", user).Where("is_admin = ?",true).Find(&UserInDB).RowsAffected
	check_user := config.DB.Where("username = ?", user).Where("is_active = ?",true).Where("is_admin = ?",false).Find(&UserInDB).RowsAffected
	if check_admin != 0{
		c.JSON(http.StatusOK, gin.H{"username": user, "is_admin": true, "is_superuser":true, "is_active":true})
		return
	}
	if check_user != 0{
		c.JSON(http.StatusOK, gin.H{"username": user, "is_admin": false, "is_superuser":false, "is_active":true})
		return
	}else if check_user == 0{
		c.JSON(http.StatusOK, gin.H{"message":"user not found"})
		return
	}
	
}

func LogoutAccount(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: "Bad session",
			Status:  "error",
		})
		return
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed to save session",
			Status:  "error",
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "Succesfuly logged out",
		Status:  "success",
	})
}

func UpdatePassword(c *gin.Context) {}

func DeleteAccount(c *gin.Context) {}
