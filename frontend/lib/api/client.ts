import { getAuthToken } from "./token"
import { API_CONFIG } from "./config"
import {
  ApiError,
  NetworkError,
  AuthenticationError,
  AuthorizationError,
  NotFoundError,
} from "./error"

// Go/Echoの標準的なエラーレスポンス型
type ApiErrorResponse = {
  message: string
  code?: string
}

type ApiRequestOptions = RequestInit & {
  timeout?: number
  skipAuth?: boolean // ログインなど認証不要なエンドポイント用
}

/**
 * Go/EchoバックエンドのAPIクライアント
 * 認証トークン、エラーハンドリング、タイムアウトを統一管理
 */
export async function apiRequest<T>(
  endpoint: string,
  options: ApiRequestOptions = {}
): Promise<T> {
  const { timeout = API_CONFIG.timeout, skipAuth = false, ...fetchOptions } = options

  // 認証トークン取得 (skipAuthの場合は取得しない)
  const token = skipAuth ? undefined : await getAuthToken()

  const headers: HeadersInit = {
    "Content-Type": "application/json",
    ...(token && { Authorization: `Bearer ${token}` }),
    ...fetchOptions.headers,
  }

  // タイムアウト制御
  const controller = new AbortController()
  const timeoutId = setTimeout(() => controller.abort(), timeout)

  try {
    const res = await fetch(`${API_CONFIG.baseURL}${endpoint}`, {
      ...fetchOptions,
      headers,
      signal: controller.signal,
    })

    clearTimeout(timeoutId)

    // Go/Echoのエラーレスポンスをハンドリング
    if (!res.ok) {
      const errorData: ApiErrorResponse = await res.json().catch(() => ({
        message: `HTTP Error: ${res.status}`,
      }))

      // ステータスコードごとに適切なエラーをthrow
      switch (res.status) {
        case 401:
          throw new AuthenticationError(errorData.message)
        case 403:
          throw new AuthorizationError(errorData.message)
        case 404:
          throw new NotFoundError(errorData.message)
        default:
          throw new ApiError(
            errorData.message || `HTTP Error: ${res.status}`,
            res.status,
            errorData.code
          )
      }
    }

    // 204 No Contentの場合
    if (res.status === 204) {
      return undefined as T
    }

    return res.json()
  } catch (error) {
    clearTimeout(timeoutId)

    // Abort errorはタイムアウト
    if (error instanceof Error && error.name === "AbortError") {
      throw new NetworkError("Request timeout")
    }

    // ネットワークエラー
    if (error instanceof TypeError) {
      throw new NetworkError()
    }

    // その他のエラーはそのまま再スロー
    throw error
  }
}
