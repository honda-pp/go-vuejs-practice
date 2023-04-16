<script setup>
import { inject } from 'vue';
import router from '@/router';

const cookies = inject('$cookies');

const logout = () => {
  cookies.remove('id');
  cookies.remove('username');
  router.push({ name: 'login' });
}
</script>

<template>
  <button class="logout-btn" @click="logout">Logout</button>
</template>

<style scoped>
.logout-btn {
  position: absolute;
  top: 10px;
  right: 10px;
  padding: 6px;
}
</style>
