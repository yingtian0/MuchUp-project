import React, { createContext, useCallback, useContext, useMemo, useState } from "react";

export type ToastVariant = "error" | "success" | "info";

export interface ToastItem {
  id: string;
  message: string;
  variant: ToastVariant;
}

interface ToastContextValue {
  push: (message: string, variant?: ToastVariant) => void;
}

const ToastContext = createContext<ToastContextValue | null>(null);

export const useToast = () => {
  const ctx = useContext(ToastContext);
  if (!ctx) {
    throw new Error("ToastProvider is missing");
  }
  return ctx;
};

const variantStyles: Record<ToastVariant, string> = {
  error: "border-[#f3c1b0] bg-[#fdece6] text-[#9b3d24]",
  success: "border-[#bcd9c0] bg-[#eef7ef] text-[#2f5a3c]",
  info: "border-[#c7d6e6] bg-[#eef4fb] text-[#2a445f]",
};

export const ToastProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [toasts, setToasts] = useState<ToastItem[]>([]);

  const push = useCallback((message: string, variant: ToastVariant = "info") => {
    const id = `${Date.now()}-${Math.random().toString(16).slice(2)}`;
    setToasts(prev => [...prev, { id, message, variant }]);

    window.setTimeout(() => {
      setToasts(prev => prev.filter(item => item.id !== id));
    }, 4000);
  }, []);

  const value = useMemo(() => ({ push }), [push]);

  return (
    <ToastContext.Provider value={value}>
      {children}
      <div className="fixed right-6 top-6 z-50 flex w-[320px] flex-col gap-3">
        {toasts.map(toast => (
          <div
            key={toast.id}
            className={`rounded-xl border px-4 py-3 text-sm shadow-lg ${
              variantStyles[toast.variant]
            }`}
          >
            {toast.message}
          </div>
        ))}
      </div>
    </ToastContext.Provider>
  );
};
