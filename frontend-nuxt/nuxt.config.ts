// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  runtimeConfig: {
    // 環境変数 NUXT_USE_MOCK から読み取り、デフォルトは true
    // 環境変数は文字列で読み込まれるため、デフォルト値をbooleanで設定
    useMock: process.env.NUXT_USE_MOCK !== 'false',

    public: {
      appName: 'NAISEI',
      apiBase: process.env.NUXT_PUBLIC_API_BASE || 'http://localhost:8080',

      // 認証だけ別にモックしたい場合のフラグ
      mockAuth: process.env.NUXT_PUBLIC_MOCK_AUTH === 'true',
    },
  },
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },
  modules: [
    '@nuxt/eslint', 
    '@nuxt/fonts', 
    '@nuxt/icon', 
    '@nuxtjs/tailwindcss',
  ],
  tailwindcss: {
    exposeConfig: true,  // VSCode補完用
    viewer: true         // DevToolsで確認可能
  },
  // https://nuxt.com/docs/4.x/directory-structure/app/components
  components: {
    dirs: [
      {
        path: '~/components',
        pathPrefix: false,
      },
    ],
  },
})