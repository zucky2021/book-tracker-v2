# Storage

## 用途

### 画像の保存

## 使用技術

### LocalStack (Local env)

- 確認方法

```bash
# インタラクティブモードでコンテナに接続
docker-compose exec -it localstack bash
```

```bash
# バケット一覧
awslocal s3 ls

# バケット内のファイル一覧
awslocal s3 ls s3://local-book-tracker

# バケット内の全てのファイルを一覧
awslocal s3 ls s3://local-book-tracker --recursive
```
