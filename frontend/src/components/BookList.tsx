import { memo, useState } from "react";
import { GoogleBook } from "../types/google_book";
import { ExternalLink, ShoppingBasket } from "lucide-react";
import Memo from "./Memo";

type BookListProps = {
  books: GoogleBook[];
  userId: string;
};

const BookList = ({ books, userId }: BookListProps) => {
  // TODO:userIdはコンテキストを使った方がいい？
  const [openMemo, setOpenMemo] = useState<string | null>(null);

  return (
    <ul className="mx-auto my-2 w-[90%]">
      {books ? (
        books.map((book) => (
          <li className="rounded-2xl p-3 shadow-xl" key={book.id}>
            {book.volumeInfo.imageLinks?.thumbnail && (
              <img
                src={book.volumeInfo.imageLinks.thumbnail}
                alt={book.volumeInfo.title}
                loading="lazy"
                className="m-1 mx-auto"
              />
            )}

            {book.saleInfo.saleability === "FOR_SALE" && (
              <a
                href={book.saleInfo.buyLink}
                className="m-2 mx-auto flex w-fit items-center rounded-md bg-red-500 p-1 text-white hover:bg-red-600"
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
                className="mt-2 flex items-center justify-center text-xl font-bold hover:opacity-30"
              >
                {book.volumeInfo.title}
                <ExternalLink size={20} />
              </a>
            </h3>

            <p>{book.volumeInfo.authors?.join(", ") || "著者不明"}</p>
            <p className="m-1 mx-auto w-fit rounded-md bg-gray-500 p-1 text-white">
              {book.volumeInfo.categories?.join(", ") || "カテゴリ不明"}
            </p>
            <p className="text-left">
              {book.volumeInfo.description || "説明不明"}
            </p>

            {/* FIXME: Change icon. */}
            <button
              onClick={() => setOpenMemo(book.id)}
              aria-label="Watch memo"
              className="m-1"
            >
              Memo
            </button>
            {openMemo && (
              <Memo
                bookId={book.id}
                onClose={() => setOpenMemo(null)}
                userId={userId}
              />
            )}
          </li>
        ))
      ) : (
        <li className="my-3 font-serif text-3xl">
          書籍が見つかりませんでした。
        </li>
      )}
    </ul>
  );
};

export default memo(BookList);
