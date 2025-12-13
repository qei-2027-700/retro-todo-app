<script setup lang="ts">
interface MenuItem {
  id: string
  label: string
  icon?: string
  badge?: string
  shortcut?: string
  action: () => void
}

interface Props {
  items: MenuItem[]
  position?: 'left' | 'right'
  darkMode?: boolean
}

const { items, position = 'left', darkMode = true } = defineProps<Props>()

const isOpen = ref(false)
const dropdownRef = ref<HTMLElement | null>(null)

const toggleDropdown = () => {
  isOpen.value = !isOpen.value
}

const closeDropdown = () => {
  isOpen.value = false
}

const handleItemClick = (item: MenuItem) => {
  item.action()
  closeDropdown()
}

// 外側クリックで閉じる
const handleClickOutside = (event: MouseEvent) => {
  if (dropdownRef.value && !dropdownRef.value.contains(event.target as Node)) {
    closeDropdown()
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})

// Expose methods for parent component
defineExpose({
  toggleDropdown,
  closeDropdown,
})
</script>

<template>
  <div ref="dropdownRef" class="relative inline-block">
    <!-- Trigger slot -->
    <div @click="toggleDropdown">
      <slot name="trigger" :is-open="isOpen" />
    </div>

    <!-- ドロップダウンメニュー -->
    <Transition
      enter-active-class="transition ease-out duration-100"
      enter-from-class="transform opacity-0 scale-95"
      enter-to-class="transform opacity-100 scale-100"
      leave-active-class="transition ease-in duration-75"
      leave-from-class="transform opacity-100 scale-100"
      leave-to-class="transform opacity-0 scale-95"
    >
      <div
        v-if="isOpen"
        :class="[
          // 位置とサイズ
          'absolute',
          'mt-2',
          'min-w-[240px]',
          'z-50',
          // 左右の配置
          position === 'right' ? 'right-0' : 'left-0',
          // 背景とボーダー
          darkMode
            ? ['bg-gray-800', 'border', 'border-gray-700']
            : ['bg-white', 'border', 'border-gray-200'],
          // 角丸と影
          'rounded-lg',
          'shadow-lg',
          'overflow-hidden',
        ]"
      >
        <div class="py-1">
          <button
            v-for="item in items"
            :key="item.id"
            @click="handleItemClick(item)"
            :class="[
              'w-full',
              'flex',
              'items-center',
              'justify-between',
              'gap-3',
              'px-4',
              'py-2.5',
              'text-left',
              'text-sm',
              'transition-colors',
              darkMode
                ? ['text-gray-200', 'hover:bg-gray-700']
                : ['text-gray-700', 'hover:bg-gray-100'],
            ]"
          >
            <div class="flex items-center gap-3 flex-1">
              <!-- アイコン -->
              <Icon v-if="item.icon" :name="item.icon" class="w-5 h-5" />
              <!-- ラベル -->
              <span>{{ item.label }}</span>
              <!-- バッジ -->
              <span
                v-if="item.badge"
                :class="[
                  'px-2',
                  'py-0.5',
                  'text-xs',
                  'rounded',
                  darkMode
                    ? ['bg-gray-700', 'text-gray-300']
                    : ['bg-gray-200', 'text-gray-700'],
                ]"
              >
                {{ item.badge }}
              </span>
            </div>

            <!-- ショートカット -->
            <div v-if="item.shortcut" class="flex items-center gap-1">
              <kbd
                v-for="(key, index) in item.shortcut.split(' | ')"
                :key="index"
                :class="[
                  'px-2',
                  'py-1',
                  'text-xs',
                  'font-semibold',
                  'rounded',
                  'border',
                  darkMode
                    ? ['bg-gray-700', 'text-gray-300', 'border-gray-600']
                    : ['bg-white', 'text-gray-600', 'border-gray-300'],
                ]"
              >
                {{ key }}
              </kbd>
            </div>
          </button>
        </div>
      </div>
    </Transition>
  </div>
</template>
