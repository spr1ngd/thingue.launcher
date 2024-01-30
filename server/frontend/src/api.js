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

export function collectLogs(id) {
    return request({
        url: `/instance/collectLogs?clientId=${id}`,
        method: 'GET',
        responseType: "blob"
    });
}

export function downloadLogs(id) {
    return request({
        url: `/instance/downloadLogs?traceId=${id}`,
        method: 'GET',
        responseType: "blob"
    });
}
