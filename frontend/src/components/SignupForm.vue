<template>
  <div>
    <h1>Sign up</h1>
    <form @submit.prevent="signup">
      <div>
      <label for="username">Username:</label>
      <input type="text" id="username" v-model="user.username" />
    </div>
    <div>
      <label for="email">Email:</label>
      <input type="email" id="email" v-model="user.email" />
    </div>
    <div>
      <label for="password">Password:</label>
      <input type="password" id="password" v-model="user.password" />
    </div>
    <button type="submit">Sign up</button>
  </form>
</div>
</template>
  
<script>
export default {
  name: 'SignupForm',
  data() {
    return {
      user: {
        username: '',
        email: '',
        password: '',
      },
    };
  },
  methods: {
    async signup() {
      this.$emit('signup', this.user);
    },
  },
};
</script>