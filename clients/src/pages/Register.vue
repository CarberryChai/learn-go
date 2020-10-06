<template>
  <div class="register-container">
    <div class="title">注册</div>
    <form class="form" @submit.prevent="submit">
      <div class="item">
        <label for="email">邮箱：</label>
        <input
          id="email"
          type="email"
          name="email"
          title="请填写有效的邮箱"
          pattern="^([^\x00-\x20\x22\x28\x29\x2c\x2e\x3a-\x3c\x3e\x40\x5b-\x5d\x7f-\xff]+|\x22([^\x0d\x22\x5c\x80-\xff]|\x5c[\x00-\x7f])*\x22)(\x2e([^\x00-\x20\x22\x28\x29\x2c\x2e\x3a-\x3c\x3e\x40\x5b-\x5d\x7f-\xff]+|\x22([^\x0d\x22\x5c\x80-\xff]|\x5c[\x00-\x7f])*\x22))*\x40([^\x00-\x20\x22\x28\x29\x2c\x2e\x3a-\x3c\x3e\x40\x5b-\x5d\x7f-\xff]+|\x5b([^\x0d\x5b-\x5d\x80-\xff]|\x5c[\x00-\x7f])*\x5d)(\x2e([^\x00-\x20\x22\x28\x29\x2c\x2e\x3a-\x3c\x3e\x40\x5b-\x5d\x7f-\xff]+|\x5b([^\x0d\x5b-\x5d\x80-\xff]|\x5c[\x00-\x7f])*\x5d))*$"
          required
        />
      </div>
      <div class="item">
        <label for="password">密码：</label>
        <input id="password" type="password" name="password" required />
      </div>
      <div class="item">
        <label for="confirmPassword">确认密码：</label>
        <input
          id="confirmPassword"
          type="password"
          name="confirmPassword"
          required
        />
      </div>
      <div class="form-sb">
        <Button html-type="submit" class-name="sb-btn" :loading="loading"
          >提交</Button
        >
      </div>
    </form>
  </div>
</template>

<script>
import { defineComponent, getCurrentInstance, ref } from 'vue'
import Button from '/@/components/Button.vue'
import { validate } from '/@/common/util'

export default defineComponent({
  name: 'Register',
  components: {
    Button,
  },
  setup(props, ctx) {
    const loading = ref(false)
    const notify = getCurrentInstance().appContext.config.globalProperties
      .$notify
    const submit = e => {
      const eles = e.target.elements
      const data = ['email', 'password', 'confirmPassword'].reduce(
        (t, c) => ((t[c] = eles[c].value), t),
        {}
      )
      if (!formValidate(data)) return
      loading.value = true
      console.log(data)
    }
    const formValidate = data => {
      for (const key of Object.keys(data)) {
        if (!validate(data[key], 'require')) {
          notify('请检查表单是否填写完整~')
          return false
        }
      }
      if (!validate(data['email'], 'email')) {
        notify('邮箱格式错误~')
        return false
      }
      if (data['confirmPassword'] !== data['password']) {
        notify('确认密码与原密码不同~')
        return false
      }
      return true
    }
    return {
      submit,
      loading,
    }
  },
})
</script>

<style lang="scss" scoped>
.register-container {
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
      width: 250px;
      height: 1px;
      background-color: #ccc;
      margin: 0 10px;
      vertical-align: middle;
    }
  }
  .form {
    margin-top: 100px;
    .item {
      margin-top: 20px;
      display: flex;
      align-items: center;
      label {
        display: block;
        margin-right: 20px;
        flex: 0.2;
        text-align: center;
      }
      input {
        border: none;
        outline: none;
        font-size: 14px;
        line-height: 1.5;
        border: 1px solid #d9d9d9;
        display: inline-block;
        flex: 0.8;
        padding: 0.4em 1em;
        border-radius: 5px;
        transition: all 0.3s;
        &:hover {
          border-color: var(--btn-primary);
        }
        &:focus {
          border-color: #40a9ff;
          box-shadow: 0 0 0 2px rgba(24, 144, 255, 0.2);
        }
        &:placeholder-shown {
          text-overflow: ellipsis;
        }
      }
    }
    .form-sb {
      margin-top: 30px;
      text-align: center;
      .sb-btn {
        padding-left: 8em;
        padding-right: 8em;
        background-color: var(--btn-primary);
        color: white;
      }
    }
  }
}
</style>
