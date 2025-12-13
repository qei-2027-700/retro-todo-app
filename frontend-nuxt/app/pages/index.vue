<script setup lang="ts">
const today = new Date()
const dateStr = `${today.getMonth() + 1}月 ${today.getDate()}日 ${['日', '月', '火', '水', '木', '金', '土'][today.getDay()]}曜日`

const tasks = ref([
  { id: 1, title: '健康保険ハガキ連絡', completed: false, tag: '2510-4' },
  { id: 2, title: '部屋の片付け', completed: false, tag: 'Personal Sprint' },
  { id: 3, title: '部屋の片付け', completed: false, tag: 'バックログ' },
  { id: 4, title: 'r-部屋作り', completed: false, tag: 'バックログ' },
  { id: 5, title: '税関係の書類確認', completed: false, tag: 'Personal Sprint' },
  { id: 6, title: 'r-ホワイトニング', completed: false, tag: 'Personal Sprint' },
])

const projects = [
  { id: 1, name: '2510-3', color: 'bg-purple-500' },
  { id: 2, name: 'バックログ', color: 'bg-blue-500' },
  { id: 3, name: '2510-4', color: 'bg-purple-500' },
  { id: 4, name: 'Personal Sprint', color: 'bg-purple-500' },
]

const toggleTask = (taskId: number) => {
  const task = tasks.value.find(t => t.id === taskId)
  if (task) {
    task.completed = !task.completed
  }
}

const addTask = () => {
  console.log('タスクを追加')
}

const addProject = () => {
  console.log('スプリントを追加')
}
</script>

<template>
  <div class="w-full py-8 px-4 lg:px-8">
    <div class="max-w-6xl">
      <!-- ウェルカムセクション -->
      <div class="mb-8">
        <p class="text-sm text-gray-400 mb-2">{{ dateStr }}</p>
        <h1 class="text-3xl font-bold text-white mb-4">keiji さん、こんばんは</h1>

        <div class="flex items-center gap-4">
          <button class="px-4 py-2 bg-gray-700 hover:bg-gray-600 rounded-lg text-sm text-gray-300 transition-colors">
            週 ▼
          </button>
          <div class="flex items-center gap-2 text-sm text-gray-400">
            <span>✓ 0 件のタスクが完了しました</span>
            <span>👥 0 人のコラボレーター</span>
          </div>
          <button class="ml-auto px-4 py-2 bg-gray-700 hover:bg-gray-600 rounded-lg text-sm text-gray-300 transition-colors">
            カスタマイズ
          </button>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- マイタスクセクション -->
        <div class="bg-gray-800 rounded-lg p-6 border border-gray-700">
          <div class="flex items-center gap-3 mb-6">
            <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-purple-600 rounded-full flex items-center justify-center">
              <span class="text-white text-xl">👤</span>
            </div>
            <div>
              <h2 class="text-xl font-semibold text-white flex items-center gap-2">
                マイタスク
                <svg class="w-4 h-4 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                </svg>
              </h2>
              <div class="flex gap-4 text-sm text-gray-400 mt-1">
                <button class="hover:text-white transition-colors">今日</button>
                <button class="hover:text-white transition-colors">期限超過 (4)</button>
                <button class="hover:text-white transition-colors">完了</button>
              </div>
            </div>
          </div>

          <!-- タスクリスト -->
          <div class="space-y-2 mb-4">
            <div
              v-for="task in tasks"
              :key="task.id"
              class="flex items-center gap-3 p-2 rounded hover:bg-gray-700 transition-colors group"
            >
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
              <span :class="['flex-1 text-gray-300', task.completed ? 'line-through opacity-50' : '']">
                {{ task.title }}
              </span>
              <span class="text-xs px-2 py-1 bg-purple-500/20 text-purple-300 rounded">
                {{ task.tag }}
              </span>
              <button class="opacity-0 group-hover:opacity-100 text-gray-500 hover:text-white transition-all">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
                </svg>
              </button>
            </div>
          </div>

          <button
            @click="addTask"
            class="text-sm text-gray-400 hover:text-white transition-colors"
          >
            + タスクを作成
          </button>

          <button class="mt-4 text-sm text-gray-400 hover:text-white transition-colors">
            もっと表示
          </button>
        </div>

        <!-- スプリントセクション -->
        <div class="bg-gray-800 rounded-lg p-6 border border-gray-700">
          <div class="flex items-center justify-between mb-6">
            <h2 class="text-xl font-semibold text-white">スプリント</h2>
            <button class="text-sm text-gray-400 hover:text-white transition-colors">
              最近 ▼
            </button>
          </div>

          <!-- スプリント追加ボタン -->
          <button
            @click="addProject"
            class="w-full p-4 mb-4 border-2 border-dashed border-gray-700 rounded-lg hover:border-gray-600 transition-colors"
          >
            <div class="flex items-center justify-center gap-2 text-gray-400">
              <span class="text-2xl">+</span>
              <span class="text-sm">スプリントを作成</span>
            </div>
          </button>

          <!-- スプリントカード -->
          <div class="space-y-3">
            <div
              v-for="project in projects"
              :key="project.id"
              class="p-4 bg-gray-700/50 rounded-lg hover:bg-gray-700 transition-colors cursor-pointer group"
            >
              <div class="flex items-center gap-3">
                <div :class="['w-8 h-8 rounded flex items-center justify-center', project.color]">
                  <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
                  </svg>
                </div>
                <span class="flex-1 font-medium text-white">{{ project.name }}</span>
                <button class="opacity-0 group-hover:opacity-100 text-gray-500 hover:text-white transition-all">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 非公開のメモ帳セクション -->
      <div class="mt-6 bg-gray-800 rounded-lg p-6 border border-gray-700">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-xl font-semibold text-white flex items-center gap-2">
            非公開のメモ帳
            <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
            </svg>
          </h2>
          <button class="text-gray-500 hover:text-white transition-colors">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        <p class="text-gray-400 text-sm mb-4">
          簡単なメモや重要な参考資料へのリンクを追加しましょう。
        </p>
        <textarea
          class="w-full bg-gray-900 border border-gray-700 rounded-lg p-3 text-gray-300 placeholder-gray-500 focus:outline-none focus:border-gray-600 resize-none"
          rows="3"
          placeholder="メモを入力..."
        />
      </div>
    </div>
  </div>
</template>
