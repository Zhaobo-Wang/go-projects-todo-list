import { defineStore } from 'pinia';
import { todoApi } from '@/services/api';
import type { TodoState, Todo } from '@/types';

export const useTodoStore = defineStore('todo', {
    state: (): TodoState => ({
        todos: [],
        loading: false,
        error: null,
    }),

    getters: {
        completedTodos: (state) => state.todos.filter(todo => todo.completed),
        pendingTodos: (state) => state.todos.filter(todo => !todo.completed),
    },

    actions: {
        async fetchTodos() {
            this.loading = true;
            this.error = null;
            try {
                const response = await todoApi.getAll();
                this.todos = response.data;
            } catch (error: any) {
                this.error = error.message;
                console.error('Failed to fetch todos:', error);
            } finally {
                this.loading = false;
            }
        },

        async fetchTodo(id: number) {
            this.loading = true;
            this.error = null;
            try {
                const response = await todoApi.get(id);
                return response.data;
            } catch (error: any) {
                this.error = error.message;
                console.error(`Failed to fetch todo ${id}:`, error);
                return null;
            } finally {
                this.loading = false;
            }
        },

        async createTodo(title: string, description: string) {
            this.loading = true;
            this.error = null;
            try {
                const response = await todoApi.create({ title, description });
                console.log(response);
                const newTodo = response.data.data;
                if (!Array.isArray(this.todos)) {
                    console.log('todos is not an array, initializing:', this.todos);
                    // 重新初始化 todos 为空数组
                    this.todos = [];
                }
                this.todos.push(newTodo);
                return newTodo;
            } catch (error: any) {
                this.error = error.message;
                console.error('Failed to create todo:', error);
                throw error;
            } finally {
                this.loading = false;
            }
        },

        async updateTodo(id: number, updates: { title?: string; description?: string; completed?: boolean }) {
            this.loading = true;
            this.error = null;
            try {
                const response = await todoApi.update(id, updates);
                const updatedTodo = response.data;

                // 更新本地状态
                const index = this.todos.findIndex(todo => todo.id === id);
                if (index !== -1) {
                    this.todos[index] = updatedTodo;
                }

                return updatedTodo;
            } catch (error: any) {
                this.error = error.message;
                console.error(`Failed to update todo ${id}:`, error);
                throw error;
            } finally {
                this.loading = false;
            }
        },

        async deleteTodo(id: number) {
            this.loading = true;
            this.error = null;
            try {
                await todoApi.delete(id);

                // 从本地状态中移除
                this.todos = this.todos.filter(todo => todo.id !== id);

                return true;
            } catch (error: any) {
                this.error = error.message;
                console.error(`Failed to delete todo ${id}:`, error);
                throw error;
            } finally {
                this.loading = false;
            }
        },

        async toggleTodoCompletion(id: number) {
            const todo = this.todos.find(t => t.id === id);
            if (!todo) return;

            return this.updateTodo(id, { completed: !todo.completed });
        }
    },
});