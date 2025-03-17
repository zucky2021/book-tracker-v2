import React, { memo, useCallback, useEffect, useRef, useState } from "react";
import { GoogleBook } from "../types/google_book";
import { BOOKSHELF_IDS, MAX_RESULTS } from "../constants/google_book";
// import "../../scss/components/Bookshelf.scss"; FIXME: tailwind css
import BookList from "./BookList";
import { CircleX, Filter, LibraryBig, Save } from "lucide-react";
import Loading from "./Loading";
import InfiniteScroll from "react-infinite-scroll-component";

const Bookshelf = () => {
  const [books, setBooks] = useState<GoogleBook[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [userId, setUserId] = useState<string | null>(null);
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [shelfId, setShelfId] = useState(0);
  const [filteredBooks, setFilteredBooks] = useState<GoogleBook[]>([]);
  const [filterText, setFilterText] = useState("");
  const [filterCategory, setFilterCategory] = useState("");
  const [categories, setCategories] = useState<string[]>([]);
  const [startIndex, setStartIndex] = useState(0);
  const [totalItems, setTotalItems] = useState(0);
  const [hasMore, setHasMore] = useState(true);

  // ユーザーIDをローカルストレージから取得 無ければdialogを表示して登録を促す
  useEffect(() => {
    const storedUserId = localStorage.getItem("googleBooksUserId");
    if (storedUserId) {
      setUserId(storedUserId);
    } else {
      setIsDialogOpen(true);
    }
  }, []);

  useEffect(() => {
    const fetchBookshelfInfo = async () => {
      try {
        const controller = new AbortController();
        const timeoutId = setTimeout(() => controller.abort(), 5000);
        const response = await fetch(
          `https://www.googleapis.com/books/v1/users/${userId}/bookshelves/${shelfId}`,
          { signal: controller.signal }
        );
        clearTimeout(timeoutId);

        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        setTotalItems(data.volumeCount || 0);
      } catch (err) {
        setError("書棚情報の取得に失敗しました");
        console.error(err);
      }
    };

    if (userId && shelfId) {
      fetchBookshelfInfo();
    }
  }, [userId, shelfId]);

  // ユーザーIDが取得できたら、Google Books APIを利用して読み込み中を表示
  useEffect(() => {
    const fetchBooks = async () => {
      try {
        const response = await fetch(
          `https://www.googleapis.com/books/v1/users/${userId}/bookshelves/${shelfId}/volumes?startIndex=${startIndex}&maxResults=${MAX_RESULTS}`
        );
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        const fetchedBooks = data.items || [];

        setBooks((prevBooks) => {
          const newBooks = [...prevBooks, ...fetchedBooks];
          const uniqueBooks = Array.from(
            new Set(newBooks.map((book) => book.id))
          ).map((id) => newBooks.find((book) => book.id === id));
          return uniqueBooks;
        });

        setHasMore(books.length < totalItems);
      } catch (err) {
        setError("書籍データの取得に失敗しました");
        console.error(err);
      }
    };

    if (userId) {
      fetchBooks();
    }
  }, [userId, shelfId, startIndex, books.length, totalItems]);

  const handleSetUserId = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const newUserId = formData.get("userId") as string;

    if (newUserId) {
      localStorage.setItem("googleBooksUserId", newUserId);
      setUserId(newUserId);
      setIsDialogOpen(false);
    }
  };

  const removeUserId = () => {
    if (!userId) return;
    if (!confirm(`ユーザーIDを削除しますか？\nID: ${userId}`)) return;

    localStorage.removeItem("googleBooksUserId");
    setUserId(null);
    setIsDialogOpen(true);
  };

  const dialogRef = useRef<HTMLDialogElement>(null);
  const inputRef = useRef<HTMLInputElement>(null);
  useEffect(() => {
    if (isDialogOpen) {
      dialogRef.current?.showModal();
      inputRef.current?.focus();
      document.body.classList.add("dialog-open");
    } else {
      dialogRef.current?.close();
      document.body.classList.remove("dialog-open");
    }
  }, [isDialogOpen]);

  /**
   * フィルター処理(メモ化)
   */
  const handleFilterBooks = useCallback(() => {
    const searchText = filterText.toLowerCase();
    const filtered = books.filter((book) => {
      const title = book.volumeInfo.title.toLowerCase();
      const author = book.volumeInfo.authors?.join(" ") || "";
      const description = book.volumeInfo.description || "";
      const categories = book.volumeInfo.categories || [];

      const isMatchedText =
        title.includes(searchText) ||
        author.toLowerCase().includes(searchText) ||
        description.toLowerCase().includes(searchText);

      const isMatchedCategory = filterCategory
        ? categories.includes(filterCategory)
        : true;

      return isMatchedText && isMatchedCategory;
    });

    setFilteredBooks(filtered);
  }, [filterText, filterCategory, books]);

  useEffect(() => {
    handleFilterBooks();
  }, [filterText, filterCategory, books, handleFilterBooks]);

  // filteredBooksの内容が変わるたびにcategoriesを更新
  useEffect(() => {
    const uniqueCategories = Array.from(
      new Set(
        filteredBooks.reduce((acc: string[], book) => {
          const bookCategories = book.volumeInfo.categories || [];
          return acc.concat(bookCategories);
        }, [])
      )
    );
    setCategories(uniqueCategories);
  }, [filteredBooks]);

  const MemorizedBookList = memo(BookList);

  return (
    <section className="bookshelf" id="bookshelf">
      <h2>
        <LibraryBig size={24} />
        書棚
      </h2>

      {isDialogOpen || (
        <button
          onClick={removeUserId}
          className="bookshelf__deleteBtn"
          aria-label="Local StorageからGoogle BooksユーザーID削除"
        >
          ユーザーID 削除
          <CircleX size={18} />
        </button>
      )}

      <dialog ref={dialogRef} className="bookshelf__dialog" aria-modal="true">
        <h3 aria-labelledby="dialog-title">
          Google BooksのユーザーIDを入力してください。
        </h3>
        <p aria-labelledby="dialog-description">
          <a
            href="https://books.google.com/books"
            target="_blank"
            rel="noopener noreferrer"
            aria-label="Google Booksサイトを新しいタブで開く"
          >
            Google Booksサイト
          </a>
          のページURLから取得できます。
          <br />
          <code>
            https://books.google.com/books?uid=<span>123456789</span>
          </code>
        </p>
        <form onSubmit={handleSetUserId}>
          <input
            type="text"
            inputMode="numeric"
            pattern="\d+"
            name="userId"
            placeholder="123456789"
            required
            aria-required="true"
            ref={inputRef}
          />
          <button type="submit" aria-label="Google BooksユーザーID設定">
            <Save size={16} />
          </button>
        </form>
      </dialog>

      <form className="bookshelf__choice">
        <fieldset>
          <legend>
            書棚選択
            <LibraryBig size={16} />
          </legend>
          <select
            value={shelfId}
            onChange={(e) => {
              setShelfId(Number(e.target.value));
              setBooks([]);
              setStartIndex(0);
            }}
            aria-label="書棚選択"
          >
            {Object.entries(BOOKSHELF_IDS).map(([id, name]) => (
              <option key={id} value={id}>
                {name}
              </option>
            ))}
          </select>
        </fieldset>
      </form>

      <form className="bookshelf__filter">
        <fieldset>
          <legend>
            書棚内検索
            <Filter size={16} />
          </legend>

          <label htmlFor="filter-text">Title / Author / Description</label>
          <input
            type="text"
            placeholder="キーワードを入力"
            id="filter-text"
            value={filterText}
            onChange={(e) => setFilterText(e.target.value)}
          />

          <label htmlFor="filter-category">Category</label>
          <select
            id="filter-category"
            value={filterCategory}
            onChange={(e) => setFilterCategory(e.target.value)}
          >
            <option value="">選択してください</option>
            {categories.map((category) => (
              <option key={category} value={category}>
                {category}
              </option>
            ))}
          </select>
        </fieldset>
      </form>

      <InfiniteScroll
        dataLength={books.length}
        next={() => setStartIndex((prevIndex) => prevIndex + MAX_RESULTS)}
        hasMore={hasMore}
        loader={<Loading />}
        endMessage={<p className="bookshelf__fetch-all-msg">That's all</p>}
      >
        <MemorizedBookList books={filteredBooks} />
      </InfiniteScroll>

      {error && (
        <div className="bookshelf__error" role="alert" aria-live="assertive">
          {error}
        </div>
      )}
    </section>
  );
};

export default Bookshelf;
