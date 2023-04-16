<script setup>
import { inject, ref } from 'vue';
import LogoutButton from '@/components/LogoutButton.vue';

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
</script>

<template>
  <div>
    <LogoutButton />
    <div class="header">
      <h1>{{ message }}</h1>
    </div>
  </div>
</template>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  background-color: #eee;
}
</style>
