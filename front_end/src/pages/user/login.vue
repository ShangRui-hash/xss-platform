<template>
  <div class="container flex flex-dir-col flex-center">
    <el-card class="box-card m-tb-2 m-rl-auto p-2">
      <div class="flex">
        <!-- 左边 -->
        <div class="login-card">
          <h1>XSS Platform</h1>
          <span>Author:XDU Rick Shang</span>
          <el-form ref="form" :model="form">
            <el-form-item label="username">
              <el-input v-model="login_form.username"></el-input>
            </el-form-item>
            <el-form-item label="password">
              <el-input v-model="login_form.password" show-password></el-input>
            </el-form-item>
            <el-form-item>
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
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="onClickLoginBtn" id="grecaptcha">登录</el-button>
              <el-button @click="is_show_register_dialog=true">注册</el-button>
            </el-form-item>
          </el-form>
        </div>
        <!-- 右边 -->
        <Logo></Logo>
      </div>
    </el-card>
    <el-dialog width="25%" title="注册" :visible.sync="is_show_register_dialog">
      <register-form :form="register_form"></register-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="is_show_register_dialog=false">取 消</el-button>
        <el-button type="primary" @click="onClickRegisterBtn">确 认</el-button>
      </div>
      <!-- 注册的ReCAPTCHA -->
      <vue-programmatic-invisible-google-recaptcha
        ref="invisibleRigsterRecaptcha"
        :sitekey="config.ReCAPTCHA.sitekey"
        :elementId="'invisibleRigsterRecaptcha'"
        :badgePosition="'left'"
        :showBadgeMobile="false"
        :showBadgeDesktop="false"
        @recaptcha-callback="recvRegisterReCAPTCHAToken"
      ></vue-programmatic-invisible-google-recaptcha>
    </el-dialog>
  </div>
</template>

<script>
import Logo from "@/components/Logo";
import RegisterForm from "@/components/RegisterForm";
export default {
  components: {
    Logo,
    RegisterForm
  },
  data() {
    return {
      login_form: {
        username: "",
        password: "",
        g_recaptcha_response: ""
      },
      register_form: {
        password: "",
        repassword: "",
        g_recaptcha_response: ""
      },
      is_show_register_dialog: false
    };
  },
  methods: {
    //等待接收google返回的token
    recvLoginReCAPTCHAToken(token) {
      this.login_form.g_recaptcha_response = token;
      this.login();
    },
    //等待接收用于注册的ReCAPTCHA token
    recvRegisterReCAPTCHAToken(token) {
      this.register_form.g_recaptcha_response = token;
      this.register();
    },
    //发送登录请求
    login() {
      this.$axios({
        method: "POST",
        url: "/api/v1/login",
        data: this.login_form
      })
        .then(resp => {
          //1.存储token
          let token = resp.data.data;
          localStorage.setItem("user_token", token);
          //2.跳转页面
          this.$router.push("/");
        })
        .catch(err => {
          console.log(err);
        });
    },
    //发送注册请求
    register() {
      this.$axios({
        method: "POST",
        url: "/api/v1/register",
        data: {
          password: this.register_form.password,
          g_recaptcha_response: this.register_form.g_recaptcha_response
        }
      })
        .then(resp => {
          //存储token
          let token = resp.data.data;
          localStorage.setItem("user_token", token);
          //通知用户
          let user_info = this.$parse.GetUserInfo(token);
          this.$alert(
            `您的用户名为${user_info.username},请妥善保管。`,
            "注册成功",
            {}
          );
          //跳转页面
          this.$router.push(`/`);
        })
        .catch(err => {
          console.log(err);
        });
    },
    //登录的前端效验
    loginValidate() {
      if (this.login_form.username.length == 0) {
        this.$message({
          type: "warning",
          message: "请输入用户名"
        });
        return false;
      }
      if (this.login_form.password.length == 0) {
        this.$message({
          type: "warning",
          message: "请输入密码"
        });
        return false;
      }
      return true;
    },
    //当点击登录按钮时
    onClickLoginBtn() {
      //前端效验
      if (!this.loginValidate()) {
        return;
      }
      //执行reCAPTCHA 人机效验,效验成功后会返回token并调用回调函数
      this.$refs.invisibleRecaptcha1.execute();
    },
    //注册功能的前端效验
    registerValidate() {
      if (this.register_form.password.length == 0) {
        this.$message({
          type: "warning",
          message: "请输入密码"
        });
        return false;
      }
      if (this.register_form.password != this.register_form.repassword) {
        this.$message({
          type: "warning",
          message: "两次密码输入不一致"
        });
        return false;
      }
      return true;
    },
    //点击注册按钮
    onClickRegisterBtn() {
      //前端效验
      if (!this.registerValidate()) {
        return;
      }
      //执行reCAPTCHA 人机效验,效验成功后会返回token并调用回调函数
      this.$refs.invisibleRigsterRecaptcha.execute();
    }
  }
};
</script>

<style scoped>
.login-card {
  width: 40%;
  max-width: 400px;
  min-width: 250px;
  margin-right: 25px;
}
</style>>
