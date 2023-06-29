import PropTypes from 'prop-types'
import React from 'react'
import styled from 'styled-components'

// ----------------------------------------------------------------------------

function _Navbar(props) {
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
            <div className="logo">Logs Service</div>
        </div>
    )
}

// ----------------------------------------------------------------------------
// Component PropTypes and default props
// ----------------------------------------------------------------------------

_Navbar.propTypes = {
    className: PropTypes.string.isRequired,
}

_Navbar.defaultProps = {}

// ----------------------------------------------------------------------------

const Navbar = styled(_Navbar)`
    & {
        width: 100%;
        height: 60px;
        background: white;

        box-shadow: inset 0px -1px 0px #e2e2ea;

        display: flex;
        align-items: center;
        padding: 0px 30px;

        .logo {
            font-weight: 700;
            font-size: 22px;
        }
    }
`
// ----------------------------------------------------------------------------

export default Navbar
