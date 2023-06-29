import PropTypes from 'prop-types'
import React, { useEffect, useRef, useState } from 'react'
import InfiniteScroll from 'react-infinite-scroll-component'
import styled from 'styled-components'

import LogItem from './LogItem'

// ----------------------------------------------------------------------------

function _Terminal(props) {
    // -------------------------------------
    // Props destructuring
    // -------------------------------------

    const { className, logs, loadNextPage, onLogSelected } = props

    // -------------------------------------
    // Hooks (e.g. useState, ...)
    // -------------------------------------

    const scrollRef = useRef(null)
    const logsWrapperRef = useRef(null)

    useScrollHandler({
        scrollRef,
        logsWrapperRef,
    })

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
        <div
            ref={scrollRef}
            id="scrollTarget"
            className={`${className}`}
            style={{
                overflow: 'auto',
                display: 'flex',
                flexDirection: 'column-reverse',
            }}
        >
            <InfiniteScroll
                inverse
                hasMore
                next={loadNextPage}
                dataLength={logs?.length}
                style={{
                    width: '100%',
                    display: 'flex',
                    flexDirection: 'column-reverse',
                }}
                scrollableTarget={'scrollTarget'}
            >
                <div ref={logsWrapperRef} className="logs-wrapper">
                    {logs?.map((log) => {
                        return (
                            <div key={log._id}>
                                <LogItem
                                    log={log}
                                    onClick={() => onLogSelected(log._id)}
                                />
                                <div className="logs-sep" />
                            </div>
                        )
                    })}
                </div>
            </InfiniteScroll>
        </div>
    )
}

// ----------------------------------------------------------------------------
// Component PropTypes and default props
// ----------------------------------------------------------------------------

_Terminal.propTypes = {
    className: PropTypes.string.isRequired,
}

_Terminal.defaultProps = {}

// ----------------------------------------------------------------------------

const Terminal = styled(_Terminal)`
    & {
        flex: 1;
        display: flex;

        position: absolute;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;

        overflow: scroll;

        background: #212121;

        .logs-wrapper {
            width: 100%;
            display: flex;
            flex-direction: column-reverse;
            justify-content: center;
            padding: 0px 12px;
            height: fit-content;
            padding-top: 10px;
            .logs-sep {
                width: 100%;
                background: white;
                align-self: center;
                border-bottom: 1px solid #a1a1a1;
            }
        }
    }
`
// ----------------------------------------------------------------------------

export default Terminal

const useScrollHandler = ({ scrollRef, logsWrapperRef }) => {
    const [userScrolledToBottom, setUserScrolledToBottom] = useState(true)

    useEffect(() => {
        /** @type {HTMLDivElement} */
        const scrollContainer = scrollRef.current
        scrollContainer.scrollTo({
            top: scrollContainer.scrollHeight,
            left: 0,
            behavior: 'auto',
        })
    }, [])

    useEffect(() => {
        /** @type {HTMLDivElement} */
        const scrollContainer = scrollRef.current
        /** @type {HTMLDivElement} */
        const logsContainer = logsWrapperRef.current

        if (!scrollContainer || !logsContainer) return

        function onWrapperResize() {
            if (userScrolledToBottom) {
                const options = {
                    top: scrollContainer.scrollHeight,
                    behavior: 'auto',
                    left: 0,
                }
                scrollContainer.scrollTo(options)
            }
        }

        function onContainerScroll() {
            const { clientHeight, scrollHeight, scrollTop } = scrollContainer

            // console.log(
            //     scrollTop + clientHeight + 1,
            //     scrollHeight,
            //     scrollTop + clientHeight + 1 >= scrollHeight
            // )

            // Save if user has scrolled to the bottom of the terminal
            // this +1 fixes some weird edge-case where sometimes the LHS resolves to RHS-0.5
            if (scrollTop + clientHeight + 1 >= scrollHeight) {
                // console.log('User scrolled to bottom of container')
                setUserScrolledToBottom(true)
            } else {
                setUserScrolledToBottom(false)
            }
        }

        const observer = new ResizeObserver(onWrapperResize)
        observer.observe(logsContainer)

        scrollContainer.addEventListener('scroll', onContainerScroll)

        return () => {
            observer.disconnect()
            scrollContainer.removeEventListener('scroll', onContainerScroll)
        }
    }, [userScrolledToBottom])
}
