import axios from 'axios';
import {Notify} from 'quasar';

const instance = axios.create({
    baseURL: window.location.pathname.slice(0, location.pathname.lastIndexOf("/"))
        .replace("/static", "") + "/api"
});

instance.interceptors.response.use(
    (response) => {
        if (response.config.responseType === "blob") {
            return response
        } else {
            if (response.data.code && response.data.msg) {
                if (response.data.code === 200) {
                    Notify.create({type: 'positive', position: 'top', message: response.data.msg});
                } else {
                    Notify.create({type: 'warning', position: 'top', message: response.data.msg});
                }
            }
            return response.data;
        }
    },
    (error) => {
        Notify.create({type: 'negative', position: 'top', message: '接口请求失败'});
        return Promise.reject(error);
    }
);

export default instance;
