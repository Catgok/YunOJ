import {createStore} from 'vuex'

const store = createStore({
    state() {
        return {
            userinfo: '',
            username: '',
            loginStatus: false,
        }
    },
    mutations: {}
})
export default store