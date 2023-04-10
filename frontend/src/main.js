import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import axios from 'axios';
import VueCookies from 'vue-cookies';

const app = createApp(App);

const axiosInstance = axios.create({
  baseURL: 'http://localhost:8080/api',
  withCredentials: true,
});

app.use(VueCookies);
app.provide('$axios', axiosInstance);
app.use(router);

app.mount('#app');
