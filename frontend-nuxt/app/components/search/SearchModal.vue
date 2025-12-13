<script setup lang="ts">
interface Props {
  modelValue: boolean
  darkMode?: boolean
  searchQuery?: string
}

interface SearchTab {
  id: string
  label: string
  icon: string
}

interface SearchResultItem {
  id: string
  title: string
  subtitle?: string
  icon: string
  type: 'task' | 'project' | 'member' | 'message' | 'saved-search'
  completed?: boolean
}

const props = defineProps<Props>()
const emit = defineEmits<{
  'update:modelValue': [value: boolean]
}>()

const activeTab = ref('all')

// タブ定義
const tabs: SearchTab[] = [
  { id: 'all', label: 'すべて', icon: 'heroicons:magnifying-glass' },
  { id: 'task', label: 'タスク', icon: 'heroicons:check' },
  { id: 'project', label: 'スプリント', icon: 'heroicons:clipboard-document-list' },
  { id: 'member', label: 'メンバー', icon: 'heroicons:user' },
  { id: 'message', label: 'メッセージ', icon: 'heroicons:chat-bubble-left-right' },
]

// APIから検索結果を取得
// const { data: recentResults } = await useFetch('/api/tasks/search/recent')
// const { data: savedSearches } = await useFetch('/api/tasks/search/saved')

const closeModal = () => {
  emit('update:modelValue', false)
}

const handleResultClick = (item: SearchResultItem) => {
  console.log('結果をクリック:', item)
  closeModal()
}

</script>

<template>
  <!-- オーバーレイ背景 -->
  <Transition
    enter-active-class="transition-opacity duration-200"
    enter-from-class="opacity-0"
    enter-to-class="opacity-100"
    leave-active-class="transition-opacity duration-200"
    leave-from-class="opacity-100"
    leave-to-class="opacity-0"
  >
    <div
      v-if="modelValue"
      class="fixed inset-0 z-40 bg-black/50 backdrop-blur-sm"
      @click="closeModal"
    />
  </Transition>

  <!-- モーダルコンテンツ -->
  <Transition
    enter-active-class="transition-all duration-200 ease-out"
    enter-from-class="opacity-0 scale-95 translate-y-4"
    enter-to-class="opacity-100 scale-100 translate-y-0"
    leave-active-class="transition-all duration-150 ease-in"
    leave-from-class="opacity-100 scale-100 translate-y-0"
    leave-to-class="opacity-0 scale-95 translate-y-4"
  >
    <div
      v-if="modelValue"
      class="fixed top-16 left-1/2 -translate-x-1/2 z-40 w-full max-w-3xl mx-auto px-4"
      @click.stop
    >
      <div
        class="rounded-lg shadow-2xl overflow-hidden"
        :class="[
          darkMode ? 'bg-gray-800' : 'bg-white',
        ]"
      >
        <!-- タブ -->
        <div
          class="px-6 py-4 border-b"
          :class="[darkMode ? 'border-gray-700' : 'border-gray-200']"
        >
          <div class="flex gap-2 overflow-x-auto">
            <button
              v-for="tab in tabs"
              :key="tab.id"
              @click="activeTab = tab.id"
              class="flex items-center gap-2 px-4 py-2 rounded-lg text-sm font-medium whitespace-nowrap transition-colors"
              :class="[
                activeTab === tab.id
                  ? darkMode
                    ? 'bg-gray-700 text-white'
                    : 'bg-gray-200 text-gray-900'
                  : darkMode
                    ? 'text-gray-400 hover:bg-gray-700/50 hover:text-gray-300'
                    : 'text-gray-600 hover:bg-gray-100 hover:text-gray-900',
              ]"
            >
              <Icon :name="tab.icon" class="w-4 h-4" />
              <span>{{ tab.label }}</span>
            </button>
          </div>
        </div>

        <div
          class="max-h-[60vh] overflow-y-auto"
          :class="[darkMode ? 'bg-gray-800' : 'bg-white']"
        >
          <div class="px-6 py-4">
            <h3
              class="text-xs font-semibold uppercase tracking-wider mb-3"
              :class="[darkMode ? 'text-gray-500' : 'text-gray-500']"
            >
              最近
            </h3>
            <div class="space-y-1">
              <button
                v-for="item in recentResults"
                :key="item.id"
                @click="handleResultClick(item)"
                class="w-full flex items-center gap-3 px-3 py-2.5 rounded-lg text-left transition-colors"
                :class="[
                  darkMode
                    ? 'hover:bg-gray-700 text-gray-200'
                    : 'hover:bg-gray-100 text-gray-800',
                ]"
              >
                <div
                  class="flex-shrink-0 w-5 h-5 rounded flex items-center justify-center text-xs"
                  :class="[
                    item.completed
                      ? darkMode
                        ? 'bg-green-600 text-white'
                        : 'bg-green-500 text-white'
                      : darkMode
                        ? 'border border-gray-600'
                        : 'border border-gray-300',
                  ]"
                >
                  <span v-if="item.completed">✓</span>
                </div>

                <div class="flex-1 min-w-0">
                  <div
                    class="text-sm font-medium truncate"
                    :class="[
                      item.completed && (darkMode ? 'text-gray-500' : 'text-gray-500'),
                    ]"
                  >
                    {{ item.title }}
                  </div>
                  <div
                    v-if="item.subtitle"
                    class="text-xs flex items-center gap-1.5 mt-0.5"
                    :class="[darkMode ? 'text-gray-500' : 'text-gray-500']"
                  >
                    <span class="w-2 h-2 rounded-full bg-purple-500" />
                    <span>{{ item.subtitle }}</span>
                  </div>
                </div>

                <img
                  src="https://i.pravatar.cc/150?img=3"
                  alt="User"
                  class="w-8 h-8 rounded-full flex-shrink-0"
                />
              </button>
            </div>
          </div>

          <!-- <div class="px-6 py-4 border-t" :class="[darkMode ? 'border-gray-700' : 'border-gray-200']">
            <h3
              class="text-xs font-semibold uppercase tracking-wider mb-3"
              :class="[darkMode ? 'text-gray-500' : 'text-gray-500']"
            >
              保存された検索
            </h3>
            <div class="space-y-1">
              <button
                v-for="item in savedSearches"
                :key="item.id"
                @click="handleResultClick(item)"
                class="w-full flex items-center gap-3 px-3 py-2.5 rounded-lg text-left transition-colors"
                :class="[
                  darkMode
                    ? 'hover:bg-gray-700 text-gray-200'
                    : 'hover:bg-gray-100 text-gray-800',
                ]"
              >
                <div
                  class="flex-shrink-0 w-5 h-5 rounded flex items-center justify-center"
                  :class="[darkMode ? 'text-gray-400' : 'text-gray-500']"
                >
                  <Icon :name="item.icon" class="w-4 h-4" />
                </div>

                <div class="flex-1 text-sm font-medium">
                  {{ item.title }}
                </div>

                <svg
                  class="w-4 h-4 flex-shrink-0"
                  :class="[darkMode ? 'text-gray-600' : 'text-gray-400']"
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
              </button>
            </div>
          </div> -->
        </div>
      </div>
    </div>
  </Transition>
</template>
