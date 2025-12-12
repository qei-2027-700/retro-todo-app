<script setup lang="ts">
interface Props {
  type?: 'button' | 'submit' | 'reset'
  variant?: 'primary' | 'secondary' | 'success' | 'danger' | 'outline' | 'ghost'
  size?: 'sm' | 'md' | 'lg'
  disabled?: boolean
  fullWidth?: boolean
  darkMode?: boolean
}

const {
  type = 'button',
  variant = 'primary',
  size = 'md',
  disabled = false,
  fullWidth = false,
  darkMode = false,
} = defineProps<Props>()

const emit = defineEmits<{
  click: [event: MouseEvent]
}>()

const handleClick = (event: MouseEvent) => {
  if (!disabled) {
    emit('click', event)
  }
}

// バリアントごとのスタイル
const variantClasses = computed(() => {
  const dark = darkMode

  const variants = {
    primary: dark
      ? [
          'bg-orange-500',
          'hover:bg-orange-600',
          'text-white',
          'border-transparent',
        ]
      : [
          'bg-blue-600',
          'hover:bg-blue-700',
          'text-white',
          'border-transparent',
        ],
    secondary: dark
      ? [
          'bg-gray-700',
          'hover:bg-gray-600',
          'text-white',
          'border-gray-600',
        ]
      : [
          'bg-gray-200',
          'hover:bg-gray-300',
          'text-gray-800',
          'border-gray-300',
        ],
    success: [
      'bg-green-600',
      'hover:bg-green-700',
      'text-white',
      'border-transparent',
    ],
    danger: [
      'bg-red-600',
      'hover:bg-red-700',
      'text-white',
      'border-transparent',
    ],
    outline: dark
      ? [
          'bg-transparent',
          'hover:bg-gray-700',
          'text-gray-200',
          'border-gray-600',
        ]
      : [
          'bg-white',
          'hover:bg-gray-50',
          'text-gray-800',
          'border-gray-300',
        ],
    ghost: dark
      ? [
          'bg-transparent',
          'hover:bg-gray-800',
          'text-gray-200',
          'border-transparent',
        ]
      : [
          'bg-transparent',
          'hover:bg-gray-100',
          'text-gray-700',
          'border-transparent',
        ],
  }

  return variants[variant]
})

// サイズごとのスタイル
const sizeClasses = computed(() => {
  const sizes = {
    sm: ['px-3', 'py-1.5', 'text-sm'],
    md: ['px-4', 'py-2', 'text-sm'],
    lg: ['px-6', 'py-3', 'text-base'],
  }

  return sizes[size]
})

// 全体のクラス
const buttonClasses = computed(() => [
  // 基本スタイル
  'inline-flex',
  'items-center',
  'justify-center',
  'gap-2',
  'rounded-lg',
  'font-medium',
  'border',
  'transition-all',
  'duration-200',
  // フォーカススタイル
  'focus:outline-none',
  'focus:ring-2',
  darkMode ? 'focus:ring-orange-400' : 'focus:ring-blue-400',
  'focus:ring-offset-1',
  // バリアント
  ...variantClasses.value,
  // サイズ
  ...sizeClasses.value,
  // 幅
  fullWidth ? 'w-full' : '',
  // 無効化
  disabled
    ? ['opacity-60', 'cursor-not-allowed', 'pointer-events-none']
    : 'cursor-pointer',
])
</script>

<template>
  <button :type="type" :disabled="disabled" :class="buttonClasses" @click="handleClick">
    <slot />
  </button>
</template>
