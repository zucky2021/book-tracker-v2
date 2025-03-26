import React, { memo, useCallback, useEffect, useRef, useState } from "react";
import { GoogleBook } from "../types/google_book";
import { BOOKSHELF_IDS, MAX_RESULTS } from "../constants/google_book";
import BookList from "./BookList";
import { CircleX, Filter, LibraryBig, Save } from "lucide-react";
import Loading from "./Loading";
import InfiniteScroll from "react-infinite-scroll-component";
import { USER_ID } from "../constants/storage_key";

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
    const storedUserId = localStorage.getItem(USER_ID);
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
          `http://localhost:8080/api/bookshelves?userId=${userId}&shelfId=${shelfId}`,
          { signal: controller.signal },
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
          `http://localhost:8080/api/books?userId=${userId}&shelfId=${shelfId}&startIndex=${startIndex}&maxResults=${MAX_RESULTS}`,
        );
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        const fetchedBooks = data.items || [];

        setBooks((prevBooks) => {
          const newBooks = [...prevBooks, ...fetchedBooks];
          const uniqueBooks = Array.from(
            new Set(newBooks.map((book) => book.id)),
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
    if (!/^\d+$/.test(newUserId)) {
      setError("User ID is invalid");
      return;
    }

    if (newUserId) {
      localStorage.setItem(USER_ID, newUserId);
      setUserId(newUserId);
      setIsDialogOpen(false);
      setError(null);
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
        }, []),
      ),
    );
    setCategories(uniqueCategories);
  }, [filteredBooks]);

  return (
    <section id="bookshelf">
      <h2 className="my-3 flex items-center justify-center font-serif text-3xl">
        <LibraryBig size={24} aria-hidden="true" />
        書棚
      </h2>

      {isDialogOpen || (
        <button
          onClick={removeUserId}
          className="mx-auto my-2 flex w-36 cursor-pointer items-center justify-around rounded-2xl bg-red-500 p-1 text-white hover:bg-red-600"
          aria-label="Local StorageからGoogle BooksユーザーID削除"
        >
          Delete user ID
          <CircleX size={18} aria-hidden="true" />
        </button>
      )}

      <dialog
        ref={dialogRef}
        className="left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 transform rounded-lg p-4 shadow-lg"
        aria-modal="true"
      >
        <h3 aria-labelledby="dialog-title" className="font-serif text-xl">
          Google BooksのユーザーIDを入力してください。
        </h3>
        <p aria-labelledby="dialog-description" className="my-2 text-lg">
          <a
            href="https://books.google.com/books"
            target="_blank"
            rel="noopener noreferrer"
            aria-label="Google Booksサイトを新しいタブで開く"
            className="text-blue-500"
          >
            Google Booksサイト
          </a>
          のページURLから取得できます。
          <br />
          <code className="mx-auto my-2 block w-fit rounded-md bg-gray-200 p-1 text-sm">
            https://books.google.com/books?uid=
            <span className="text-lg text-red-500">123456789</span>
          </code>
        </p>
        <form
          onSubmit={handleSetUserId}
          className="flex items-center justify-center"
        >
          <input
            type="text"
            inputMode="numeric"
            pattern="\d+"
            name="userId"
            placeholder="123456789"
            required
            aria-required="true"
            ref={inputRef}
            className="w-60 rounded-md border border-gray-300 p-2 text-center text-sm"
          />
          <button
            type="submit"
            aria-label="Google BooksユーザーID設定"
            className="ml-2 cursor-pointer rounded-xl bg-sky-300 p-2 shadow-xl"
          >
            <Save size={20} />
          </button>
        </form>
      </dialog>

      <form>
        <fieldset className="mx-auto my-2 w-[80%] min-w-72 rounded-lg border border-gray-300 p-3">
          <legend
            id="bookshelf-legend"
            className="flex items-center font-serif text-2xl"
          >
            書棚選択
            <LibraryBig size={24} aria-hidden="true" />
          </legend>
          <label htmlFor="bookshelf-select" className="sr-only">
            書棚を選択してください
          </label>
          <select
            id="bookshelf-select"
            value={shelfId}
            onChange={(e) => {
              setShelfId(Number(e.target.value));
              setBooks([]);
              setStartIndex(0);
            }}
            aria-label="書棚選択"
            aria-labelledby="bookshelf-legend"
            className="w-36 rounded-md border border-gray-300 p-2 text-center"
          >
            {Object.entries(BOOKSHELF_IDS).map(([id, name]) => (
              <option key={id} value={id}>
                {name}
              </option>
            ))}
          </select>
        </fieldset>
      </form>

      <form>
        <fieldset className="mx-auto my-2 w-[80%] min-w-72 rounded-lg border border-gray-300 p-3">
          <legend className="flex items-center font-serif text-2xl">
            書棚内検索
            <Filter size={24} />
          </legend>

          <label htmlFor="filter-text" className="block font-serif">
            Title / Author / Description
          </label>
          <input
            type="text"
            placeholder="キーワードを入力"
            id="filter-text"
            value={filterText}
            onChange={(e) => setFilterText(e.target.value)}
            className="mx-auto block w-40 rounded-md border border-gray-300 p-2 text-center text-sm"
          />

          <label htmlFor="filter-category" className="font-serif">
            Category
          </label>
          <select
            id="filter-category"
            value={filterCategory}
            onChange={(e) => setFilterCategory(e.target.value)}
            className="mx-auto block w-40 rounded-md border border-gray-300 p-2 text-center text-sm"
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
        endMessage={
          <p className="my-3 text-center font-serif text-xl">That's all</p>
        }
      >
        <BookList books={filteredBooks} />
      </InfiniteScroll>

      {error && (
        <div
          className="my-2 text-center text-red-500"
          role="alert"
          aria-live="assertive"
        >
          {error}
        </div>
      )}
    </section>
  );
};

export default memo(Bookshelf);
