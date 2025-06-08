# コーディングルール

## 共通ルール

- ボーイスカウトルールに従う：コードを見つけたときよりも綺麗にする
- YAGNI(You Aren't Gonna Need It)の原則に従う：必要になるまで機能を実装しない
- DRY(Don't Repeat Yourself)の原則に従う：再利用可能なロジックを関数に抽出し、コードの重複を避ける
- KISS(Keep It Simple Stupid)の原則に従う：シンプルで分かりやすいコードを書く
- SOLIDの原則に従う
- タイポを検出し、修正を提案する

## TypeScript/JavaScript

- インデントには2スペースを使用
- Google JavaScript Style Guideに準拠
- MDNにリストされている非推奨または廃止された機能の使用を避ける
- プロパティ名にはキャメルケースを使用し、必要に応じてPropTypesを適用
- フックの使用に関する公式Reactガイドラインに従う
- インラインスタイルを避ける
- 'var'の代わりに'let'または'const'を使用
- パフォーマンス最適化（メモ化、レンダリング最適化）を適切に実装
- 不要な再レンダリングを防ぐために、関数コンポーネントにはReact.memo()を、クラスコンポーネントにはPureComponentを使用
- コード分割と遅延ロードを実装
- セキュリティ脆弱性（特にXSS）に注意
- ユーザー入力をサニタイズし、dangerouslySetInnerHTMLの使用を避ける

## Go

- インデントにはタブを使用（Go言語の標準的な方法）
- gofmtツールを使用してコードをフォーマット
- goimportsを使用してインポートを管理
- go vetを使用して疑わしい構造を検出
- golintを使用してスタイルエラーを検出
- テストとベンチマークの実行
- Go Style Guideに従う
- クリーンアーキテクチャの原則に従う
- DDD（Domain-Driven Design）の原則を適用

## React (TSX)

- セマンティックなHTML要素を使用
- Reactのベストプラクティスに従う
- WAI-ARIA属性を実装
- 適切なキーボードナビゲーションサポート
- 可視テキストのない要素にはaria-labelまたはaria-labelledbyを使用
- モーダルやその他の動的コンテンツにフォーカス管理を実装

## テスト

- テストフレームワークのベストプラクティスに従う
- テストの説明を詳細に記述
- 適切なアサーションとマッチャーを使用
- 重要なパスとエッジケースのテストカバレッジを実装
- テストを独立させ、テスト間で状態を共有しない

## Docker

- Dockerfile Best Practicesに従う
- マルチステージビルドを活用
- 適切なベースイメージを選択
- ビルドキャッシュを効果的に利用
- 最小権限の原則に従う
- .dockerignoreファイルを使用
- RUN命令を結合
- 適切な環境変数の設定
- 必要最小限のファイルのみをコピー
- ヘルスチェックの実装
- Linterツールの使用を検討

## Markdown

- [markdownlint-cli2](https://github.com/DavidAnson/markdownlint-cli2)に準拠していること
