package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Tx struct{
	gorm.Model
	UserID uint
	User Users
	TxBTC string `json:"txbtc" form:"txbtc" xml:"txbtc"`
	Is_Valid sql.NullBool `json:"is_valid" form:"is_valid" xml:"is_valid" gorm:"default:false"`
	WhichStuffID uint
	WhichStuff Stuff
}