# Retro Todo App

個人用の振り返り・TODO管理アプリケーション

## 技術スタック

### Backend
- **Go** + **Echo Framework**
- **PostgreSQL**
- **JWT認証**
- **Swagger/OpenAPI** ドキュメント生成

### Frontend
- **Next.js 16** (App Router)
- **React 19**
- **TypeScript**
- **TailwindCSS 4**

## 開発環境のセットアップ

### Backend

```bash
cd backend

# Swagger ドキュメントを生成
swag init

# 開発サーバーを起動 (Live reload with Air)
air

# または直接実行
GO_ENV=dev go run main.go
```

### Frontend

```bash
cd frontend

# 依存関係をインストール
pnpm install

# 型定義を生成 (Swaggerから自動生成)
pnpm run generate:types

# 開発サーバーを起動
pnpm run dev
```

## 型定義の自動生成

Backend の Swagger ドキュメントから Frontend の TypeScript 型定義を自動生成できます。

### 生成手順（推奨: Taskコマンド）

スプリントルートから以下のコマンドを実行:

```bash
# Backend の Swagger ドキュメントを生成
task be:swag

# Frontend の TypeScript 型定義を生成
task fe:swag

# または両方を一度に実行
task be:swag && task fe:swag
```

### 生成手順（手動実行）

1. Backend で Swagger ドキュメントを生成:
```bash
cd backend
swag init -g cmd/api/main.go -o docs --outputTypes go,json,yaml
```

2. Frontend で型定義を生成:
```bash
cd frontend
pnpm run generate:types
```

これにより、以下のファイルが生成されます:
- `backend/docs/openapi.json` - Swagger 2.0 を OpenAPI 3.0 に変換したファイル
- `frontend/lib/types/api.d.ts` - TypeScript 型定義ファイル

### 型定義の使用例

```typescript
import type { components } from '@/lib/types/api'

// モデル型を使用
type Todo = components['schemas']['model.Todo']
type User = components['schemas']['model.User']
type LoginRequest = components['schemas']['model.LoginRequest']

const todo: Todo = {
  id: 1,
  title: 'Sample Todo',
  description: 'This is a sample',
  completed: false,
  sprint_id: 1,
  created_at: '2024-01-01T00:00:00Z',
  updated_at: '2024-01-01T00:00:00Z',
}
```

### 注意点

- 生成されたファイル (`openapi.json`, `api.d.ts`) は `.gitignore` に含まれています
- Backend の Swagger ドキュメントを更新したら、必ず `task be:swag && task fe:swag` を実行してください
- 型定義は自動生成されるため、直接編集しないでください

## Taskコマンド一覧

スプリントルートで利用可能なTaskコマンド:

### Backend関連
- `task be:dev` - 開発サーバーを起動（ホットリロード）
- `task be:swag` - Swaggerドキュメントを生成
- `task be:migrate` - データベースマイグレーションを実行
- `task be:test` - テストを実行
- `task be:build` - アプリケーションをビルド

### Frontend関連
- `task fe:dev` - 開発サーバーを起動
- `task fe:swag` - TypeScript型定義を生成（Swaggerから）

### Docker関連
- `task start` - すべてのコンテナを起動
- `task down` - すべてのコンテナを停止
- `task ps` - 起動中のコンテナを表示
- `task logs` - すべてのログを表示
- `task logs:be` - Backendのログのみ表示

全てのコマンドを確認するには:
```bash
task --list
```

## デフォルトユーザー

開発環境では以下のテストユーザーが利用可能です:

- **ユーザー名**: `testuser`
- **パスワード**: `password123`
