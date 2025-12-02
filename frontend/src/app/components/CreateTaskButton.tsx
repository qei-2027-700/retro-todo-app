"use client"

import { useState } from "react"
import { createTodo } from "@/lib/actions/todos"
import { useRouter } from "next/navigation"

interface CreateTaskButtonProps {
  variant?: "primary" | "secondary"
  size?: "sm" | "md" | "lg"
  className?: string
  onClick?: () => void
}

export default function CreateTaskButton({
  variant = "primary",
  size = "md",
  className = ""
}: CreateTaskButtonProps) {
  const [isOpen, setIsOpen] = useState(false)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState("")
  const router = useRouter()

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    setError("")
    setLoading(true)

    const formData = new FormData(e.currentTarget)
    const title = formData.get("title") as string
    const description = formData.get("description") as string

    if (!title.trim()) {
      setError("タイトルを入力してください")
      setLoading(false)
      return
    }

    try {
      await createTodo({
        title: title.trim(),
        description: description.trim(),
        completed: false,
      })

      setIsOpen(false)
      router.refresh()
    }
    catch (err) {
      setError("タスクの作成に失敗しました")
      console.error("Create task error:", err)
    }
    finally {
      setLoading(false)
    }
  }

  const buttonSizeClasses = {
    sm: "px-3 py-1.5 text-sm",
    md: "px-4 py-2 text-base",
    lg: "px-6 py-3 text-lg",
  }

  const buttonVariantClasses = {
    primary: "bg-blue-600 hover:bg-blue-700 text-white",
    secondary: "bg-gray-600 hover:bg-gray-700 text-white",
  }

  return (
    <>
      <button
        onClick={() => setIsOpen(true)}
        className={`${buttonVariantClasses[variant]} ${buttonSizeClasses[size]} rounded font-medium transition-colors cursor-pointer ${className}`}
      >
        + タスクを作成
      </button>

      {isOpen && (
        <div className="fixed inset-0 bg-opacity-50 flex items-center justify-center p-4 z-50">
          <div className="bg-white rounded-lg shadow-xl max-w-md w-full p-6">
            <div className="flex justify-between items-center mb-4">
              <h2 className="text-xl font-bold text-gray-800">新しいタスクを作成</h2>
              <button
                onClick={() => setIsOpen(false)}
                className="text-gray-500 hover:text-gray-700 text-5xl leading-none cursor-pointer"
                disabled={loading}
              >
                ×
              </button>
            </div>

            {error && (
              <div className="mb-4 p-3 bg-red-100 border border-red-400 text-red-700 rounded">
                {error}
              </div>
            )}

            <form onSubmit={handleSubmit} className="space-y-4">
              <div>
                <label htmlFor="title" className="block text-sm font-medium text-gray-700 mb-1">
                  タイトル <span className="text-red-500">*</span>
                </label>
                <input
                  id="title"
                  name="title"
                  type="text"
                  required
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-black"
                  placeholder="タスクのタイトル"
                  disabled={loading}
                />
              </div>

              <div>
                <label htmlFor="description" className="block text-sm font-medium text-gray-700 mb-1">
                  説明
                </label>
                <textarea
                  id="description"
                  name="description"
                  rows={4}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-black resize-none"
                  placeholder="タスクの詳細説明（任意）"
                  disabled={loading}
                />
              </div>

              <div className="flex gap-3 justify-end">
                <button
                  type="button"
                  onClick={() => setIsOpen(false)}
                  disabled={loading}
                  className="px-4 py-2 bg-gray-300 hover:bg-gray-400 text-gray-800 rounded font-medium transition-colors disabled:bg-gray-200 disabled:cursor-not-allowed"
                >
                  キャンセル
                </button>
                <button
                  type="submit"
                  disabled={loading}
                  className="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded font-medium transition-colors disabled:bg-blue-400 disabled:cursor-not-allowed cursor-pointer"
                >
                  {loading ? "作成中..." : "作成"}
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </>
  )
}
