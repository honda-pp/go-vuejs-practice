<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld :msg ="this.apiData" />
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import axios from 'axios'
export default {
  name: 'App',
  components: {
    HelloWorld
  },
  data() {
    return {
      apiData: this.apiData,
    };
  },
  mounted(){
    axios
      .get('http://localhost:8080/', { withCredentials: true })
      .then(response => {
        this.apiData = response.data["message"]
        console.log(response)
      })
      .catch((error) => {
        console.log(error)
      });
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
