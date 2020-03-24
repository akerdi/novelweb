import axios from 'axios'
import { Message } from 'element-ui'
import Store from '@/store'
import { USER_SIGNOUT, USER_CHECK } from '@/store/mutation-types'
import router from '../router'

axios.defaults.timeout = 10000
axios.defaults.router = 3
axios.defaults.retryDelay = 800

if (process.env.NODE_ENV === "development") {
  axios.defaults.baseURL = "/api"
}

axios.interceptors.request.use(
  config => {
    return config
  },
  error => {
    Message.error({message: error.response.data || "加载失败"})
    return Promise.reject(error.response.data)
  }
)
axios.interceptors.response.use(
  response => {
    return response
  },
  error => {
    if (error.response) {
      switch (error.response.status) {
        case 401:
        case 403:
          if (error.response.data === 'Login Required') {
            Store.commit('user/' + USER_SIGNOUT)
            Store.commit('user/' + USER_CHECK, true)
            if (router.history.current.path !== '/') {
              return router.push({ name: 'login' })
            } else {
              return Promise.reject(error)
            }
          }
          if (router.history.current.path === '/' || router.history.current.path === '/oauth/authorize') {
            return Message.error(error.response.data)
          }
      }
    }
    let isCheck = Store.getters['user/isCheck']
    if (isCheck) Message.error({message: error.response.data || '请求失败'})
    return Promise.reject(error)
  }
)

export default axios