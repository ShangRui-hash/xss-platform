<template>
  <div class="flex el-header">
    <div class="logo">XSS Platform</div>
    <div class="flex-grow" style="text-align:right">
      <el-dropdown @command="onClickDropdownMenuItem">
        <span class="el-dropdown-link" style="cursor:pointer">
          {{username}}
          <i class="el-icon-arrow-down el-icon--right"></i>
        </span>
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item>修改密码</el-dropdown-item>
          <el-dropdown-item command="logout">退出登录</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
  </div>
</template>

<script>
export default {
  props: ["is_admin"],
  data() {
    return {
      username: "",
      token_name: "",
      token: "",
      resp: ""
    };
  },
  mounted() {
    this.token_name = this.is_admin ? "admin_token" : "user_token";
    this.token = localStorage.getItem(this.token_name);
    let userinfo = this.$parse.GetUserInfo(this.token);
    this.username = userinfo.username;
  },
  methods: {
    onClickDropdownMenuItem(command) {
      switch (command) {
        case "logout":
          this.logout();
          break
        default:
          alert(command);
      }
    },
    logout() {
      //清除前端维护的数据
      localStorage.removeItem(this.token_name);
      //清除后端维护的数据
      this.$axios({
        method: "POST",
        url: "/api/v1/logout",
        data: { token: this.token }
      })
        .then(resp => {
          this.resp = resp;
          let login_url = this.is_admin ? "/admin/login" : "/login";
          this.$router.push(login_url);
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>
<style scoped>
.el-header {
  line-height: 60px;
  border-bottom: 1px solid green;
  overflow: hidden;
}
.logo {
  font-size: 30px;
}
</style>