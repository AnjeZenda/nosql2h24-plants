import { createRouter, createWebHistory } from 'vue-router'

const routes = [
    { path: '/', redirect: '/plants/start'},
    { path: '/plants/start', name: 'Start', component: () => import('./components/pages/StartPage.vue')},
    { path: '/plants/main', name: 'Main', component: () => import('./components/pages/MainPage.vue')}
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router