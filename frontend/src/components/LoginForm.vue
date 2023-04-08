<template>
  <div>
    <h1>Login</h1>
    <form @submit.prevent='submit'>
    <div>
      <label for='username'>Username</label>
      <input type='text' id='username' v-model='username' />
    </div>
    <div>
      <label for='password'>Password</label>
      <input type='password' id='password' v-model='password' />
    </div>
    <button type='submit'>Login</button>
    </form>
  </div>
</template>
  
<script>
export default {
  data() {
    return {
      username: '',
      password: '',
    };
  },
  methods: {
    submit() {
      const credentials = {
        username: this.username,
        password: this.password,
      };
      this.$emit('login', credentials);
    },
  },
};
</script>

<style>
form {
  text-align: center;
}
form > div {
  display: flex;
  flex-direction: column;
  margin-bottom: 10px;
}
form > div label {
  margin-bottom: 5px;
}
form > div input {
  padding: 5px;
  border-radius: 3px;
  border: 1px solid #000000;
  width: 50%;
  min-width: 200px;
  margin: 0 auto;
}
</style>