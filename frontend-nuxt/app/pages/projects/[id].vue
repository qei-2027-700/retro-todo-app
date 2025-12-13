<script setup lang="ts">
// TODO: 将来、projects APIが実装されたら有効化
// import { extractIdFromParam } from '~/utils/slug'
// const route = useRoute()
// const projectId = extractIdFromParam(route.params.id)

// 現在は未実装のため、404ページにリダイレクト
const router = useRouter()
router.push('/404')
</script>

<template>
  <div />
</template>
