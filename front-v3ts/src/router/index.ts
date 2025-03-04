import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    component: () => import('../views/Main_view.vue')
  },
  {
    path: '/search',
    name: "search",
    component: () => import('../views/Search_view.vue')
  },
  {
    path: '/admin',
    name: "admin",
    component: () => import('../views/Admin_view.vue')
  },
  {
    path: '/note',
    name: "note",
    component: () => import('../views/Note_view.vue')
  },
  {
    path : '/vedio/:vedioId/:episoNum',
    name : "vedio",
    component : () => import("../views/Vedio_view.vue")
  },
  {
    path: '/log',
    name: "log",
    component : () => import("../views/Log_view.vue")
  },
  {
    path: '/test',
    name: "test",
    component : () => import("../views/Test_view.vue")
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
