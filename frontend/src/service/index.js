import axios from './http'

export const search = params => { return axios.post(`/api/novel/search/${params.q}/${params.p}`)}
export const searchRecommand = params => { return axios.post(`/api/novel/searchRecommand/${params.q}`)}
export const chapter = params => { return axios.post(`/api/novel/chapter/${params.md5}`)}
export const content = params => { return axios.post(`/api/novel/content/${params.md5}/${params.index}`)}