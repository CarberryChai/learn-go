import { createApp } from 'vue'
import App from './App.vue'
import { get, post } from './common/util'
import notify from './components/Notify'
import './index.scss'
import router from './router'

const app = createApp(App)
app.config.globalProperties.$notify = notify
app.config.globalProperties.$get = get
app.config.globalProperties.$post = post
app.use(router)
app.mount('#app')
