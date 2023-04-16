import { createRouter, createWebHistory } from 'vue-router';
import VueCookies from 'vue-cookies';
import Login from '../views/Login.vue';
import Home from '../views/Home.vue';

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
      meta: { requiresAuth: true },
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
      meta: { requiresAuth: false },
    },
  ],
});

router.beforeEach((to, from, next) => {
  const isLoggedIn = VueCookies.get('id') != null && VueCookies.get('username') != null;
  if (!isLoggedIn && to.name !== 'login') {
    next({ name: 'login' })
  } else {
    next()
  }
});

export default router;
