<script setup lang="ts">
// バックログは特別な固定プロジェクト
const project = ref({
  id: 'backlog',
  title: 'バックログ',
  type: 'permanent',
  description: 'すべてのタスクが集まる場所',
})

const activeTab = ref('list')
const filterCount = ref(0)

// APIからバックログのタスクを取得
// バックログはスプリントID=2として管理
const { data: sections } = await useFetch('/api/sprints/2/tasks')

const tabs = [
  { id: 'overview', label: '概要', icon: 'heroicons:clipboard-document-list' },
  { id: 'board', label: 'ボード', icon: 'heroicons:view-columns' },
  { id: 'list', label: 'リスト', icon: 'heroicons:list-bullet' },
  { id: 'timeline', label: 'タイムライン', icon: 'heroicons:chart-bar' },
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
            <div class="w-10 h-10 bg-blue-500 rounded flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4" />
              </svg>
            </div>

            <!-- プロジェクト名 -->
            <div>
              <h1 class="text-2xl font-semibold text-white">{{ project.title }}</h1>
              <p class="text-sm text-gray-400">{{ project.description }}</p>
            </div>
          </div>

          <div class="flex items-center gap-2">
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
        </div>
      </div>
    </div>

    <!-- ツールバー -->
    <div class="bg-gray-800 border-b border-gray-700 px-8 py-3">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <button
            @click="handleAddTask"
            class="flex items-center gap-2 px-4 py-2 bg-blue-600 hover:bg-blue-700 rounded text-sm text-white transition-colors"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            <span>タスクを追加</span>
          </button>

          <button
            @click="handleAddSection"
            class="flex items-center gap-2 px-4 py-2 bg-gray-700 hover:bg-gray-600 rounded text-sm text-gray-300 transition-colors"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            <span>セクションを追加</span>
          </button>
        </div>

        <div class="flex items-center gap-3">
          <!-- フィルター -->
          <button
            v-if="filterCount > 0"
            class="flex items-center gap-2 px-3 py-2 bg-blue-600 hover:bg-blue-700 rounded text-sm text-white transition-colors"
          >
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
        <div class="text-sm font-medium text-gray-400">タスク名</div>
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
