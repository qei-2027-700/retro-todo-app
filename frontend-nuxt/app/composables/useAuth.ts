interface User {
  id: number
  username: string
  email?: string
  createdAt?: string
  updatedAt?: string
}

interface LoginRequest {
  username: string
  password: string
}

interface LoginResponse {
  token: string
  user: User
}

const TOKEN_KEY = 'auth_token'
const USER_KEY = 'auth_user'

export const useAuth = () => {
  const { $publicApi } = useNuxtApp()
  const router = useRouter()
  const config = useRuntimeConfig()

  // 認証状態（リアクティブ）
  const token = useState<string | null>('auth_token', () => null)
  const user = useState<User | null>('auth_user', () => null)
  const isAuthenticated = computed(() => !!token.value)

  // 初期化：localStorageからトークンとユーザー情報を復元
  const initialize = () => {
    if (!import.meta.client) return

    const savedToken = localStorage.getItem(TOKEN_KEY)
    const savedUser = localStorage.getItem(USER_KEY)

    if (savedToken) {
      token.value = savedToken
    }

    if (savedUser) {
      try {
        user.value = JSON.parse(savedUser)
      }
      catch (error) {
        console.error('Failed to parse user data:', error)
        localStorage.removeItem(USER_KEY)
      }
    }
  }

  const saveSession = (res: LoginResponse) => {
    token.value = res.token
    user.value = res.user

    if (import.meta.client) {
      localStorage.setItem(TOKEN_KEY, res.token)
      localStorage.setItem(USER_KEY, JSON.stringify(res.user))
    }
  }

  const login = async (credentials: LoginRequest) => {
    try {
      let response: LoginResponse

      if (config.public.mockAuth === true) {
        // Nitro 側のモックAPI (/api/auth/login など) を叩くパターン
        // 例: server/api/auth/login.post.ts でモックレスポンスを返す実装
        response = await $fetch<LoginResponse>('/api/auth/login', {
          method: 'POST',
          body: credentials,
        })
      }
      else {
        // 本番 API
        response = await $publicApi<LoginResponse>('/login', {
          method: 'POST',
          body: credentials,
        })
      }

      saveSession(response)
      return { success: true, data: response }
    }
    catch (error: any) {
      console.error('Login error:', error)
      return {
        success: false,
        error: error?.data?.message || error?.message || 'ログインに失敗しました',
      }
    }
  }

  const logout = async () => {
    token.value = null
    user.value = null

    if (import.meta.client) {
      localStorage.removeItem(TOKEN_KEY)
      localStorage.removeItem(USER_KEY)
    }

    await router.push('/login')
  }

  // トークンをリフレッシュ（必要に応じて実装）
  const refreshToken = async () => {
    // TODO: バックエンドにリフレッシュエンドポイントがあれば実装
  }

  return {
    token: readonly(token),
    user: readonly(user),
    isAuthenticated,
    initialize,
    login,
    logout,
    refreshToken,
  }
}
