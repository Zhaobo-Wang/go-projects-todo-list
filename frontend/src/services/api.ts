import axios from 'axios';

const API_URL = 'http://localhost:8080/api/v1';

// 创建 axios 实例
const apiClient = axios.create({
    baseURL: API_URL,
    headers: {
        'Content-Type': 'application/json',
    },
});

// 请求拦截器 - 添加 token 到请求头
apiClient.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('token');
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// 响应拦截器 - 处理 401 错误
apiClient.interceptors.response.use(
    (response) => {
        return response;
    },
    (error) => {
        if (error.response && error.response.status === 401) {
            // 清除 token 并重定向到登录页
            localStorage.removeItem('token');
            window.location.href = '/login';
        }
        return Promise.reject(error);
    }
);

// 用户认证相关 API
export const authApi = {
    register: (username: string, email: string, password: string) => {
        return apiClient.post('/register', { username, email, password });
    },
    login: (username: string, password: string) => {
        return apiClient.post('/login', { username, password });
    },
    logout: () => {
        return apiClient.post('/logout');
    },
    getCurrentUser: () => {
        return apiClient.get('/user-profile');
    },
};

// Todo 相关 API
export const todoApi = {
    getAll: () => {
        return apiClient.get('/todos');
    },
    get: (id: number) => {
        return apiClient.get(`/todos/${id}`);
    },
    create: (todo: { title: string; description: string }) => {
        return apiClient.post('/todos', todo);
    },
    update: (id: number, todo: { title?: string; description?: string; completed?: boolean }) => {
        return apiClient.put(`/todos/${id}`, todo);
    },
    delete: (id: number) => {
        return apiClient.delete(`/todos/${id}`);
    },
};