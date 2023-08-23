export function bytesToSize(bytes) {
    if (bytes === 0) return '0 B';

    var k = 1024;

    var sizes = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];

    var i = Math.floor(Math.log(bytes) / Math.log(k));

    return (bytes / Math.pow(k, i)).toPrecision(2) + ' ' + sizes[i];
    //toPrecision(3) 后面保留一位小数，如1.0GB                                                                                                                  //return (bytes / Math.pow(k, i)).toPrecision(3) + ' ' + sizes[i];
}

export function processStateToText(val) {
    if (val === 0) {
        return '未启动';
    } else if (val === 1) {
        return '正在运行';
    } else if (val === -1) {
        return '异常退出';
    } else {
        return '未知';
    }
}

export function getPaksName(paks) {
    if (paks) {
        return paks.map(pak => pak.name)
    } else {
        return []
    }
}