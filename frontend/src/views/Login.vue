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
      this.$axios.post('/login', credentials)
        .then(response => {
          console.log(response);
          //this.$router.push({ name: 'home' });
        })
        .catch(error => {
          console.error(error);
        });
    },
  },
};
</script>
