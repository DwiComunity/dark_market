package models

import "gorm.io/gorm"

type Stuff struct {
	gorm.Model
	Code  string `json:"Code" form:"Code" xml:"Code" gorm:"unique"`
	Img   string `json:"img" form:"img" xml:"img"`
	Title string `json:"title" form:"title" xml:"title" gorm:"sort:asc collate:utf8"`
	Desc  string `json:"desc" form:"desc" xml:"desc"`
	Price int    `json:"price" form:"price" xml:"price"`
}
