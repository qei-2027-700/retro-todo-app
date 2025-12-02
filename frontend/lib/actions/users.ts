// import { revalidateTag } from "next/cache"
import { apiRequest } from "../api/client"
import { User, CreateUserData, UpdateUserData } from "../types/user"

/**
 * 全ユーザーを取得
 * GET /api/users
 */
export async function getUsers(): Promise<User[]> {
  return apiRequest<User[]>("/api/users", {
    next: { tags: ["users"] },
  })
}

/**
 * IDでユーザーを取得
 * GET /api/users/:id
 */
export async function getUserById(id: string): Promise<User> {
  return apiRequest<User>(`/api/users/${id}`, {
    next: { tags: ["users", `user-${id}`] },
  })
}

/**
 * 新しいユーザーを作成
 * POST /api/users
 */
export async function createUser(data: CreateUserData): Promise<User> {
  const user = await apiRequest<User>("/api/users", {
    method: "POST",
    body: JSON.stringify(data),
  })

  // キャッシュを再検証
  // TODO: Next.js 16でのrevalidateTag使い方を確認
  // revalidateTag("users")

  return user
}

/**
 * ユーザー情報を更新
 * PUT /api/users/:id
 */
export async function updateUser(id: string, data: UpdateUserData): Promise<User> {
  const user = await apiRequest<User>(`/api/users/${id}`, {
    method: "PUT",
    body: JSON.stringify(data),
  })

  // キャッシュを再検証
  // TODO: Next.js 16でのrevalidateTag使い方を確認
  // revalidateTag("users")
  // revalidateTag(`user-${id}`)

  return user
}

/**
 * ユーザーを削除
 * DELETE /api/users/:id
 */
export async function deleteUser(id: string): Promise<void> {
  await apiRequest<void>(`/api/users/${id}`, {
    method: "DELETE",
  })

  // キャッシュを再検証
  // TODO: Next.js 16でのrevalidateTag使い方を確認
  // revalidateTag("users")
  // revalidateTag(`user-${id}`)
}
