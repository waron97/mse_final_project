export function getStatusColor(level) {
    switch (level) {
        case 'debug':
            return '#B0DDFF'
        case 'info':
            return '#3BABFF'
        case 'warning':
            return '#FFC13B'
        case 'error':
            return '#FD572A'
        case 'critical':
            return '#FF0000'
        default:
            return 'white'
    }
}
