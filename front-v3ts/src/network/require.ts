import axios, { AxiosInstance } from 'axios'

// 必须以 '/' 结尾
// export const baseURL = "http://127.0.0.1:8888"

export const baseURL = "http://220.167.104.236:21266"

// 创建 axios 实例
const request: AxiosInstance = axios.create({
    baseURL: baseURL,
    timeout: 5000, // 添加请求超时时间
    withCredentials: true // 表示请求可以携带cookie
})
export default request