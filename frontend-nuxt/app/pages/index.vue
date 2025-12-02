<script setup lang="ts">
const isModalOpen = ref(false)

const handleClick = () => {
  isModalOpen.value = true
}

const handleClose = () => {
  isModalOpen.value = false
}
</script>

<template>
  <h1 class="text-2xl font-bold mb-4">テストページ</h1>
  <p>Welcome to your dashboard!</p>

  <BaseButton @click="handleClick" class="mt-4">
    slotの内容
  </BaseButton>

  <BaseModal v-if="isModalOpen" @close="handleClose">
    slotのコンテンツ
  </BaseModal>
</template>
