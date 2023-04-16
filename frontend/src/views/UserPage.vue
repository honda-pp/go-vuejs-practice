<script setup>
import { inject, ref } from 'vue';
import { useRoute } from 'vue-router';
import LogoutButton from '@/components/LogoutButton.vue';

const axios = inject('$axios');
const message = ref('');
const route = useRoute();
let id = route.params.id;

const getUserInfo = (id) => {
  axios.get('/userInfo/' + id)
    .then(response => {
      message.value = response.data.userInfo.Username + '\'s Page';
    })
    .catch(error => {
      console.error(error);
      message.value = 'エラーが発生しました';
    });
}

getUserInfo(id);
</script>

<template>
  <div>
    <div class="header">
      <h1>{{ message }}</h1>
    </div>
    <LogoutButton />
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
</style>
