<template>
  <h2>{{ count }}</h2>
  <Button :loading="loading" @click="inc">Click</Button>
  <Button @click="sayHello">open</Button>
  <input type="file" @change="upload" />
</template>

<script>
import Button from './components/Button.vue'
import {
  defineComponent,
  ref,
  watchEffect,
  createVNode,
  render,
  getCurrentInstance,
} from 'vue'
export default defineComponent({
  name: 'App',
  components: {
    Button,
  },
  setup(props, ctx) {
    const count = ref(0)
    const loading = ref(false)
    const app = getCurrentInstance()
    watchEffect(() => {
      if (count.value >= 10) {
        loading.value = true
      }
    })
    const sayHello = () =>
      app.appContext.config.globalProperties.$notify('操作成功')
    const inc = () => count.value++
    const upload = e => {
      const file = e.target.files[0]
      const data = new FormData()
      data.append('file', file)
      data.append('name', 'carberry')
      window
        .fetch('/api/upload', {
          method: 'post',
          body: data,
        })
        .then(r => r.json())
        .then(console.log)
    }
    return {
      sayHello,
      count,
      upload,
      loading,
      inc,
    }
  },
})
</script>
