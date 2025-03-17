export const FAVORITE = 0;
export const TO_READ = 2;
export const READING_NOW = 3;
export const HAVE_READ = 4;

/** 書籍IDと表示名のマッピング */
export const BOOKSHELF_IDS: { [key: number]: string } = {
  [FAVORITE]: "お気に入り",
  [TO_READ]: "読みたい本",
  [READING_NOW]: "読書中",
  [HAVE_READ]: "既読",
};

/** Google Books APIの取得件数(最大値) */
export const MAX_RESULTS = 40;
