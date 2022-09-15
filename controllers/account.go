package controllers

import (
	"fmt"
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
	var UserInDB models.Users
	session := sessions.Default(c)
	user := session.Get(userkey)
	check_admin := config.DB.Where("username = ?", user).Where("is_admin = ?", true).Find(&UserInDB).RowsAffected
	check_user := config.DB.Where("username = ?", user).Where("is_active = ?", true).Where("is_admin = ?", false).Find(&UserInDB).RowsAffected
	if check_admin != 0 {
		var allAccount []models.Users
		getall := config.DB.Find(&allAccount).RowsAffected
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
				Data:    allAccount,
			})
			return
		}
	}
	if check_user != 0 {
		c.JSON(http.StatusForbidden, models.Response{
			Code:    http.StatusForbidden,
			Message: "You not have access",
			Status:  "error",
		})
		return
	} else if check_user == 0 {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: "You are not logged in",
			Status:  "error",
		})
		return
	}
}

func GetAccountUsername(c *gin.Context) {
	var getaccountusername models.Users
	// var getusernameonly models.UserGetUsername

	if err := config.DB.Where("username = ?", c.Param("username")).First(&getaccountusername).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Status:  "error",
		})
		return
	}
	get1 := make(map[string]interface{})
	get1["username"] = getaccountusername.Username
	get1["is_active"] = getaccountusername.Is_Active
	get1["is_admin"] = getaccountusername.Is_Admin
	fmt.Println(get1)
	c.JSON(http.StatusOK, models.UsersUsernameResponseAny{
		Code:    http.StatusOK,
		Message: "Data Found !",
		Status:  "success",
		Data:    get1,
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
			Message: e.Error(),
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
	check_admin := config.DB.Where("username = ?", user).Where("is_admin = ?", true).Find(&UserInDB).RowsAffected
	check_user := config.DB.Where("username = ?", user).Where("is_active = ?", true).Where("is_admin = ?", false).Find(&UserInDB).RowsAffected
	if check_admin != 0 {
		c.JSON(http.StatusOK, gin.H{"username": user, "is_admin": true, "is_superuser": true, "is_active": true})
		return
	}
	if check_user != 0 {
		c.JSON(http.StatusOK, gin.H{"username": user, "is_admin": false, "is_superuser": false, "is_active": true})
		return
	} else if check_user == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "user not found"})
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

func UpdatePassword(c *gin.Context) {
	var UserInDB models.Users
	session := sessions.Default(c)
	user := session.Get(userkey)
	var inputPassword models.RequestUsersChangePassword
	if user == nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: "Bad session",
			Status:  "error",
		})
		return
	}
	query := config.DB.Where("username = ?", user).Find(&UserInDB)
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
	check := bcrypt.CompareHashAndPassword([]byte(UserInDB.Password), []byte(inputPassword.Old_password))
	if check != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: "Incorrect password",
			Status:  "error",
		})
		return
	}
	if err := c.Bind(&inputPassword); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Status:  "error",
		})
		return
	}
	update := make(map[string]interface{})
	update["old_password"] = inputPassword.Old_password
	update["new_password"] = inputPassword.New_password
	update["confirm_new_password"] = inputPassword.Confirm_new_password

	if err := config.DB.Where("username = ?", user).First(&UserInDB).Updates(update).Error; err != nil {
		c.JSON(http.StatusNotFound, &models.Response{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Status:  "error",
		})
		return
	}
	c.JSON(http.StatusAccepted, &models.Response{
		Code:    http.StatusAccepted,
		Message: "Succesfuly Change",
		Status:  "success",
	})
}

func UpdateAdmin(c *gin.Context) {
	var UserInDB models.Users
	session := sessions.Default(c)
	user := session.Get(userkey)
	check_superuser := config.DB.Where("username = ?", user).Where("is_admin = ?", true).Where("is_super_user = ?", true).Find(&UserInDB).RowsAffected
	check_user := config.DB.Where("username = ?", user).Where("is_active = ?", true).Where("is_admin = ?", false).Find(&UserInDB).RowsAffected
	if check_superuser != 0 {
		var getUsername models.Users
		if err := config.DB.Where("username = ?", c.Param("username")).First(&getUsername).Update("is_admin", true).Error; err != nil {
			c.JSON(http.StatusBadRequest, models.Response{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Status:  "error",
			})
			return
		}
		c.JSON(http.StatusOK, models.Response{
			Code:    http.StatusOK,
			Message: "Succesfully Update Admin",
			Status:  "success",
		})
		return
	}
	if check_user != 0 {
		c.JSON(http.StatusForbidden, models.Response{
			Code:    http.StatusForbidden,
			Message: "You not have access",
			Status:  "error",
		})
		return
	} else if check_user == 0 {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: "You are not logged in",
			Status:  "error",
		})
		return
	}
}

func InactiveAccount(c *gin.Context) {
	var UserInDB models.Users
	session := sessions.Default(c)
	user := session.Get(userkey)
	check_superuser := config.DB.Where("username = ?", user).Where("is_admin = ?", true).Where("is_super_user = ?", true).Find(&UserInDB).RowsAffected
	fmt.Println(check_superuser)
	check_user := config.DB.Where("username = ?", user).Where("is_active = ?", true).Where("is_admin = ?", false).Find(&UserInDB).RowsAffected
	if check_superuser != 0 {
		var getUsername models.Users
		if err := config.DB.Where("username = ?", c.Param("username")).First(&getUsername).Update("is_active", false).Error; err != nil {
			c.JSON(http.StatusBadRequest, models.Response{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Status:  "error",
			})
			return
		}
		c.JSON(http.StatusOK, models.Response{
			Code:    http.StatusOK,
			Message: "Succesfully Inactived Users",
			Status:  "success",
		})
		return
	}
	if check_user != 0 {
		c.JSON(http.StatusForbidden, models.Response{
			Code:    http.StatusForbidden,
			Message: "You not have access",
			Status:  "error",
		})
		return
	} else if check_user == 0 {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: "You are not logged in",
			Status:  "error",
		})
		return
	}
}

func ActivateAccount(c *gin.Context) {
	var UserInDB models.Users
	session := sessions.Default(c)
	user := session.Get(userkey)
	check_superuser := config.DB.Where("username = ?", user).Where("is_admin = ?", true).Where("is_super_user = ?", true).Find(&UserInDB).RowsAffected
	fmt.Println(check_superuser)
	check_user := config.DB.Where("username = ?", user).Where("is_active = ?", true).Where("is_admin = ?", false).Find(&UserInDB).RowsAffected
	if check_superuser != 0 {
		var getUsername models.Users
		if err := config.DB.Where("username = ?", c.Param("username")).First(&getUsername).Update("is_active", true).Error; err != nil {
			c.JSON(http.StatusBadRequest, models.Response{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Status:  "error",
			})
			return
		}
		c.JSON(http.StatusOK, models.Response{
			Code:    http.StatusOK,
			Message: "Succesfully Activate users",
			Status:  "success",
		})
		return
	}
	if check_user != 0 {
		c.JSON(http.StatusForbidden, models.Response{
			Code:    http.StatusForbidden,
			Message: "You not have access",
			Status:  "error",
		})
		return
	} else if check_user == 0 {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: "You are not logged in",
			Status:  "error",
		})
		return
	}
}

func DeleteAccount(c *gin.Context) {
	var UserInDB models.Users
	session := sessions.Default(c)
	user := session.Get(userkey)
	check_superuser := config.DB.Where("username = ?", user).Where("is_admin = ?", true).Where("is_super_user = ?", true).Find(&UserInDB).RowsAffected
	check_user := config.DB.Where("username = ?", user).Where("is_active = ?", true).Where("is_admin = ?", false).Find(&UserInDB).RowsAffected
	if check_superuser != 0 {
		var getUsername models.Users
		if err := config.DB.Where("username = ?", c.Param("username")).First(&getUsername).Error; err != nil {
			c.JSON(http.StatusNotFound, models.Response{
				Code:    http.StatusNotFound,
				Message: err.Error(),
				Status:  "error",
			})
			return
		}
		config.DB.Delete(&getUsername)
		c.JSON(http.StatusOK, models.Response{
			Code:    http.StatusOK,
			Message: "Succesfully Deleted",
			Status:  "success",
		})
		return
	}
	if check_user != 0 {
		c.JSON(http.StatusForbidden, models.Response{
			Code:    http.StatusForbidden,
			Message: "You not have access",
			Status:  "error",
		})
		return
	} else if check_user == 0 {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: "You are not logged in",
			Status:  "error",
		})
		return
	}
}
