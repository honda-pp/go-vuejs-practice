<script setup>
import { reactive, inject } from 'vue';
import router from '@/router';
import LogoutButton from '@/components/LogoutButton.vue';

const axios = inject('$axios');
const state = reactive({
  userList: [],
});

const getUserList = async () => {
  try {
    const response = await axios.get('/userList');
    state.userList = JSON.parse(response.data.userList);
  } catch (error) {
    console.error(error);
  }
};
getUserList();


const goToUserPage = (id) => {
  router.push({ name: 'userPage' , params: { id } });
};

</script>

<template>
  <div>
    <LogoutButton />
    <div>
      <ul>
        <li v-for="user in state.userList" :key="user.ID" @click="goToUserPage(user.ID)">
          {{ user.Username }}
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
ul {
  list-style: none;
  padding: 0;
  text-align: center;
  width: 50%;
  margin: 0 auto;
}
li {
  padding: 5px;
  padding-left: 5%;
  margin-bottom: 5px;
  text-align: left;
  background-color: #eee;
}

</style>