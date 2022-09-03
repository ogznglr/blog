package models

import "blog/admin/database"

type SwiperSlide struct {
	Postid uint `gorm:"primarykey"`
}

func (SwiperSlide) GetAll() []SwiperSlide {
	var ss []SwiperSlide
	database.DB.Find(&ss)
	return ss
}

func (ss SwiperSlide) Add() {
	database.DB.Create(&ss)
}

func (ss SwiperSlide) Remove() {
	database.DB.Delete(&ss)
}
