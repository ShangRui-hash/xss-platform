<template>
  <div>
    <el-button @click="is_show_add_module_dialog=true" size="small">添加模块</el-button>
    <el-dialog title="添加模块" :visible.sync="is_show_add_module_dialog">
      <ModuleForm :form="form" :disabled="false"></ModuleForm>
      <div slot="footer" class="dialog-footer">
        <el-button @click="is_show_add_module_dialog=false">取 消</el-button>
        <el-button type="primary" @click="addModuleHandler">确 认</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import ModuleForm from "@/components/ModuleForm";
export default {
  props: ["add_module_url"],
  components: {
    ModuleForm
  },
  data: () => {
    return {
      is_show_add_module_dialog: false,
      form: {
        id: "",
        name: "",
        desc: "",
        xss_payload: "",
        is_common:true,
        param_list: [],
        option_list: []
      }
    };
  },
  methods: {
    //添加模块
    addModuleHandler() {
      this.$axios({
        method: "POST",
        url: this.add_module_url,
        data: this.form
      })
        .then(res => {
          this.$message({
            message: "添加成功",
            type: "success"
          });
          //向前端维护的数据中添加数据
          let data = res.data.data;
          let module = {
            id: data.id,
            name: data.name,
            desc: data.desc,
            created_at: "刚刚",
            updated_at: "刚刚",
            username: data.username,
          };
          this.$emit("addModule", module);
          //关闭对话框
          this.is_show_add_module_dialog = false;
          //清空form
          this.form = {
            id: "",
            name: "",
            desc: "",
            xss_payload: "",
            param_list: [],
            option_list: [],
            is_common:true,
          };
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>