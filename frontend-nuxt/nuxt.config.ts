// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  runtimeConfig: {
    useMock: true,

    public: {
      appName: 'NAISEI',
      apiBase: 'http://localhost:8080',
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