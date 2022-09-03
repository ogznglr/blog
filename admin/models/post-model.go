package models

import (
	"blog/admin/database"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title, Slug, Description, Content, PictureURL string
	CategoryId                                    int
}

func (post Post) Add() {
	database.DB.Create(&post)
}

func (post Post) Get(getid int) Post {
	database.DB.First(&post, getid)
	return post
}
func (post Post) GetFromSlug(getSlug string) Post {
	database.DB.Where("slug = ?", getSlug).First(&post)
	return post
}
func (post Post) GetFromCategory(getCategory int) []Post {
	var posts []Post
	database.DB.Where("category_id = ?", getCategory).Find(&posts)
	return posts
}

func (post Post) GetAll() []Post {
	var posts []Post
	database.DB.Find(&posts)
	return posts
}

func (post Post) Update(column string, value interface{}) Post {
	database.DB.Model(&post).Update(column, value)
	return post
}

func (post Post) Updates(newPost Post) Post {
	database.DB.Model(&post).Updates(&newPost)
	return post
}

func (post Post) Delete() {
	database.DB.Unscoped().Delete(&post)
}
