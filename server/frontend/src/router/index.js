import {createRouter, createWebHashHistory} from 'vue-router'

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: '/',
            name: 'home',
            // component: () => import('@/views/HomeView.vue'),
            component: () => import('@/components/InstanceList.vue'),
        }, {
            path: '/client',
            name: 'client',
            component: () => import('@/views/ClientView.vue'),
        }, {
            path: '/instance',
            name: 'instance',
            component: () => import('@/views/InstanceView.vue'),
        }, {
            path: '/sync',
            name: 'sync',
            component: () => import('@/views/SyncView.vue')
        }, {
            path: '/relay-setting',
            name: 'relay-setting',
            component: () => import('@/views/RelaySettingView.vue'),
        }
    ]
})
router.beforeEach((to, from) => {
    // ...
    // 返回 false 以取消导航
    return true
})

export default router
