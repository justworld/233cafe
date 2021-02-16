import ELEMENT from 'elementui'
import Vue from 'vue'
import Vant from 'vant'
import router from '@/router'
import App from './App.vue'
import store from '@/store'
import constants from '@/common/constant'
import api from '@/api'
import tool from '@/utils/tool'

Vue.use(ELEMENT)
Vue.use(Vant)

Vue.prototype.$constants = constants // constants
Vue.prototype.$api = api // api
Vue.prototype.$tool = tool // tool
Vue.prototype.$set = Vue.set // set

let vm = new Vue({
  router,
  store,
  el: '#app',
  render: h => h(App)
})

Vue.use({
  vm
})
