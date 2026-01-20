/*
 * @Date: 2026-01-20 10:07:39
 * @LastEditors: zxd shisanlailin@outlook.com
 * @LastEditTime: 2026-01-20 10:12:02
 * @FilePath: \front-backend-demo\front\src\apis\index.ts
 * @Author: tingfeng
 */
import type {
  AxiosError,
  AxiosResponse,
  InternalAxiosRequestConfig,
} from 'axios'
import axios from 'axios'
import { ApiEventEmitter, getEnvData, getToken } from '@/utils'

const { VITE_API_BASE_URL, VITE_API_UPLOAD_URL } = getEnvData()

const codeWhite = [20000, 200, 1001]
const tokenWhite = ['/login', '/register']

const baseAxios = axios.create({
  timeout: 1000 * 60 * 0.5,
})

// 设置 baseAxios 基本 URL
if (VITE_API_BASE_URL) {
  baseAxios.defaults.baseURL = `${VITE_API_BASE_URL}`
}

baseAxios.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const path = `${config.url}${new Date().valueOf()}`

    // 处理文件上传请求
    if (/\/fileUpload/.test(config.url!)) {
      config.url = config.url?.replace('/fileUpload', '')
      config.baseURL = VITE_API_UPLOAD_URL
    }

    // 从 localStorage 或 Cookie 获取 token
    const token = getToken()

    if (token && !tokenWhite.some((item) => new RegExp(item).test(path))) {
      // 添加 Bearer 前缀
      config.headers.Authorization = `Bearer ${token}`
      config.headers.cancelToken = path
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  },
)

baseAxios.interceptors.response.use(
  (response: AxiosResponse) => {
    const res = response.data
    const url = response.config.url

    if (url?.includes('/apiStat')) {
      return response
    }

    // 上传文件成功的特殊处理
    if (url?.includes('/uploadFile')) {
      if (res.msg === '操作成功') {
        return response
      }
    }

    const { status } = response
    if (!codeWhite.includes(status)) {
      ApiEventEmitter.emit('noInCodeWhiteErr', { res: response })
      return Promise.reject(new Error('Error'))
    }
    return response
  },
  (error) => {
    ApiEventEmitter.emit('responseErr', {
      res: error as unknown as AxiosError,
    })
    return Promise.reject(error)
  },
)

export default baseAxios
