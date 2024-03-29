import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../views/Home.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import(/* webpackChunkName: "login" */ '../views/Login.vue'),
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import(/* webpackChunkName: "register" */ '../views/Register.vue'),
  },
  {
    path: '/upload',
    name: 'Upload',
    component: () => import(/* webpackChunkName: "upload" */ '../views/Upload.vue'),
  },
  {
    path: '/watch/:slug',
    name: 'Watch',
    component: () => import(/* webpackChunkName: "view" */ '../views/Watch.vue'),
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
