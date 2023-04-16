<script setup>
import { inject, ref } from 'vue';

const axiosInstance = inject('$axios');
const user = ref({
  username: '',
  email: '',
  password: '',
});
const signupMessage = ref('');

const signup = async () => {
  axiosInstance.post('/signup', user.value)
    .then(response => {
      signupMessage.value = response.data.message
    })
    .catch(error => {
      signupMessage.value = error.response.data.message;
    })
};
</script>

<template>
  <div>
    <h1>Sign up</h1>
    <form @submit.prevent="signup">
      <div>
      <label for="username">Username:</label>
      <input type="text" id="username" v-model="user.username" />
    </div>
    <div>
      <label for="email">Email:</label>
      <input type="email" id="email" v-model="user.email" />
    </div>
    <div>
      <label for="password">Password:</label>
      <input type="password" id="password" v-model="user.password" />
    </div>
    <div v-if="signupMessage" class='signup-message'>
        {{ signupMessage }}
      </div>
    <button type="submit">Sign up</button>
  </form>
  </div>
</template>
