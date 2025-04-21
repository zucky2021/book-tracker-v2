# Setup (環境構築)

## 1 ルートディレクトリ作成

``` bash
mkdir book-tracker-v2 && cd book-tracker-v2
```

## 2 フロントエンド作成

```bash
npm create vite@latest frontend -- --template react-swc-ts \
&& cd frontend \
&& touch Dockerfile \
&& npm run dev
```

## Githooks 共有化

### 実行権限を付与

```bash
chmod +x .githooks/*
```

### Git の設定で上記のディレクトリを有効化

```bash
git config core.hooksPath .githooks
```

## DB接続確認

```bash
docker exec -it book-tracker-db mysql -u root -pRoot12345
```

## AWS (localstack)

- 実行権限を付与

```bash
chmod +x localstack/init-s3.sh
```

- S3バケット確認

```bash
docker-compose exec localstack bash

# バケット一覧
awslocal s3 ls

# バケット内のファイル一覧
awslocal s3 ls s3://local-book-tracker
```
