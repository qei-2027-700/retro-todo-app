# レスポンシブ対応ガイド (React + Tailwind CSS v4)

## 現状

- **Tailwind CSS v4** を使用
- カスタムブレイクポイント未設定
- VueUse的なReact Hooksなし

## 推奨アプローチ

### 方法1: `usehooks-ts` (VueUseスタイル) ⭐ 推奨

VueUseの`useBreakpoints`に相当する`useMediaQuery`を提供。

#### インストール
```bash
pnpm add usehooks-ts
```

#### 使い方
```typescript
'use client'

import { useMediaQuery } from 'usehooks-ts'

export default function ResponsiveComponent() {
  const isMobile = useMediaQuery('(max-width: 640px)')
  const isTablet = useMediaQuery('(min-width: 641px) and (max-width: 1024px)')
  const isDesktop = useMediaQuery('(min-width: 1025px)')

  return (
    <div>
      {isMobile && <MobileNav />}
      {isTablet && <TabletNav />}
      {isDesktop && <DesktopNav />}
    </div>
  )
}
```

---

### 方法2: カスタムフック (自作)

VueUseスタイルのカスタムフックを自作。

#### 実装例: `lib/hooks/useBreakpoint.ts`
```typescript
'use client'

import { useEffect, useState } from 'react'

type Breakpoint = 'sm' | 'md' | 'lg' | 'xl' | '2xl'

const breakpoints = {
  sm: 640,
  md: 768,
  lg: 1024,
  xl: 1280,
  '2xl': 1536,
}

export function useBreakpoint() {
  const [currentBreakpoint, setCurrentBreakpoint] = useState<Breakpoint>('sm')

  useEffect(() => {
    const handleResize = () => {
      const width = window.innerWidth

      if (width >= breakpoints['2xl']) setCurrentBreakpoint('2xl')
      else if (width >= breakpoints.xl) setCurrentBreakpoint('xl')
      else if (width >= breakpoints.lg) setCurrentBreakpoint('lg')
      else if (width >= breakpoints.md) setCurrentBreakpoint('md')
      else setCurrentBreakpoint('sm')
    }

    handleResize()
    window.addEventListener('resize', handleResize)
    return () => window.removeEventListener('resize', handleResize)
  }, [])

  return {
    breakpoint: currentBreakpoint,
    isMobile: currentBreakpoint === 'sm',
    isTablet: currentBreakpoint === 'md',
    isDesktop: currentBreakpoint === 'lg' || currentBreakpoint === 'xl' || currentBreakpoint === '2xl',
    isLargeDesktop: currentBreakpoint === 'xl' || currentBreakpoint === '2xl',
  }
}
```

#### 使い方
```typescript
'use client'

import { useBreakpoint } from '@/lib/hooks/useBreakpoint'

export default function Header() {
  const { isMobile, isDesktop } = useBreakpoint()

  return (
    <header>
      {isMobile && <MobileMenu />}
      {isDesktop && <DesktopMenu />}
    </header>
  )
}
```

---

### 方法3: Tailwind CSSのみ (Server Component対応)

Client Componentを使わず、Tailwindのレスポンシブユーティリティのみで対応。

```tsx
export default function Header() {
  return (
    <header>
      {/* モバイル: ハンバーガーメニュー */}
      <div className="block lg:hidden">
        <MobileMenu />
      </div>

      {/* デスクトップ: 通常メニュー */}
      <div className="hidden lg:flex">
        <DesktopMenu />
      </div>
    </header>
  )
}
```

**Tailwind v4のブレイクポイント (デフォルト):**
```
sm:  640px  (min-width)
md:  768px  (min-width)
lg:  1024px (min-width)
xl:  1280px (min-width)
2xl: 1536px (min-width)
```

---

## 認証ボタンの実装

### ヘッダーにログイン/ログアウトボタンを追加

```tsx
// src/app/components/Header.tsx
'use client'

export default function Header() {
  // 仮: 認証状態 (将来はContext/Stateで管理)
  const isAuthenticated = true

  return (
    <header className="w-full bg-gray-800 text-white p-4 flex items-center justify-between">
      <h1 className="text-xl font-bold">Retro Todo App</h1>

      <div className="flex gap-4">
        {isAuthenticated ? (
          <button className="px-4 py-2 bg-red-600 hover:bg-red-700 rounded">
            ログアウト
          </button>
        ) : (
          <button className="px-4 py-2 bg-blue-600 hover:bg-blue-700 rounded">
            ログイン
          </button>
        )}
      </div>
    </header>
  )
}
```

---

## 推奨構成

### 1. インストール
```bash
pnpm add usehooks-ts
```

### 2. ファイル構成
```
src/
├── app/
│   └── components/
│       ├── Header.tsx          # レスポンシブヘッダー
│       ├── MobileNav.tsx       # モバイルナビ
│       └── DesktopNav.tsx      # デスクトップナビ
└── lib/
    └── hooks/
        └── useBreakpoint.ts    # カスタムフック (オプション)
```

### 3. 使い分け

| 方法 | 用途 | Server Component対応 |
|------|------|---------------------|
| Tailwindのみ | シンプルな表示切替 | ✅ Yes |
| `usehooks-ts` | 条件付きレンダリング | ❌ No (Client Component必須) |
| カスタムフック | 複雑な状態管理 | ❌ No (Client Component必須) |

**推奨**: 基本はTailwind、動的な制御が必要な箇所のみ`usehooks-ts`を使用
