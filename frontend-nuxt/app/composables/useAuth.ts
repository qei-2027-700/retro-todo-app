interface User {
  id: number
  name: string
  email: string
  avatar: string
}

// composableの外で状態を定義することで、グローバルな状態として共有される
const user = ref<User | null>(null)

export const useAuth = () => {
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

  const login = async (credentials: { username: string; password: string }): Promise<{ success: boolean; error?: string }> => {
    // モック認証処理
    console.log('ログイン:', credentials)
    
    // Promise.resolveを使って確実に結果を返す
    return new Promise((resolve) => {
      try {
        // 簡単な認証チェック（モック）
        if (credentials.username === 'testuser' && credentials.password === 'password123') {
          user.value = {
            id: 1,
            name: 'テストユーザー',
            email: 'test@example.com',
            avatar: 'https://i.pravatar.cc/150?img=3'
          }

          const result = { success: true }
          console.log('Login success, returning:', result) // デバッグ用
          resolve(result)
        } else {
          const result = { success: false, error: 'ユーザー名またはパスワードが正しくありません' }
          console.log('Login failed, returning:', result) // デバッグ用
          resolve(result)
        }
      } catch (error) {
        console.error('ログインエラー:', error)
        const result = { success: false, error: 'ログイン処理中にエラーが発生しました' }
        console.log('Login error, returning:', result) // デバッグ用
        resolve(result)
      }
    })
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