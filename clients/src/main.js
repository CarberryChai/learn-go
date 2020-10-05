import { createApp } from 'vue'
import App from './App.vue'
import notify from './components/Notify'
import './index.css'

const app = createApp(App)
app.config.globalProperties.$notify = notify
app.mount('#app')
