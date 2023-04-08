<template>
  <div>
    <h1>{{ message }}</h1>
  </div>
</template>

<script>
import axios from 'axios'

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
      axios.get('http://localhost:8080/api/', { withCredentials: true })
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