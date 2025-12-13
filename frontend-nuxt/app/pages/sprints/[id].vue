<script setup lang="ts">
// ルートパラメータからIDを抽出（例: "1-2025-12-week1" → 1）
const route = useRoute()
const sprintId = extractIdFromParam(route.params.id)

// TODO: API から sprint データを取得
// const { fetchSprint } = useSprints()
// const sprint = await fetchSprint(sprintId)

// 仮のデータ（後でAPI呼び出しに置き換え）
const sprint = ref({
  id: sprintId,
  title: 'Personal Sprint',
  completed: false,
})

const activeTab = ref('list')
const filterCount = ref(1)

// APIからセクションデータを取得
const { data: sections } = await useFetch(`/api/sprints/${sprintId}/tasks`)

const tabs = [
  { id: 'overview', label: '概要', icon: 'heroicons:clipboard-document-list' },
  { id: 'board', label: 'ボード', icon: 'heroicons:view-columns' },
  { id: 'list', label: 'リスト', icon: 'heroicons:list-bullet' },
  { id: 'timeline', label: 'タイムライン', icon: 'heroicons:chart-bar' },
  { id: 'dashboard', label: 'ダッシュボード', icon: 'heroicons:presentation-chart-line' },
  { id: 'gantt', label: 'ガント', icon: 'heroicons:chart-bar-square' },
  { id: 'workload', label: 'ワークロード', icon: 'heroicons:scale' },
  { id: 'calendar', label: 'カレンダー', icon: 'heroicons:calendar' },
]

const toggleSection = (sectionId: string) => {
  if (!sections.value) return

  const section = sections.value.find(s => s.id === sectionId)
  if (section) {
    section.expanded = !section.expanded
  }
}

const toggleTask = (taskId: string) => {
  if (!sections.value) return
  
  sections.value.forEach(section => {
    const task = section.tasks.find(t => t.id === taskId)
    if (task) {
      task.completed = !task.completed
    }
  })
}

const addTask = (sectionId?: string) => {
  console.log('タスクを追加:', sectionId)
}

const handleAddTask = () => {
  console.log('タスクを追加')
}

const handleAddApprovalRequest = () => {
  console.log('承認リクエストを追加')
}

const handleAddMilestone = () => {
  console.log('マイルストーンを追加')
}

const handleAddSection = () => {
  console.log('セクションを追加')
}
</script>

<template>
  <div class="w-full">
    <!-- プロジェクトヘッダー -->
    <div class="border-b border-gray-700 bg-gray-800">
      <div class="px-8 py-4">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center gap-3">
            <!-- プロジェクトアイコン -->
            <div class="w-10 h-10 bg-purple-500 rounded flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
              </svg>
            </div>

            <!-- プロジェクト名 -->
            <h1 class="text-2xl font-semibold text-white">{{ sprint.title }}</h1>

            <!-- お気に入り -->
            <button class="text-yellow-400 hover:text-yellow-300 transition-colors">
              <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z" />
              </svg>
            </button>

            <!-- ステータス設定 -->
            <button class="flex items-center gap-2 px-3 py-1 bg-gray-700 hover:bg-gray-600 rounded text-sm text-gray-300 transition-colors">
              <span class="w-2 h-2 bg-gray-500 rounded-full"></span>
              <span>ステータスを設定</span>
            </button>
          </div>

          <div class="flex items-center gap-2">
            <!-- メンバー -->
            <div class="flex -space-x-2">
              <div class="w-8 h-8 rounded-full bg-gradient-to-br from-blue-400 to-purple-600 border-2 border-gray-800">
                <img src="https://i.pravatar.cc/150?img=3" alt="User" class="w-full h-full rounded-full object-cover" />
              </div>
            </div>

            <!-- 共有ボタン -->
            <button class="p-2 hover:bg-gray-700 rounded transition-colors">
              <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
              </svg>
            </button>

            <!-- カスタマイズボタン -->
            <button class="px-4 py-2 bg-gray-700 hover:bg-gray-600 rounded text-sm text-gray-300 transition-colors">
              カスタマイズ
            </button>
          </div>
        </div>

        <!-- タブナビゲーション -->
        <div class="flex items-center gap-6 overflow-x-auto">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            @click="activeTab = tab.id"
            :class="[
              'flex items-center gap-2 px-3 py-2 text-sm whitespace-nowrap transition-colors border-b-2',
              activeTab === tab.id
                ? 'text-white border-white'
                : 'text-gray-400 border-transparent hover:text-gray-200',
            ]"
          >
            <Icon :name="tab.icon" class="w-4 h-4" />
            <span>{{ tab.label }}</span>
          </button>
          <button class="text-gray-400 hover:text-white transition-colors">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- ツールバー -->
    <div class="bg-gray-800 border-b border-gray-700 px-8 py-3">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <TaskAddDropdown
            :dark-mode="true"
            @add-task="handleAddTask"
            @add-approval-request="handleAddApprovalRequest"
            @add-milestone="handleAddMilestone"
            @add-section="handleAddSection"
          />
        </div>

        <div class="flex items-center gap-3">
          <!-- フィルター -->
          <button class="flex items-center gap-2 px-3 py-2 bg-blue-600 hover:bg-blue-700 rounded text-sm text-white transition-colors">
            <span>フィルター: {{ filterCount }}</span>
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>

          <!-- ソート -->
          <button class="flex items-center gap-2 px-3 py-2 hover:bg-gray-700 rounded text-sm text-gray-300 transition-colors">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4h13M3 8h9m-9 4h6m4 0l4-4m0 0l4 4m-4-4v12" />
            </svg>
            <span>ソート</span>
          </button>

          <!-- グループ -->
          <button class="flex items-center gap-2 px-3 py-2 hover:bg-gray-700 rounded text-sm text-gray-300 transition-colors">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
            <span>グループ</span>
          </button>

          <!-- オプション -->
          <button class="flex items-center gap-2 px-3 py-2 hover:bg-gray-700 rounded text-sm text-gray-300 transition-colors">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
            </svg>
            <span>オプション</span>
          </button>

          <!-- 検索 -->
          <button class="p-2 hover:bg-gray-700 rounded text-gray-300 transition-colors">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <ListView
      v-if="activeTab === 'list'"
      :sections="sections"
      @toggle-section="toggleSection"
      @toggle-task="toggleTask"
      @add-task="addTask"
    />

    <BoardView
      v-else-if="activeTab === 'board'"
      :sections="sections"
      @toggle-section="toggleSection"
      @toggle-task="toggleTask"
      @add-task="addTask"
    />
  </div>
</template>
