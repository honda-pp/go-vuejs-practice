<template>
  <div>
    <h1>Login</h1>
    <form @submit.prevent='submit'>
      <div>
        <label for='username'>Username</label>
        <input v-model='username' type='text' id='username' />
      </div>
      <div>
        <label for='password'>Password</label>
        <input v-model='password' type='password' id='password' />
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