<template>
  <div class="register-container">
    <div class="title">注册</div>
    <form class="form" @submit.prevent="submit">
      <div class="item">
        <label for="email">邮箱：</label>
        <input id="email" type="text" name="email" />
      </div>
      <div class="item">
        <label for="password">密码：</label>
        <input id="password" type="password" name="password" />
      </div>
      <div class="item">
        <label for="confirmPassword">确认密码：</label>
        <input id="confirmPassword" type="password" name="confirmPassword" />
      </div>
      <Button html-type="submit">提交</Button>
    </form>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import Button from '/@/components/Button.vue'
export default defineComponent({
  name: 'Register',
  components: {
    Button,
  },
  setup(props, ctx) {
    const submit = e => {
      console.log(e)
    }
    return {
      submit,
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
        padding: 0.3em 1em;
        border-radius: 5px;
        transition: all 0.3s;
        &:hover {
          border-color: #40a9ff;
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
  }
}
</style>
