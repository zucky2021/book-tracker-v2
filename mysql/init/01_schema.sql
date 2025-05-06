CREATE TABLE
  IF NOT EXISTS memos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL COMMENT 'Google books user id',
    book_id VARCHAR(255) NOT NULL COMMENT 'Google books book id',
    text TEXT NOT NULL,
    img_file_name VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    UNIQUE KEY idx_user_book (user_id, book_id)
  ) DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT 'Memos table';