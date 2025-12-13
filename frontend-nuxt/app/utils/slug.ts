/**
 * タイトルからURLセーフなスラッグを生成
 * @param title - 元のタイトル
 * @returns URLセーフな文字列
 */
export function slugify(title: string): string {
  return title
    .toLowerCase()
    .trim()
    // 日本語文字をそのまま残す（URLエンコードされる）
    .replace(/\s+/g, '-') // スペースをハイフンに
    .replace(/[^\w\u3040-\u309F\u30A0-\u30FF\u4E00-\u9FAF-]/g, '') // 英数字、ひらがな、カタカナ、漢字、ハイフンのみ残す
    .replace(/--+/g, '-') // 連続するハイフンを1つに
    .replace(/^-+|-+$/g, '') // 前後のハイフンを削除
}

/**
 * ハイブリッドID形式からIDを抽出
 * @param param - ルートパラメータ（例: "1-2025-12-week1"）
 * @returns 数値ID
 * @throws パラメータが無効な場合にエラー
 */
export function extractIdFromParam(param: string | string[] | undefined): number {
  if (!param) {
    throw new Error('Parameter is required')
  }

  if (Array.isArray(param)) {
    param = param[0]
  }

  // ハイフンの前の数字部分を抽出
  const match = param.match(/^(\d+)/)
  if (!match) {
    throw new Error('Invalid ID format')
  }

  return Number(match[1])
}

/**
 * スプリントのハイブリッドURL（ID-スラッグ）を生成
 * @param id - スプリントID
 * @param title - スプリントタイトル
 * @returns ハイブリッド形式のパス（例: "1-2025-12-week1"）
 */
export function generateSprintPath(id: number, title: string): string {
  const slug = slugify(title)
  return slug ? `${id}-${slug}` : `${id}`
}
