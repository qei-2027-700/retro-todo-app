// import type { paths } from '~/types/api.d.ts'

export default defineEventHandler(async (event) => {
  // モックフラグ
  const config = useRuntimeConfig()
  const useMock = config.useMock === true

  if (useMock) {
    // OpenAPI型から正確にモックデータ生成
    // const mockTodos: paths['/todos']['get']['responses'][200]['content']['application/json'] = [
    const mockTodos = [
      {
        id: 1,
        title: 'モックTODO1',
        description: 'テスト用TODO',
        completed: false,
        sprint_id: 1,
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString()
      },
      {
        id: 2,
        title: 'モックTODO2',
        description: '完了済み',
        completed: true,
        sprint_id: 1,
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString()
      }
    ]
    return mockTodos
  }

  // // 本番: Go backendへプロキシ
  // const backendUrl = config.public.apiBase
  // // const { data } = await $fetch<paths['/todos']['get']['responses'][200]['content']['application/json']>(
  // const { data } = await $fetch(
  //   '/todos',
  //   { baseURL: backendUrl }
  // )
  // return data
})
