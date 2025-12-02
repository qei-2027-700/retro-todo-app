import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'

const AUTH_TOKEN_KEY = 'auth_token'

// 認証が必要なパス
const protectedPaths = [
  '/dashboard', 
  '/todos', 
  '/sprints'
] as const

// 認証済みユーザーがアクセスできないパス
const authPaths = [
  '/login', 
  '/register'
] as const

export function middleware(request: NextRequest) {
  const { pathname } = request.nextUrl
  const token = request.cookies.get(AUTH_TOKEN_KEY)?.value

  // 認証が必要なページへのアクセス
  if (protectedPaths.some(path => pathname.startsWith(path))) {
    if (!token) {
      const loginUrl = new URL('/login', request.url)
      return NextResponse.redirect(loginUrl)
    }
  }

  // ログイン・登録ページへのアクセス
  if (authPaths.some(path => pathname.startsWith(path))) {
    if (token) {
      const dashboardUrl = new URL('/dashboard', request.url)
      return NextResponse.redirect(dashboardUrl)
    }
  }

  return NextResponse.next()
}

// middlewareを適用するパスを指定
export const config = {
  matcher: [
    /*
     * 以下を除くすべてのパスにマッチ:
     * - api (APIルート)
     * - _next/static (静的ファイル)
     * - _next/image (画像最適化ファイル)
     * - favicon.ico (ファビコン)
     */
    '/((?!api|_next/static|_next/image|favicon.ico).*)',
  ],
}
