<template>
  <div>
    <h1>Login Page</h1>
    <LoginForm v-on:login="handleLogin" />
  </div>
</template>

<script>
import { inject } from 'vue';
import LoginForm from '@/components/LoginForm.vue';
import router from '../router';

export default {
  name: 'Login',
  components: {
    LoginForm,
  },
  setup() {
    const axiosInstance = inject('$axios');
    const $cookies = inject('$cookies');

    const handleLogin = async (credentials) => {
      try {
        const response = await axiosInstance.post('/login', credentials);
        $cookies.set('id', response.data.id);
        $cookies.set('username', response.data.username);
        router.push({ name: 'home' });
      } catch (error) {
        console.error(error);
      }
    };

    return {
      handleLogin,
    };
  },
};
</script>
