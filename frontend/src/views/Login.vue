<template>
  <div>
    <h1>Login Page</h1>
    <LoginForm v-on:login="handleLogin" />
  </div>
</template>
  
<script>
import LoginForm from '@/components/LoginForm.vue';

export default {
  name: 'Login',
  components: {
    LoginForm,
  },
  methods: {
    handleLogin(credentials) {
      console.log(credentials);
    },
  },
};
</script>