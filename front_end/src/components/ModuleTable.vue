<template>
  <div>
    <el-table :data="modules" style="width: 100% border:0px">
      <el-table-column prop="id" label="#" width="50" type="index"></el-table-column>
      <el-table-column prop="name" label="模块名称" width="150"></el-table-column>
      <el-table-column prop="desc" label="模块描述" width="300"></el-table-column>
      <el-table-column prop="created_at" label="添加时间" width="160"></el-table-column>
      <el-table-column prop="updated_at" label="更新时间" width="160"></el-table-column>
      <el-table-column prop="username" label="添加用户"></el-table-column>
      <el-table-column label="是否公开" align="center">
        <template slot-scope="scope">{{scope.row.is_common == true?"是":"否"}}</template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="150">
        <template slot-scope="scope">
          <el-button @click="onClickConfigBtn(scope.row)" type="primary" size="small">配置</el-button>
          <span style="margin-right:10px"></span>
          <el-popconfirm title="您确定要删除吗" @confirm="deleteModule(scope.row.id)">
            <el-button slot="reference" type="warning" size="small">删除</el-button>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <!-- 配置模块的模态框 -->
    <el-dialog title="配置模块" :visible.sync="is_show_update_module_dialog">
      <ModuleForm :form="form"></ModuleForm>
      <div slot="footer" class="dialog-footer">
        <el-button @click="is_show_update_module_dialog=false">取 消</el-button>
        <el-button type="primary" @click="onConfirmUpdate">确 认</el-button>
      </div>
    </el-dialog>
    <load-more-btn :loading="loading" :loadmore_btn_text="loadmore_btn_text" @loadmore="loadmore"></load-more-btn>
  </div>
</template>

<script>
import ModuleForm from "@/components/ModuleForm";
import LoadMoreBtn from "@/components/LoadMoreBtn";
export default {
  props: [
    "get_modules_url",
    "del_module_url",
    "update_module_url",
    "is_show_username",
    "modules"
  ],
  components: {
    ModuleForm,
    LoadMoreBtn
  },
  data: () => {
    return {
      form: {
        id: "",
        name: "",
        desc: "",
        xss_payload: "",
        param_list: [],
        option_list: [],
        is_common: ""
      },
      is_show_update_module_dialog: false,
      loadmore_btn_text: "",
      loading: false,
      offset: 0,
      count: 10
    };
  },
  mounted() {
    this.loadmore_btn_text = this.config.loadmore_text;
    this.getModules();
  },
  methods: {
    loadmore() {
      this.loading = true;
      this.getModules();
    },
    //获取模块列表
    getModules() {
      this.$axios({
        method: "GET",
        url: `${this.get_modules_url}/${this.offset}/${this.count}`
      })
        .then(res => {
          let modules = res.data.data;
          for (let i = 0; i < modules.length; i++) {
            let index = this.modules.findIndex(item => {
              return item.id == modules[i].id;
            });
            if (-1 == index) {
              this.modules.push(modules[i]);
            } else {
              this.modules[index] = modules[i];
            }
          }

          this.offset = this.modules.length;
          this.loading = false;
          if (res.data.data.length < this.count) {
            this.loadmore_btn_text = this.config.nomore_text;
          }
        })
        .catch(err => {
          console.log(err);
        });
    },
    //点击配置按钮时
    onClickConfigBtn(row) {
      this.is_show_update_module_dialog = true;
      this.$axios({
        methods: "GET",
        url: `${this.update_module_url}/${row.id}`
      })
        .then(res => {
          this.form = res.data.data;
        })
        .catch(err => {
          console.log(err);
        });
    },
    //删除模块
    deleteModule(id) {
      //删除前端维护的数据
      this.modules.splice(
        this.modules.findIndex(item => {
          return item.id == id;
        }),
        1
      );
      //删除后端维护的数据
      this.$axios({
        method: "DELETE",
        url: `${this.del_module_url}/${id}`
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
    //确认修改
    onConfirmUpdate() {
      //修改前端维护的数据
      let index = this.modules.findIndex(item => {
        return item.id == this.form.id;
      });
      this.modules[index].name = this.form.name;
      this.modules[index].desc = this.form.desc;
      this.modules[index].is_common = this.form.is_common;
      //修改后端维护的数据
      this.$axios({
        method: "PUT",
        url: this.update_module_url,
        data: this.form
      })
        .then(() => {
          this.$message({
            message: "修改成功",
            type: "success"
          });
          this.is_show_update_module_dialog = false;
        })
        .catch(err => {
          console(err);
        });
    }
  }
};
</script>