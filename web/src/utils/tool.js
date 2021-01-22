import store from '@/store'

const tool = {
  // 字典转换为列表
  dictToArray (d) {
    if (!d) {
      return []
    }

    let r = []
    for (var i in d) {
      r.push(d[i])
    }
    return r
  },
  // 列表转为字典
  arrayToDict (l, map) {
    if (!l) {
      return {}
    }

    let r = {}
    for (var i in l) {
      let key = map ? map(l[i]) : i
      r[key] = l[i]
    }
    return r
  }
}

export default tool
