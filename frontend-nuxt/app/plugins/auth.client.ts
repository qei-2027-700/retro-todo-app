export default defineNuxtPlugin(() => {
  const { initialize } = useAuth()
  
  // クライアントサイドでのみ認証状態を初期化
  initialize()
})