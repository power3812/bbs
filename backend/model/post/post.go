package post

import (
	connect "backend/database"
	model "backend/model"
)

type Post struct {
	model.Model

	Body string `json:"body" validate:"required"`
}

func GetAll() ([]Post, error) {
	p := []Post{}

	if err := connect.Db.Find(&p).Error; err != nil {
		return p, err
	}

	return p, nil
}

func (p *Post) Create() (*Post, error) {
	if err := connect.Db.Create(&p).Error; err != nil {
		return p, err
	}

	return p, nil
}
