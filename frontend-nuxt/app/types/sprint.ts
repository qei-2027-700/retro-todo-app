
export interface Sprint {
  id: number
  name: string
  color: string
  is_favorite: boolean
  created_at?: string
  updated_at?: string
}

export interface CreateSprintRequest {
  name: string
  color: string
  is_favorite?: boolean
}

/**
 * スプリント更新リクエスト
 */
export interface UpdateSprintRequest {
  name?: string
  color?: string
  is_favorite?: boolean
}
