package book

import "go-exercise/db/tables"

type Service interface {
	FindAll() ([]tables.Book, error)
	FindById(ID int) (tables.Book, error)
	Create(request BookRequest) (tables.Book, error)
}

type service struct {
	repository Repository
}

func BookService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]tables.Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindById(ID int) (tables.Book, error) {
	return s.repository.FindById(ID)
}

func (s *service) Create(request BookRequest) (tables.Book, error) {
	price, _ := request.Price.Int64()
	newBook := tables.Book{
		Title:    request.Title,
		Price:    int(price),
		Discount: request.Discount,
	}
	return s.repository.Create(newBook)
}
