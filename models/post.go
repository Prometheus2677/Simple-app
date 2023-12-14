package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Title string `gorm:"type:varchar(255);"json:"title"`
	Body  string `gorm:"type:text" json:"body"`
}

func GetPost() Post {
	var post Post
	return post
}

func GetPosts() []Post {
	var posts []Post
	return posts
}
