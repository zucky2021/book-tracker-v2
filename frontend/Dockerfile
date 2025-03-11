# ベースイメージの指定
FROM node:22.14-alpine

# 作業ディレクトリ指定
WORKDIR /app

# Docker キャッシュを活用するために、最初に package.json と package-lock.json をコピー
COPY package.json package-lock.json ./

# 依存関係をインストール
RUN npm install

# プロジェクトの全てのファイルをコピー
COPY . .

# アプリケーションをビルド
RUN npm run build

EXPOSE 3000

# アプリケーションを実行
CMD ["npm", "run", "start"]