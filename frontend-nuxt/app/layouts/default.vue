<script lang="ts" setup>
// PC用: グループ化されたナビゲーション
const items = [
  {
    label: 'メイン',
    items: [
      { label: 'ホーム', to: '/', icon: 'HomeIcon' },
      { label: 'タスク一覧', to: '/todos', icon: 'ListBulletIcon' },
      { label: 'sprint一覧', to: '/sprints', icon: 'ListBulletIcon' },
      { label: '設定', to: '/settings', icon: 'Cog6ToothIcon' },
    ],
  },
]

// SP用: 主要な5つのタブ
const mainTabs = [
  { label: 'ホーム', to: '/', icon: 'HomeIcon' },
  { label: 'タスク', to: '/todos', icon: 'ListBulletIcon' },
  { label: 'Sprint', to: '/sprints', icon: 'ListBulletIcon' },
  { label: '設定', to: '/settings', icon: 'Cog6ToothIcon' },
]
</script>

<template>
  <Header />

  <main class="flex py-2 px-2 min-h-[90vh]">
    <!-- 左カラム: サイドナビ (PC: md以上で表示) -->
    <SideNavigation
      :nav-groups="items"
      class="hidden md:flex w-64 flex-shrink-0"
    />

    <!-- 右カラム: ページ内容 -->
    <div class="flex-1 md:ml-4 mb-16 md:mb-0">
      <slot></slot>
    </div>
  </main>

  <!-- 下部タブバー (SP: md未満で表示) -->
  <BottomTabBar
    :tabs="mainTabs"
    class="md:hidden"
  />

  <Footer />
</template>