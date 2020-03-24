import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/pages/home'
import Search from '@/pages/search'
import Chapter from '@/pages/chapter'
import Content from '@/pages/content'

Vue.use(Router)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/search',
    name: 'Search',
    component: Search
  },
  {
    path: "/chapter",
    name: "Chapter",
    component: Chapter
  },
  {
    path: "/content",
    name: "Content",
    component: Content
  }
]

const router = new Router({
  routes,
  mode: 'history',
  strict: process.env.NODE_ENV !== "production"
})

router.beforeEach((to, from, next) => {
  (async() => {
    try {
      const { meta } = to
      next()
    } catch (error) {
      next()
    }
  })()
})

export default router
