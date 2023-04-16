<script setup>
import { ref, inject } from 'vue';
import router from '@/router';

const axiosInstance = inject('$axios');
const cookies = inject('$cookies');
const username = ref('');
const password = ref('');
const errorMessage = ref('');

const submit = () => {
  const credentials = {
    username: username.value,
    password: password.value,
  };
  axiosInstance.post('/login', credentials)
    .then(response => {
      cookies.set('id', response.data.id);
      cookies.set('username', response.data.username);
      router.push({ name: 'home' });
    })
    .catch(error => {
      errorMessage.value = error.response.data.message;
    });
};
</script>

<template>
  <div>
    <h1>Login</h1>
    <form @submit.prevent='submit'>
      <div>
        <label for='username'>Username</label>
        <input v-model='username' type='text' id='username' />
      </div>
      <div>
        <label for='password'>Password</label>
        <input v-model='password' type='password' id='password' />
      </div>
      <div v-if="errorMessage" class='error-message'>
        {{ errorMessage }}
      </div>
      <button type='submit'>Login</button>
    </form>
  </div>
</template>
