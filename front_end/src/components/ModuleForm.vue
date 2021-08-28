<template>
  <div>
    <el-form :model="form">
      <el-form-item label="名称" :label-width="formLabelWidth">
        <el-input
          v-model="form.name"
          autocomplete="off"
          :disabled="disabled"
          maxlength="50"
          show-word-limit
        ></el-input>
      </el-form-item>
      <el-form-item label="描述" :label-width="formLabelWidth">
        <el-input
          type="textarea"
          :autosize="{minRows:5,maxRows:8}"
          v-model="form.desc"
          autocomplete="off"
          :disabled="disabled"
          maxlength="300"
          show-word-limit
        ></el-input>
      </el-form-item>
      <FormTags
        @update_list="updateParamList"
        label="传参"
        :closable="disabled?false:true"
        :list="form.param_list"
        :formLabelWidth="formLabelWidth"
      ></FormTags>
      <FormTags
        @update_list="updateOptionList"
        label="配置项"
        :closable="disabled?false:true"
        :list="form.option_list"
        :formLabelWidth="formLabelWidth"
      ></FormTags>
      <el-form-item label="代码" :label-width="formLabelWidth">
        <el-input
          type="textarea"
          :autosize="{ minRows: 10, maxRows: 20}"
          placeholder="xss payload"
          v-model="form.xss_payload"
          :disabled="disabled"
        ></el-input>
      </el-form-item>
      <el-form-item label="是否公开" :label-width="formLabelWidth" >
        <el-switch v-model="form.is_common" :disabled="disabled"></el-switch>
      </el-form-item>
    </el-form>
    <el-alert title="注意事项" :description="notice" type="info" show-icon></el-alert>
  </div>
</template>

<script>
import FormTags from "@/components/FormTags";
export default {
  props: ["form", "disabled"],
  components: {
    FormTags
  },
  data: () => {
    return {
      formLabelWidth: "80px",
      notice: ""
    };
  },
  mounted() {
    this.$axios({
      method: "GET",
      url: "/api/v1/module/notice"
    })
      .then(resp => {
        this.notice = resp.data.data;
      })
      .catch(err => {
        console.log(err);
      });
  },
  methods: {
    updateParamList(list) {
      this.form.param_list = list;
    },
    updateOptionList(list) {
      this.form.option_list = list;
    }
  }
};
</script>

<style scoped>
</style>