package book

type BookRepository interface {
	GetBooks(userId int, shelfId string, startIndex int, maxResult int) ([]Book, error)
}
