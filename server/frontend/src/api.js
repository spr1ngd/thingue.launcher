import request from '@/request';

export function getClientList() {
    return request({
        url: `/instance/clientList`,
        method: 'GET'
    });
}

export function controlProcess(streamerId, command) {
    return request({
        url: `/instance/processControl`,
        method: 'POST',
        data: {
            streamerId,
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
