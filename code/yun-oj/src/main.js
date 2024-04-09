import {createApp} from 'vue'
import router from './router'
import 'element-plus/dist/index.css'
import App from './App.vue'
import axios from 'axios'
// import {authed} from "@/utils/utils"
// axios.defaults.baseURL = 'https://utiool-api.hhxxblog.top/';
axios.defaults.baseURL = 'http://localhost:8888/v1/';

const app = createApp(App)
app.use(router)
router.beforeEach((to, from, next) => {
    document.title = to.meta.title || 'YunOj'
    next()
})

// app.use(VueMathJax)
app.config.globalProperties.$axios = axios
app.mount('#app')
export default app
