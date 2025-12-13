<script setup lang="ts">
const sidebarOpen = ref(true)
const mainContentShift = ref(true) // メインコンテンツのシフト状態（遅延させる）
const darkMode = ref(true)

const toggleSidebar = () => {
  if (sidebarOpen.value) {
    // 閉じる場合：先にサイドバーを閉じて、アニメーション完了後にメインコンテンツをシフト
    sidebarOpen.value = false
    setTimeout(() => {
      mainContentShift.value = false
    }, 180)
  }
  else {
    mainContentShift.value = true

    setTimeout(() => {
      sidebarOpen.value = true
    }, 100)
  }
}

const closeSidebar = () => {
  // モバイルでのみサイドバーを閉じる
  if (window.innerWidth < 1024) {
    sidebarOpen.value = false

    setTimeout(() => {
      mainContentShift.value = false
    }, 200)
  }
}

// レスポンシブ対応: PC版では自動的にサイドバーを開く
onMounted(() => {
  const handleResize = () => {
    if (window.innerWidth >= 1024) {
      sidebarOpen.value = true
      mainContentShift.value = true
    }
  }

  window.addEventListener('resize', handleResize)
  handleResize()

  onUnmounted(() => {
    window.removeEventListener('resize', handleResize)
  })
})
</script>

<template>
  <div :class="['min-h-screen', darkMode ? 'bg-gray-900' : 'bg-gray-50']">
    <AppHeader :dark-mode="darkMode" @toggle-sidebar="toggleSidebar" />

    <SideNavigation :is-open="sidebarOpen" :dark-mode="darkMode" @close="closeSidebar" />

    <main
      class="transition-all duration-200 pt-16"
      :class="[
        // PC版ではサイドバーの幅分マージン（mainContentShiftで制御）
        mainContentShift ? 'lg:ml-64' : 'lg:ml-0',
        'pt-16',
      ]"
    >
      <slot />
    </main>
  </div>
</template>
