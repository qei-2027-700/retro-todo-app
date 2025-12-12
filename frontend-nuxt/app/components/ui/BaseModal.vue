<script setup lang="ts">
const emit = defineEmits<{
  (e: 'close'): void
}>()

const handleOverlayClick = () => {
  emit('close')
}
</script>

<template>
  <!-- 画面全体の半透明オーバーレイ -->
  <div
    class="fixed inset-0 z-50 flex items-center justify-center bg-black/50"
    @click="handleOverlayClick"
  >
    <!-- モーダル本体。クリックが外側に伝播しないようにする -->
    <div
      class="bg-white dark:bg-neutral-100 rounded-lg shadow-xl max-w-lg w-full p-4 relative"
      @click.stop
    >
      <button
        type="button"
        class="absolute right-3 top-3 text-gray-400 hover:text-gray-600"
        @click="emit('close')"
      >
        ×
      </button>

      <slot />
    </div>
  </div>
</template>
