import axios from 'axios'
import notify from '/@/components/Notify'
axios.interceptors.response.use(
  function (response) {
    const { data } = response
    if (data?.code < 0) {
      notify(data.msg)
    }
    return data
  },
  function (error) {
    // Any status codes that falls outside the range of 2xx cause this function to trigger
    // Do something with response error
    return Promise.reject(error)
  }
)

export function get(url, data) {
  return axios.get(url, {
    params: data,
  })
}
export function post(url, data) {
  return axios.post(url, data)
}
/**
 *表单验证
 *
 * @export
 * @param {string} value
 * @param {string} type
 * @return {boolean}
 */
export function validate(value, type) {
  value =
    value === null
      ? ''
      : (value + '').replace(/^[\s\uFEFF\xA0]+|[\s\uFEFF\xA0]+$/g, '')
  // 非空验证
  if (type === 'require') {
    return !!value
  }
  // 手机号验证
  if (type === 'mobile') {
    return /^1[3456789][0-9]{9}$/.test(value)
  }
  // 邮箱格式验证
  if (type === 'email') {
    return /^(\w)+(\.\w+)*@(\w)+((\.\w{2,3}){1,3})$/.test(value)
  }
  return false
}
