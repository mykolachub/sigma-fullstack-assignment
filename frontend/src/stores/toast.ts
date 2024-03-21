import { create } from 'zustand';

const RESET_TIMEOUT = 3000; // 2 seconds

export enum ToastType {
  info = 'info',
  warning = 'warning',
  error = 'error',
}

interface Toast {
  type: ToastType | null;
  message: string | null;
}

interface ToastState {
  toasts: Toast[];
  counter: number;
  addToastWarning: (message: string) => void;
  addToastInfo: (message: string) => void;
  addToastError: (message: string) => void;
  reset: () => void;
}

const useToastStore = create<ToastState>((set) => ({
  hasToast: false,
  toasts: [],
  counter: 0,
  addToastWarning: (message: string): void => {
    set((state) => ({
      toasts: [...state.toasts, { type: ToastType.warning, message }],
      counter: state.counter + 1,
    }));
    useToastStore.getState().reset();
  },
  addToastInfo: (message: string): void => {
    set((state) => ({
      toasts: [...state.toasts, { type: ToastType.info, message }],
      counter: state.counter + 1,
    }));
    useToastStore.getState().reset();
  },
  addToastError: (message: string): void => {
    set((state) => ({
      toasts: [...state.toasts, { type: ToastType.error, message }],
      counter: state.counter + 1,
    }));
    useToastStore.getState().reset();
  },
  reset: (): void => {
    setTimeout(() => {
      set((state) => ({
        toasts: state.toasts.slice(1),
        counter: state.counter - 1,
      }));
    }, RESET_TIMEOUT);
  },
}));

export default useToastStore;
