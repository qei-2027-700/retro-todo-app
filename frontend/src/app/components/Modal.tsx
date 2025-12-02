"use client"

import { ReactNode } from "react"

interface ModalProps {
  isOpen: boolean
  onClose: () => void
  title?: string
  children: ReactNode
  maxWidth?: "sm" | "md" | "lg" | "xl" | "2xl"
  showCloseButton?: boolean
}

export default function Modal({
  isOpen,
  onClose,
  title,
  children,
  maxWidth = "md",
  showCloseButton = true,
}: ModalProps) {
  if (!isOpen) return null

  const maxWidthClasses = {
    sm: "max-w-sm",
    md: "max-w-md",
    lg: "max-w-lg",
    xl: "max-w-xl",
    "2xl": "max-w-2xl",
  }

  const handleBackdropClick = (e: React.MouseEvent<HTMLDivElement>) => {
    if (e.target === e.currentTarget) {
      onClose()
    }
  }

  return (
    <div
      className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50"
      onClick={handleBackdropClick}
    >
      <div className={`bg-white rounded-lg shadow-xl ${maxWidthClasses[maxWidth]} w-full p-6`}>
        {(title || showCloseButton) && (
          <div className="flex justify-between items-center mb-4">
            {title && <h2 className="text-xl font-bold text-gray-800">{title}</h2>}
            {showCloseButton && (
              <button
                onClick={onClose}
                className="text-gray-500 hover:text-gray-700 text-5xl leading-none cursor-pointer ml-auto"
                aria-label="閉じる"
              >
                ×
              </button>
            )}
          </div>
        )}
        {children}
      </div>
    </div>
  )
}
