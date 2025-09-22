<template>
  <div class="home-container">
    <header class="app-header">
      <h1>Todo App</h1>
      <div class="user-controls">
        <span v-if="authStore.user">Welcome, {{ authStore.user.username }}</span>
        <button class="logout-btn" @click="logout">Logout</button>
      </div>
    </header>
    
    <main class="main-content">
      <div class="todo-form-container">
        <h2>Add New Todo</h2>
        <form @submit.prevent="createTodo" class="todo-form">
          <div class="form-group">
            <label for="title">Title</label>
            <input 
              id="title"
              v-model="newTodo.title"
              type="text" 
              required
              placeholder="Enter todo title"
            />
          </div>
          
          <div class="form-group">
            <label for="description">Description</label>
            <textarea 
              id="description"
              v-model="newTodo.description"
              placeholder="Enter todo description"
              rows="3"
            ></textarea>
          </div>
          
          <button type="submit" class="submit-btn" :disabled="todoStore.loading">
            {{ todoStore.loading ? 'Adding...' : 'Add Todo' }}
          </button>
        </form>
      </div>
      
      <div class="todos-container">
        <div v-if="todoStore.loading && !todoStore.todos.length" class="loading">
          Loading todos...
        </div>
        
        <div v-else-if="todoStore.error" class="error-message">
          {{ todoStore.error }}
        </div>
        
        <div v-else-if="!todoStore.todos.length" class="empty-state">
          No todos yet. Add your first todo above!
        </div>
        
        <div v-else>
          <h2>My Todos</h2>
          
          <div class="filter-controls">
            <button 
              @click="filter = 'all'" 
              :class="{ active: filter === 'all' }"
            >
              All ({{ todoStore.todos.length }})
            </button>
            <button 
              @click="filter = 'active'" 
              :class="{ active: filter === 'active' }"
            >
              Active ({{ todoStore.pendingTodos.length }})
            </button>
            <button 
              @click="filter = 'completed'" 
              :class="{ active: filter === 'completed' }"
            >
              Completed ({{ todoStore.completedTodos.length }})
            </button>
          </div>
          
          <ul class="todo-list">
            <li v-for="todo in filteredTodos" :key="todo.id" class="todo-item">
              <div class="todo-content">
                <div class="todo-header">
                  <div class="todo-title" :class="{ completed: todo.completed }">
                    <input 
                      type="checkbox"
                      :checked="todo.completed"
                      @change="toggleTodo(todo.id)"
                    />
                    <span>{{ todo.title }}</span>
                  </div>
                  
                  <div class="todo-actions">
                    <button @click="editTodo(todo)" class="edit-btn">Edit</button>
                    <button @click="deleteTodo(todo.id)" class="delete-btn">Delete</button>
                  </div>
                </div>
                
                <p class="todo-description">{{ todo.description }}</p>
                
                <div class="todo-meta">
                  <span class="todo-date">Created: {{ formatDate(todo.created_at) }}</span>
                </div>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </main>
    
    <!-- 编辑 Todo 对话框 -->
    <div v-if="editingTodo" class="modal-overlay">
      <div class="modal-content">
        <h2>Edit Todo</h2>
        
        <form @submit.prevent="updateTodo">
          <div class="form-group">
            <label for="edit-title">Title</label>
            <input 
              id="edit-title"
              v-model="editingTodo.title"
              type="text" 
              required
            />
          </div>
          
          <div class="form-group">
            <label for="edit-description">Description</label>
            <textarea 
              id="edit-description"
              v-model="editingTodo.description"
              rows="3"
            ></textarea>
          </div>
          
          <div class="form-group checkbox-group">
            <label>
              <input 
                type="checkbox"
                v-model="editingTodo.completed"
              />
              Completed
            </label>
          </div>
          
          <div class="modal-actions">
            <button type="button" @click="cancelEdit" class="cancel-btn">Cancel</button>
            <button type="submit" class="submit-btn">Save Changes</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';
import { useTodoStore } from '@/stores/todo';
import type { Todo } from '@/types';

// 状态
const newTodo = ref({
  title: '',
  description: '',
});
const filter = ref('all');
const editingTodo = ref<Todo | null>(null);

// Store
const authStore = useAuthStore();
const todoStore = useTodoStore();
const router = useRouter();

// 计算属性
const filteredTodos = computed(() => {
  switch (filter.value) {
    case 'active':
      return todoStore.pendingTodos;
    case 'completed':
      return todoStore.completedTodos;
    default:
      return todoStore.todos;
  }
});

// 方法
const createTodo = async () => {
  try {
    await todoStore.createTodo(newTodo.value.title, newTodo.value.description);
    // 重置表单
    newTodo.value.title = '';
    newTodo.value.description = '';
  } catch (error) {
    console.error('Failed to create todo:', error);
  }
};

const toggleTodo = async (id: number) => {
  try {
    await todoStore.toggleTodoCompletion(id);
  } catch (error) {
    console.error('Failed to toggle todo:', error);
  }
};

const editTodo = (todo: Todo) => {
  editingTodo.value = { ...todo };
};

const cancelEdit = () => {
  editingTodo.value = null;
};

const updateTodo = async () => {
  if (!editingTodo.value) return;
  
  try {
    await todoStore.updateTodo(editingTodo.value.id, {
      title: editingTodo.value.title,
      description: editingTodo.value.description,
      completed: editingTodo.value.completed,
    });
    editingTodo.value = null;
  } catch (error) {
    console.error('Failed to update todo:', error);
  }
};

const deleteTodo = async (id: number) => {
  if (!confirm('Are you sure you want to delete this todo?')) return;
  
  try {
    await todoStore.deleteTodo(id);
  } catch (error) {
    console.error('Failed to delete todo:', error);
  }
};

const logout = async () => {
  await authStore.logout();
  router.push('/login');
};

const formatDate = (dateString: string) => {
  const date = new Date(dateString);
  return date.toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  });
};

// 生命周期钩子
onMounted(async () => {
  if (!authStore.isAuthenticated) {
    router.push('/login');
    return;
  }
  
  if (!authStore.user) {
    try {
      await authStore.fetchCurrentUser();
    } catch (error) {
      router.push('/login');
      return;
    }
  }
  
  await todoStore.fetchTodos();
});
</script>

<style scoped>
.home-container {
  max-width: 1000px;
  margin: 0 auto;
  padding: 2rem 1rem;
}

.app-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  padding-bottom: 1rem;
  border-bottom: 1px solid #eee;
}

.user-controls {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.logout-btn {
  padding: 0.5rem 1rem;
  background-color: #f44336;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.logout-btn:hover {
  background-color: #d32f2f;
}

.main-content {
  display: grid;
  grid-template-columns: 1fr 2fr;
  gap: 2rem;
}

@media (max-width: 768px) {
  .main-content {
    grid-template-columns: 1fr;
  }
}

.todo-form-container {
  padding: 1.5rem;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.todo-form {
  margin-top: 1rem;
}

.form-group {
  margin-bottom: 1rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  font-size: 0.9rem;
  color: #555;
}

input, textarea {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
}

textarea {
  resize: vertical;
}

.submit-btn {
  width: 100%;
  padding: 0.75rem;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
}

.submit-btn:hover {
  background-color: #45a049;
}

.submit-btn:disabled {
  background-color: #cccccc;
  cursor: not-allowed;
}

.todos-container {
  padding: 1.5rem;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.filter-controls {
  display: flex;
  margin-bottom: 1rem;
  gap: 0.5rem;
}

.filter-controls button {
  padding: 0.5rem 1rem;
  background-color: #f5f5f5;
  border: 1px solid #ddd;
  border-radius: 4px;
  cursor: pointer;
}

.filter-controls button.active {
  background-color: #e0e0e0;
  font-weight: bold;
}

.todo-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.todo-item {
  padding: 1rem;
  border: 1px solid #eee;
  border-radius: 4px;
  margin-bottom: 1rem;
}

.todo-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.5rem;
}

.todo-title {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-weight: bold;
}

.todo-title.completed span {
  text-decoration: line-through;
  color: #888;
}

.todo-actions {
  display: flex;
  gap: 0.5rem;
}

.edit-btn, .delete-btn {
  padding: 0.25rem 0.5rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.8rem;
}

.edit-btn {
  background-color: #2196f3;
  color: white;
}

.edit-btn:hover {
  background-color: #1976d2;
}

.delete-btn {
  background-color: #f44336;
  color: white;
}

.delete-btn:hover {
  background-color: #d32f2f;
}

.todo-description {
  margin: 0.5rem 0;
  color: #555;
}

.todo-meta {
  font-size: 0.8rem;
  color: #888;
  margin-top: 0.5rem;
}

.loading, .empty-state {
  text-align: center;
  padding: 2rem;
  color: #888;
}

.error-message {
  padding: 1rem;
  background-color: #ffebee;
  color: #c62828;
  border-radius: 4px;
  margin-bottom: 1rem;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 10;
}

.modal-content {
  background-color: white;
  padding: 2rem;
  border-radius: 8px;
  width: 90%;
  max-width: 500px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  margin-top: 1.5rem;
}

.cancel-btn {
  padding: 0.75rem 1.5rem;
  background-color: #f5f5f5;
  border: 1px solid #ddd;
  border-radius: 4px;
  cursor: pointer;
}

.checkbox-group {
  display: flex;
  align-items: center;
}

.checkbox-group label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  cursor: pointer;
}

.checkbox-group input {
  width: auto;
}
</style>