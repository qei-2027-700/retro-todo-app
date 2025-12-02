/**
 * User型定義
 * GoバックエンドのUser構造体に対応
 */
export type User = {
  id: string
  name: string
  email: string
  created_at: string
  updated_at: string
}

/**
 * User作成時のデータ
 */
export type CreateUserData = Pick<User, "name" | "email">

/**
 * User更新時のデータ
 */
export type UpdateUserData = Partial<CreateUserData>
