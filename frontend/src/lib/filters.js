import Vue from 'vue'
import moment from 'moment'

Vue.filter('timestamp', (msec, formatString) => {
  if (!msec) {
    return ''
  }
  formatString = formatString || 'YYYY-MM-DD HH:mm'
  return moment(msec).format(formatString)
})

export default {}