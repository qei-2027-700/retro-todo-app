'use client'

import Link from 'next/link';
import { logout } from '@/lib/actions/auth';

const Header = () => {
  // TODO: 認証状態管理を実装後、実際の状態を使用
  const isAuthenticated = true;

  const handleLogout = async () => {
    try {
      await logout();
    }
    catch (error) {
      console.error('Logout error:', error);
    }
  };

  return (
    <header className="w-full bg-gray-800 text-white p-4">
      <div className="flex items-center justify-between">
        {/* 左側: メニューボタン(モバイルのみ) + ロゴ */}
        <div className="flex items-center gap-4">
          {/* モバイルメニューボタン (将来実装) */}
          <button
            className="md:hidden p-2 hover:bg-gray-700 rounded"
            aria-label="メニュー"
          >
            <svg
              className="w-6 h-6"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M4 6h16M4 12h16M4 18h16"
              />
            </svg>
          </button>

          {/* ロゴ/タイトル */}
          <Link href="/">
            <h1 className="text-lg sm:text-xl font-bold cursor-pointer hover:text-gray-300 transition-colors">
              Retro Todo App
            </h1>
          </Link>
        </div>

        {/* 右側: 認証ボタン */}
        <div className="flex items-center gap-2 sm:gap-4">
          {/* ユーザー名 (モバイルでは非表示) */}
          {isAuthenticated && (
            <span className="hidden sm:inline text-sm text-gray-300">
              myuser
            </span>
          )}

          {/* ログイン/ログアウトボタン */}
          {isAuthenticated ? (
            <button
              onClick={handleLogout}
              className="px-3 sm:px-4 py-2 bg-red-600 hover:bg-red-700 rounded text-xs sm:text-sm font-medium transition-colors cursor-pointer"
            >
              ログアウト
            </button>
          ) : (
            <Link href="/login">
              <button className="px-3 sm:px-4 py-2 bg-blue-600 hover:bg-blue-700 rounded text-xs sm:text-sm font-medium transition-colors">
                ログイン
              </button>
            </Link>
          )}
        </div>
      </div>
    </header>
  );
};

export default Header;