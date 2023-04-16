<script setup>
import router from '@/router';

const userList = () => {
  router.push({ name: 'userList' });
}
</script>

<template>
  <button class="userList-btn" @click="userList">User List</button>
</template>

<style scoped>
.userList-btn {
  position: absolute;
  top: 10px;
  right: 70px;
  padding: 6px;
}
</style>
