<template>
  <div>
    <h1>{{ message }}</h1>
  </div>
</template>

<script>
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
      this.$axios.get('/')
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