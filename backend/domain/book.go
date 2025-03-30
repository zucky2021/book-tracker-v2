package domain

type Book struct {
	ID     string
	Title  string
	Author string
}

type BookRepository interface {
	GetBooks(userId string, shelfId string, startIndex int, maxResults int) ([]Book, error)
}
