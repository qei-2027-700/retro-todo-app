<script setup lang="ts">
interface Task {
  id: string
  name: string
  completed: boolean
}
interface Section {
  id: string
  name: string
  tasks: Task[]
}

const props = defineProps<{
  sections: Section[]
}>()

const emit = defineEmits<{
  (e: 'toggle-task', taskId: string): void
  (e: 'add-task', sectionId?: string): void
}>()
</script>

<template>
  <div>List View Component</div>

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
</template>
