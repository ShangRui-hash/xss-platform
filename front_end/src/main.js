import Vue from 'vue'
import App from './App.vue'
import router from './router.js'
import './plugins/element.js' //专门用于导入element-ui 相关的组件
import './plugins/request.js' //专用用来发送http请求
import "./plugins/reCATCHA.js"//google recatpcha invisable组件

import parse from './plugins/parse.js' //用来解析token
import config from  './config.js'

Vue.config.productionTip = false
Vue.prototype.config = config
Vue.prototype.$parse = parse



new Vue({
  render: h => h(App),
  router
}).$mount('#app')

