<script setup lang="ts">
interface MenuItem {
  id: string
  label: string
  icon?: string
  action: () => void
}

interface Props {
  items: MenuItem[]
  darkMode?: boolean
}

const { items, darkMode = true } = defineProps<Props>()

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
</script>

<template>
  <div ref="dropdownRef" class="relative">
    <!-- トリガーボタン -->
    <button
      @click="toggleDropdown"
      class="p-1 rounded transition-colors"
      :class="[
        darkMode ? 'hover:bg-gray-800' : 'hover:bg-gray-200',
      ]"
    >
      <span :class="['text-lg', darkMode ? 'text-gray-400' : 'text-gray-600']">+</span>
    </button>

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
        class="absolute right-0 mt-1 w-48 z-50 rounded-lg shadow-lg overflow-hidden"
        :class="[
          darkMode
            ? ['bg-gray-800', 'border', 'border-gray-700']
            : ['bg-white', 'border', 'border-gray-200'],
        ]"
      >
        <div class="py-1">
          <button
            v-for="item in items"
            :key="item.id"
            @click="handleItemClick(item)"
            class="w-full flex items-center gap-3 px-4 py-2 text-left text-sm transition-colors"
            :class="[
              darkMode
                ? ['text-gray-200', 'hover:bg-gray-700']
                : ['text-gray-700', 'hover:bg-gray-100'],
            ]"
          >
            <Icon v-if="item.icon" :name="item.icon" class="w-4 h-4" />
            <span>{{ item.label }}</span>
          </button>
        </div>
      </div>
    </Transition>
  </div>
</template>