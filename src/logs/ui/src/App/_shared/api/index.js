import axios from 'axios'
import formUrlEncoded from 'form-urlencoded'

function getBaseUrl() {
    return window.location.origin
}

export function getApiUrl(endpoint, params) {
    let url = getBaseUrl() + endpoint
    if (params) {
        url += `?${formUrlEncoded(params)}`
    }
    return url
}

export function getLogs(params, apiKey) {
    // return {
    //     data: [
    //         {
    //             _id: Math.random().toString(),
    //             level: 'debug',
    //             location: 'r1',
    //             message: 'Random 1',
    //             date: dayjs().toISOString(),
    //         },
    //         {
    //             _id: Math.random().toString(),
    //             level: 'info',
    //             location: 'r1',
    //             message: 'Random 2',
    //             date: dayjs().toISOString(),
    //         },
    //         {
    //             _id: Math.random().toString(),
    //             level: 'warning',
    //             location: 'r1',
    //             message: 'Random 3',
    //             date: dayjs().toISOString(),
    //         },
    //         {
    //             _id: Math.random().toString(),
    //             level: 'error',
    //             location: 'r1',
    //             message: 'Random 4',
    //             date: dayjs().toISOString(),
    //         },
    //         {
    //             _id: Math.random().toString(),
    //             level: 'critical',
    //             location: 'r2',
    //             message: 'Random 5',
    //             date: dayjs().toISOString(),
    //         },
    //     ],
    // }
    const url = getApiUrl('/logs', { ...params })
    return sendRequest(url, 'GET', null, apiKey)
}

export function getAppIds(apiKey) {
    const url = getApiUrl('/logs/app-ids')
    return sendRequest(url, 'GET', null, apiKey)
}

export function validateApiKey(apiKey) {
    const url = getApiUrl('/validate-key')
    return sendRequest(url, 'GET', null, apiKey)
}

async function sendRequest(url, method, body, apiKey) {
    const { data } = await axios({
        method,
        url,
        data: body,
        headers: { Authorization: `apiKey ${apiKey}` },
    })

    return data
}
