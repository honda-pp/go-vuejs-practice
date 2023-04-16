<script setup>
import { reactive, inject } from 'vue';
import LogoutButton from '@/components/LogoutButton.vue';

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

</script>

<template>
  <div>
    <LogoutButton />
    <div>
      <ul>
        <li v-for="user in state.userList" :key="user.Id">{{ user.Username }}</li>
      </ul>
    </div>
  </div>
</template>

