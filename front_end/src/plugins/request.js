import Vue from 'vue'
import axios from 'axios'
import config from '../config.js'
import router from '../router.js'
Vue.prototype.$axios = axios

//设置请求的baseURL
axios.defaults.baseURL = config.axios.baseURL;

//设置请求的超时时间
axios.defaults.timeout = config.axios.timeout;

//拦截发送的请求
axios.interceptors.request.use((request) => {
    let token_name
    if(request.url.indexOf("admin")!=-1){
        token_name= "admin_token"
    }else{
        token_name = "user_token"
    }
    const token = localStorage.getItem(token_name);
    if (token) {
        request.headers.Authorization = `Bearer ${token}`;
    }
    return request;
}, (err) => {
    return Promise.reject(err);
});

//拦截接收到的响应
axios.interceptors.response.use((resp) => {
    //如果需要登录
    let code = resp.data.code
    if (code  == config.axios.NEED_LOGIN_CODE || code == config.axios.INVALID_TOKEN ) {
        if (resp.request.responseURL.indexOf("admin") != -1) {
            router.push("/admin/login")
        } else {
            router.push("/login")
        }
        return Promise.reject("需要登录")
    }
    //返回的响应码不是成功
    if (code != config.axios.SUCCESS_CODE) {
        let msg
        if (typeof resp.data.msg == "string") {
            msg = resp.data.msg
        } else {
            msg = Object.values(resp.data.msg)[0]
        }
        Vue.prototype.$message({
            message: msg,
            type: "warning",
        })
        return Promise.reject(resp.data.msg)
    }
    return resp
}, err => {
    return Promise.reject(err)
})