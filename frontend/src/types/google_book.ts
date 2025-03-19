import { FAVORITE, HAVE_READ, READING_NOW, TO_READ } from "../constants/google_book";

export type GoogleBook = {
  id: string;
  saleInfo: {
    /** 購入リンク */
    buyLink: string;
    /** 販売状態 */
    saleability: "FOR_SALE" | "FREE" | "NOT_FOR_SALE" | "FOR_PREORDER";
  };
  volumeInfo: {
    title: string;
    authors?: string[];
    description?: string;
    categories?: string[];
    imageLinks?: {
      thumbnail: string;
    };
    infoLink: string;
  };
};

export type BookshelfId = typeof FAVORITE | typeof TO_READ | typeof READING_NOW | typeof HAVE_READ;
