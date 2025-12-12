<script setup lang="ts">
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

const insightExpanded = ref(true)
const favoriteExpanded = ref(true)
const projectExpanded = ref(true)

const mainNavItems = [
  { id: 'home', label: 'ホーム', icon: '🏠', to: '/', active: false },
  { id: 'tasks', label: 'マイタスク', icon: '✓', to: '/tasks', active: true },
  { id: 'inbox', label: '受信トレイ', icon: '🔔', to: '/inbox', active: false },
]

const insightItems = [
  { id: 'report', label: 'レポート', icon: '📊', to: '/report', active: false },
  { id: 'portfolio', label: 'ポートフォリオ', icon: '📁', to: '/portfolio', active: false },
  { id: 'goal', label: 'ゴール', icon: '🎯', to: '/goal', active: false },
]

const favoriteItems = [
  { id: 'personal-sprint', label: 'Personal Sprint', color: 'bg-purple-500', to: '/projects/personal-sprint', active: false },
]

const projectItems = [
  { id: 'personal-sprint-2', label: 'Personal Sprint', color: 'bg-purple-500', to: '/projects/personal-sprint', active: false },
  { id: '2510-3', label: '2510-3', color: 'bg-purple-500', to: '/projects/2510-3', active: false },
  { id: 'backlog', label: 'バックログ', color: 'bg-blue-500', to: '/projects/backlog', active: false },
  { id: '2510-4', label: '2510-4', color: 'bg-purple-500', to: '/projects/2510-4', active: false },
]

const teamItems = [
  { id: 'it', label: 'IT', icon: '👥', to: '/teams/it', hasChildren: true },
]

const handleNavClick = (id: string) => {
  console.log('ナビゲーション:', id)
}

const handleAddInsight = () => {
  console.log('インサイトを追加')
}

const handleAddProject = () => {
  console.log('プロジェクトを追加')
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
              <span class="text-lg">{{ item.icon }}</span>
              <span>{{ item.label }}</span>
            </NuxtLink>
          </div>

          <!-- インサイト -->
          <div class="mb-6">
            <div class="flex items-center justify-between px-3 py-2">
              <button
                @click="insightExpanded = !insightExpanded"
                class="flex items-center gap-2 text-sm font-semibold"
                :class="[
                  darkMode ? 'text-gray-400' : 'text-gray-600',
                ]"
              >
                <svg
                  class="w-4 h-4 transition-transform"
                  :class="[
                    insightExpanded ? 'rotate-90' : '',
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
                <span>インサイト</span>
              </button>
              <button
                @click="handleAddInsight"
                class="p-1 rounded transition-colors"
                :class="[
                  darkMode ? 'hover:bg-gray-800' : 'hover:bg-gray-200',
                ]"
              >
                <span :class="['text-lg', darkMode ? 'text-gray-400' : 'text-gray-600']">+</span>
              </button>
            </div>
            <div v-if="insightExpanded" class="space-y-1 mt-1">
              <NuxtLink
                v-for="item in insightItems"
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
                <span class="text-lg">{{ item.icon }}</span>
                <span>{{ item.label }}</span>
              </NuxtLink>
            </div>
          </div>

          <!-- お気に入り -->
          <div class="mb-6">
            <button
              @click="favoriteExpanded = !favoriteExpanded"
              class="flex items-center gap-2 px-3 py-2 text-sm font-semibold"
              :class="[
                darkMode ? 'text-gray-400' : 'text-gray-600',
              ]"
            >
              <svg
                class="w-4 n-4 transition-transform"
                :class="[
                  favoriteExpanded ? 'rotate-90' : '',
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
              <span>お気に入り</span>
            </button>
            <div v-if="favoriteExpanded" class="space-y-1 mt-1">
              <NuxtLink
                v-for="item in favoriteItems"
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

          <!-- プロジェクト -->
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
                <span>プロジェクト</span>
              </button>
              <button
                @click="handleAddProject"
                class="p-1 rounded transition-colors"
                :class="[
                  darkMode ? 'hover:bg-gray-800' : 'hover:bg-gray-200',
                ]"
              >
                <span :class="['text-lg', darkMode ? 'text-gray-400' : 'text-gray-600']">+</span>
              </button>
            </div>
            <div v-if="projectExpanded" class="space-y-1 mt-1">
              <NuxtLink
                v-for="item in projectItems"
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
