import {createApp} from 'vue'
import router from './router'
import 'element-plus/dist/index.css'
import App from './App.vue'
import apiService from "@/utils/apiService";

export {_ResizeObserver} from "@/utils/utils"
// import {authed} from "@/utils/utils"

const app = createApp(App)
app.use(router)
router.beforeEach((to, from, next) => {
    document.title = to.meta.title || 'YunOj'
    next()
})

// app.use(VueMathJax)
app.config.globalProperties.$axios = apiService
app.mount('#app')
export default app
