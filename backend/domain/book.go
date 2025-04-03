package domain

type Book struct {
	ID     string
	Title  string
	Author string
}

type BookRepository interface {
	// google books apiを使用して本を取得する
	FindAll(userId string, shelfId int, startIndex int, maxResult int) ([]Book, error)
}
