import { createRouter, createWebHistory } from 'vue-router'
import AuthCallback from '@/AuthCallback.vue'

const routes = [
  {
    path: '/callback',
    name: 'Callback',
    component: AuthCallback
  }
]

const router = createRouter({
  history: createWebHistory('/'),
  routes
})

export default router
