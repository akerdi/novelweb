import Cookies from 'js-cookie'
import { userRoles, accessLevels } from './routingConfig'

const roleTitle = (user) => {
  if (!(user && user.role && user.role.title)) { return 'anon' }
  return user.role.title
}

const authorize = (user, accessLevel) => {
  let ref
  return accessLevel.bitMask & ((ref = userRoles[roleTitle(user)]) != null ? ref.bitMask : void 0)
}

const isLoggedIn = (user) => {
  return roleTitle(user) !== 'anon'
}

const isAdmin = (user) => {
  return roleTitle(user) === 'admin'
}

const roleAttr = (user) => {
  return userRoles[roleTitle(user)]
}

export default {
  access: () => {
    return accessLevels
  },
  uid: (state) => {
    return state._id
  },
  roleTitle: (state) => {
    return roleTitle(state)
  },
  authorize: (state) => (accessLevel) => {
    return authorize(state, accessLevel)
  },
  isLoggedIn: (state) => {
    return isLoggedIn(state)
  },
  isAdmin: (state) => {
    return isAdmin(state)
  },
  roleAttr: (state) => {
    return roleAttr(state)
  },
  isCheck: (state) => {
    return !!state.isCheck
  },
  getXSRFTOKEN: () => {
    let token = Cookies.get('XSRF-TOKEN')
    return token
  },
  getCookie: () => {
    return Cookies.get('')
  }
}
