export const useAuth = () => {
  const user = ref(null)
  const isAuthenticated = computed(() => !!user.value)

  const initialize = () => {
    // モックユーザー情報を設定
    user.value = {
      id: 1,
      name: 'テストユーザー',
      email: 'test@example.com',
      avatar: 'https://i.pravatar.cc/150?img=3'
    }
  }

  const login = async (credentials: { email: string; password: string }) => {
    // モック認証処理
    console.log('ログイン:', credentials)
    
    user.value = {
      id: 1,
      name: 'テストユーザー',
      email: credentials.email,
      avatar: 'https://i.pravatar.cc/150?img=3'
    }

    // ダッシュボードにリダイレクト
    await navigateTo('/dashboard')
  }

  const logout = async () => {
    user.value = null
    await navigateTo('/login')
  }

  const register = async (userData: { name: string; email: string; password: string }) => {
    // モック登録処理
    console.log('登録:', userData)
    
    user.value = {
      id: 1,
      name: userData.name,
      email: userData.email,
      avatar: 'https://i.pravatar.cc/150?img=3'
    }

    await navigateTo('/dashboard')
  }

  return {
    user: readonly(user),
    isAuthenticated,
    initialize,
    login,
    logout,
    register,
  }
}