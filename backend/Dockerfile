# ベースイメージの指定
FROM golang:1.24 AS builder

# 作業ディレクトリ指定
WORKDIR /app

# 依存関係のダウンロード
COPY go.mod go.sum ./
RUN go mod download

# Air: ホットリロード
# Swag: APIドキュメントの生成
RUN go install github.com/air-verse/air@latest && \
  go install github.com/swaggo/swag/cmd/swag@latest

# プロジェクトの全てのファイルをコピー
COPY . .

# airの設定ファイルをコピー
COPY .air.toml .

# airを使ってアプリケーションを実行
CMD ["air"]
