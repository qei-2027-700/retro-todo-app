<script setup lang="ts">
// このページは存在しないすべてのルートをキャッチします
// デフォルトでdefaultレイアウトが使用されます
</script>

<template>
  <div class="flex items-center justify-center min-h-[calc(100vh-64px)] p-8">
    <div class="text-center max-w-md">
      <h1 class="text-6xl font-bold text-gray-800 dark:text-gray-100 mb-4">
        404
      </h1>
      <h2 class="text-2xl font-semibold text-gray-700 dark:text-gray-200 mb-4">
        ページが見つかりません
      </h2>
      <p class="text-gray-600 dark:text-gray-400 mb-8">
        お探しのページは存在しないか、移動された可能性があります。
      </p>
      <NuxtLink
        to="/"
        class="inline-flex items-center px-6 py-3 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-lg transition-colors"
      >
        ホームに戻る
      </NuxtLink>
    </div>
  </div>
</template>