import Vue from 'vue'
import VueRouter from 'vue-router'
import PageAdminIndex from './pages/admin/index.vue'
import PageAdminLogin from './pages/admin/login.vue'
import PageUserIndex from  './pages/user/index.vue'
import PageAdminHome from './pages/admin/left/home.vue'
import PageAdminModules from './pages/admin/left/modules.vue'
import PageAdminUser from './pages/admin/left/user.vue'
import PageMyProject from './pages/user/left/project.vue'
import PageMyModule from './pages/user/left/module.vue'
import PageCommonModule from './pages/user/left/common_module.vue'
import PageUserLogin from "./pages/user/login.vue"
import Page404 from "./pages/404.vue"
Vue.use(VueRouter)

const router = new VueRouter({
    routes:[
        //后台
        {
            path:"/admin/index/",
            component:PageAdminIndex,
            children:[
                {
                    path:"home",
                    component:PageAdminHome,
                },
                {
                    path:"modules",
                    component:PageAdminModules,
                },
                {
                    path:"/",
                    component:PageAdminUser,
                },
            ]
        },
        {
            path:"/admin/login",
            component:  PageAdminLogin,
        },
        //前台
        {
            path:"/login",
            component:  PageUserLogin,
        },
        {
            path:"/",
            component:PageUserIndex,
            children:[
                {
                    path:"/",
                    component:PageMyProject
                },
                {
                    path:"module",
                    component:PageMyModule,
                },
                {
                    path:"common_module",
                    component:PageCommonModule,
                }
            ]
        },
        //404
        {
            path:"*",
            component:Page404,
        }
    ]
})
export default router