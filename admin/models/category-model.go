package models

import (
	"blog/admin/database"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name, Slug string
}

func (category Category) Add() {
	database.DB.Create(&category)
}

func (category Category) Get(getid int) Category {
	database.DB.First(&category, getid)
	return category
}
func (category Category) GetByName(getName string) Category {
	database.DB.Where("name = ?", getName).First(&category)
	return category
}
func (category Category) GetAll() []Category {
	var categories []Category
	database.DB.Find(&categories)
	return categories
}
func (category Category) Update(column string, value string) Category {
	database.DB.Model(&category).Update(column, value)
	return category
}
func (category Category) Updates(newCategory Category) Category {
	database.DB.Model(&category).Updates(&newCategory)
	return category
}
func (category Category) Delete() {
	database.DB.Unscoped().Delete(&category)
}
