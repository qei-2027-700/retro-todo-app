<script setup lang="ts">
import { useWindowSize } from '@vueuse/core';
// フロントエンド用のSprint型定義
interface Sprint {
  id: number
  name: string
  color: string
  created_at?: string
  updated_at?: string
}

interface Props {
  isOpen?: boolean
  darkMode?: boolean
}

const {
  isOpen = true,
  darkMode = true
} = defineProps<Props>()

const emit = defineEmits<{
  close: []
}>()

const route = useRoute()

// TailwindCSSのlgブレークポイント（1024px）に合わせた判定
const { width } = useWindowSize()
const isLargeScreen = computed(() => width.value >= 1024)

const projectExpanded = ref(true)

const { data: sprintsData } = await useFetch<Sprint[]>('/api/sprints')

const mainNavItems = computed(() => [
  { id: 'home', label: 'ホーム', icon: 'heroicons:home', to: '/dashboard', active: route.path === '/dashboard' },
  { id: 'tasks', label: 'マイタスク', icon: 'heroicons:check', to: '/tasks', active: route.path === '/tasks' },
  { id: 'inbox', label: '通知', icon: 'heroicons:inbox', to: '/inbox', active: route.path === '/inbox' },
])

const sprintMenuItems = [
  {
    id: 'new-sprint',
    label: '新しいスプリント',
    icon: 'heroicons:clipboard-document-list',
    action: () => console.log('新しいスプリントを作成'),
  },
  {
    id: 'import-sprint',
    label: 'スプリントをインポート',
    icon: 'heroicons:arrow-down-tray',
    action: () => console.log('スプリントをインポート'),
  },
]

const sprintItems = computed(() => {
  if (!sprintsData.value) return []

  return sprintsData.value.map(sprint => ({
    id: sprint.id.toString(),
    label: sprint.name,
    color: sprint.color,
    to: sprint.name === 'バックログ' ? '/sprints/backlog' : `/sprints/${sprint.id}`,
    active: route.path === (sprint.name === 'バックログ' ? '/sprints/backlog' : `/sprints/${sprint.id}`),
  }))
})

const handleNavClick = (id: string) => {
  console.log('ナビゲーション:', id)
  // SPの場合（lgブレークポイント未満）、ナビゲーション移動時にドロワーを閉じる
  if (!isLargeScreen.value) {
    emit('close')
  }
}
</script>

<template>
  <!-- オーバーレイ（モバイル用） -->
  <Transition
    enter-active-class="transition-opacity duration-200"
    enter-from-class="opacity-0"
    enter-to-class="opacity-100"
    leave-active-class="transition-opacity duration-200"
    leave-from-class="opacity-100"
    leave-to-class="opacity-0"
  >
    <div
      v-if="isOpen"
      class="fixed inset-0 bg-black bg-opacity-50 z-30 lg:hidden"
      @click="emit('close')"
    />
  </Transition>

  <!-- サイドバー -->
  <Transition
    enter-active-class="transition-transform duration-200"
    enter-from-class="-translate-x-full"
    enter-to-class="translate-x-0"
    leave-active-class="transition-transform duration-200"
    leave-from-class="translate-x-0"
    leave-to-class="-translate-x-full"
  >
    <aside
      v-if="isOpen"
      class="fixed left-0 top-14 w-64 z-30 overflow-y-auto h-[calc(100vh-3.5rem)] border-r lg:translate-x-0"
      :class="[
        darkMode ? 'bg-gray-900' : 'bg-gray-50',
        darkMode ? 'border-gray-700' : 'border-gray-200',
      ]"
    >
      <div class="flex flex-col h-full">
        <!-- ナビゲーションコンテンツ -->
        <nav class="flex-1 px-3 py-4 space-y-1">
          <!-- メインナビゲーション -->
          <div class="space-y-1 mb-6">
            <NuxtLink
              v-for="item in mainNavItems"
              :key="item.id"
              :to="item.to"
              @click="handleNavClick(item.id)"
              class="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-sm transition-colors"
              :class="[
                item.active
                  ? darkMode
                    ? ['bg-gray-700', 'text-white']
                    : ['bg-gray-200', 'text-gray-900']
                  : darkMode
                    ? ['text-gray-300', 'hover:bg-gray-800']
                    : ['text-gray-700', 'hover:bg-gray-100'],
              ]"
            >
              <Icon :name="item.icon" class="w-5 h-5" />
              <span>{{ item.label }}</span>
            </NuxtLink>
          </div>

          <!-- スプリント -->
          <div class="mb-6">
            <div class="flex items-center justify-between px-3 py-2">
              <button
                @click="projectExpanded = !projectExpanded"
                class="flex items-center gap-2 text-sm font-semibold"
                :class="[
                  darkMode ? 'text-gray-400' : 'text-gray-600',
                ]"
              >
                <svg
                  class="w-4 h-4 transition-transform"
                  :class="[
                    projectExpanded ? 'rotate-90' : '',
                  ]"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M9 5l7 7-7 7"
                  />
                </svg>
                <span>スプリント</span>
              </button>
              <SimpleDropdown
                :items="sprintMenuItems"
                :dark-mode="darkMode"
              />
            </div>
            <div>
              <NuxtLink
                :to="'/sprints'"
                @click="handleNavClick('sprints')"
                class="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-sm transition-colors"
                :class="[
                  darkMode
                    ? ['text-gray-300', 'hover:bg-gray-800']
                    : ['text-gray-700', 'hover:bg-gray-100'],
                ]"
              >
                <div :class="['w-3', 'h-3', 'rounded', 'bg-gray-500']" />
                <span>sprint一覧</span>
              </NuxtLink>
            </div>
            <div v-if="projectExpanded" class="space-y-1 mt-1">
              <NuxtLink
                v-for="item in sprintItems"
                :key="item.id"
                :to="item.to"
                @click="handleNavClick(item.id)"
                class="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-sm transition-colors"
                :class="[
                  darkMode
                    ? ['text-gray-300', 'hover:bg-gray-800']
                    : ['text-gray-700', 'hover:bg-gray-100'],
                ]"
              >
                <div :class="['w-3', 'h-3', 'rounded', item.color]" />
                <span>{{ item.label }}</span>
              </NuxtLink>
            </div>
          </div>

          <div class="mb-6">
            <NuxtLink
              to="/sample"
              @click="handleNavClick('sample')"
              class="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-sm transition-colors"
              :class="[
                darkMode
                  ? ['text-gray-300', 'hover:bg-gray-800']
                  : ['text-gray-700', 'hover:bg-gray-100'],
              ]"
            >
              <Icon name="heroicons:archive" class="w-5 h-5" />
              <span>サンプル</span>
            </NuxtLink>
          </div>
        </nav>
      </div>
    </aside>
  </Transition>
</template>

<style scoped>
aside::-webkit-scrollbar {
  width: 6px;
}
aside::-webkit-scrollbar-thumb {
  background: #4b5563;
  border-radius: 4px;
}
aside::-webkit-scrollbar-thumb:hover {
  background: #6b7280;
}
</style>
