const config = {
  accessLevels: {
    'admin': 1,
    'editor': 1 << 5,
    'vip2': 1 << 10,
    'vip1': 1 << 11,
    'user': 1 << 28,
    'anon': 1 << 29,
    'public': 1 << 30
  },
  roles: {
    'anon': {
      accessLevels: ['anon']
    },
    'user': {
      accessLevels: ['public', 'user']
    },
    'vip1': {
      accessLevels: ['public', 'user', 'vip1']
    },
    'vip2': {
      accessLevels: ['public', 'user', 'vip1', 'vip2']
    },
    'editor': {
      accessLevels: ['public', 'user', 'vip1', 'vip2', 'editor']
    },
    'admin': {
      accessLevels: ['public', 'user', 'vip1', 'vip2', 'editor', 'admin']
    }
  }
}

const buildRoles = () => {
  let roleAccess = {}
  for (let role in config.roles) {
    let bitMask = 0
    let roleAttr = config.roles[role]
    for (let index in roleAttr.accessLevels) {
      let level = roleAttr.accessLevels[index]
      if (config.accessLevels.hasOwnProperty(level)) {
        bitMask = bitMask | config.accessLevels[level]
      } else {
        console.log('Access level ' + level + 'not found!')
      }
    }
    bitMask = bitMask | config.accessLevels['public']
    roleAccess[role] = {
      title: role,
      bitMask: bitMask
    }
  }
  return roleAccess
}

const buildAccessLevels = () => {
  let accessLevels = {}
  for (let level in config.accessLevels) {
    accessLevels[level] = {
      level: level,
      bitMask: config.accessLevels[level]
    }
  }
  return accessLevels
}

export const userRoles = buildRoles()
export const accessLevels = buildAccessLevels()
