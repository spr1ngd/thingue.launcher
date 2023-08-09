export function bytesToSize(bytes) {
  if (bytes === 0) return '0 B';

  var k = 1024;

  var sizes = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];

  var i = Math.floor(Math.log(bytes) / Math.log(k));

  return (bytes / Math.pow(k, i)).toPrecision(2) + ' ' + sizes[i];
  //toPrecision(3) 后面保留一位小数，如1.0GB                                                                                                                  //return (bytes / Math.pow(k, i)).toPrecision(3) + ' ' + sizes[i];
}

export function statusToText(val) {
  if (val === null) {
    return '未就绪';
  } else if (val === 'Starting') {
    return '正在启动';
  } else if (val === 'Started') {
    return '已启动';
  } else if (val === 'Stoped') {
    return '已停止';
  } else {
    return '未知';
  }
}
