export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const useMock = config.useMock === true

  if (useMock) {
    const mockSections = [
      {
        id: 'section1',
        name: '2510-4 - IT',
        count: 1,
        expanded: true,
        tasks: [
          {
            id: 'task1',
            name: '健康保険ハガキ連絡',
            completed: false,
            dueDate: '',
            collaborators: [],
            project: '2510-4',
            projectColor: 'bg-purple-500',
            tags: ['IT'],
          },
        ],
      },
      {
        id: 'section2',
        name: 'Personal Sprint - IT',
        count: 3,
        expanded: true,
        tasks: [
          {
            id: 'task2',
            name: '部屋の片付け',
            completed: false,
            dueDate: '',
            collaborators: ['コラボレータ...'],
            project: 'Personal Sprint',
            projectColor: 'bg-purple-500',
            tags: ['バックログ', 'IT'],
            hasSubtasks: true,
            subtaskCount: 6,
          },
          {
            id: 'task3',
            name: '税関係の書類確認',
            completed: false,
            dueDate: '',
            collaborators: [],
            project: 'Personal Sprint',
            projectColor: 'bg-purple-500',
            tags: ['バックログ', 'IT'],
          },
          {
            id: 'task4',
            name: 'r-ホワイトニング',
            completed: false,
            dueDate: '',
            collaborators: [],
            project: 'Personal Sprint',
            projectColor: 'bg-purple-500',
            tags: ['バックログ', 'IT'],
          },
        ],
      },
      {
        id: 'section3',
        name: 'バックログ - IT',
        count: 5,
        expanded: true,
        tasks: [
          {
            id: 'task5',
            name: '部屋の片付け',
            completed: false,
            dueDate: '',
            collaborators: [],
            project: 'バックログ',
            projectColor: 'bg-blue-500',
            tags: ['IT'],
            hasSubtasks: true,
            subtaskCount: 2,
          },
          {
            id: 'task6',
            name: 'r-部屋作り',
            completed: false,
            dueDate: '',
            collaborators: [],
            project: 'バックログ',
            projectColor: 'bg-blue-500',
            tags: ['IT'],
          },
          {
            id: 'task7',
            name: '税関係の書類確認',
            completed: false,
            dueDate: '',
            collaborators: [],
            project: 'Personal Sprint',
            projectColor: 'bg-purple-500',
            tags: ['バックログ', 'IT'],
          },
          {
            id: 'task8',
            name: 'r-ホワイトニング',
            completed: false,
            dueDate: '',
            collaborators: [],
            project: 'Personal Sprint',
            projectColor: 'bg-purple-500',
            tags: ['バックログ', 'IT'],
          },
          {
            id: 'task9',
            name: '靴を洗う',
            completed: false,
            dueDate: '',
            collaborators: [],
            project: 'バックログ',
            projectColor: 'bg-blue-500',
            tags: [],
          },
        ],
      },
      {
        id: 'section4',
        name: 'プロジェクトなし',
        count: 0,
        expanded: true,
        tasks: [
          {
            id: 'task10',
            name: 'PayPay発券',
            completed: false,
            dueDate: '4月 5日',
            collaborators: [],
            project: '',
            projectColor: '',
            tags: ['自分だけ'],
          },
        ],
      },
    ]

    return mockSections
  }

  // 本番: Go backendへプロキシ
  // const backendUrl = config.public.apiBase
  // const data = await $fetch('/tasks/sections', { baseURL: backendUrl })
  // return data
})
