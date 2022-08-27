package book

import (
	"go-exercise/db/tables"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]tables.Book, error)
	FindById(ID int) (tables.Book, error)
	Create(book tables.Book) (tables.Book, error)
}

type repository struct {
	db *gorm.DB
}

func RepoToDb(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]tables.Book, error) {
	var book []tables.Book

	err := r.db.Find(&book).Error

	return book, err
}

func (r *repository) FindById(ID int) (tables.Book, error) {
	var book tables.Book

	err := r.db.Find(&book, ID).Error

	return book, err
}

func (r *repository) Create(book tables.Book) (tables.Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}
