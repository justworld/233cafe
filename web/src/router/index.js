import Vue from 'vue'
import Router from 'vuerouter'
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
        // 首页
        {
          path: '/',
          name: 'home',
          components: {
            header: CommonHeader,
            content: () => import('@/views/HomePage'),
            footer: CommonFooter
          },
          meta: {
            menuIndex: 'home'
          }
        },
        {
          path: '/j',
          name: 'j',
          components: {
            header: CommonHeader,
            content: () => import('@/views/HomePage'),
            footer: CommonFooter
          },
          meta: {
            menuIndex: 'j'
          }
        },
        {
          path: '/k',
          name: 'k',
          components: {
            header: CommonHeader,
            content: () => import('@/views/HomePage'),
            footer: CommonFooter
          },
          meta: {
            menuIndex: 'k'
          }
        },
        {
          path: '/s',
          name: 's',
          components: {
            header: CommonHeader,
            content: () => import('@/views/HomePage'),
            footer: CommonFooter
          },
          meta: {
            menuIndex: 's'
          }
        },
        {
          path: '/t',
          name: 't',
          components: {
            header: CommonHeader,
            content: () => import('@/views/HomePage'),
            footer: CommonFooter
          },
          meta: {
            menuIndex: 't'
          }
        },
        {
          path: '/y',
          name: 'y',
          components: {
            header: CommonHeader,
            content: () => import('@/views/HomePage'),
            footer: CommonFooter
          },
          meta: {
            menuIndex: 'y'
          }
        },
        {
          path: 'info/:id',
          name: 'info',
          components: {
            header: CommonHeader,
            content: () => import('@/views/InfoPage'),
            footer: CommonFooter
          },
          meta: {
            menuIndex: 'home'
          }
        }
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
