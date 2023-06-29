import PropTypes from 'prop-types'
import React, { useMemo, useState } from 'react'
import styled from 'styled-components'

import Filters, { LogLevels } from './components/Filters'
import LogDrawer from './components/LogDrawer'
import Terminal from './components/Terminal'
import useLogs from './hooks/useLogs'

// ----------------------------------------------------------------------------

function _Home(props) {
    // -------------------------------------
    // Props destructuring
    // -------------------------------------

    const { className } = props

    // -------------------------------------
    // Hooks (e.g. useState, ...)
    // -------------------------------------

    const [selectedLogId, setSelectedLogId] = useState()

    const [filters, setFilters] = useState({
        appId: undefined,
        levels: LogLevels.map((l) => l.value),
    })

    const [logs, loadNextPage] = useLogs(filters)

    const selectedLog = useMemo(
        () => logs?.find((l) => l._id === selectedLogId),
        [selectedLogId]
    )

    // -------------------------------------
    // Memoized values
    // -------------------------------------

    // -------------------------------------
    // Effects
    // -------------------------------------

    // -------------------------------------
    // Component functions
    // -------------------------------------

    function handleChange(field) {
        return (v) => {
            const value = v?.target?.value !== undefined ? v.target.value : v
            setFilters({ ...filters, [field]: value })
        }
    }

    // -------------------------------------
    // Component local variables
    // -------------------------------------

    return (
        <div className={`${className}`}>
            <div className="terminal-wrapper">
                <Filters values={filters} onChange={handleChange} />
                <div className="terminal-wrapper-inner">
                    <Terminal
                        logs={logs}
                        loadNextPage={loadNextPage}
                        onLogSelected={(id) => setSelectedLogId(id)}
                    />
                </div>
            </div>
            <LogDrawer
                open={!!selectedLogId}
                onClose={() => setSelectedLogId(null)}
                log={selectedLog}
            />
        </div>
    )
}

// ----------------------------------------------------------------------------
// Component PropTypes and default props
// ----------------------------------------------------------------------------

_Home.propTypes = {
    className: PropTypes.string.isRequired,
}

_Home.defaultProps = {}

// ----------------------------------------------------------------------------

const Home = styled(_Home)`
    & {
        display: flex;
        flex-direction: column;
        align-items: center;
        height: 100%;
        flex: 1;
        .terminal-wrapper {
            width: 70%;
            flex: 1;

            display: flex;
            flex-direction: column;
            position: relative;

            .terminal-wrapper-inner {
                flex: 1;
                display: flex;
                flex-direction: column;
                position: relative;
            }
        }
    }
`
// ----------------------------------------------------------------------------

export default Home
