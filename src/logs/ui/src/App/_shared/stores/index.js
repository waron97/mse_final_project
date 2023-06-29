import create from 'zustand'
import { persist } from 'zustand/middleware'

/** @type {() => {apiKey?: string; setApiKey: (key: string) => void}} */
export const useAuth = create(
    persist((set) => ({
        apiKey: null,
        setApiKey: (key) => set({ apiKey: key }),
    }))
)
