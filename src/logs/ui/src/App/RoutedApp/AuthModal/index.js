import { Button, Input, Modal, message } from 'antd'
import PropTypes from 'prop-types'
import React, { useState } from 'react'
import styled from 'styled-components'

import { validateApiKey } from '../../_shared/api'
import { useAuth } from '../../_shared/stores'

// ----------------------------------------------------------------------------

function _AuthModal(props) {
    // -------------------------------------
    // Props destructuring
    // -------------------------------------

    const { className } = props

    // -------------------------------------
    // Hooks (e.g. useState, ...)
    // -------------------------------------

    const [value, setValue] = useState('')

    const { apiKey, setApiKey } = useAuth()

    // -------------------------------------
    // Memoized values
    // -------------------------------------

    // -------------------------------------
    // Effects
    // -------------------------------------

    // -------------------------------------
    // Component functions
    // -------------------------------------

    function handleSubmit(e) {
        e.preventDefault()

        validateApiKey(value)
            .then(() => {
                setApiKey(value)
            })
            .catch(() => {
                message.error('Invalid API key')
            })
    }

    // -------------------------------------
    // Component local variables
    // -------------------------------------

    return (
        <Modal open={!apiKey} footer={false} title={false} closable={false}>
            <div className={`${className}`}>
                <h4 className="mb-1 font-semibold">Authentication required</h4>
                <form
                    onSubmit={handleSubmit}
                    className="flex flex-row items-center gap-2"
                >
                    <Input
                        className="block"
                        value={value}
                        placeholder="Insert service API key"
                        onChange={(e) => setValue(e.target.value)}
                    />
                    <Button
                        type="primary"
                        htmlType="submit"
                        className="block bg-blue-400 border-blue-400"
                    >
                        Confirm
                    </Button>
                </form>
            </div>
        </Modal>
    )
}

// ----------------------------------------------------------------------------
// Component PropTypes and default props
// ----------------------------------------------------------------------------

_AuthModal.propTypes = {
    className: PropTypes.string.isRequired,
}

_AuthModal.defaultProps = {}

// ----------------------------------------------------------------------------

const AuthModal = styled(_AuthModal)`
    & {
    }
`
// ----------------------------------------------------------------------------

export default AuthModal
