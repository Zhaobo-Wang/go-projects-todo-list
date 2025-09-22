<template>
  <div class="login-container">
    <div class="login-card">
      <h1>{{ isRegister ? 'Register' : 'Login' }}</h1>
      
      <div v-if="error" class="error-message">
        {{ error }}
      </div>
      
      <form @submit.prevent="submitForm">
        <div class="form-group" v-if="isRegister">
          <label for="username">Username</label>
          <input 
            id="username"
            v-model="username"
            type="text" 
            required
            placeholder="Enter username"
          />
        </div>
        
        <div class="form-group" v-if="isRegister">
          <label for="email">Email</label>
          <input 
            id="email"
            v-model="email"
            type="email" 
            required
            placeholder="Enter email"
          />
        </div>
        
        <div class="form-group">
          <label for="login-username">{{ isRegister ? 'Username' : 'Username or Email' }}</label>
          <input 
            id="login-username"
            v-model="loginUsername"
            type="text" 
            required
            placeholder="Enter username or email"
          />
        </div>
        
        <div class="form-group">
          <label for="password">Password</label>
          <input 
            id="password"
            v-model="password"
            type="password" 
            required
            placeholder="Enter password"
          />
        </div>
        
        <button type="submit" class="submit-btn" :disabled="loading">
          {{ loading ? 'Processing...' : (isRegister ? 'Register' : 'Login') }}
        </button>
      </form>
      
      <div class="toggle-form">
        {{ isRegister ? 'Already have an account?' : "Don't have an account?" }}
        <a href="#" @click.prevent="toggleForm">
          {{ isRegister ? 'Login' : 'Register' }}
        </a>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';

// 状态
const isRegister = ref(false);
const username = ref('');
const email = ref('');
const loginUsername = ref('');
const password = ref('');
const error = ref('');
const loading = ref(false);

// 服务
const authStore = useAuthStore();
const router = useRouter();

// 方法
const toggleForm = () => {
  isRegister.value = !isRegister.value;
  error.value = '';
};

const submitForm = async () => {
  error.value = '';
  loading.value = true;
  
  try {
    if (isRegister.value) {
      await authStore.register(username.value, email.value, password.value);
      // 注册后自动登录
      await authStore.login(username.value, password.value);
    } else {
      await authStore.login(loginUsername.value, password.value);
    }
    
    // 登录成功，重定向到首页
    router.push('/');
  } catch (err: any) {
    error.value = err.message;
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f5f5f5;
}

.login-card {
  width: 100%;
  max-width: 400px;
  padding: 2rem;
  background: white;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

h1 {
  margin-bottom: 1.5rem;
  color: #333;
  text-align: center;
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

input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
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
  margin-top: 1rem;
}

.submit-btn:hover {
  background-color: #45a049;
}

.submit-btn:disabled {
  background-color: #cccccc;
  cursor: not-allowed;
}

.error-message {
  padding: 0.75rem;
  margin-bottom: 1rem;
  background-color: #ffebee;
  color: #c62828;
  border-radius: 4px;
  font-size: 0.9rem;
}

.toggle-form {
  margin-top: 1.5rem;
  text-align: center;
  font-size: 0.9rem;
}

.toggle-form a {
  color: #2196f3;
  text-decoration: none;
  margin-left: 0.25rem;
}

.toggle-form a:hover {
  text-decoration: underline;
}
</style>