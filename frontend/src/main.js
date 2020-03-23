// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import Vuex from 'vuex'
import store from './store'
import ElementUI from 'element-ui'

import '@/assets/styles/element-variables.scss'

import '@/lib/filters'
import 'font-awesome/css/font-awesome.min.css'
import '@/assets/styles/common.scss'
import '@/assets/fonts/iconfont.css'

Vue.config.productionTip = false

Vue.use(ElementUI)
Vue.use(Vuex)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
