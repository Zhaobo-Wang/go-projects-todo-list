<template>
  <div class="todo-detail-container">
    <div class="back-link">
      <router-link to="/">&larr; Back to Todos</router-link>
    </div>
    
    <div v-if="loading" class="loading">
      Loading todo details...
    </div>
    
    <div v-else-if="error" class="error-message">
      {{ error }}
    </div>
    
    <div v-else-if="!todo" class="not-found">
      Todo not found.
    </div>
    
    <div v-else class="todo-detail-card">
      <div class="todo-header">
        <h1 :class="{ completed: todo.completed }">{{ todo.title }}</h1>
        <div class="todo-status">
          <span :class="['status-badge', todo.completed ? 'completed' : 'pending']">
            {{ todo.completed ? 'Completed' : 'Pending' }}
          </span>
        </div>
      </div>
      
      <div class="todo-content">
        <div class="todo-description">
          <h3>Description</h3>
          <p>{{ todo.description || 'No description provided.' }}</p>
        </div>
        
        <div class="todo-meta">
          <div class="meta-item">
            <strong>Created:</strong> {{ formatDate(todo.created_at) }}
          </div>
          <div class="meta-item">
            <strong>Last Updated:</strong> {{ formatDate(todo.updated_at) }}
          </div>
        </div>
      </div>
      
      <div class="todo-actions">
        <button 
          @click="toggleTodoCompletion" 
          class="toggle-btn"
          :class="todo.completed ? 'mark-incomplete' : 'mark-complete'"
        >
          {{ todo.completed ? 'Mark as Incomplete' : 'Mark as Complete' }}
        </button>
        <button @click="editTodo" class="edit-btn">Edit</button>
        <button @click="deleteTodo" class="delete-btn">Delete</button>
      </div>
    </div>
    
    <!-- 编辑 Todo 对话框 -->
    <div v-if="isEditing" class="modal-overlay">
      <div class="modal-content">
        <h2>Edit Todo</h2>
        
        <form @submit.prevent="updateTodo">
          <div class="form-group">
            <label for="edit-title">Title</label>
            <input 
              id="edit-title"
              v-model="editForm.title"
              type="text" 
              required
            />
          </div>
          
          <div class="form-group">
            <label for="edit-description">Description</label>
            <textarea 
              id="edit-description"
              v-model="editForm.description"
              rows="3"
            ></textarea>
          </div>
          
          <div class="form-group checkbox-group">
            <label>
              <input 
                type="checkbox"
                v-model="editForm.completed"
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
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useTodoStore } from '@/stores/todo';
import type { Todo } from '@/types';

// 状态
const todo = ref<Todo | null>(null);
const loading = ref(true);
const error = ref('');
const isEditing = ref(false);
const editForm = ref({
  title: '',
  description: '',
  completed: false,
});

// 服务
const route = useRoute();
const router = useRouter();
const todoStore = useTodoStore();

// 方法
const fetchTodo = async () => {
  loading.value = true;
  error.value = '';
  
  try {
    const id = parseInt(route.params.id as string);
    if (isNaN(id)) {
      error.value = 'Invalid todo ID';
      return;
    }
    
    const fetchedTodo = await todoStore.fetchTodo(id);
    todo.value = fetchedTodo;
    
    if (!todo.value) {
      error.value = 'Todo not found';
    }
  } catch (err: any) {
    error.value = err.message || 'Failed to load todo';
  } finally {
    loading.value = false;
  }
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

const toggleTodoCompletion = async () => {
  if (!todo.value) return;
  
  try {
    await todoStore.toggleTodoCompletion(todo.value.id);
    // 更新本地状态
    todo.value.completed = !todo.value.completed;
  } catch (err: any) {
    error.value = err.message || 'Failed to update todo status';
  }
};

const editTodo = () => {
  if (!todo.value) return;
  
  editForm.value = {
    title: todo.value.title,
    description: todo.value.description,
    completed: todo.value.completed,
  };
  
  isEditing.value = true;
};

const cancelEdit = () => {
  isEditing.value = false;
};

const updateTodo = async () => {
  if (!todo.value) return;
  
  try {
    const updatedTodo = await todoStore.updateTodo(todo.value.id, {
      title: editForm.value.title,
      description: editForm.value.description,
      completed: editForm.value.completed,
    });
    
    // 更新本地状态
    if (updatedTodo) {
      todo.value = updatedTodo;
    }
    
    isEditing.value = false;
  } catch (err: any) {
    error.value = err.message || 'Failed to update todo';
  }
};

const deleteTodo = async () => {
  if (!todo.value) return;
  
  if (!confirm('Are you sure you want to delete this todo?')) return;
  
  try {
    await todoStore.deleteTodo(todo.value.id);
    router.push('/');
  } catch (err: any) {
    error.value = err.message || 'Failed to delete todo';
  }
};

// 生命周期钩子
onMounted(() => {
  fetchTodo();
});
</script>

<style scoped>
.todo-detail-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem 1rem;
}

.back-link {
  margin-bottom: 1.5rem;
}

.back-link a {
  color: #2196f3;
  text-decoration: none;
  display: inline-flex;
  align-items: center;
}

.back-link a:hover {
  text-decoration: underline;
}

.loading, .not-found {
  text-align: center;
  padding: 3rem;
  color: #888;
}

.error-message {
  padding: 1rem;
  background-color: #ffebee;
  color: #c62828;
  border-radius: 4px;
  margin-bottom: 1rem;
}

.todo-detail-card {
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  padding: 2rem;
}

.todo-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
  padding-bottom: 1rem;
  border-bottom: 1px solid #eee;
}

.todo-header h1 {
  margin: 0;
  font-size: 1.8rem;
}

.todo-header h1.completed {
  text-decoration: line-through;
  color: #888;
}

.status-badge {
  padding: 0.25rem 0.75rem;
  border-radius: 16px;
  font-size: 0.8rem;
  font-weight: bold;
}

.status-badge.completed {
  background-color: #e8f5e9;
  color: #2e7d32;
}

.status-badge.pending {
  background-color: #fff8e1;
  color: #ff8f00;
}

.todo-content {
  margin-bottom: 2rem;
}

.todo-description h3 {
  margin-top: 0;
  margin-bottom: 0.5rem;
  color: #555;
}

.todo-description p {
  margin-top: 0;
  line-height: 1.5;
}

.todo-meta {
  margin-top: 1.5rem;
  padding-top: 1rem;
  border-top: 1px dashed #eee;
}

.meta-item {
  margin-bottom: 0.5rem;
  color: #666;
  font-size: 0.9rem;
}

.todo-actions {
  display: flex;
  gap: 1rem;
}

.todo-actions button {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
}

.toggle-btn.mark-complete {
  background-color: #4caf50;
  color: white;
}

.toggle-btn.mark-complete:hover {
  background-color: #45a049;
}

.toggle-btn.mark-incomplete {
  background-color: #ff9800;
  color: white;
}

.toggle-btn.mark-incomplete:hover {
  background-color: #f57c00;
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

/* 模态框样式继承自 HomeView */
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

.submit-btn {
  padding: 0.75rem 1.5rem;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.submit-btn:hover {
  background-color: #45a049;
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