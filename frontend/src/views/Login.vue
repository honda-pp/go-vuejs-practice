<script setup>
import { inject, ref } from 'vue';
import LoginForm from '@/components/LoginForm.vue';
import SignupForm from '@/components/SignupForm.vue';

const axiosInstance = inject('$axios');
const switchStatus = ref(false);
const buttonText = { false: 'Don\'t have an account? Sign up', true: 'Already have an account? Login' };

const handleSignup = async (user) => {
  try {
    await axiosInstance.post('/signup', user);
  } catch (error) {
    console.error(error);
  }
};

const switchForm = async () => {
  switchStatus.value = !switchStatus.value;
};

</script>

<template>
  <h1>Login Page</h1>
  <div>
    <LoginForm v-if='!switchStatus'></LoginForm>
    <SignupForm v-else @signup='handleSignup' />
    <button @click='switchForm'>
      {{ buttonText[switchStatus] }}
    </button>
  </div>
</template>

<style>
  form {
    text-align: center;
  }
  form > div {
    display: flex;
    flex-direction: column;
    margin-bottom: 10px;
  }
  form > div label {
    margin-bottom: 5px;
  }
  form > div input {
    padding: 5px;
    border-radius: 3px;
    border: 1px solid #000000;
    width: 200px;
    margin: 0 auto;
  }
  button {
    margin-top: 10px;
  }
</style>
