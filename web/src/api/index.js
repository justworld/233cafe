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
  // 获取验证码
  getCaptcha (username, type) {
    return httpRequest.adornUrl('common/captcha') + '?username=' + username + '&type=' + type + '&r=' + new Date().getTime()
  },
  // 获取国际化配置
  getI18n (handler, data, err) {
    httpRequest.get(httpRequest.adornUrl('common/i18n'), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 注册
  reg (handler, data, err) {
    httpRequest.post(httpRequest.adornUrl('user/reg'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 登陆
  login (handler, data, err) {
    httpRequest.post(httpRequest.adornUrl('user/login'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // oauth登陆
  oauthLogin (handler, data, err) {
    httpRequest.post(httpRequest.adornUrl('user/oauthlogin'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 退出登陆
  logout (handler, data, err) {
    httpRequest.post(httpRequest.adornUrl('user/logout'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 用户信息
  userInfo (handler, data, err) {
    httpRequest.get(httpRequest.adornUrl('user/info'), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 改变主体
  changeClient (handler, data, err) {
    httpRequest.post(httpRequest.adornUrl('user/changeclient'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 修改密码
  changePasswd (handler, data, err) {
    httpRequest.post(httpRequest.adornUrl('user/changepasswd'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 获取aws token
  awsToken (handler, data, err) {
    httpRequest.post(httpRequest.adornUrl('common/temptoken'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 获取iot概览数据
  getIotOverview (handler, data, err) {
    httpRequest.get(httpRequest.adornUrl('product/overview'), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 获取产品类型数据
  getProductTypes (handler, data, err) {
    httpRequest.get(httpRequest.adornUrl('product/typelist'), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 产品列表
  getProducts (handler, data, err) {
    httpRequest.get(httpRequest.adornUrl('product'), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 获取可选产品
  getSelectProducts (handler, data, err) {
    httpRequest.get(httpRequest.adornUrl('product/selectlist'), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 创建产品
  addProduct (handler, data, err) {
    httpRequest.post(httpRequest.adornUrl('product'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 编辑产品
  updateProduct (handler, data, err) {
    let productId = data.id
    delete data.id
    httpRequest.put(httpRequest.adornUrl('product/' + productId), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 获取产品
  getProductInfo (handler, data, err) {
    let productId = data.id
    delete data.id
    httpRequest.get(httpRequest.adornUrl('product/' + productId), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  publishProduct (handler, data, err) {
    let productId = data.id
    delete data.id
    httpRequest.post(httpRequest.adornUrl('product/' + productId + '/publish'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 删除产品
  deleteProduct (handler, data, err) {
    let productId = data.id
    delete data.id
    httpRequest.delete(httpRequest.adornUrl('product/' + productId), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 获取固件代码
  getProductCode (handler, data, err) {
    let productId = data.id
    delete data.id
    httpRequest.get(httpRequest.adornUrl('product/' + productId + '/codelist'), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 上传固件代码
  uploadProductCode (handler, data, err) {
    let productId = data.id
    delete data.id
    httpRequest.post(httpRequest.adornUrl('product/' + productId + '/uploadcode'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 编辑固件代码
  updateProductCode (handler, data, err) {
    let id = data.id
    delete data.id
    httpRequest.post(httpRequest.adornUrl('product/' + id + '/updatecode'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 删除固件代码
  deleteProductCode (handler, data, err) {
    let id = data.id
    delete data.id
    httpRequest.post(httpRequest.adornUrl('product/' + id + '/deletecode'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 获取固件发布列表
  getProductCodeRelease (handler, data, err) {
    let id = data.id
    delete data.id
    httpRequest.get(httpRequest.adornUrl('product/' + id + '/coderelease'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 新建固件发布
  addProductCodeRelease (handler, data, err) {
    let id = data.id
    delete data.id
    httpRequest.post(httpRequest.adornUrl('product/' + id + '/addcoderelease'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 删除固件发布
  deleteProductCodeRelease (handler, data, err) {
    let id = data.id
    delete data.id
    httpRequest.post(httpRequest.adornUrl('product/' + id + '/deletecoderelease'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 获取固件发布信息
  getProductCodeReleaseInfo (handler, data, err) {
    let id = data.id
    delete data.id
    httpRequest.get(httpRequest.adornUrl('product/' + id + '/codereleaseinfo'), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 获取设备列表
  getDevices (handler, data, err) {
    httpRequest.get(httpRequest.adornUrl('device'), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 获取设备详情
  getDeviceInfo (handler, data, err) {
    let id = data.id
    delete data.id
    httpRequest.get(httpRequest.adornUrl('device/' + id), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 获取可量产产品
  getMassProducts (handler, data, err) {
    httpRequest.get(httpRequest.adornUrl('mass/products'), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 获取烧录固件列表
  getBurnCodes (handler, data, err) {
    let productId = data.id
    delete data.id
    httpRequest.get(httpRequest.adornUrl('mass/' + productId + '/burncodes'), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 生成授权码
  changeBurningVersion (handler, data, err) {
    let productId = data.id
    delete data.id
    httpRequest.post(httpRequest.adornUrl('mass/' + productId + '/changeversion'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 获取授权码列表
  getAuthCodes (handler, data, err) {
    let productId = data.id
    delete data.id
    httpRequest.get(httpRequest.adornUrl('mass/' + productId + '/authcodes'), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 生成授权码
  buyAuthCode (handler, data, err) {
    let productId = data.id
    delete data.id
    httpRequest.post(httpRequest.adornUrl('mass/' + productId + '/buyauthcode'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 帐号列表
  getAccounts (handler, data, err) {
    httpRequest.get(httpRequest.adornUrl('account'), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 角色列表
  getRoles (handler, data, err) {
    let userId = data.id
    delete data.id
    httpRequest.get(httpRequest.adornUrl('account/' + userId + '/roles'), { params: data }).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 增加授权
  addAccount (handler, data, err) {
    httpRequest.post(httpRequest.adornUrl('account/add'), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 删除授权
  delAccount (handler, data, err) {
    let userId = data.id
    delete data.id
    httpRequest.delete(httpRequest.adornUrl('account/' + userId), data).then(res => {
      resHandle(res, handler, err)
    })
  },
  // 用户授权
  permitAccount (handler, data, err) {
    let userId = data.id
    delete data.id
    httpRequest.post(httpRequest.adornUrl('account/' + userId + '/permit'), data).then(res => {
      resHandle(res, handler, err)
    })
  }
}

export default api
