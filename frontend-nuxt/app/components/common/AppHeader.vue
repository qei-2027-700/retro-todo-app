<script setup lang="ts">
interface Props {
  darkMode?: boolean
}

const { darkMode = true } = defineProps<Props>()

const emit = defineEmits<{
  toggleSidebar: []
}>()

const userMenuOpen = ref(false)
const userMenuRef = ref<HTMLElement | null>(null)

const isOpenSearchMenu = ref(false)
const searchQuery = ref('')
const searchInputRef = ref<HTMLInputElement | null>(null)

const createMenuItems = [
  {
    id: 'task',
    label: 'タスク',
    icon: '✓',
    action: () => console.log('タスクを作成'),
  },
  {
    id: 'project',
    label: 'プロジェクト',
    icon: '📋',
    action: () => console.log('プロジェクトを作成'),
  },
  {
    id: 'message',
    label: 'メッセージ',
    icon: '💬',
    action: () => console.log('メッセージを作成'),
  },
]

const toggleUserMenu = () => {
  userMenuOpen.value = !userMenuOpen.value
}

const handleSearch = () => {
  console.log('検索を実行')
  isOpenSearchMenu.value = true
}

const handleHelpClick = () => {
  console.log('ヘルプを表示')
}

const handleLogout = () => {
  console.log('ログアウト')
  userMenuOpen.value = false
}

const handleProfile = () => {
  console.log('プロフィール表示')
  userMenuOpen.value = false
}

const handleSettings = () => {
  console.log('設定表示')
  userMenuOpen.value = false
}

// 外側クリックで閉じる
const handleClickOutside = (event: MouseEvent) => {
  if (userMenuRef.value && !userMenuRef.value.contains(event.target as Node)) {
    userMenuOpen.value = false
  }
}

// ESCキーで検索モーダルを閉じる
const handleKeydown = (event: KeyboardEvent) => {
  if (event.key === 'Escape' && isOpenSearchMenu.value) {
    isOpenSearchMenu.value = false
    searchQuery.value = ''
  }
}

// 検索モーダルが閉じたときに検索クエリをクリア
watch(isOpenSearchMenu, (isOpen) => {
  if (!isOpen) {
    searchQuery.value = ''
  }
})

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
  document.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
  document.removeEventListener('keydown', handleKeydown)
})
</script>

<template>
  <header
    class="fixed top-0 left-0 right-0 z-50 border-b"
    :class="[
      // 背景
      darkMode ? 'bg-gray-800' : 'bg-white',
      // ボーダー
      darkMode ? 'border-gray-700' : 'border-gray-200',
    ]"
  >
    <div class="flex items-center justify-between px-4 py-3 md:px-6">
      <!-- 左側: ハンバーガーメニュー + 作成ボタン -->
      <div class="flex items-center gap-3">
        <!-- ハンバーガーメニュー -->
        <button
          @click="() => { console.log('ハンバーガーメニュークリック'); emit('toggleSidebar'); }"
          class="flex items-center justify-center w-8 h-8"
          :class="[
            darkMode
              ? ['text-gray-300', 'hover:bg-gray-700']
              : ['text-gray-600', 'hover:bg-gray-100'],
          ]"
          aria-label="メニュー"
        >
          <svg
            class="w-6 h-6"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4 6h16M4 12h16M4 18h16"
            />
          </svg>
        </button>

        <!-- 作成ボタン (PC版: ボタン + ラベル, SP版: ボタンのみ) -->
        <div class="hidden sm:block">
          <CreateDropdown :items="createMenuItems" :dark-mode="darkMode" />
        </div>
        <div class="block sm:hidden">
          <CreateDropdown
            :items="createMenuItems"
            button-label=""
            :dark-mode="darkMode"
          />
        </div>
      </div>

      <!-- 中央: 検索バー -->
      <div class="flex-1 max-w-2xl mx-4">
        <div class="relative">
          <div
            class="flex items-center w-full px-4 py-2 rounded-lg transition-colors"
            :class="[
              darkMode
                ? ['bg-gray-700', 'hover:bg-gray-600']
                : ['bg-gray-100', 'hover:bg-gray-200'],
            ]"
          >
            <!-- 検索アイコン -->
            <svg
              class="'w-5 h-5 mr-3"
              :class="[darkMode ? 'text-gray-400' : 'text-gray-500']"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
              />
            </svg>

            <!-- 検索入力 -->
            <input
              ref="searchInputRef"
              v-model="searchQuery"
              type="text"
              placeholder="検索"
              @focus="handleSearch"
              class="flex-1 bg-transparent outline-none text-sm"
              :class="[
                darkMode ? 'text-gray-200 placeholder-gray-500' : 'text-gray-800 placeholder-gray-400',
              ]"
            />

            <!-- キーボードショートカット (PC版のみ) -->
            <div class="hidden md:flex items-center gap-1 ml-3">
              <kbd
                class="px-2 py-1 text-xs font-semibold rounded border"
                :class="[
                  darkMode
                    ? ['bg-gray-600', 'text-gray-300', 'border-gray-500']
                    : ['bg-white', 'text-gray-600', 'border-gray-300'],
                ]"
              >
                ⌘
              </kbd>
              <kbd
                class="px-2 py-1 text-xs font-semibold rounded border"
                :class="[
                  darkMode
                    ? ['bg-gray-600', 'text-gray-300', 'border-gray-500']
                    : ['bg-white', 'text-gray-600', 'border-gray-300'],
                ]"
              >
                K
              </kbd>
            </div>
          </div>
        </div>
      </div>

      <!-- 右側: ヘルプ + ユーザーメニュー -->
      <div class="flex items-center gap-2">
        <!-- ヘルプボタン -->
        <button
          @click="handleHelpClick"
          class="flex items-center justify-center w-8 h-8 rounded-full transition-colors"
          :class="[
            darkMode
              ? ['text-gray-300', 'hover:bg-gray-700']
              : ['text-gray-600', 'hover:bg-gray-100'],
          ]"
          aria-label="ヘルプ"
        >
          <span class="text-lg">?</span>
        </button>

        <!-- ユーザーメニュー -->
        <div ref="userMenuRef" class="relative">
          <button
            @click="toggleUserMenu"
            class="flex items-center gap-1 rounded-full transition-colors"
            :class="[
              darkMode ? 'hover:bg-gray-700' : 'hover:bg-gray-100',
              'p-1',
            ]"
            aria-label="ユーザーメニュー"
          >
            <!-- アバター -->
            <div
              class="w-8 h-8 rounded-full bg-gradient-to-br from-purple-400 to-pink-600 flex items-center justify-center text-white text-sm font-semibold"
            >
              <img
                src="https://i.pravatar.cc/150?img=3"
                alt="User"
                class="w-full h-full rounded-full object-cover"
              />
            </div>

            <!-- ドロップダウンアイコン -->
            <svg
              class="w-4 h-4 transition-transform"
              :class="[
                userMenuOpen ? 'rotate-180' : '',
                darkMode ? 'text-gray-400' : 'text-gray-600',
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

          <!-- ユーザードロップダウンメニュー -->
          <Transition
            enter-active-class="transition ease-out duration-100"
            enter-from-class="transform opacity-0 scale-95"
            enter-to-class="transform opacity-100 scale-100"
            leave-active-class="transition ease-in duration-75"
            leave-from-class="transform opacity-100 scale-100"
            leave-to-class="transform opacity-0 scale-95"
          >
            <div
              v-if="userMenuOpen"
              class="absolute right-0 mt-2 w-56 rounded-lg shadow-lg overflow-hidden border"
              :class="[
                darkMode
                  ? ['bg-gray-800', 'border-gray-700']
                  : ['bg-white', 'border-gray-200'],
              ]"
            >
              <div class="py-1">
                <button
                  @click="handleProfile"
                  class="w-full flex items-center gap-3 px-4 py-2.5 text-left text-sm transition-colors"
                  :class="[
                    darkMode
                      ? ['text-gray-200', 'hover:bg-gray-700']
                      : ['text-gray-700', 'hover:bg-gray-100'],
                  ]"
                >
                  <span class="text-lg">👤</span>
                  <span>プロフィール</span>
                </button>
                <button
                  @click="handleSettings"
                  class="w-full flex items-center gap-3 px-4 py-2.5 text-left text-sm transition-colors"
                  :class="[
                    darkMode
                      ? ['text-gray-200', 'hover:bg-gray-700']
                      : ['text-gray-700', 'hover:bg-gray-100'],
                  ]"
                >
                  <span class="text-lg">⚙️</span>
                  <span>設定</span>
                </button>
                <hr :class="['my-1', darkMode ? 'border-gray-700' : 'border-gray-200']" />
                <button
                  @click="handleLogout"
                  class="w-full flex items-center gap-3 px-4 py-2.5 text-left text-sm transition-colors"
                  :class="[
                    darkMode
                      ? ['text-gray-200', 'hover:bg-gray-700']
                      : ['text-gray-700', 'hover:bg-gray-100'],
                  ]"
                >
                  <span class="text-lg">🚪</span>
                  <span>ログアウト</span>
                </button>
              </div>
            </div>
          </Transition>
        </div>
      </div>
    </div>
  </header>

  <!-- 検索モーダル -->
  <SearchModal
    v-model="isOpenSearchMenu"
    :dark-mode="darkMode"
    :search-query="searchQuery"
  />
</template>
