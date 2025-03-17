# Setup (環境構築)

1. ルートディレクトリ作成
``` bash
mkdir book-tracker-v2 && cd book-tracker-v2
```

2. フロントエンド作成
```bash
npm create vite@latest frontend -- --template react-swc-ts \
&& cd frontend \
&& touch Dockerfile \
&& npm run dev
```

