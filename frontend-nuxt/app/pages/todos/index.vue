<script lang="ts" setup>
import type { components } from "../../../types/api.d.ts"

type Todo = components["schemas"]["model.Todo"]

const todos = ref<Todo[]>([])

const fetchTodos = async () => {
  try {
    const data = await $fetch<Todo[]>('/api/todos')
    console.log(data)
    todos.value = data
  }
  catch (error) {
    console.error('Error fetching TODOs:', error)
  }
}
await fetchTodos()
</script>

<template>
  <h1 class="text-2xl font-bold mb-4">TODO</h1>
  <p>TODO一覧です</p>
  <ul class="mt-4 space-y-2">
    <li
      v-for="todo in todos"
      :key="todo.id"
      class="p-4 border rounded shadow"
    >
      <h2 class="text-lg font-semibold">{{ todo.title }}</h2>
      <p class="text-gray-600">{{ todo.description }}</p>
      <p class="text-sm text-blue-500">Status: {{ todo.completed ? 'Completed' : 'Pending' }}</p>
    </li>
  </ul>
</template>
