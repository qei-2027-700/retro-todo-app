import type { Metadata } from "next";
import Header from "../components/Header";
import Footer from "../components/Footer";
import SideNav from "../components/SideNav";

export const metadata: Metadata = {
  title: "Retro Todo App",
  description: "振り返りとTODOを管理するアプリケーション。",
};

/**
 * 認証が必要なページ用のレイアウト
 * Header/Footer/SideNavを含む
 */
export default function ProtectedLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className="flex flex-col min-h-screen bg-white dark:bg-black">
      <Header />
      <div className="flex flex-1">
        <SideNav />
        <main className="flex-1 p-4 md:p-6">
          {children}
        </main>
      </div>
      <Footer />
    </div>
  );
}
