package modules

import "github.com/jinzhu/gorm"

type BaseModel struct{
	gorm.Model
	db *gorm.DB
	dbName string
}

