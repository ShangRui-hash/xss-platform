<template>
  <div class="container flex flex-dir-col flex-center">
    <div class="login m-tb-2 m-rl-auto p-2">
      <form action="javascript:;" method="post">
        <div class="m-tb-2 flex flex-space-between">
          <label>username</label>
          <input class="input-yellow" v-model="form.username" type="text" name="username" />
        </div>
        <div class="m-tb-2 flex flex-space-between">
          <label>password</label>
          <input class="input-yellow" v-model="form.password" type="password" name="password" />
        </div>
        <div class="flex flex-center">
          <button class="btn-green" @click="onClickLoginBtn">Login</button>
        </div>
      </form>
      <div class="alert">{{alert}}</div>
      <!-- This is in the component you want to have the reCAPTCHA -->
      <vue-programmatic-invisible-google-recaptcha
        ref="invisibleRecaptcha1"
        :sitekey="config.ReCAPTCHA.sitekey"
        :elementId="'invisibleRecaptcha1'"
        :badgePosition="'left'"
        :showBadgeMobile="false"
        :showBadgeDesktop="false"
        @recaptcha-callback="recvLoginReCAPTCHAToken"
      ></vue-programmatic-invisible-google-recaptcha>
    </div>
  </div>
</template>

<script>
export default {
  data: () => {
    return {
      form: {
        username: "",
        password: "",
        g_recaptcha_response :"",
      },
      alert: ""
    };
  },
  methods: {
    //发送登录请求
    login() {
      this.$axios({
        method: "post",
        url: "/api/v1/admin/login",
        data: this.form
      })
        .then(res => {
          //登录成功
          this.alert = "";
          localStorage.setItem("admin_token", res.data.data);
          this.$router.push(`/admin/index/home`);
        })
        .catch(e => {
          console.log("err:", e);
        });
    },
    //前端效验
    validate() {
      if (this.form.username.length == 0 || this.form.password.length == 0) {
        this.alert = "Please input username and password.";
        return false;
      }
      return true;
    },
    //当点击登录按钮时
    onClickLoginBtn() {
      if (!this.validate()) {
        return;
      }
      //执行reCAPTCHA 人机效验,效验成功后会返回token并调用回调函数
      this.$refs.invisibleRecaptcha1.execute();
    },
    //等待接收google返回的token
    recvLoginReCAPTCHAToken(token) {
      this.form.g_recaptcha_response = token;
      this.login();
    }
  }
};
</script>

<style>
</style>
