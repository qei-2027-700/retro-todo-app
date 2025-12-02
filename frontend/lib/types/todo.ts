/**
 * Todo型定義
 * GoバックエンドのTodo構造体に対応
 */
export type Todo = {
  id: string
  title: string
  description: string
  completed: boolean
  sprint_id: string | null
  created_at: string
  updated_at: string
}

/**
 * Todo作成時のデータ
 */
export type CreateTodoData = Pick<Todo, "title" | "description"> & {
  sprint_id?: string | null
}

/**
 * Todo更新時のデータ
 */
export type UpdateTodoData = Partial<CreateTodoData> & {
  completed?: boolean
}
