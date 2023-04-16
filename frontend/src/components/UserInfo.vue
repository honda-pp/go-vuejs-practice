<script setup>
import { defineProps, inject, ref } from 'vue';
import LogoutButton from '@/components/LogoutButton.vue';

const axios = inject('$axios');
const message = ref('');
const cookies = inject('$cookies');
const props = defineProps({
    id: String,
});

const getUserInfo = () => {
  let id = props.id
  if (id == null) {
    id = cookies.get('id')
  }
  axios.get('/userInfo/' + id)
    .then(response => {
      message.value = response.data.userInfo.Username + '\'s Page';
    })
    .catch(error => {
      console.error(error);
      message.value = 'エラーが発生しました';
    });
}

getUserInfo();
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
