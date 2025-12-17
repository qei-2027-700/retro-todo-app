export default defineNuxtRouteMiddleware((to) => {
  const { isAuthenticated } = useAuth()

  // 認証不要のパブリックページ（パスベース）
  const publicPages = ['/login', '/register']
  const isPublicPageByPath = publicPages.includes(to.path)

  // メタデータベースでのパブリックページ判定
  const isPublicPageByMeta = to.meta.auth === false

  const isPublicPage = isPublicPageByPath || isPublicPageByMeta

  // パブリックページの場合
  if (isPublicPage) {
    // 認証済みユーザーがログインページにアクセスした場合はダッシュボードにリダイレクト
    if (isAuthenticated.value && (to.path === '/login' || to.path === '/register')) {
      return navigateTo('/dashboard')
    }
    return // パブリックページはそのまま通す
  }

  // プライベートページ（認証が必要）
  if (!isAuthenticated.value) {
    return navigateTo('/login')
  }
})