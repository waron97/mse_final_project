import PropTypes from 'prop-types'
import React from 'react'
import { RouterProvider, createHashRouter } from 'react-router-dom'
import styled from 'styled-components'

import Navbar from '../_shared/components/Navbar'
import AuthModal from './AuthModal'
import Home from './Home'

// ----------------------------------------------------------------------------

const router = createHashRouter([{ path: '/', element: <Home /> }])

function _RoutedApp(props) {
    // -------------------------------------
    // Props destructuring
    // -------------------------------------

    const { className } = props

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
        <div className={`${className}`}>
            <Navbar />
            <main>
                <RouterProvider router={router} />
            </main>
            <AuthModal />
        </div>
    )
}

// ----------------------------------------------------------------------------
// Component PropTypes and default props
// ----------------------------------------------------------------------------

_RoutedApp.propTypes = {
    className: PropTypes.string.isRequired,
}

_RoutedApp.defaultProps = {}

// ----------------------------------------------------------------------------

const RoutedApp = styled(_RoutedApp)`
    & {
        display: flex;
        flex-direction: column;
        width: 100vw;
        max-width: 100vw;
        min-height: 100vh;
        main {
            flex: 1;
            background: #ededed;
            padding: 20px;

            display: flex;
            flex-direction: column;
        }
    }
`
// ----------------------------------------------------------------------------

export default RoutedApp
