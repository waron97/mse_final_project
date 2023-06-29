import { Drawer } from 'antd'
import dayjs from 'dayjs'
import PropTypes from 'prop-types'
import React from 'react'
import ReactJson from 'react-json-view'
import styled from 'styled-components'

import { getStatusColor } from '../../../../_shared/constants'

// ----------------------------------------------------------------------------

function _LogDrawer(props) {
    // -------------------------------------
    // Props destructuring
    // -------------------------------------

    const { className, open, onClose, log } = props

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
        <Drawer
            open={open}
            onClose={onClose}
            // placement="bottom"
            size="large"
            title={`${log?.location} - ${dayjs(log?.date).format(
                'D MMM YYYY HH:mm'
            )}`}
        >
            <div className={`${className}`}>
                <div className="flex flex-col items-start">
                    <StatusBadge className="mb-2" status={log?.level} />
                    <p>{log?.message}</p>
                </div>
                <div className="sep" />
                <ReactJson src={log?.detail ?? {}} />
            </div>
        </Drawer>
    )
}

// ----------------------------------------------------------------------------
// Component PropTypes and default props
// ----------------------------------------------------------------------------

_LogDrawer.propTypes = {
    className: PropTypes.string.isRequired,
}

_LogDrawer.defaultProps = {}

// ----------------------------------------------------------------------------

const LogDrawer = styled(_LogDrawer)`
    & {
        p {
            display: block;
        }
        .sep {
            width: 100%;
            border-bottom: 1px solid #a5a5a5;
            margin: 12px 0px;
        }
    }
`
// ----------------------------------------------------------------------------

export default LogDrawer

const Container = styled.div`
    color: white;
    font-weight: 700;
    padding: 0px 6px;
    font-size: 11px;
    height: 20px;
    border-radius: 20px;
    display: flex;
    align-items: center;
    justify-content: center;
`
// eslint-disable-next-line
const StatusBadge = ({ status, className }) => {
    const bgColor = getStatusColor(status)

    return (
        <Container className={className} style={{ background: bgColor }}>
            {status}
        </Container>
    )
}
