<script setup lang="ts">
const appName = useRuntimeConfig().public.appName
</script>

<template>
  <footer class="py-4 px-4 text-center bg-slate-900 text-slate-50">
    <h1>{{ appName }}</h1>
    <p>&copy;2026~</p>
  </footer>
</template>
