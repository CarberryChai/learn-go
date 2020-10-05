<template>
  <transition name="fade" appear @after-leave="leave">
    <div v-if="show" class="notify-container">
      <span>{{ msg }}</span>
    </div>
  </transition>
</template>

<script>
import { defineComponent } from 'vue'
export default defineComponent({
  name: 'Notify',
  props: {
    duration: {
      type: Number,
      required: true,
    },
    msg: {
      type: String,
      required: true,
    },
    onClose: {
      type: Function,
      default: () => {},
    },
  },
  data() {
    return {
      show: true,
    }
  },
  mounted() {
    setTimeout(() => {
      this.show = false
    }, this.duration)
  },
  methods: {
    leave() {
      console.log('leave')
      this.onClose()
    },
  },
})
</script>
<style lang="scss" scoped>
.notify-container {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  display: inline-block;
  padding: 0.6em 1.4em;
  background-color: rgba(0, 0, 0, 0.5);
  border-radius: 0.4em;
  color: white;
  font-size: 0.8rem;
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
