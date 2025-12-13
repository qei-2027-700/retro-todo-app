/**
 * 統一APIクライアントプラグイン
 *
 * このプラグインは、アプリケーション全体で使用する統一されたAPIクライアントを提供します。
 * - 自動的にAuthorizationヘッダーにトークンを付与
 * - 401エラー時の自動ログアウト処理
 */
export default defineNuxtPlugin(() => {
  const config = useRuntimeConfig()
  const { token, logout } = useAuth()

  // $apiClient: 認証が必要なAPIリクエスト用
  const $apiClient = $fetch.create({
    baseURL: config.public.apiBase,
    onRequest({ options }) {
      // トークンが存在する場合、Authorizationヘッダーに追加
      if (token.value) {
        const headers = new Headers(options.headers)
        headers.set('Authorization', `Bearer ${token.value}`)
        options.headers = headers
      }
    },
    onResponseError({ response }) {
      // 401エラー（未認証）の場合、ログアウト処理
      if (response.status === 401) {
        logout()
      }
    },
  })

  // $publicApi: 認証不要なAPIリクエスト用（ログイン、登録など）
  const $publicApi = $fetch.create({
    baseURL: config.public.apiBase,
  })

  return {
    provide: {
      apiClient: $apiClient,
      publicApi: $publicApi,
    },
  }
})
