<template>
  <div>
    <el-form :model="form">
      <el-form-item label="密   码" :label-width="formLabelWidth">
        <el-input v-model="form.password" autocomplete="off" show-password></el-input>
      </el-form-item>
      <el-form-item label="确认密码" :label-width="formLabelWidth">
        <el-input v-model="form.repassword" show-password></el-input>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  props: ["form"],
  data: () => {
    return {
      formLabelWidth: "80px"
    };
  }
};
</script>

<style scoped>
</style>