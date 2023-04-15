<template>
  <div>
    <div class="header">
      <h1>{{ message }}</h1>
      <button class="logout-btn" @click="logout">ログアウト</button>
    </div>
  </div>
</template>

<script setup>
import { inject, ref } from 'vue';
import router from '../router';

const cookies = inject('$cookies');
const axios = inject('$axios');

const message = ref('');

const fetchMessage = () => {
  axios.get('/')
    .then(response => {
      message.value = response.data.message;
      console.log(response);
    })
    .catch(error => {
      console.error(error);
      message.value = 'エラーが発生しました';
    });
}
fetchMessage();

const logout = () => {
  cookies.remove('id');
  cookies.remove('username');
  router.push({ name: 'login' });
}
</script>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  background-color: #eee;
}
.logout-btn {
  padding: 6px;
}
</style>
