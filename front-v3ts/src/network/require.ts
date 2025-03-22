import axios, { AxiosInstance } from 'axios'
import { baseURL } from '@/config'

// 创建 axios 实例
const request: AxiosInstance = axios.create({
    baseURL: baseURL,
    timeout: 5000, // 添加请求超时时间
    withCredentials: true // 表示请求可以携带cookie
})
export default request