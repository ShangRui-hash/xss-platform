<template>
  <div>
    <el-table :data="modules" style="width: 100% border:0px">
      <el-table-column prop="id" label="#" width="50" type="index"></el-table-column>
      <el-table-column prop="name" label="模块名称" width="100"></el-table-column>
      <el-table-column prop="desc" label="模块描述"></el-table-column>
      <el-table-column prop="created_at" label="添加时间"></el-table-column>
      <el-table-column prop="updated_at" label="更新时间"></el-table-column>
      <el-table-column fixed="right" label="操作" width="150">
        <template slot-scope="scope">
          <el-button @click="onClickDetailBtn(scope.row)" type="primary" size="small">详情</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog title="模块详情" :visible.sync="is_show_module_dialog">
      <ModuleForm :form="form" disabled="true"></ModuleForm>
    </el-dialog>
    <load-more-btn :loading="loading" :loadmore_btn_text="loadmore_btn_text" @loadmore="loadmore"></load-more-btn>
  </div>
</template>

<script>
import ModuleForm from "@/components/ModuleForm";
import LoadMoreBtn from "@/components/LoadMoreBtn";
export default {
  components: {
    ModuleForm,
    LoadMoreBtn
  },
  data: () => {
    return {
      modules: [],
      offset: 0,
      count: 10,
      is_show_module_dialog: false,
      form: "",
      loadmore_btn_text: "",
      loading: false
    };
  },
  created() {
    this.loadmore_btn_text = this.config.loadmore_text;
    this.getCommonModules();
  },
  methods: {
    //获取公共模块列表
    getCommonModules() {
      this.$axios({
        method: "GET",
        url: `/api/v1/modules/common/${this.offset}/${this.count}`
      })
        .then(resp => {
          this.modules.push(...resp.data.data);
          this.loading = false;
          this.offset = this.modules.length
          if (resp.data.data.length < this.count) {
            this.loadmore_btn_text = this.config.nomore_text;
          }
        })
        .catch(err => {
          console.log(err);
        });
    },
    loadmore() {
      this.loading = true;
      this.getCommonModules();
    },
    onClickDetailBtn(row) {
      this.$axios({
        method: "GET",
        url: `/api/v1/module/${row.id}`
      })
        .then(resp => {
          this.form = resp.data.data;
        })
        .catch(err => {
          console.log(err);
        });
      this.is_show_module_dialog = true;
    }
  }
};
</script>