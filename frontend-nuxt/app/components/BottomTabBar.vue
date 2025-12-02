<script setup lang="ts">
interface TabItem {
  label: string
  icon?: any
  to?: string
}
defineProps<{
  tabs: TabItem[]
}>()
</script>

<template>
  <nav class="fixed bottom-0 left-0 right-0 bg-white border-t border-gray-200 z-50">
    <ul class="flex justify-around items-center h-16">
      <li
        v-for="tab in tabs"
        :key="tab.label"
        class="flex-1"
      >
        <NuxtLink
          :to="tab.to"
          class="flex flex-col items-center justify-center h-full gap-1 text-gray-600 hover:text-gray-900 transition"
          active-class="text-blue-600"
        >
          <component v-if="tab.icon" :is="tab.icon" class="w-6 h-6" />
          <span class="text-xs">{{ tab.label }}</span>
        </NuxtLink>
      </li>
    </ul>
  </nav>
</template>

<style scoped>
/* アクティブ状態のスタイル調整 */
a.router-link-active {
  @apply text-blue-600 font-semibold;
}
</style>
