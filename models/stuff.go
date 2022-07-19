package models

import "gorm.io/gorm"

type Stuff struct {
	gorm.Model
	Code  string 	`json:"code" form:"code" xml:"code" gorm:"unique"`
	Img   string 	`json:"img" form:"img" xml:"img"`
	Title string 	`json:"title" form:"title" xml:"title"`
	Desc  string 	`json:"desc" form:"desc" xml:"desc"`
	Stock uint		`json:"stock" form:"stock" xml:"stock" gorm:"default:0"`
	Price float64	`json:"price" form:"price" xml:"price" gorm:"default:0"`
}
