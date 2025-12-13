<script setup lang="ts">
import BaseButton from '~/components/ui/BaseButton.vue'

definePageMeta({
  layout: 'auth',
})

const config = useRuntimeConfig()

const { login } = useAuth()
const router = useRouter()

const username = ref('')
const password = ref('')
const isLoading = ref(false)
const errorMessage = ref('')

// TODO: リリース前には除去する
const setSampleCredentials = () => {
  username.value = 'testuser'
  password.value = 'password123'
}

const onSubmit = async () => {
  if (!username.value || !password.value) {
    errorMessage.value = 'ユーザー名とパスワードを入力してください'
    return
  }

  isLoading.value = true
  errorMessage.value = ''

  const result = await login({
    username: username.value,
    password: password.value,
  })

  isLoading.value = false

  if (result.success) {
    router.push('/dashboard')
  }
  else {
    errorMessage.value = result.error || 'ログインに失敗しました'
  }
}

const continueWithGoogle = () => {
  // TODO: Google OAuth実装
  console.log('Google OAuth is not implemented yet')
}
</script>

<template>
  <div class="min-h-screen flex flex-col">
    <!-- <main class="flex-1 flex items-center justify-center"> -->
      <div class="w-full max-w-md text-center px-4">
        <h1 class="text-3xl font-semibold mb-2">
          {{ config.public.appName }}へようこそ
        </h1>
        <p class="text-gray-500 mb-8">
          まずはサインインしてください
        </p>

        <div class="text-dark-500 mb-8">
          テスト用アカウント
          <ul>
            <li>testuser
            </li>
            <li>password123</li>
          </ul>
        </div>

        <!-- Googleで続行 -->
        <button
          type="button"
          @click="continueWithGoogle"
          class="w-full flex items-center justify-center border border-gray-300 rounded-md py-3 mb-6
                 hover:bg-gray-50 transition text-sm font-medium"
        >
          <span class="mr-2">
            <!-- 簡易 Google アイコン -->
            <span class="inline-block w-5 h-5 rounded-full bg-white border" />
          </span>
          <span>Google で続ける</span>
        </button>

        <!-- 区切り線 -->
        <div class="flex items-center my-4">
          <div class="flex-1 h-px bg-gray-200"></div>
          <span class="mx-4 text-xs text-gray-400">または</span>
          <div class="flex-1 h-px bg-gray-200"></div>
        </div>

        <!-- エラーメッセージ -->
        <div v-if="errorMessage" class="mb-4 p-3 bg-red-50 border border-red-200 rounded-md">
          <p class="text-sm text-red-600">{{ errorMessage }}</p>
        </div>
  
        <div>
          <BaseButton class="mb-4" @click="setSampleCredentials">
            テスト用アカウントのID / PWを入力
          </BaseButton>
        </div>

        <!-- ログインフォーム -->
        <form @submit.prevent="onSubmit" class="mt-4">
          <div class="mb-4 text-left">
            <label class="block text-xs font-medium text-gray-600 mb-1">
              ユーザー名
            </label>
            <input
              v-model="username"
              type="text"
              required
              :disabled="isLoading"
              class="block w-full border border-gray-300 rounded-md px-3 py-2 text-sm
                     focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500
                     disabled:bg-gray-100 disabled:cursor-not-allowed"
            />
          </div>

          <div class="mb-4 text-left">
            <label class="block text-xs font-medium text-gray-600 mb-1">
              パスワード
            </label>
            <input
              v-model="password"
              type="password"
              required
              :disabled="isLoading"
              class="block w-full border border-gray-300 rounded-md px-3 py-2 text-sm
                     focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500
                     disabled:bg-gray-100 disabled:cursor-not-allowed"
            />
          </div>

          <BaseButton
            :disabled="isLoading"
            type="submit"
            class="w-full bg-blue-600 hover:bg-blue-700 text-white text-sm font-medium
                   rounded-md py-3 mt-2 transition disabled:bg-blue-400 disabled:cursor-not-allowed"
          >
            {{ isLoading ? 'ログイン中...' : '続行' }}
          </BaseButton>
        </form>
      </div>
    <!-- </main> -->
  </div>
</template>
