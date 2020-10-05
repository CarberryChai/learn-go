import { createApp } from 'vue'
import App from './App.vue'
import notify from './components/Notify'
import './index.css'
import router from './router'

const app = createApp(App)
app.config.globalProperties.$notify = notify
app.use(router)
app.mount('#app')
