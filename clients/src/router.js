import { createRouter, createWebHashHistory } from 'vue-router'
import Home from './pages/Home.vue'

const Login = () => import('./pages/Login.vue')
const Register = () => import('./pages/Register.vue')
const routes = [
  {
    path: '/',
    component: Home,
  },
  {
    path: '/login',
    component: Login,
  },
  {
    path: '/register',
    component: Register,
  },
]

const router = createRouter({ routes, history: createWebHashHistory() })

export default router
