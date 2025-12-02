/**
 * API共通エラークラス
 */
export class ApiError extends Error {
  constructor(
    message: string,
    public statusCode?: number,
    public code?: string
  ) {
    super(message)
    this.name = "ApiError"
  }
}

/**
 * ネットワークエラー
 * タイムアウトや接続エラーなど
 */
export class NetworkError extends ApiError {
  constructor(message: string = "Network request failed") {
    super(message)
    this.name = "NetworkError"
  }
}

/**
 * 認証エラー (401)
 */
export class AuthenticationError extends ApiError {
  constructor(message: string = "Authentication failed") {
    super(message, 401)
    this.name = "AuthenticationError"
  }
}

/**
 * 認可エラー (403)
 */
export class AuthorizationError extends ApiError {
  constructor(message: string = "Authorization failed") {
    super(message, 403)
    this.name = "AuthorizationError"
  }
}

/**
 * リソース未検出エラー (404)
 */
export class NotFoundError extends ApiError {
  constructor(message: string = "Resource not found") {
    super(message, 404)
    this.name = "NotFoundError"
  }
}
