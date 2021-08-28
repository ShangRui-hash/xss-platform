<template>
  <div style="height:100%" class="flex flex-dir-col">
    <el-header>
      <HeaderBar></HeaderBar>
    </el-header>
    <el-container>
      <el-aside>
        <NavBar :navbar="navbar"></NavBar>
      </el-aside>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
    <el-footer>
      <i>{{config.foot_text}}</i>
    </el-footer>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar";
import HeaderBar from "@/components/HeaderBar";
export default {
  components: {
    NavBar,
    HeaderBar
  },
  created() {
    //连接服务器的websocket
    let user_token = localStorage.getItem("user_token");
    this.conn = new WebSocket(
      `${this.config.ws.protocol}://${this.config.ws.domain}:${this.config.ws.port}/${this.config.ws.path}/${user_token}`
    );
    this.conn.onopen = this.onOpen;
    //当服务器websocket关闭时
    this.conn.onclose = this.onClose;
    //当收到服务器websocket发来的消息时
    this.conn.onmessage = this.onMessage;
  },
  beforeDestroy() {
    //关闭连接，下线
    this.closeWebsocket();
  },
  data: () => {
    return {
      conn: "",
      navbar: [
        {
          index: 1,
          title: "我的项目",
          to: "/",
          icon: "el-icon-files"
        },
        {
          index: 2,
          title: "我的模块",
          to: "/module",
          icon: "el-icon-star-off"
        },
        {
          index: 3,
          title: "公共模块",
          to: "/common_module",
          icon: "el-icon-toilet-paper"
        }
      ]
    };
  },
  methods: {
    //websocket连接成功
    onOpen(evt) {
      if ("development" == process.env.NODE_ENV) {
        this.$notify({
          message: "ws连接成功"
        });
        console.log(evt)
      }
    },
    //websocket连接关闭
    onClose(evt) {
      if ("development" == process.env.NODE_ENV)
        this.$notify({
          message: "服务端关闭ws连接"
        });
        console.log(evt)
    },
    //当收到服务端推送的消息时
    onMessage(evt) {
      let loot = JSON.parse(evt.data);
      this.$notify.info({
        title: `项目${loot.url_key}战利品+1`,
        message: this.$createElement(
          "i",
          { style: "word-break: break-all" },
          loot.content
        ),
        duration: 0
      });
    },
    //关闭websocket连接
    closeWebsocket() {
      this.conn.close();
      if ("development" == process.env.NODE_ENV) {
        this.$notify({
          message: "客户端主动关闭ws连接"
        });
      }
    }
  }
};
</script>
