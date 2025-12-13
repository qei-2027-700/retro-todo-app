export const isMacOs = () => {
  if (typeof navigator === 'undefined') return false
  return /mac os x/i.test(navigator.userAgent)
}
