import axios from 'axios';
import { Notify } from 'quasar';

const instance = axios.create({
  baseURL: './'
});

instance.interceptors.response.use(
  (response) => {
    if (response.data.code && response.data.msg) {
      if (response.data.code === 200) {
        Notify.create({ type: 'positive', position: 'top', message: response.data.msg });
      } else {
        Notify.create({ type: 'warning', position: 'top', message: response.data.msg });
      }
    }
    return response.data;
  },
  (err) => {
    Notify.create({ type: 'negative', position: 'top', message: '接口请求失败' });
    return Promise.reject(err);
  }
);

export default instance;
