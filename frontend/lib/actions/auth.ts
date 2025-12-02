"use server"

import { redirect } from "next/navigation"
import { apiRequest } from "../api/client"
import { setAuthToken, clearAuthToken } from "../api/token"
import { AuthResponse, LoginCredentials, RegisterCredentials, User } from "../types/auth"

/**
 * ログイン
 * POST /login
 */
export async function login(credentials: LoginCredentials): Promise<User> {
  const response = await apiRequest<AuthResponse>("/login", {
    method: "POST",
    body: JSON.stringify(credentials),
    skipAuth: true, // ログインAPIは認証不要
  })

  // トークンをCookieに保存
  await setAuthToken(response.token)

  return response.user
}

/**
 * フォームアクション用のログイン
 * <form action={loginAction}> で使用
 */
export async function loginAction(formData: FormData) {
  const credentials: LoginCredentials = {
    username: formData.get("username") as string,
    password: formData.get("password") as string,
  }

  try {
    await login(credentials)
    redirect("/todos")
  } catch (error) {
    // エラーハンドリングは呼び出し側で行う
    throw error
  }
}

/**
 * 新規登録
 * POST /register
 */
export async function register(credentials: RegisterCredentials): Promise<User> {
  const response = await apiRequest<AuthResponse>("/register", {
    method: "POST",
    body: JSON.stringify(credentials),
    skipAuth: true, // 登録APIは認証不要
  })

  // トークンをCookieに保存
  await setAuthToken(response.token)

  return response.user
}

/**
 * フォームアクション用の新規登録
 * <form action={registerAction}> で使用
 */
export async function registerAction(formData: FormData) {
  const credentials: RegisterCredentials = {
    username: formData.get("username") as string,
    email: formData.get("email") as string,
    password: formData.get("password") as string,
  }

  try {
    await register(credentials)
    redirect("/todos")
  } catch (error) {
    throw error
  }
}

/**
 * ログアウト
 * トークンを削除してログインページにリダイレクト
 */
export async function logout() {
  await clearAuthToken()
  redirect("/login")
}

/**
 * 現在のユーザー情報を取得
 * 認証チェック用 (将来実装)
 */
export async function getCurrentUser(): Promise<User | null> {
  try {
    // TODO: バックエンドに /auth/me エンドポイントを追加後、実装
    // const user = await apiRequest<User>("/auth/me")
    // return user
    return null
  } catch (error) {
    return null
  }
}
