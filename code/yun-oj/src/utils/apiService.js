import axios from 'axios';
import {eventBus} from "@/utils/eventBus";

const apiService = axios.create({
    // baseURL: 'http://localhost:8888/v1/',
    baseURL: 'https://api.yunoj.hhxxblog.top/v1/',
    timeout: 3000, // 请求超时时间ms
});

// 请求拦截器
apiService.interceptors.request.use(config => {
    const token = localStorage.getItem('U-Token');
    if (token) {
        config.headers['U-Token'] = token;
    }
    return config;
}, error => {
    return Promise.reject(error);
});

// 响应拦截器
apiService.interceptors.response.use(response => {
    if (response.status === 401) {
        localStorage.clear();
        window.location.href = '/login';
        return;
    }
    if (response.data.code !== 0) {
        // const noticeData = {type: "error", message: response.data.message, duration: 2500}
        // eventBus.emit('globalNotice', noticeData)
        // response.data.code = -1
        // return response;
    }
    return response;
}, error => {
    // console.log(error);
    // 响应错误处理
    if (error.response.status === 401) {
        localStorage.clear();
        window.location.href = '/login';
    }
});


export default apiService;
