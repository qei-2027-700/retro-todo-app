import { getTodos } from "../../../../lib/actions/todos";

export default async function DashboardPage() {
  const todos = await getTodos();

  // 完了/未完了のタスクを分ける
  const incompleteTodos = todos.filter(todo => !todo.completed);
  const completedTodos = todos.filter(todo => todo.completed);

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto">
        {/* ヘッダーエリア */}
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-2">
            ダッシュボード
          </h1>
          <p className="text-gray-600">
            今日のタスクと進捗状況を確認しましょう
          </p>
        </div>

        {/* 統計カード */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
          <div className="bg-white rounded-lg shadow p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">全タスク</p>
                <p className="text-3xl font-bold text-gray-900 mt-2">
                  {todos.length}
                </p>
              </div>
              <div className="h-12 w-12 bg-blue-100 rounded-full flex items-center justify-center">
                <svg className="h-6 w-6 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
                </svg>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">未完了</p>
                <p className="text-3xl font-bold text-orange-600 mt-2">
                  {incompleteTodos.length}
                </p>
              </div>
              <div className="h-12 w-12 bg-orange-100 rounded-full flex items-center justify-center">
                <svg className="h-6 w-6 text-orange-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">完了</p>
                <p className="text-3xl font-bold text-green-600 mt-2">
                  {completedTodos.length}
                </p>
              </div>
              <div className="h-12 w-12 bg-green-100 rounded-full flex items-center justify-center">
                <svg className="h-6 w-6 text-green-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
            </div>
          </div>
        </div>

        {/* タスク一覧 */}
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
          {/* 未完了タスク */}
          <div className="bg-white rounded-lg shadow">
            <div className="p-6 border-b border-gray-200">
              <h2 className="text-xl font-semibold text-gray-900">
                未完了のタスク
              </h2>
            </div>
            <div className="p-6">
              {incompleteTodos.length === 0 ? (
                <p className="text-gray-500 text-center py-8">
                  未完了のタスクはありません
                </p>
              ) : (
                <ul className="space-y-3">
                  {incompleteTodos.slice(0, 5).map((todo) => (
                    <li
                      key={todo.id}
                      className="flex items-start p-3 hover:bg-gray-50 rounded-lg transition-colors"
                    >
                      <div className="flex-shrink-0 mt-1">
                        <div className="h-5 w-5 rounded border-2 border-gray-300"></div>
                      </div>
                      <div className="ml-3 flex-1">
                        <p className="text-sm font-medium text-gray-900">
                          {todo.title}
                        </p>
                        {todo.description && (
                          <p className="text-sm text-gray-500 mt-1">
                            {todo.description}
                          </p>
                        )}
                      </div>
                    </li>
                  ))}
                </ul>
              )}
              {incompleteTodos.length > 5 && (
                <div className="mt-4 text-center">
                  <a href="/todos" className="text-sm text-blue-600 hover:text-blue-700 font-medium">
                    すべて表示 ({incompleteTodos.length}件)
                  </a>
                </div>
              )}
            </div>
          </div>

          {/* 完了済みタスク */}
          <div className="bg-white rounded-lg shadow">
            <div className="p-6 border-b border-gray-200">
              <h2 className="text-xl font-semibold text-gray-900">
                完了済みタスク
              </h2>
            </div>
            <div className="p-6">
              {completedTodos.length === 0 ? (
                <p className="text-gray-500 text-center py-8">
                  完了済みのタスクはありません
                </p>
              ) : (
                <ul className="space-y-3">
                  {completedTodos.slice(0, 5).map((todo) => (
                    <li
                      key={todo.id}
                      className="flex items-start p-3 hover:bg-gray-50 rounded-lg transition-colors"
                    >
                      <div className="flex-shrink-0 mt-1">
                        <div className="h-5 w-5 rounded bg-green-500 flex items-center justify-center">
                          <svg className="h-3 w-3 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={3} d="M5 13l4 4L19 7" />
                          </svg>
                        </div>
                      </div>
                      <div className="ml-3 flex-1">
                        <p className="text-sm font-medium text-gray-500 line-through">
                          {todo.title}
                        </p>
                        {todo.description && (
                          <p className="text-sm text-gray-400 mt-1 line-through">
                            {todo.description}
                          </p>
                        )}
                      </div>
                    </li>
                  ))}
                </ul>
              )}
              {completedTodos.length > 5 && (
                <div className="mt-4 text-center">
                  <a href="/todos" className="text-sm text-blue-600 hover:text-blue-700 font-medium">
                    すべて表示 ({completedTodos.length}件)
                  </a>
                </div>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
