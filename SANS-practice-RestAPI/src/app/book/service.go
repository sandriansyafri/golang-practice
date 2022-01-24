package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	FindLIKE(title string) ([]Book, error)
	Create(book BookRequest) (Book, error)
	Update(book BookRequest, id int) (Book, error)
	Delete(id int) (Book, error)
}

type service struct {
	repository Repository
}

func NewServiceBook(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindLIKE(title string) ([]Book, error) {
	books, err := s.repository.FindLIKE(title)

	return books, err
}

func (s *service) FindByID(id int) (Book, error) {
	book, err := s.repository.FindByID(id)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	book := Book{
		Title:  bookRequest.Title,
		Price:  bookRequest.Price,
		Rating: bookRequest.Rating,
	}
	book, err := s.repository.Create(book)
	return book, err
}
func (s *service) Update(bookRequest BookRequest, id int) (Book, error) {
	book, _ := s.FindByID(id)
	book.Title = bookRequest.Title
	book.Price = bookRequest.Price
	book.Rating = bookRequest.Rating

	updatedBook, err := s.repository.Update(book)

	return updatedBook, err

}

func (s *service) Delete(id int) (Book, error) {
	book, _ := s.repository.FindByID(id)
	deletedBook, err := s.repository.Delete(book)

	return deletedBook, err

}
