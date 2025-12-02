// import { revalidateTag } from "next/cache"
import { apiRequest } from "../api/client"
import { Todo, CreateTodoData, UpdateTodoData } from "../types/todo"

/**
 * 全Todoを取得
 * GET /todos
 */
export async function getTodos(): Promise<Todo[]> {
  return apiRequest<Todo[]>("/todos", {
    next: { tags: ["todos"] },
  })
}

/**
 * IDでTodoを取得
 * GET /todos/:id
 */
export async function getTodoById(id: string): Promise<Todo> {
  return apiRequest<Todo>(`/todos/${id}`, {
    next: { tags: ["todos", `todo-${id}`] },
  })
}

/**
 * 新しいTodoを作成
 * POST /todos
 */
export async function createTodo(data: CreateTodoData): Promise<Todo> {
  const todo = await apiRequest<Todo>("/todos", {
    method: "POST",
    body: JSON.stringify(data),
  })

  // キャッシュを再検証
  // TODO: Next.js 16でのrevalidateTag使い方を確認
  // await revalidateTag("todos")

  return todo
}

/**
 * Todoを更新
 * PUT /todos/:id
 */
export async function updateTodo(id: string, data: UpdateTodoData): Promise<Todo> {
  const todo = await apiRequest<Todo>(`/todos/${id}`, {
    method: "PUT",
    body: JSON.stringify(data),
  })

  // キャッシュを再検証
  // TODO: Next.js 16でのrevalidateTag使い方を確認
  // await revalidateTag("todos")
  // await revalidateTag(`todo-${id}`)

  return todo
}

/**
 * Todoを削除
 * DELETE /todos/:id
 */
export async function deleteTodo(id: string): Promise<void> {
  await apiRequest<void>(`/todos/${id}`, {
    method: "DELETE",
  })

  // キャッシュを再検証
  // TODO: Next.js 16でのrevalidateTag使い方を確認
  // await revalidateTag("todos")
  // await revalidateTag(`todo-${id}`)
}
