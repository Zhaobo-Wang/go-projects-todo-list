import { defineStore } from 'pinia';
import { authApi } from '@/services/api';
import type { AuthState, User } from '@/types';

export const useAuthStore = defineStore('auth', {
    state: (): AuthState => ({
        user: null,
        token: localStorage.getItem('token'),
        isAuthenticated: !!localStorage.getItem('token'),
    }),

    actions: {
        async register(username: string, email: string, password: string) {
            try {
                const response = await authApi.register(username, email, password);
                return response.data;
            } catch (error: any) {
                throw new Error(error.response?.data?.error || 'Registration failed');
            }
        },

        async login(username: string, password: string) {
            try {
                const response = await authApi.login(username, password);
                this.token = response.data.token;
                localStorage.setItem('token', response.data.token);
                this.isAuthenticated = true;
                await this.fetchCurrentUser();
                return response.data;
            } catch (error: any) {
                throw new Error(error.response?.data?.error || 'Login failed');
            }
        },

        async logout() {
            try {
                await authApi.logout();
            } catch (error) {
                console.error('Logout error:', error);
            } finally {
                this.token = null;
                this.user = null;
                this.isAuthenticated = false;
                localStorage.removeItem('token');
            }
        },

        async fetchCurrentUser() {
            try {
                const response = await authApi.getCurrentUser();
                this.user = response.data;
                return response.data;
            } catch (error) {
                this.user = null;
                throw error;
            }
        },

        initAuth() {
            if (this.token) {
                this.fetchCurrentUser().catch(() => {
                    this.logout();
                });
            }
        },
    },
});