/**
 * 認証関連の型定義
 * GoバックエンドのUser/Auth構造体に対応
 */

export type User = {
  id: number
  username: string
  email: string
  external_id?: string | null
  provider: string
  is_active: boolean
  created_at: string
  updated_at: string
}

/**
 * ログインリクエスト
 */
export type LoginCredentials = {
  username: string
  password: string
}

/**
 * 登録リクエスト
 */
export type RegisterCredentials = {
  username: string
  email: string
  password: string
}

/**
 * ログイン/登録レスポンス
 */
export type AuthResponse = {
  token: string
  user: User
}
