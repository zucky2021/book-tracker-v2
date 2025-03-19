# ベースイメージ
FROM node:22.14-alpine as builder

# 作業ディレクトリ指定
WORKDIR /app

# package.jsonとpackage-lock.jsonをコピーして依存関係をインストール
COPY package.json package-lock.json ./
RUN npm install

# プロジェクト全体をコピー
COPY . .

# 環境変数でホットリロードを有効化
ENV CHOKIDAR_USEPOLLING=true

# Create non-root user
RUN addgroup -g 1000 appuser && \
  adduser -u 1000 -G appuser -D -H appuser
USER appuser

EXPOSE 3000

# 開発モードでアプリケーションを実行
CMD ["npm", "run", "dev"]
