export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const useMock = config.useMock === true
  const sprintId = getRouterParam(event, 'id')

  if (useMock) {
    const mockSections = [
      {
        id: 'section1',
        name: '1006-101: 承認リクエスト',
        expanded: true,
        tasks: [
          {
            id: 'task1',
            name: 'mindmap',
            completed: false,
            dueDate: '10月 11日',
            hasSubtasks: false,
          },
        ],
      },
      {
        id: 'section2',
        name: 'バックログ',
        expanded: true,
        tasks: [
          {
            id: 'task2',
            name: 'r-ホワイトニング',
            completed: false,
            dueDate: '',
            hasSubtasks: false,
          },
          {
            id: 'task3',
            name: 'ヘアアイロン検討',
            completed: false,
            dueDate: '10月 5日',
            hasSubtasks: false,
          },
          {
            id: 'task4',
            name: 'エアコンマニュアル',
            completed: false,
            dueDate: '',
            hasSubtasks: false,
          },
          {
            id: 'task5',
            name: '部屋の片付け',
            completed: false,
            dueDate: '',
            hasSubtasks: true,
            subtaskCount: 3,
          },
          {
            id: 'task6',
            name: '冷凍庫の中身を消費する：餃子・ブロッコリー',
            completed: false,
            dueDate: '',
            hasSubtasks: false,
          },
          {
            id: 'task7',
            name: 'mcp: claudecodeとgithubなどを連携する',
            completed: false,
            dueDate: '',
            hasSubtasks: false,
          },
          {
            id: 'task8',
            name: 'ポモドーロタイマーをデバイスで買う',
            completed: false,
            dueDate: '',
            hasSubtasks: false,
          },
          {
            id: 'task9',
            name: '自炊計画',
            completed: false,
            dueDate: '',
            hasSubtasks: true,
            subtaskCount: 3,
          },
        ],
      },
    ]

    return mockSections
  }

  // 本番: Go backendへプロキシ
  const backendUrl = config.public.apiBase
  const data = await $fetch(`/sprints/${sprintId}/tasks`, { baseURL: backendUrl })
  return data
})
