export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const useMock = config.useMock === true

  if (useMock) {
    const mockSprints = [
      { id: 2, name: 'バックログ', color: 'bg-blue-500' },
      { id: 3, name: '2510-4', color: 'bg-purple-500' },
      { id: 4, name: 'Personal Sprint', color: 'bg-purple-500' },
    ]

    return mockSprints
  }

  // 本番: Go backendへプロキシ
  const backendUrl = config.public.apiBase
  const data = await $fetch('/sprints', { baseURL: backendUrl })
  return data
})
