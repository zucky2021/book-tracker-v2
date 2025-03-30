package domain

type Bookshelf struct {
	ID    string
	Title string
}

type BookshelfRepository interface {
	GetBookshelves(userId string) ([]Bookshelf, error)
}
