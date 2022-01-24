package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	FindLIKE(title string) ([]Book, error)
	Create(book Book) (Book, error)
	Update(book Book) (Book, error)
	Delete(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryBook(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindLIKE(title string) ([]Book, error) {
	books := []Book{}

	err := r.db.Where("title LIKE ?", "%"+title+"%").Find(&books).Error

	return books, err
}

func (r *repository) FindAll() ([]Book, error) {
	books := []Book{}

	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) FindByID(id int) (Book, error) {
	book := Book{}
	err := r.db.First(&book, id).Error

	return book, err
}

func (r *repository) Create(book Book) (Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) Update(book Book) (Book, error) {
	err := r.db.Save(&book).Error

	return book, err
}

func (r *repository) Delete(book Book) (Book, error) {
	err := r.db.Delete(book).Error

	return book, err
}
