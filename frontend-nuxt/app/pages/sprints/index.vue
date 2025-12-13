<script setup lang="ts">
const activeTab = ref('list')
const filterCount = ref(1)

// TODO: 後でルートパラメータに応じて動的に取得
// 現在は固定のスプリントIDを使用
const sprintId = 'personal-sprint'
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
    <!-- スプリントヘッダー -->
    <div class="border-b border-gray-700 bg-gray-800">
      <div class="px-8 py-4">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center gap-3">
            <!-- スプリントアイコン -->
            <div class="w-10 h-10 bg-purple-500 rounded flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
              </svg>
            </div>

            <!-- スプリント名 -->
            <h1 class="text-2xl font-semibold text-white">Personal Sprint</h1>

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

    <!-- タスクリスト -->
    <div class="bg-gray-900">
      <!-- テーブルヘッダー -->
      <div class="grid grid-cols-[1fr,200px,50px] border-b border-gray-700 px-8 py-3 bg-gray-800 sticky top-14">
        <div class="text-sm font-medium text-gray-400">名前</div>
        <div class="text-sm font-medium text-gray-400">期日</div>
        <div></div>
      </div>

      <!-- セクションリスト -->
      <div>
        <div v-for="section in sections" :key="section.id" class="border-b border-gray-800">
          <!-- セクションヘッダー -->
          <div class="px-8 py-3 bg-gray-850 hover:bg-gray-800 transition-colors cursor-pointer" @click="toggleSection(section.id)">
            <div class="flex items-center gap-2">
              <svg
                :class="['w-4 h-4 text-gray-400 transition-transform', section.expanded ? 'rotate-90' : '']"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
              <span class="text-sm font-semibold text-white">{{ section.name }}</span>
            </div>
          </div>

          <!-- タスク -->
          <div v-if="section.expanded">
            <div
              v-for="task in section.tasks"
              :key="task.id"
              class="grid grid-cols-[1fr,200px,50px] px-8 py-3 hover:bg-gray-800 transition-colors group border-b border-gray-800/50"
            >
              <!-- タスク名 -->
              <div class="flex items-center gap-3">
                <!-- 折りたたみアイコン（サブタスクがある場合） -->
                <button v-if="task.hasSubtasks" class="text-gray-500 hover:text-gray-300">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                  </svg>
                </button>

                <!-- チェックボックス -->
                <button
                  @click="toggleTask(task.id)"
                  :class="[
                    'w-5 h-5 rounded border-2 transition-colors flex-shrink-0',
                    task.completed
                      ? 'bg-green-500 border-green-500'
                      : 'border-gray-500 hover:border-gray-400',
                  ]"
                >
                  <svg
                    v-if="task.completed"
                    class="w-full h-full text-white"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="3"
                      d="M5 13l4 4L19 7"
                    />
                  </svg>
                </button>

                <!-- タスク名 -->
                <span :class="['text-sm text-gray-300', task.completed ? 'line-through opacity-50' : '']">
                  {{ task.name }}
                </span>

                <!-- サブタスクカウント -->
                <span v-if="task.hasSubtasks" class="text-xs text-gray-500">{{ task.subtaskCount }} 🔗</span>
              </div>

              <!-- 期日 -->
              <div class="flex items-center">
                <span v-if="task.dueDate" class="text-sm text-orange-400">{{ task.dueDate }}</span>
              </div>

              <!-- アクション -->
              <div class="flex items-center justify-end opacity-0 group-hover:opacity-100 transition-opacity">
                <button class="p-1 hover:bg-gray-700 rounded">
                  <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
                  </svg>
                </button>
              </div>
            </div>

            <!-- タスク追加ボタン -->
            <button
              @click="addTask(section.id)"
              class="w-full px-8 py-3 text-left text-sm text-gray-500 hover:text-gray-300 hover:bg-gray-800 transition-colors"
            >
              タスクを追加...
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
