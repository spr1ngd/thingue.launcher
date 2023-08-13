export function RunnerStateCodeToString(code) {
    switch (code) {
        case 0:
            return "停止"
        case 1:
            return "正在运行"
        case -1:
            return "异常退出"
        default:
            return "未知"
    }
}

export default {RunnerStateCodeToString}