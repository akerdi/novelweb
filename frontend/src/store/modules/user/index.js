import _ from 'lodash'
import Vue from 'vue'
import Cookies from 'js-cookie'
import { getProfile } from '@/service/user'
import { USER_SIGNIN, USER_SIGNOUT, USER_INFO, USER_CHECK } from '../../mutation-types'
import getters from './getters'

const CookieKey = "novelserver"
const state = JSON.parse(Cookies.get(CookieKey) || '{}')
state.isCheck = false

export default {
  namespaced: true,
  state,
  getters,
  mutation: {

  },
  actions: {

  }
}