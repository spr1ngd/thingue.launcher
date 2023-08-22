import request from '@/request';

export function getNodeList(params) {
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

export function sendPakControl(action) {
    return request({
        url: `/ue/pak/control`,
        method: 'POST',
        data: action
    });
}

export function saveInstanceConfig(config) {
    return request({
        url: `/ue/instance/config/save`,
        method: 'POST',
        data: config
    });
}
