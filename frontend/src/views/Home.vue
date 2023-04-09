<template>
  <div>
    <h1>{{ message }}</h1>
  </div>
</template>

<script>
import { inject } from 'vue';

export default {
  name: 'Home',
  data() {
    return {
      message: '',
    }
  },
  mounted() {
    this.fetchMessage()
  },
  methods: {
    fetchMessage() {
      inject('$axios').get('/')
        .then(response => {
          this.message = response.data.message
        })
        .catch(error => {
          console.error(error)
          this.message = 'エラーが発生しました'
        })
    }
  }
}
</script>