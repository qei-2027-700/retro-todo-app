"use client"

import { useState } from "react"
import { login } from "@/lib/actions/auth"
import { useRouter } from "next/navigation"
import { AuthenticationError } from "@/lib/api/error"

export default function LoginPage() {
  const [error, setError] = useState("")
  const [loading, setLoading] = useState(false)
  const router = useRouter()

  const [username, setUsername] = useState("testuser")
  const [password, setPassword] = useState("password123")

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    setError("")
    setLoading(true)

    const formData = new FormData(e.currentTarget)

    try {
      await login({
        username: formData.get("username") as string,
        password: formData.get("password") as string,
      })

      // ログイン成功後、ダッシュボードにリダイレクト
      router.push("/dashboard")
    }
    catch (err) {
      if (err instanceof AuthenticationError) {
        setError("ユーザー名またはパスワードが正しくありません")
      }
      else {
        setError("ログインに失敗しました。もう一度お試しください。")
      }
      console.error("Login error:", err)
    }
    finally {
      setLoading(false)
    }
  }

  return (
    <div className="flex min-h-screen items-center justify-center bg-gray-100 p-4">
      <div className="w-full max-w-md bg-white rounded-lg shadow-md p-8">
        <h1 className="text-2xl font-bold text-gray-800 mb-6 text-center">
          Retro Todo App
        </h1>
        <h2 className="text-xl font-semibold text-gray-700 mb-6 text-center">
          ログイン
        </h2>

        {error && (
          <div className="mb-4 p-3 bg-red-100 border border-red-400 text-red-700 rounded">
            {error}
          </div>
        )}

        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label htmlFor="username" className="block text-sm font-medium text-gray-700 mb-1">
              ユーザー名
            </label>
            <input
              id="username"
              name="username"
              type="text"
              required
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-black"
              placeholder="username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              disabled={loading}
            />
          </div>

          <div>
            <label htmlFor="password" className="block text-sm font-medium text-gray-700 mb-1">
              パスワード
            </label>
            <input
              id="password"
              name="password"
              type="password"
              required
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-black"
              placeholder="••••••••"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              disabled={loading}
            />
          </div>

          <button
            type="submit"
            disabled={loading}
            className="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-md transition-colors disabled:bg-blue-400 disabled:cursor-not-allowed cursor-pointer"
          >
            {loading ? "ログイン中..." : "ログイン"}
          </button>
        </form>

        {/* TODO: 新規登録実装 */}
        {/* <div className="mt-6 text-center">
          <p className="text-sm text-gray-600">
            アカウントをお持ちでない方は{" "}
            <a href="/register" className="text-blue-600 hover:underline">
              新規登録
            </a>
          </p>
        </div> */}

        <div className="mt-4 text-center">
          <p className="text-xs text-gray-500">
            テストユーザー: myuser / password123
          </p>
        </div>
      </div>
    </div>
  )
}
