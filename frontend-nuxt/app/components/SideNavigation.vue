<script setup lang="ts">
interface NavItem {
  label: string
  icon?: any
  to?: string
}

interface NavGroup {
  label: string
  items: NavItem[]
}

defineProps<{
  navGroups: NavGroup[]
}>()

const config = useRuntimeConfig()
const appName = config.public.appName
</script>

<template>
  <nav class="w-64 bg-gray-50 h-screen p-4 border-r border-gray-200 flex flex-col">
    <!-- Logo -->
    <div class="flex items-center mb-6 px-2">
      <slot name="logo">
        <h1 class="text-xl font-bold">{{ appName }}</h1>
      </slot>
    </div>

    <!-- Navigation Groups -->
    <div class="flex-1 space-y-6 overflow-y-auto">
      <div v-for="group in navGroups" :key="group.label">
        <p class="text-xs text-gray-500 font-semibold mb-2 px-2 uppercase tracking-wide">
          {{ group.label }}
        </p>

        <ul class="space-y-1">
          <li
            v-for="item in group.items"
            :key="item.label"
          >
            <NuxtLink
              :to="item.to"
              class="w-full flex items-center gap-3 px-3 py-2 rounded-xl text-gray-700 hover:bg-gray-200 transition"
            >
              <component v-if="item.icon" :is="item.icon" class="w-5 h-5" />
              <span>{{ item.label }}</span>
            </NuxtLink>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<style scoped>
nav::-webkit-scrollbar {
  width: 6px;
}
nav::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 4px;
}
</style>
