<template>
  <div class="login-container">
    <div class="title">Login</div>
    <form class="form" @submit.prevent="submit">
      <div class="item">
        <label for="email">邮箱</label>
        <input id="email" type="email" required />
      </div>
      <div class="item">
        <label for="password">密码</label>
        <input id="password" type="password" required />
      </div>
      <div class="form-sb">
        <Button html-type="submit" class-name="sb-btn" :loading="loading"
          >登录</Button
        >
      </div>
    </form>
  </div>
</template>

<script>
import { defineComponent, ref, getCurrentInstance } from 'vue'
import Button from '/@/components/Button.vue'
export default defineComponent({
  name: 'Login',
  components: {
    Button,
  },
  setup(props, ctx) {
    const loading = ref(false)
    const { $post } = getCurrentInstance().appContext.config.globalProperties
    const submit = e => {
      $post('/api/login').then(console.log)
    }
    return {
      loading,
      submit,
    }
  },
})
</script>

<style lang="scss">
.login-container {
  margin: 50px auto 0;
  max-width: 600px;
  .title {
    text-align: center;
    font-size: 24px;
    font-weight: 500;
    &::before,
    &::after {
      content: '';
      display: inline-block;
      width: 41%;
      height: 1px;
      background-color: #ccc;
      margin: 0 10px;
      vertical-align: middle;
    }
  }
}
</style>
