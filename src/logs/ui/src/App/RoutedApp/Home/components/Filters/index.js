import { Input, Select } from 'antd'
import PropTypes from 'prop-types'
import React, { useEffect } from 'react'
import { useQuery } from 'react-query'
import styled from 'styled-components'

import { getAppIds } from '../../../../_shared/api'
import { useAuth } from '../../../../_shared/stores'

// ----------------------------------------------------------------------------

export const LogLevels = [
    { label: 'Debug', value: 'debug' },
    { label: 'Info', value: 'info' },
    { label: 'Warning', value: 'warning' },
    { label: 'Error', value: 'error' },
    { label: 'Critical', value: 'critical' },
]

function _Filters(props) {
    // -------------------------------------
    // Props destructuring
    // -------------------------------------

    const { className, values, onChange } = props

    // -------------------------------------
    // Hooks (e.g. useState, ...)
    // -------------------------------------

    const { apiKey } = useAuth()

    const { data: appIds = [] } = useQuery(
        ['app-ids', apiKey],
        () => getAppIds(apiKey),
        { enabled: !!apiKey }
    )

    // -------------------------------------
    // Memoized values
    // -------------------------------------

    // -------------------------------------
    // Effects
    // -------------------------------------

    useEffect(() => {
        if (!values.appId && appIds?.[0]) {
            onChange('appId')(appIds[0])
        }
    }, [appIds])

    // -------------------------------------
    // Component functions
    // -------------------------------------

    // -------------------------------------
    // Component local variables
    // -------------------------------------

    const labelClass = 'block'
    const inputClass = 'block'

    return (
        <div
            className={`${className} flex flex-col items-start flex-wrap gap-4`}
        >
            <div className="flex flex-row items-center justify-between w-full gap-3">
                <label className={`${labelClass} flex-1`}>
                    <span className="font-medium">Service</span>
                    <Select
                        className={`${inputClass} w-full`}
                        options={appIds.map((appId) => ({
                            label: appId,
                            value: appId,
                        }))}
                        value={values.appId}
                        onChange={onChange('appId')}
                    />
                </label>
                <label className={`${labelClass} flex-1`}>
                    <span className="font-medium">Text search</span>
                    <Input
                        className={`${inputClass} w-full`}
                        value={values.text}
                        onChange={onChange('text')}
                    />
                </label>
                {/* <label className={`${labelClass} flex-1`}>
                    <span className="font-medium">A partire da</span>
                    <DatePicker
                        className={`${inputClass} w-full`}
                        value={values?.since ? moment(values.since) : null}
                        onChange={(v) => onChange('since')(v)}
                        showTime
                        format={'D MMM - HH:mm'}
                    />
                </label> */}
            </div>
            <label className="block w-full">
                <span className="font-medium">Levels</span>
                <Select
                    className={`${inputClass} w-full`}
                    options={LogLevels}
                    value={values.levels}
                    onChange={onChange('levels')}
                    mode="multiple"
                />
            </label>
        </div>
    )
}

// ----------------------------------------------------------------------------
// Component PropTypes and default props
// ----------------------------------------------------------------------------

_Filters.propTypes = {
    className: PropTypes.string.isRequired,
}

_Filters.defaultProps = {}

// ----------------------------------------------------------------------------

const Filters = styled(_Filters)`
    & {
        width: 100%;
        /* height: 50px; */
        border-bottom: 1px solid #c1c1c1;
        background: white;
        padding: 10px 15px;

        .app-id {
        }
    }
`
// ----------------------------------------------------------------------------

export default Filters
