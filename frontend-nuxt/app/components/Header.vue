<script setup lang="ts">
const appName = useRuntimeConfig().public.appName
</script>

<template>
  <header class="py-4 px-4 bg-gray-800 text-white flex items-center justify-between">
    <h1
      @click="navigateTo('/')"
      class="cursor-pointer"
    >{{ appName }}</h1>
    <p>aaa</p>
  </header>
</template>
