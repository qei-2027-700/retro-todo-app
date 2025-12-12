<script setup lang="ts">
definePageMeta({
  layout: 'default',
})

const email = ref('')
const onSubmit = () => {
  // ここでAPI呼び出しなど
  console.log(email)
}
const continueWithGoogle = () => {
  // Google OAuth など
}
</script>

<template>
  <div class="min-h-screen flex flex-col">
    <!-- 上部ロゴ -->
    <!-- <header class="p-6">
      <div class="flex items-center space-x-2">
        <div class="w-8 h-8 rounded-full bg-pink-500" />
        <span class="text-2xl font-semibold tracking-tight">asana</span>
      </div>
    </header> -->

    <!-- 中央コンテンツ -->
    <!-- <main class="flex-1 flex items-center justify-center"> -->
      <div class="w-full max-w-md text-center px-4">
        <h1 class="text-3xl font-semibold mb-2">
          Asana へようこそ
        </h1>
        <p class="text-gray-500 mb-8">
          まずはサインインしてください
        </p>

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

        <!-- メールフォーム -->
        <form @submit.prevent="onSubmit" class="mt-4">
          <div class="mb-4 text-left">
            <label class="block text-xs font-medium text-gray-600 mb-1">
              メールアドレス
            </label>
            <input
              v-model="email"
              type="email"
              required
              class="block w-full border border-gray-300 rounded-md px-3 py-2 text-sm
                     focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>

          <button
            type="submit"
            class="w-full bg-blue-600 hover:bg-blue-700 text-white text-sm font-medium
                   rounded-md py-3 mt-2 transition"
          >
            続行
          </button>
        </form>
      </div>
    <!-- </main> -->
  </div>
</template>
