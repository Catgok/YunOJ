import {createRouter, createWebHistory} from 'vue-router'

const routes = [{
    path: '/',
    name: 'main',
    component: () => import('@/views/main/main.vue'),
    meta: {},
}, {
    path: '/login',
    name: 'login',
    component: () => import('@/views/login/login.vue'),
    meta: {},
}, {
    path: '/problem',
    name: 'problem',
    component: () => import('@/views/problem/problemList.vue'),
    meta: {},
}, {
    path: '/problem/:problemId',
    name: 'problemContent',
    component: () => import('@/views/problem/problemContent.vue'),
    meta: {}
}, {
    path: '/contest',
    name: 'competition',
    component: () => import('@/views/contest/contest.vue'),
    meta: {},
}, {
    path: '/announcement',
    name: 'announcement',
    component: () => import('@/views/announcement/announcement.vue'),
    meta: {},
}, {
    path: '/onlineIDE',
    name: 'onlineIDE',
    component: () => import('@/views/onlineIDE/onlineIDE.vue'),
    meta: {},
}
]
const router = createRouter({
    history: createWebHistory(),
    routes // `routes: routes` 的缩写
})
export default router
