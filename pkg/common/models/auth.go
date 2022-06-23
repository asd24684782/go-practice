package models

import "gorm.io/gorm"

type Auth struct{
	gorm.Model
	Name		string `json:"name"`
	Password 	string `json:"password"`
}