import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import './assets/main.css'
import App from './App.vue'
import Home from './views/Home.vue'
import Features from './views/Features.vue'
import Pricing from './views/Pricing.vue'
import Docs from './views/Docs.vue'

const routes = [
  { path: '/', name: 'Home', component: Home },
  { path: '/features', name: 'Features', component: Features },
  { path: '/pricing', name: 'Pricing', component: Pricing },
  { path: '/docs', name: 'Docs', component: Docs },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  }
})

const app = createApp(App)
app.use(router)
app.mount('#app')
