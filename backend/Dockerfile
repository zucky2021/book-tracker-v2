# ベースイメージの指定
FROM golang:1.24-alpine AS base

# 非rootユーザーの作成
RUN addgroup --gid 1000 appuser && \
  adduser --uid 1000 --gid 1000 --disabled-password --gecos "" appuser

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

# 非ルートユーザーに切り替え
USER appuser

# airを使ってアプリケーションを実行
CMD ["air"]

# 本番用のビルドステージ
FROM base AS production

# 本番用の依存関係のビルド
COPY go.mod go.sum ./
RUN go mod download

# 本番用のビルド
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# WORKDIRを設定
WORKDIR /app

# 実行可能ファイルを指定
CMD ["./main"]
