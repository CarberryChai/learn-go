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
