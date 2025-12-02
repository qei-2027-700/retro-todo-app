import { getTodos } from "@/lib/actions/todos";
import CreateTaskButton from "@/app/components/CreateTaskButton";
// import Modal from "@/app/components/Modal";
// import { useState } from "react";

export default async function TodosPage() {
  const todos = await getTodos();

  // const [isOpen, setIsOpen] = useState(false);

  return (
    <div className="flex min-h-screen bg-gray-100 p-8">
      <div className="w-full max-w-4xl">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-4xl font-bold text-gray-800">
            Todos
          </h1>
          <CreateTaskButton />
          {/* <CreateTaskButton onClick={() => setIsOpen(true)} />
           {isOpen && (
          <Modal isOpen={isOpen} onClose={() => setIsOpen(false)}>
            aaa
          </Modal> */}
        {/* )} */}
        </div>

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