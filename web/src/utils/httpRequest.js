import axios from 'axios'
import { Message } from 'elementui'

const CancelToken = axios.CancelToken
const pending = []
const http = axios.create({
  timeout: 1000 * 30,
  headers: {
    'Content-Type': 'application/json;charset=utf-8'
  },
  cancelToken: new CancelToken(function executor (c) {
    pending.push(c)
  })
})

// http request 拦截器
http.interceptors.request.use(
  config => {
    return config
  },
  err => {
    return Promise.reject(err)
  }
)

// http response 拦截器
http.interceptors.response.use(
  response => {
    return response
  },
  error => {
    if (error.message === 'Network Error') {
      while (pending.length > 0) {
        pending.pop()('请求中断')
      }
      Message({
        message: '没有网络连接',
        type: 'error'
      })
    } else if (error.response) {
      switch (error.response.status) {
        case 500:
          Message({
            message: '服务器异常, 请稍后重试',
            type: 'error'
          })
          break
      }
    }
    return Promise.reject(error.response.data)
  }
)

// 请求地址处理
http.adornUrl = (actionName) => {
  return window.SITE_CONFIG.baseUrl + actionName + '/'
}

export default http
