/**
 * API設定
 * 環境変数から設定を読み込む
 */
export const API_CONFIG = {
  baseURL: process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8080",
  timeout: Number(process.env.API_TIMEOUT) || 10000,
  retryAttempts: 3,
  retryDelay: 1000,
} as const
