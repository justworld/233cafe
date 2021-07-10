import httpRequest from '@/utils/httpRequest'
import { Message } from 'elementui'

function resHandle (res, succ, err) {
  if (res && res.data.code === 403) {
    return
  }
  if (res && res.data.result) {
    succ(res.data)
  } else if (err) {
    err(res.data)
  } else {
    // 接口失败默认提示
    Message({
      message: res.data.msg,
      duration: 5000,
      showClose: true,
      type: 'error'
    })
  }
}

const api = {
  // common
  get (url, handler, data, err) {
    httpRequest.get(url, { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  post (url, handler, data, err) {
    httpRequest.post(url, data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 获取文章列表
  getArticles (handler, data, err) {
    httpRequest.get(httpRequest.adornUrl('article'), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 获取文章
  getArticle (handler, data, err) {
    let id = data.id
    delete data.id
    httpRequest.get(httpRequest.adornUrl('article/' + id), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  reg (handler, data, err) {
    httpRequest.post(httpRequest.adornUrl('user/reg'), data).then(res => {
      resHandle(res, handler, err)
    })
  }
}

export default api
