"use server"

import { cookies } from "next/headers"

const AUTH_TOKEN_KEY = "auth_token"

/**
 * Cookieから認証トークンを取得
 * lib/api/client.ts から呼ばれる
 */
export async function getAuthToken(): Promise<string | undefined> {
  const cookieStore = await cookies()
  return cookieStore.get(AUTH_TOKEN_KEY)?.value
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
