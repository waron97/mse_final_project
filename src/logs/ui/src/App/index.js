import PropTypes from 'prop-types'
import React from 'react'
import { QueryClient, QueryClientProvider } from 'react-query'
import styled from 'styled-components'

import RoutedApp from './RoutedApp'

const client = new QueryClient()

// ----------------------------------------------------------------------------

function _App() {
    // -------------------------------------
    // Props destructuring
    // -------------------------------------

    // -------------------------------------
    // Hooks (e.g. useState, ...)
    // -------------------------------------

    // -------------------------------------
    // Memoized values
    // -------------------------------------

    // -------------------------------------
    // Effects
    // -------------------------------------

    // -------------------------------------
    // Component functions
    // -------------------------------------

    // -------------------------------------
    // Component local variables
    // -------------------------------------

    return (
        <QueryClientProvider client={client}>
            <RoutedApp />
        </QueryClientProvider>
    )
}

// ----------------------------------------------------------------------------
// Component PropTypes and default props
// ----------------------------------------------------------------------------

_App.propTypes = {
    className: PropTypes.string.isRequired,
}

_App.defaultProps = {}

// ----------------------------------------------------------------------------

const App = styled(_App)`
    & {
    }
`
// ----------------------------------------------------------------------------

export default App
