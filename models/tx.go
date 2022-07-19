package models

import "gorm.io/gorm"

type Tx struct{
	gorm.Model
	User Users `json:"buyers" form:"buyers" xml:"buyers"`
	TxBTC string `json:"txbtc" form:"txbtc" xml:"txbtc"`
	Is_Valid bool `json:"is_valid" form:"is_valid" xml:"is_valid" gorm:"default:false"`
	WhichStuff Stuff
}