<script setup>
import { reactive, inject } from 'vue';
import router from '@/router';
import LogoutButton from '@/components/LogoutButton.vue';

const axios = inject('$axios');
const cookies = inject('$cookies');
const followMessage = reactive({
  message: '',
  id: -1
});
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

const goToUserPage = (id) => {
  router.push({ name: 'userPage' , params: { id } });
};

const followRequest = async (followeeId) => {
  followMessage.id = parseInt(followeeId);
  try {
    const response = await axios.post('/follow', {'followee_id': followeeId});
    followMessage.message = response.data.message;
  } catch (error) {
    followMessage.message = error.response.data.message;
  }
};

getUserList();
</script>

<template>
  <div class="container">
    <LogoutButton />
    <div class="user-list-container">
      <ul class="user-list">
        <li v-for="user in state.userList" :key="user.ID">
          <div class="user-info" @click="goToUserPage(user.ID)">
            {{ user.Username }}
          </div>
          <button v-if="cookies.get('id') != user.ID" class="follow-button" @click="followRequest(user.ID)">
            Follow
          </button>
          <div v-if="followMessage.id == user.ID" class="follow-message">
            {{ followMessage.message }}
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>
<style scoped>
.container {
  text-align: center;
  width: 50%;
  margin: 0 auto;
}

.user-list {
  list-style: none;
  padding: 0;
}

li {
  padding: 5px;
  padding-left: 5%;
  margin-bottom: 5px;
  text-align: left;
  background-color: #eee;
}

.user-info {
  cursor: pointer;
}

.follow-button {
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 5px;
  padding: 5px 10px;
  margin-left: 10px;
  cursor: pointer;
}

.follow-message {
  margin-top: 5px;
  font-weight: bold;
  color: #007bff;
}
</style>
