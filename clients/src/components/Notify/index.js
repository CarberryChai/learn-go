import { createVNode, render } from 'vue'
import Notify from './Notify.vue'
export default function notify(msg, config = {}) {
  const div = document.createElement('div')
  const vnode = createVNode(
    Notify,
    {
      duration: 1500,
      msg,
      onClose() {
        div.remove()
      },
      ...config,
    },
    null
  )
  render(vnode, div)
  document.body.appendChild(div)
}
