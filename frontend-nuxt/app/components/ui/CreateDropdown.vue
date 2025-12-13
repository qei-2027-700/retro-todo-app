<script setup lang="ts">
import BaseDropdown from './BaseDropdown.vue'
import CreateButton from './CreateButton.vue'


interface MenuItem {
  id: string
  label: string
  icon: string
  action: () => void
}

interface Props {
  items: MenuItem[]
  buttonLabel?: string
  buttonIcon?: string
  position?: 'left' | 'right'
  size?: 'sm' | 'md' | 'lg'
  darkMode?: boolean
}

const {
  items,
  buttonLabel = '作成',
  buttonIcon = 'heroicons:plus',
  position = 'left',
  size = 'md',
  darkMode = true,
} = defineProps<Props>()
</script>

<template>
  <BaseDropdown :items="items" :position="position" :dark-mode="darkMode">
    <template #trigger>
      <CreateButton :label="buttonLabel" :icon="buttonIcon" :size="size" :dark-mode="darkMode" />
    </template>
  </BaseDropdown>
</template>
