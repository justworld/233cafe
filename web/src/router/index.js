import Vue from 'vue'
import Router from 'vuerouter'
import { MessageBox } from 'elementui'
import store from '@/store'
import IndexPage from '@/views/IndexPage'

import CommonHeader from '@/components/headers/CommonHeader'
import CommonFooter from '@/components/footers/CommonFooter'

Vue.use(Router)

let router = new Router({
  fallback: false,
  mode: 'history',
  scrollBehavior: () => ({ y: 0 }),
  routes: [
    {
      path: '/',
      name: 'index',
      component: IndexPage,
      children: [
      ]
    }
  ]
})

router.beforeEach((to, from, next) => {
  next()
})

router.afterEach((to, from, next) => {
})

export default router
