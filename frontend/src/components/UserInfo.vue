<script setup>
import { defineProps, inject, ref } from 'vue';
import LogoutButton from '@/components/LogoutButton.vue';
import UserListButton from '@/components/UserListButton.vue';

const axios = inject('$axios');
const message = ref('');
const cookies = inject('$cookies');
const mypage = ref(false)
const props = defineProps({
    id: String,
});
const followMessage = ref('');
const getUserInfo = () => {
  let id = props.id
  mypage.value = id == cookies.get('id')
  axios.get('/userInfo/' + id)
    .then(response => {
      message.value = response.data.userInfo.Username + '\'s Page';
    })
    .catch(error => {
      console.error(error);
      message.value = 'エラーが発生しました';
    });
}
const followRequest = async () => {
  try {
    const response = await axios.post('/follow', {'followee_id': parseInt(props.id)});
    followMessage.value = response.data.message;
  } catch (error) {
    followMessage.value = error.response.data.message;
  }
};

getUserInfo();
</script>

<template>
  <div class="btn">
    <LogoutButton />
    <UserListButton />
  </div>
  <div class="header">
    <h1>{{ message }}</h1>
    <button v-if="!mypage" class="follow-button" @click="followRequest()">
      Follow
    </button>
    <div>
      {{ followMessage }}
    </div>
  </div>
</template>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 5px;
  background-color: #eee;
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
</style>
