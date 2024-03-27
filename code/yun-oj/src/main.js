import {createApp} from 'vue'
import router from './router'
import 'element-plus/dist/index.css'
import App from './App.vue'
import axios from 'axios'
// import {authed} from "@/utils/utils"


axios.defaults.baseURL = 'https://utiool-api.hhxxblog.top/';
// axios.defaults.baseURL = 'http://localhost:8080/';

const app = createApp(App)
app.use(router)
app.config.globalProperties.$axios = axios;
app.mount('#app')

export default app
