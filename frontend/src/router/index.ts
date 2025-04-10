import { createRouter, createWebHistory } from 'vue-router'
import LandingPage from '@/views/LandingPage.vue'

const routes = [
  { path: '/', component: LandingPage, name: 'LandingPage' },
  { path: '/loading', name: 'LoadingPage', component: () => import('@/views/LoadingPage.vue')},
  { path: '/result', name: 'ResultsPage', component: () => import('@/views/ResultsPage.vue')},
  { path: '/about', name: 'AboutPage', component: () => import('@/views/AboutPage.vue')},
  { path: '/contact', name: 'ContactPage', component: () => import('@/views/ContactPage.vue')},
  { path: '/error', name: 'ErrorPage', component: () => import('@/views/ErrorPage.vue')},
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router;
