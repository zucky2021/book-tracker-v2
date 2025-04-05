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

func (gb *GetBookshelf) GetBookshelf(userId string, shelfId int) (*domain.Bookshelf, error) {
	bookshelf, err := gb.bookshelfRepo.FindById(userId, shelfId)
	if err != nil {
		return nil, err
	}
	return bookshelf, nil
}
