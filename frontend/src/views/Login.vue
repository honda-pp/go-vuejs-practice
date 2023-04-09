<template>
  <div>
    <h1>Login Page</h1>
    <LoginForm v-on:login="handleLogin" />
  </div>
</template>
  
<script>
import { inject } from 'vue';
import LoginForm from '@/components/LoginForm.vue';

export default {
  name: 'Login',
  components: {
    LoginForm,
  },
  setup() {
    const axios = inject('$axios');

    const handleLogin = async (credentials) => {
      console.log("handleLogin")
      try {
        console.log(credentials)
        const response = await axios.post('/login', credentials);
        console.log(response.data);
        //this.$router.push({ name: 'home' });
        //password123
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
