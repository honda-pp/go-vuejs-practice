<template>
  <h1>Login Page</h1>
  <div>
    <LoginForm v-if="!switchStatus" v-on:login="handleLogin" />
    <SignupForm v-else v-on:signup="handleSignup" />
    <button @click="switchForm">
      {{ switchStatus ? "Already have an account? " : "Don't have an account? " }}
      {{ buttonText[switchStatus] }}
    </button>
  </div>
</template>

<script>
import { inject, ref } from 'vue';
import LoginForm from '@/components/LoginForm.vue';
import SignupForm from '@/components/SignupForm.vue';
import router from '../router';

export default {
  name: 'Login',
  components: {
    LoginForm,
    SignupForm,
  },
  setup() {
    const axiosInstance = inject('$axios');
    const cookies = inject('$cookies');
    const switchStatus = ref(false);
    const buttonText = { false: 'Sign up', true: 'Login' };

    const handleLogin = async (credentials) => {
      try {
        const response = await axiosInstance.post('/login', credentials);
        cookies.set('id', response.data.id);
        cookies.set('username', response.data.username);
        router.push({ name: 'home' });
      } catch (error) {
        console.error(error);
      }
    };

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
    return {
      handleLogin,
      handleSignup,
      switchForm,
      switchStatus,
      buttonText,
    };
  },
};
</script>

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
