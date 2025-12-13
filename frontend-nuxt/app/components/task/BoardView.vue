<script setup lang="ts">
import { RETRO_UNIT_NAME } from '~/constants/common'

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
  <div class="px-8 py-4 overflow-x-auto">
    <div class="flex items-start gap-4">
      <div
        v-for="section in props.sections"
        :key="section.id"
        class="w-72 bg-gray-850 rounded-md p-3 flex-shrink-0"
      >
        <!-- セクションヘッダー -->
        <div class="flex items-center justify-between mb-2">
          <h2 class="text-sm font-semibold text-white">
            {{ section.name }}
          </h2>
          <!-- 省略: メニューアイコンなど -->
        </div>

        <!-- タスクカード -->
        <div class="space-y-2">
          <div
            v-for="task in section.tasks"
            :key="task.id"
            class="bg-gray-800 rounded-md p-3 cursor-pointer hover:bg-gray-750 transition-colors"
          >
            <div class="flex items-start gap-2">
              <button
                @click.stop="emit('toggle-task', task.id)"
                :class="[
                  'w-4 h-4 rounded border-2 flex-shrink-0 mt-0.5',
                  task.completed ? 'bg-green-500 border-green-500' : 'border-gray-500',
                ]"
              />
              <p
                :class="[
                  'text-sm text-gray-200',
                  task.completed && 'line-through opacity-60',
                ]"
              >
                {{ task.name }}
              </p>
            </div>
          </div>
        </div>

        <!-- タスク追加 -->
        <button
          @click="emit('add-task', section.id)"
          class="mt-2 w-full text-left text-xs text-gray-400 hover:text-gray-200 hover:bg-gray-800 rounded px-2 py-1"
        >
          タスクを追加…
        </button>
      </div>

      <!-- 右端に「列を追加」ボタンなど -->
      <button
        class="w-72 flex-shrink-0 border border-dashed border-gray-600 rounded-md p-3 text-sm text-gray-400 hover:bg-gray-850 hover:text-gray-200"
      >
        新しい{{ RETRO_UNIT_NAME }}を追加
      </button>
    </div>
  </div>
</template>
