INSERT INTO
  memos (
    user_id,
    book_id,
    text,
    img_file_name,
    created_at,
    updated_at
  )
VALUES
  (
    'iniUser1',
    'iniBook1',
    'Initial test memo 1',
    '20250101120000123456789.jpg',
    NOW(),
    NOW()
  ),
  (
    'iniUser1',
    'iniBook2',
    'Initial test memo 2',
    '20250102130000123456789.png',
    NOW(),
    NOW()
  );