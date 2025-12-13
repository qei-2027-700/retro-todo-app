export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const useMock = config.useMock === true

  if (useMock) {
    const mockRecentResults = [
      {
        id: '1',
        title: '部屋の片付け',
        subtitle: 'Personal Sprint',
        icon: 'heroicons:check',
        type: 'task' as const,
        completed: true,
      },
      {
        id: '2',
        title: '健康保険ハガキ連絡',
        subtitle: '2510-4',
        icon: 'heroicons:check',
        type: 'task' as const,
        completed: true,
      },
      {
        id: '3',
        title: 'エアコンマニュアル',
        subtitle: 'Personal Sprint',
        icon: 'heroicons:check',
        type: 'task' as const,
        completed: true,
      },
    ]

    return mockRecentResults
  }

  // 本番: Go backendへプロキシ
  const backendUrl = config.public.apiBase
  const data = await $fetch('/tasks/search/recent', { baseURL: backendUrl })
  return data
})
