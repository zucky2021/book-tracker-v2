# Basic design (基本設計)

TODO:こちらGithub上で確認
![Stack architecture](stack-architecture.png)

## 使用技術
- backend
  - GO(Gin)
- frontend
  - npm
  - Vite
  - React
  - TypeScript
- API
  - Google Books API
- DB
  - MySQL

## ソフトウェアアーキテクチャ
- backend
  - DDD

## DB

```mermaid

---
title: "ER図"
---

erDiagram

  memos {
    bigint id PK
    bigint user_id "Google User ID"
    string book_id "Google Book ID"
    text text "メモ"
    timestamp crated_at
    timestamp updated_at
  }

```

## ディレクトリ構造
```
book-tracker-v2/
├── .vscode/
│   └── .settings.json
├── backend/
│   └── Dockerfile
├── frontend/
│   └── Dockerfile
├── .coderabbit.yaml
└── docker-compose.yml
```
[Reference](https://tree.nathanfriend.com/?s=(%27options!(%27fancy!true~fullPath!false~trailingSlash!true~rootDot!false)~5(%275%27book-tra4-v2*.vscode*3.settings.json*back0font0.coderabbi6do4-compose6%27)~version!%271%27)*%5Cn30end*3Do4file*3%20%204cker5source!6.yml*%0165430*)