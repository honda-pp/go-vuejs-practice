import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import axios from 'axios';

const app = createApp(App);

app.provide('$axios', axios.create({
  baseURL: 'http://localhost:8080/api',
  withCredentials: true,
}));

app.use(router);
app.mount('#app');
