"use server"

import { cookies } from "next/headers"

const AUTH_TOKEN_KEY = "auth_token"

/**
 * Cookieから認証トークンを取得
 * lib/api/client.ts から呼ばれる
 */
export async function getAuthToken(): Promise<string | undefined> {
  // TODO: 認証機能実装後はCookieから取得する
  // const cookieStore = await cookies()
  // return cookieStore.get(AUTH_TOKEN_KEY)?.value

  // 一時的な固定トークン (開発用)
  // 本番では必ず認証機能を実装してください
  return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6Im15dXNlciIsImlzcyI6InJldHJvLXRvZG8tYXBpIiwiZXhwIjoxNzY0NjUwNzQ5LCJpYXQiOjE3NjQ1NjQzNDl9.kdTpk3vLRt1X6dmpYqVNQcAc9jk5FgmbindvnPzX5Bs"
}

/**
 * Cookieに認証トークンを保存
 * lib/actions/auth.ts から呼ばれる
 */
export async function setAuthToken(token: string): Promise<void> {
  const cookieStore = await cookies()
  cookieStore.set(AUTH_TOKEN_KEY, token, {
    httpOnly: true,
    secure: process.env.NODE_ENV === "production",
    sameSite: "lax",
    maxAge: 60 * 60 * 24 * 7, // 7日間
  })
}

/**
 * Cookieから認証トークンを削除
 * lib/actions/auth.ts から呼ばれる
 */
export async function clearAuthToken(): Promise<void> {
  const cookieStore = await cookies()
  cookieStore.delete(AUTH_TOKEN_KEY)
}
