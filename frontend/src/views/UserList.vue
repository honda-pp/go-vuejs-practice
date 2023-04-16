<script setup>
import { reactive, inject } from 'vue';
import router from '@/router';

const cookies = inject('$cookies');
const axios = inject('$axios');
const state = reactive({
  userList: [],
});

const getUserList = async () => {
  try {
    const response = await axios.get('/userList');
    state.userList = JSON.parse(response.data.userList);

    console.log(state.userList[0])
  } catch (error) {
    console.error(error);
  }
};
getUserList();

const logout = () => {
  cookies.remove('id');
  cookies.remove('username');
  router.push({ name: 'login' });
}
</script>

<template>
  <button class="logout-btn" @click="logout">ログアウト</button>
  <div>
    <ul>
      <li v-for="user in state.userList" :key="user.Id">{{ user.Username }}</li>
    </ul>
  </div>
</template>

<style scoped>
.logout-btn {
  padding: 6px;
}
</style>
