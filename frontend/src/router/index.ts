import { createRouter, createWebHistory } from 'vue-router'
import LoadingPage from '@/views/LoadingPage.vue'
import ResultsPage from '@/views/ResultsPage.vue'
import LandingPage from '@/views/LandingPage.vue'
import ContactPage from '@/views/ContactPage.vue'
import AboutPage from '@/views/AboutPage.vue'
import ErrorPage from '@/views/ErrorPage.vue'

const routes = [
  { path: '/', component: LandingPage, name: 'LandingPage' },
  { path: '/loading', component: LoadingPage, name: 'LoadingPage' },
  { path: '/result', component: ResultsPage, name: 'ResultsPage' },
  { path: '/about', component: AboutPage, name: 'AboutPage' },
  { path: '/contact', component: ContactPage, name: 'ContactPage' },
  { path: '/error', component: ErrorPage, name: 'ErrorPage' },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router;
