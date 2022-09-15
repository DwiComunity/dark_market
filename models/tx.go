package models

import (
	"gorm.io/gorm"
)

type Tx struct {
	gorm.Model
	UserID       uint `json:"user_id" form:"user_id" xml:"user_id"`
	User         Users
	TxBTC        string `json:"txbtc" form:"txbtc" xml:"txbtc"`
	Is_Valid     bool   `json:"is_valid" form:"is_valid" xml:"is_valid" gorm:"default:false"`
	WhichStuffID uint   `json:"whichstuff_id" form:"whichstuff_id" xml:"whichstuff_id"`
	WhichStuff   Stuff
}
