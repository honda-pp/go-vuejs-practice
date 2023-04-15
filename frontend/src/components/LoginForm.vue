<script setup>
import { defineEmits } from 'vue';

const emit = defineEmits(['login']);

let username = '';
let password = '';

function submit() {
  const credentials = {
    username: username,
    password: password,
  };
  emit('login', credentials);
}
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
      <button type='submit'>Login</button>
    </form>
  </div>
</template>
