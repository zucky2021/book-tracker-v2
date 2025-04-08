package domain

type Bookshelf struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	VolumeCount int    `json:"volumeCount"`
}

type BookshelfRepository interface {
	FindByID(userId string, shelfId int) (*Bookshelf, error)
}
