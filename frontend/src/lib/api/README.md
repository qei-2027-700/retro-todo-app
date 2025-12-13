# API Client構成ガイド

## Go/Echo バックエンド + Next.js フロントエンド

このスプリントでは、バックエンドがGo/Echoで構築されており、
Next.jsはあくまで**フロントエンド**として動作します。

## 重要な設計原則

### ❌ Server Actionsではない
`lib/actions/`配下の関数は、**Server Actionsではありません**。
これらは単に**Server Component内で呼ばれるAPI呼び出し関数**です。

### ✅ 正しい理解
```
Browser → Next.js Server Component → lib/actions/todos.ts
          → lib/api/client.ts → Go/Echo Backend
```

## ファイル構成の役割

### `lib/api/client.ts`
- 役割: fetch()のラッパー
- "use server": **不要**
- 理由: 単なるHTTPクライアント

### `lib/actions/todos.ts`
- 役割: Server Component用のデータ取得関数
- "use server": **不要** (Server Componentから呼ばれるだけ)
- 理由: formアクションとして使われない

### `lib/actions/auth.ts` (将来実装)
- 役割: ログインフォームのアクション
- "use server": **必要**
- 理由: `<form action={login}>`として使われる

## Server Actionsが必要な例

```typescript
// lib/actions/auth.ts
"use server" // ✅ これは必要

export async function login(formData: FormData) {
  // formから送信されたデータを処理
  const username = formData.get("username")
  const password = formData.get("password")

  // Go/Echo APIを呼ぶ
  const response = await apiRequest("/login", {
    method: "POST",
    body: JSON.stringify({ username, password }),
    skipAuth: true,
  })

  // トークンを保存
  await setAuthToken(response.token)
  redirect("/todos")
}
```

使用例:
```typescript
// app/login/page.tsx
import { login } from "@/lib/actions/auth"

export default function LoginPage() {
  return (
    <form action={login}>  {/* ← Server Action */}
      <input name="username" />
      <input name="password" />
      <button type="submit">ログイン</button>
    </form>
  )
}
```

## データ取得は"use server"不要

```typescript
// lib/actions/todos.ts
// ❌ "use server" は不要

export async function getTodos(): Promise<Todo[]> {
  return apiRequest<Todo[]>("/todos")
}
```

使用例:
```typescript
// app/todos/page.tsx (Server Component)
import { getTodos } from "@/lib/actions/todos"

export default async function TodosPage() {
  const todos = await getTodos() // Server Componentから呼ぶだけ
  return <div>{/* ... */}</div>
}
```
