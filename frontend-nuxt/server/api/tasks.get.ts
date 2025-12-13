export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const useMock = config.useMock === true

  if (useMock) {
    const mockTasks = [
      { id: 1, title: '健康保険ハガキ連絡', completed: false, tag: '2510-4' },
      { id: 2, title: '部屋の片付け', completed: false, tag: 'Personal Sprint' },
      { id: 3, title: '部屋の片付け', completed: false, tag: 'バックログ' },
      { id: 4, title: 'r-部屋作り', completed: false, tag: 'バックログ' },
      { id: 5, title: '税関係の書類確認', completed: false, tag: 'Personal Sprint' },
      { id: 6, title: 'r-ホワイトニング', completed: false, tag: 'Personal Sprint' },
    ]

    return mockTasks
  }

  // 本番: Go backendへプロキシ
  const backendUrl = config.public.apiBase
  const data = await $fetch('/tasks', { baseURL: backendUrl })
  return data
})
