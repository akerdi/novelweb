import axios from './http'

export const search = params => { return axios.get(`/api/novel/search/${params.q}/${params.p}`)}
export const chapter = params => { return axios.get(`/api/novel/chapter/${params.md5}}`)}
export const content = params => { return axios.get(`/api/novel/search/${params.md5}/${params.p}`)}