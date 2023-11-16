import request from '@/request';

export function getClientList() {
    return request({
        url: `/instance/clientList`,
        method: 'GET'
    });
}

export function getInstanceList() {
    return request({
        url: `/instance/instanceList`,
        method: 'GET'
    });
}

export function listCloudRes() {
    return request({
        url: `/sync/listCloudRes`,
        method: 'GET'
    });
}

export function createCloudRes(data) {
    return request({
        url: `/sync/createCloudRes`,
        method: 'POST',
        data
    });
}

export function updateCloudRes(data) {
    return request({
        url: `/sync/updateCloudRes`,
        method: 'POST',
        data
    });
}

export function deleteCloudRes(data) {
    return request({
        url: `/sync/deleteCloudRes`,
        method: 'POST',
        data
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
