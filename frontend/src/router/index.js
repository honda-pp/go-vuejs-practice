import { createRouter, createWebHistory } from 'vue-router';
import VueCookies from 'vue-cookies';
import Login from '@/views/Login.vue';
import UserList from '@/views/UserList.vue';
import UserPage from '@/views/UserPage.vue';

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: UserPage,
      meta: { requiresAuth: true },
    },
    {
      path: '/userPage/:id',
      name: 'userPage',
      component: UserPage,
      meta: { requiresAuth: true },
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
      meta: { requiresAuth: false },
    },
    {
      path: '/userList',
      name: 'userList',
      component: UserList,
      meta: { requiresAuth: true },
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
