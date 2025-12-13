export default defineNuxtRouteMiddleware((to) => {
  const { isAuthenticated, initialize } = useAuth()

  // クライアントサイドでのみ初期化（localStorageアクセス）
  if (import.meta.client) {
    initialize()
  }

  // ログイン前専用ページ（authレイアウトを使用するページ）
  const guestOnlyPages = ['/login', '/register']

  // ログイン必須ではないページ（ランディングページなど）
  // 現在は空配列（将来LPを追加する場合はここに追加）
  const publicPages: string[] = []

  const isGuestOnlyPage = guestOnlyPages.includes(to.path)
  const isPublicPage = publicPages.includes(to.path)

  // 認証済みユーザーがログイン前専用ページにアクセスした場合
  if (isAuthenticated.value && isGuestOnlyPage) {
    return navigateTo('/')
  }

  // 未認証ユーザーがログイン必須ページにアクセスした場合
  if (!isAuthenticated.value && !isGuestOnlyPage && !isPublicPage) {
    return navigateTo('/login')
  }
})
