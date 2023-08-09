import request from '@/request';

export function queryAgent(params) {
  return request({
    url: `/ue/agent`,
    method: 'GET'
  });
}

export function sendInstanceControl(instanceId, actionType) {
  return request({
    url: `/ue/instance/control/${instanceId}/${actionType}`,
    method: 'GET'
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

export function createInstanceConfig(config) {
  return request({
    url: `/ue/instance/config/create`,
    method: 'POST',
    data: config
  });
}

export function deleteInstance(id) {
  return request({
    url: `/ue/instance/delete/${id}`,
    method: 'GET'
  });
}
