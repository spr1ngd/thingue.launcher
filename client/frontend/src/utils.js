function RunnerStateCodeToString(code) {
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

function GoTimeFormat(dateString) {
    const milliseconds = Date.parse(dateString);
    const date = new Date(milliseconds);
    if (date.getTime() < 0) {
        return "无"
    }
    const year = date.getFullYear();
    const month = date.getMonth() + 1; // 月份是从0开始的，需要加1
    const day = date.getDate();
    const hours = date.getHours();
    const minutes = date.getMinutes();
    const seconds = date.getSeconds();

    return `${year}-${month < 10 ? '0' + month : month}-${day < 10 ? '0' + day : day} ${hours < 10 ? '0' + hours : hours}:${minutes < 10 ? '0' + minutes : minutes}:${seconds < 10 ? '0' + seconds : seconds}`
}

export {RunnerStateCodeToString, GoTimeFormat}