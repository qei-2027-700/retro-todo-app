import { getTodos } from "../../../lib/actions/todos";

export default async function TodosPage() {
  const todos = await getTodos();

  return (
    <div className="flex min-h-screen items-center justify-center bg-gray-100 p-8">
      <div>
        <h1 className="text-4xl font-bold text-gray-800 mb-6">
          Todos
        </h1>

        <ul className="space-y-4">
          {todos.map((todo) => (
            <li key={todo.id} className="p-4 bg-white rounded shadow">
              <h2 className="text-xl font-semibold">{todo.title}</h2>
              <p className="text-gray-600">{todo.description}</p>
              <span className="text-sm text-gray-400">
                Completed: {todo.completed ? "Yes" : "No"}
              </span>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}