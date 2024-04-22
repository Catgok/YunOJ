import axios from 'axios';

const apiService = axios.create({
    baseURL: 'http://localhost:8888/v1/',
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
        localStorage.removeItem('U-Token');
        window.location.href = '/login';
        return;
    }
    return response;
}, error => {
    // 响应错误处理
    // return Promise.reject(error);
    if (error.response.status === 401) {
        localStorage.clear();
        window.location.href = '/login';
    }
});

export default apiService;
