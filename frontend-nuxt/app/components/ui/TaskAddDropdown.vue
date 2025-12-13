<script setup lang="ts">
interface Props {
  darkMode?: boolean
}

const { darkMode = true } = defineProps<Props>()

const emit = defineEmits<{
  addTask: []
  addApprovalRequest: []
  addMilestone: []
  addSection: []
}>()

const menuItems = [
  {
    id: 'task',
    label: 'タスク',
    icon: 'heroicons:check',
    badge: 'デフォルト',
    action: () => emit('addTask'),
  },
    {
    id: 'section',
    label: 'セクション',
    icon: 'heroicons:list-bullet',
    shortcut: 'Tab | N',
    action: () => emit('addSection'),
  },
]
</script>

<template>
  <BaseDropdown :items="menuItems" :dark-mode="darkMode" position="left">
    <template #trigger="{ isOpen }">
      <button
        class="flex items-center gap-2 px-4 py-2 rounded text-sm transition-colors"
        :class="[
          darkMode
            ? ['bg-gray-700', 'hover:bg-gray-600', 'text-gray-300']
            : ['bg-gray-200', 'hover:bg-gray-300', 'text-gray-700'],
        ]"
      >
        <Icon name="heroicons:plus" class="w-4 h-4" />
        <span>タスクを追加</span>
        <svg
          class="w-4 h-4 transition-transform"
          :class="[
            isOpen ? 'rotate-180' : '',
          ]"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M19 9l-7 7-7-7"
          />
        </svg>
      </button>
    </template>
  </BaseDropdown>
</template>
