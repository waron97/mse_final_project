import moment from 'moment'
import { useCallback, useEffect, useRef, useState } from 'react'

import { getLogs } from '../../../_shared/api'
import { useAuth } from '../../../_shared/stores'

const seconds = (n) => n * 1000

const dedupe = (logs) => {
    return logs.filter(
        (log, index, self) => self.findIndex((l) => l._id === log._id) === index
    )
}

export default function useLogs(filters) {
    const [logs, setLogs] = useState([])
    const [pagination, setPagination] = useState()

    const { apiKey } = useAuth()

    const intervalRef = useRef(null)

    const loadNextPage = useCallback(async () => {
        if (!pagination) return
        if (pagination.page < pagination.maxPages) {
            const nextPageData = await getLogs(
                {
                    ...filters,
                    page: pagination.page + 1,
                },
                apiKey
            )
            const { pagination: newPagination, data: newData } = nextPageData
            setLogs(dedupe([...logs, ...newData]))
            setPagination(newPagination)
        }
    }, [logs, pagination, filters])

    useEffect(() => {
        setPagination(null)
        setLogs([])
    }, [filters])

    useEffect(() => {
        if (!apiKey) {
            return
        }

        async function setup() {
            if (!logs.length && !pagination) {
                const { data, pagination } = await getLogs(filters, apiKey)
                setLogs(data.map((log) => ({ ...log, isFirstBatch: true })))
                setPagination(pagination)
            } else {
                intervalRef.current = setInterval(async () => {
                    const since =
                        logs?.[0] && moment(logs[0].date).toISOString()
                    const newLogs = await getLogs({ ...filters, since }, apiKey)
                    if (newLogs?.data?.length) {
                        setLogs(dedupe([...newLogs.data, ...logs]))
                    }
                }, seconds(2))
            }
        }

        setup()

        return () => {
            intervalRef.current && clearInterval(intervalRef.current)
        }
    }, [logs, apiKey, filters])

    return [logs, loadNextPage]
}
