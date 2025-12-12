<script setup lang="ts">
const sidebarOpen = ref(true) // サイドバーの表示状態
const mainContentShift = ref(true) // メインコンテンツのシフト状態（遅延させる）
const darkMode = ref(true) // ダークモード

const toggleSidebar = () => {
  console.log('toggleSidebar called, current state:', sidebarOpen.value)

  if (sidebarOpen.value) {
    // 閉じる場合：先にサイドバーを閉じて、アニメーション完了後にメインコンテンツをシフト
    sidebarOpen.value = false
    setTimeout(() => {
      mainContentShift.value = false
    }, 200) // サイドバーのアニメーション時間と同じ
  }
  else {
    // 開く場合：同時に開く
    sidebarOpen.value = true
    mainContentShift.value = true
  }

  console.log('new state:', sidebarOpen.value)
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
      :class="[
        'transition-all',
        'duration-200',
        // PC版ではサイドバーの幅分マージン（mainContentShiftで制御）
        mainContentShift ? 'lg:ml-64' : 'lg:ml-0',
        // ヘッダーの高さ分のパディング (fixed headerのため)
        'pt-16',
      ]"
    >
      <slot />
    </main>
  </div>
</template>
