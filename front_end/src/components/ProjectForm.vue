<template>
  <div>
    <el-form :model="form">
      <el-form-item label="名称" :label-width="formLabelWidth">
        <el-input v-model="form.name" autocomplete="off" maxlength="10" show-word-limit></el-input>
      </el-form-item>
      <el-form-item label="描述" :label-width="formLabelWidth">
        <el-input
          type="textarea"
          :autosize="{minRows:3,maxRows:5}"
          v-model="form.desc"
          autocomplete="off"
          maxlength="300"
          show-word-limit
        ></el-input>
      </el-form-item>
      <el-form-item label="模块" :label-width="formLabelWidth">
        <el-table
          ref="multipleTable"
          :data="form.module_list"
          tooltip-effect="dark"
          style="width: 100%"
          @select-all="onSelectAll"
          @select="onSelect"
          @expand-change="onExpandChange"
        >
          <el-table-column type="selection" width="55" align="center"></el-table-column>
          <el-table-column prop="name" label="名称" width="120"></el-table-column>
          <el-table-column prop="option_list" label="配置项" align="center">
            <template slot-scope="scope">
              <el-form v-show="scope.row.is_choosed" ref="form" label-width="80px">
                <el-form-item
                  :key="i"
                  v-for="(option,i) in scope.row.option_list"
                  :label="option.name"
                >
                  <el-input v-model="option.value" size="small"></el-input>
                </el-form-item>
              </el-form>
              <div v-show="scope.row.is_choosed==false && scope.row.option_list.length>0">...</div>
              <div v-if="scope.row.option_list.length==0">无</div>
            </template>
          </el-table-column>
          <el-table-column type="expand">
            <template slot-scope="props">
              <el-input
                type="textarea"
                :autosize="{minRows:10,maxRows:20}"
                :value="getXSSPayload(props.row.id)"
                :disabled="true"
              ></el-input>
            </template>
          </el-table-column>
        </el-table>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  props: ["form"],
  data: () => {
    return {
      formLabelWidth: "80px",
      form: "",
      is_show_module_dialog: false,
      module_details: []
    };
  },
  updated: function() {
    for (let i = 0; i < this.form.module_list.length; i++) {
      //将is_choosed 的值和一行是否选中单向绑定
      this.$refs.multipleTable.toggleRowSelection(
        this.form.module_list[i],
        this.form.module_list[i].is_choosed
      );
    }
  },
  methods: {
    getXSSPayload(id) {
      let index = this.module_details.findIndex(item => {
        return item.id == id;
      });
      if (index != -1) {
        return this.module_details[index].xss_payload;
      } else {
        return "";
      }
    },
    // 当全部选中时
    onSelectAll() {
      for (let i = 0; i < this.form.module_list.length; i++) {
        this.form.module_list[i].is_choosed = !this.form.module_list[i]
          .is_choosed;
      }
    },
    //当手动选中某一项时
    onSelect(selection, row) {
      let index = this.form.module_list.findIndex(item => {
        return item.id == row.id;
      });
      this.form.module_list[index].is_choosed = !this.form.module_list[index]
        .is_choosed;
    },
    //当点击扩展按钮时
    onExpandChange(row) {
      this.getModuleDetail(row.id)
    },
    //获取模块详情
    getModuleDetail(id) {
      this.$axios({
        method: "GET",
        url: `/api/v1/module/${id}`
      })
        .then(resp => {
          let index = this.module_details.findIndex(item => {
            return item.id == id;
          });
          if (-1 == index) {
            this.module_details.push(resp.data.data);
          }
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>

<style scoped>
.demo-table-expand {
  font-size: 0;
}
.demo-table-expand label {
  width: 90px;
  color: #99a9bf;
}
.demo-table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 50%;
}
</style>