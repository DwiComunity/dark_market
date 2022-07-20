package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username     string 		`json:"username" form:"username" xml:"username" gorm:"unique"`
	Password     string 		`json:"password" form:"password" xml:"password"`
	Is_Active    bool   `json:"is_active" form:"is_active" xml:"is_active" gorm:"default:true"`
	Is_Admin     bool   `json:"is_admin" form:"is_admin" xml:"is_admin" gorm:"default:false"`
	Is_SuperUser bool   `json:"is_superuser" form:"is_superuser" xml:"is_superuser" gorm:"default:false"`
}

type RequestUsersLogin struct {
	Username string `json:"username" form:"username" xml:"username" gorm:"unique"`
	Password string `validate:"required,min=6" form:"password" json:"password"`
}

type RequestUsersRegister struct {
	Username         string `validate:"required,min=5,max=20" json:"username" form:"username" xml:"username"`
	Password         string `validate:"required,min=6,max=50" form:"password" json:"password"`
	Confirm_password string `validate:"eqfield=Password" form:"confirm_password" json:"confirm_password"`
}

type RequestUsersChangePassword struct {
	Old_password         string `validate:"required,min=6,max=50" form:"old_password" json:"old_password"`
	New_password         string `validate:"required,min=6,max=50,nefield=Old_password" form:"new_password" json:"new_password"`
	Confirm_new_password string `validate:"eqfield=New_password" form:"confirm_new_password" json:"confirm_new_password"`
}
