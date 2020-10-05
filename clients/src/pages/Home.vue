<template>
  <header class="header-container">
    <div class="logo">OurTube</div>
    <div class="log-container">
      <span class="iconfont icon-add" />
      <router-link to="/login">登录</router-link>
      <router-link to="/register">注册</router-link>
    </div>
  </header>
</template>

<script>
import { defineComponent } from 'vue'
export default defineComponent({
  name: 'Home',
})
</script>

<style lang="scss" scoped>
.header-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  .logo {
    margin: 10px;
    font-weight: 600;
    font-size: 24px;
  }
  .log-container {
    width: 20%;
    min-width: 200px;
    display: flex;
    justify-content: space-evenly;
    align-items: center;
    .icon-add {
      font-size: 24px;
      cursor: pointer;
    }
  }
}
</style>
