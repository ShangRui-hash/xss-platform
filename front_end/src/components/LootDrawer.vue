<template>
  <el-drawer
    :title="project_name"
    :visible.sync="is_show_drawer"
    size="50%"
    @close="onCloseHandler"
  >
    <div style="padding:15px;">
      <el-table :data="loots" style="width: 100%;">
        <el-table-column prop="id" label="#" width="50" type="index"></el-table-column>
        <el-table-column prop="created_at" label="时间"></el-table-column>
        <el-table-column prop="content" label="内容" width="400">
          <template slot-scope="scope">
            <div v-for="(val,key) in JSON.parse(scope.row.content)" :key="key">
              <i style="color:yellow;">{{key}} :</i>
              <span>
                {{val}}
              </span>
            </div>
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="80">
          <template slot-scope="scope">
            <el-button @click="deleteLoot(scope.row.id)" type="warning" size="small" plain>删 除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </el-drawer>
</template>

<script>
export default {
  props: ["project_name", "is_show_drawer", "loots"],
  filters: {
    parse: function(value) {
      return JSON.parse(value);
    }
  },
  methods: {
    //当点击关闭按钮时
    onCloseHandler() {
      this.is_show_drawer = false;
      this.$emit("drawerClose");
    },
    //删除战利品
    deleteLoot(id) {
      this.$axios({
        method: "DELETE",
        url: `/api/v1/loot/${id}`
      })
        .then(() => {
          this.$message({
            type: "success",
            message: "删除成功"
          });
          //更新前端维护的数据
          let index = this.loots.findIndex(item => {
            return item.id == id;
          });
          this.loots.splice(index, 1);
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>