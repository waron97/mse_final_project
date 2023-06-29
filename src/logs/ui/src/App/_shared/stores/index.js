import create from 'zustand'
import { persist } from 'zustand/middleware'

import { LogLevels } from '../../RoutedApp/Home/components/Filters'

/** @type {() => {apiKey?: string; setApiKey: (key: string) => void}} */
export const useAuth = create(
    persist((set) => ({
        apiKey: null,
        setApiKey: (key) => set({ apiKey: key }),
    }))
)

/** @type {() => { filters: {appId: string, levels: string[] }, setFilters: (f) => {} }} */
export const useLogFilters = create(
    persist((set) => ({
        filters: {
            appId: undefined,
            levels: LogLevels.map((l) => l.value),
        },
        setFilters: (filters) => set({ filters }),
    }))
)
