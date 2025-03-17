import { GoogleBook } from "../types/google_book";
import { ExternalLink, ShoppingBasket } from "lucide-react";
// import "../../scss/components/BookList.scss"; FIXME: tailwind css

type BookListProps = {
  books: GoogleBook[];
};

const BookList = ({ books }: BookListProps) => {
  return (
    <ul className="book-list">
      {books ? (
        books.map(book => (
          <li className="book-list__item" key={book.id}>
            {book.volumeInfo.imageLinks?.thumbnail && (
              <img
                src={book.volumeInfo.imageLinks.thumbnail}
                alt={book.volumeInfo.title}
                loading="lazy"
              />
            )}

            {book.saleInfo.saleability === "FOR_SALE" && (
              <a
                href={book.saleInfo.buyLink}
                className="book-list__sale"
                target="_blank"
                rel="noopener noreferrer"
                aria-label={`セール中 - ${book.volumeInfo.title}の購入ページを新しいタブで開く`}
              >
                SALE!!
                <ShoppingBasket size={16} />
              </a>
            )}

            <h3>
              <a
                href={book.volumeInfo.infoLink}
                target="_blank"
                rel="noopener noreferrer"
                className="book-list__item-title"
              >
                {book.volumeInfo.title}
                <ExternalLink size={16} />
              </a>
            </h3>

            <p className="book-list__item-authors">
              {book.volumeInfo.authors?.join(", ") || "著者不明"}
            </p>
            <p className="book-list__item-category">
              {book.volumeInfo.categories?.join(", ") || "カテゴリ不明"}
            </p>
            <p className="book-list__item-description">
              {book.volumeInfo.description || "説明不明"}
            </p>
          </li>
        ))
      ) : (
        <li className="book-list__empty">書籍が見つかりませんでした。</li>
      )}
    </ul>
  );
};

export default BookList;