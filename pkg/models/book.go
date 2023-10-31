package models

import (
	"github.com/woudanator/go-book-app/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int) (*Book, *gorm.DB) {
	var getBook Book
	dbreturn := db.Table("books").Where("id=?", id).First(&getBook)
	return &getBook, dbreturn
}

func DeleteBook(id int) *gorm.DB {
	db := db.Where("id=?", id).Delete(&Book{})
	return db
}

func UpdateBook(id int, updateBook *Book) (*Book, error) {
	var book Book
	db := config.GetDb()
	db.Table("books").Where("id = ?", id).First(&book)
	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}
	db.Save(&book)
	return &book, nil
}
