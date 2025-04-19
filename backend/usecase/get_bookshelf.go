package usecase

import "backend/domain"

type GetBookshelf struct {
	bookshelfRepo domain.BookshelfRepository
}

func NewGetBookshelf(bookshelfRepo domain.BookshelfRepository) *GetBookshelf {
	return &GetBookshelf{
		bookshelfRepo: bookshelfRepo,
	}
}

func (gb *GetBookshelf) Execute(userId string, shelfId int) (*domain.Bookshelf, error) {
	bookshelf, err := gb.bookshelfRepo.FindByID(userId, shelfId)
	if err != nil {
		return nil, err
	}
	return bookshelf, nil
}
