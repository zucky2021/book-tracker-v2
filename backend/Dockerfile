# ベースイメージの指定
FROM golang:1.24-alpine AS build

# 非rootユーザーの作成
RUN addgroup -g 1000 appuser && \
  adduser -u 1000 -G appuser -D -H appuser

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

# ビルドキャッシュディレクトリを作成し、所有権を変更
RUN mkdir -p /home/appuser/.cache/go-build && chown -R appuser:appuser /home/appuser/.cache

# 非ルートユーザーに切り替え
USER appuser

# airを使ってアプリケーションを実行
CMD ["air"]

# 本番用のビルドステージ TODO:適宜のビルドを追加
FROM build AS production
