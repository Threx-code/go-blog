package models

import (
	"github.com/Threx-code/go-blog/package/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Blogs struct {
	gorm.Model
	Title   string `gorm: "unique" json:title`
	Author  string `json:author`
	Content string `json:content`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Blogs{})
}

func Index() []Blogs {
	var AllBlogs []Blogs
	db.Find(&AllBlogs)
	return AllBlogs
}

func Read(id int64) (*Blogs, *gorm.DB) {
	var getBlog Blogs
	db := db.Where("ID=?", id).Find(&getBlog)
	return &getBlog, db
}

func (b *Blogs) Store() *Blogs {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func Destroy(id int64) Blogs {
	var blog Blogs
	db.Where("ID=?", id).Delete(blog)
	return blog
}
