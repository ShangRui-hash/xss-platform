<template>
  <div style="height:100%;" class="flex flex-dir-col">
    <el-header>
      <HeaderBar is_admin=true></HeaderBar>
    </el-header>
    <el-container >
      <el-aside >
        <NavBar :navbar="navbar"></NavBar>
      </el-aside>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
    <el-footer><i>{{config.foot_text}}</i></el-footer>
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
  data:()=>{
    return {
       navbar: [
        {
          index: 1,
          to: "/admin/index/home",
          title: "仪表盘",
          icon: "el-icon-data-analysis"
        },
        {
          index: 2,
          to: "/admin/index/",
          title: "用户管理",
          icon: "el-icon-user"
        },
        {
          index: 3,
          to: "/admin/index/modules",
          title: "模块管理",
          icon: "el-icon-star-off"
        },
        {
          index: 4,
          to: "",
          title: "站点配置",
          icon: "el-icon-setting"
        }
      ]
    }
  }
};
</script>
