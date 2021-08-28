<template>
  <div class="content-wrapper">
    <el-table :data="users" style="width: 100% border:0px">
      <el-table-column prop="id" label="#" width="50" type="index"></el-table-column>
      <el-table-column prop="username" label="用户名" width="120"></el-table-column>
      <el-table-column prop="created_at" label="注册时间"></el-table-column>
      <el-table-column prop="updated_at" label="更新时间"></el-table-column>
      <el-table-column prop="logined_at" label="登录时间"></el-table-column>
      <el-table-column prop="is_banned" label="状态"></el-table-column>
      <el-table-column fixed="right" label="操作" width="150">
        <template slot-scope="scope">
          <el-button
            @click="changeUserStatus(scope.row.id)"
            type="warning"
            size="small"
          >{{scope.row.is_banned == "正常"?"封 号":"解 封"}}</el-button>
          <span style="margin-right:10px"></span>
          <el-popconfirm title="您确定要删除吗" @confirm="deleteUser(scope.row.id)">
            <el-button slot="reference" type="danger" size="small">删 除</el-button>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <LoadMoreBtn
      :loading ="loading"
      :loadmore_btn_text = "loadmore_btn_text"
      @loadmore = loadmore
    ></LoadMoreBtn>
  </div>
</template>
<script>
import LoadMoreBtn from '@/components/LoadMoreBtn'
export default {
  components:{
    LoadMoreBtn
  },
  data: () => {
    return {
      offset: 0,
      count: 10,
      users: [],
      loading: false,
      loadmore_btn_text:"",
    };
  },
  created() {
    this.loadmore_btn_text = this.config.loadmore_text
    this.getUsers();
  },
  methods: {
    //获取用户列表
    getUsers(){
      this.$axios({
        method: "GET",
        url: `/api/v1/admin/users/${this.offset}/${this.count}`
      })
        .then(resp => {
          let users = resp.data.data;
          users.forEach(user => {
            user.is_banned = user.is_banned == false ? "正常" : "被封";
          });
          this.users.push(...users);
          this.offset = this.users.length;
          this.loading = false;
          if(users.length < this.count){
            this.loadmore_btn_text = this.config.nomore_text
          }
        })
        .catch(err => {
          console.log(err);
        });
    },
    //删除用户
    deleteUser(id) {
      //删除前端维护的数据
      this.users.splice(
        this.users.findIndex(item => {
          return item.id == id;
        }),
        1
      );
      //删除后端维护的数据
      this.$axios({
        method: "DELETE",
        url: `/api/v1/admin/user/${id}`
      })
        .then(() => {
          this.$message({
            message: "删除成功",
            type: "success"
          });
        })
        .catch(err => {
          console.log(err);
        });
    },
    //改变用户状态
    changeUserStatus(id) {
      //修改前端维护的数据
      let index = this.users.findIndex(item => {
        return item.id == id;
      });
      this.users[index].is_banned =
        this.users[index].is_banned == "正常" ? "被封" : "正常";
      //修改后端维护的数据
      this.$axios({
        method: "PUT",
        url: `/api/v1/admin/user/${id}`
      })
        .then(() => {
          this.$message({
            message: "修改成功",
            type: "success"
          });
        })
        .catch(err => {
          console.log(err);
        });
    },
    loadmore() {
      this.loading = true;
      this.getUsers()
    }
  } //methods end
};
</script>
<style scoped>
</style>

