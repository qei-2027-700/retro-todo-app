<script setup lang="ts">
import BaseButton from './BaseButton.vue'

// Props定義
interface Props {
  label?: string
  icon?: string
  size?: 'sm' | 'md' | 'lg'
  darkMode?: boolean
  disabled?: boolean
}

const {
  label = '作成',
  icon = '➕',
  size = 'md',
  darkMode = true,
  disabled = false,
} = defineProps<Props>()

// Emit定義
const emit = defineEmits<{
  click: [event: MouseEvent]
}>()

const handleClick = (event: MouseEvent) => {
  emit('click', event)
}
</script>

<template>
  <BaseButton
    variant="primary"
    :size="size"
    :dark-mode="darkMode"
    :disabled="disabled"
    @click="handleClick"
  >
    <span class="text-lg">{{ icon }}</span>
    <span>{{ label }}</span>
  </BaseButton>
</template>
