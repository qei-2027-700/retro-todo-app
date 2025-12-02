import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "ログイン - Retro Todo App",
  description: "ログインして振り返りとTODOを管理しましょう。",
};

/**
 * 認証ページ用のレイアウト
 * Header/Footer/SideNavなし、シンプルなレイアウト
 */
export default function AuthLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return <>{children}</>;
}
