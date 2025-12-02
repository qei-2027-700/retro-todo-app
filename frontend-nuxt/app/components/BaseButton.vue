<script setup lang="ts">
const props = defineProps<{
  type?: 'button' | 'submit' | 'reset'
}>()

const type = computed(() => props.type ?? 'button')
</script>

<template>
  <button
    :type="type"
    class="inline-flex items-center justify-center px-3 py-1.5 rounded border border-gray-300
           bg-white text-sm text-gray-800 hover:bg-gray-50
           focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-1
           disabled:opacity-60 disabled:cursor-not-allowed"
  >
    <slot />
  </button>
</template>
