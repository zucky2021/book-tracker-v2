# ベースイメージ
FROM node:24-alpine AS builder

# 作業ディレクトリ指定
WORKDIR /app

# package.jsonとpackage-lock.jsonをコピーして依存関係をインストール
COPY package.json package-lock.json ./
RUN npm install

# プロジェクト全体をコピー
COPY . .

# 非ルートユーザーの作成
RUN addgroup -S appuser && \
  adduser -S -G appuser appuser && \
  chown -R appuser:appuser /app

EXPOSE 3000

# 開発ステージ
FROM builder AS local
# 非ルートユーザーに切り替え(FROMの段階でルートユーザーになるため都度指定)
USER appuser
# 環境変数でホットリロードを有効化
ENV CHOKIDAR_USEPOLLING=true
# 開発モードでアプリケーションを実行
CMD ["npm", "run", "dev"]

# テストステージ
FROM builder AS test
USER appuser
CMD ["npm", "num", "test"]

# 本番ステージ
FROM builder AS production
USER appuser
# プロダクションビルド
RUN npm run build