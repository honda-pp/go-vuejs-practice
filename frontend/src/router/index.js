import axios from 'axios';
import { createRouter, createWebHistory } from 'vue-router';
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

router.beforeEach(async (to, from, next) => {
  try {
    const response = await axios.get('http://localhost:8080/api/isLogin', { withCredentials: true });
    const userId = response.data.id;
    const isLoggedIn = !!userId;
    if (!isLoggedIn && to.name !== 'login') {
      next({ name: 'login' });
    } else {
      next();
    }
  } catch (error) {
    console.error(error);
  }
});

export default router;
