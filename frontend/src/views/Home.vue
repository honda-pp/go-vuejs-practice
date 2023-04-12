<template>
  <div>
    <div class="header">
      <h1>{{ message }}</h1>
      <button class="logout-btn" @click="logout">ログアウト</button>
    </div>
  </div>
</template>

<script>
import { inject } from 'vue';
import router from '../router';

export default {
  name: 'Home',
  setup() {
    const $cookies = inject('$cookies');
    return {
      $cookies,
    };
  },
  data() {
    return {
      message: '',
    }
  },
  mounted() {
    this.fetchMessage();
  },
  methods: {
    fetchMessage() {
      inject('$axios').get('/')
        .then(response => {
          this.message = response.data.message;
        })
        .catch(error => {
          console.error(error);
          this.message = 'エラーが発生しました';
        });
    },
    logout() {
      this.$cookies.remove('id');
      this.$cookies.remove('username');
      router.push({ name: 'login' });
    }
  }
}
</script>
