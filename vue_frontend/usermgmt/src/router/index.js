import { createRouter, createWebHistory } from 'vue-router';
import RegisterUser from '@/components/RegisterUser.vue';
// import Login from '../components/Login.vue';
import ProfileUser from '@/components/ProfileUser.vue';
import ContactUser from '../components/ContactUser.vue';

const routes = [
  { path: '/register', component: RegisterUser },
//   { path: '/login', component: Login },
  { path: '/profile', component: ProfileUser },
  { path: '/contacts', component: ContactUser }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
