import request from '@/request';

export function getNodeList() {
    return request({
        url: `/instance/nodeList`,
        method: 'GET'
    });
}

export function controlProcess(sid, command) {
    return request({
        url: `/instance/processControl`,
        method: 'POST',
        data: {
            sid,
            command
        }
    });
}

export function sendPakControl(data) {
    return request({
        url: `/instance/pakControl`,
        method: 'POST',
        data
    });
}

export function collectLogs(data) {
    return request({
        url: `/instance/collectLogs`,
        method: 'POST',
        data
    });
}

export function downloadLogs(id) {
    return request({
        url: `/instance/downloadLogs?traceId=${id}`,
        method: 'GET',
        responseType: "blob"
    });
}
