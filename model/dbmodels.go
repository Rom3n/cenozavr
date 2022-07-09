package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model

	Id      int
	Name    string
	Url     string
	Url_img string
	Price   int
}
