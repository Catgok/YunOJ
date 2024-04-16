import {createRouter, createWebHistory} from 'vue-router'

const routes = [{
    path: '/',
    name: 'main',
    component: () => import('@/views/main/main.vue'),
    meta: {
        title: 'YunOJ'
    }
}, {
    path: '/login',
    name: 'login',
    component: () => import('@/views/login/login.vue'),
    meta: {
        title: 'YunOJ'
    },
}, {
    path: '/problem',
    name: 'problem',
    component: () => import('@/views/problem/problemList.vue'),
    meta: {
        title: 'YunOj - 题库'
    },
}, {
    path: '/problem/:problemId',
    name: 'problemContent',
    component: () => import('@/views/problem/problemContent.vue'),
    meta: {
        title: 'YunOj - 题目'
    }
}, {
    path: '/problem/submission/:problemId',
    name: 'submitList',
    component: () => import('@/views/problem/submissionList.vue'),
    meta: {
        title: 'YunOj - 提交列表'
    }
}, {
    path: '/problem/submission/:problemId/:submissionId',
    name: 'submissionContent',
    component: () => import('@/views/problem/submissionContent.vue'),
    meta: {
        title: 'YunOj - 提交详情'
    }
}, {
    path: '/contest',
    name: 'contest',
    component: () => import('@/views/contest/contestList.vue'),
    meta: {
        title: 'YunOj - 竞赛'
    },
}, {
    path: '/contest/:contestId',
    name: 'contestContest',
    component: () => import('@/views/contest/contestContent.vue'),
    meta: {
        title: 'YunOj - 竞赛'
    },
}, {
    path: '/contest/:contestId/problem/:problemId',
    name: 'contestProblemContent',
    component: () => import('@/views/contest/detail/contestProblemContent.vue'),
    meta: {
        title: 'YunOj - 竞赛'
    },
},
    {
        path: '/announcement',
        name: 'announcement',
        component: () => import('@/views/announcement/announcement.vue'),
        meta: {},
    }, {
        path: '/onlineIDE',
        name: 'onlineIDE',
        component: () => import('@/views/onlineIDE/onlineIDE.vue'),
        meta: {
            title: 'YunOj - 在线IDE'
        },
    }
]
const router = createRouter({
    history: createWebHistory(),
    routes // `routes: routes` 的缩写
})
export default router
