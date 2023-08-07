export function statusCodeToString(code) {
    switch (code) {
        case 0:
            return "未启动"
        case 1:
            return "已启动"
        default:
            return "未知状态"
    }
}

export default {statusCodeToString}