package models

import "gorm.io/gorm"

type Tx struct{
	gorm.Model
	User Users `json:"buyers" form:"buyers" xml:"buyers"`
	TxBTC string `json:"txbtc" form:"txbtc" xml:"txbtc"`
	Admin Users `gorm:"check:Is_Admin <> true" json:"admin" form:"admin" xml:"admin"`
	WhichStuff Stuff
}