<template>
  <div>
    <h1>Login Page</h1>
    <LoginForm v-on:login="handleLogin" />
  </div>
</template>
  
<script>
import LoginForm from '@/components/LoginForm.vue';
import axios from 'axios'

export default {
  name: 'Login',
  components: {
    LoginForm,
  },
  methods: {
    handleLogin(credentials) {
      axios.post('http://localhost:8080/api/login', credentials)
        .then(response => {
          console.log(response.data);
        })
        .catch(error => {
          console.error(error)
        })
    }
  }
};
</script>