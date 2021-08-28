<template>
  <div class="content-wrapper">
    <AddModuleBtn add_module_url="/api/v1/admin/modules" @addModule="addModule"></AddModuleBtn>
    <ModuleTable
      get_modules_url="/api/v1/admin/modules"
      del_module_url="/api/v1/admin/module"
      update_module_url="/api/v1/admin/module"
      :modules = "modules"
    ></ModuleTable>
  </div>
</template>

<script>
import ModuleTable from "@/components/ModuleTable";
import AddModuleBtn from "@/components/AddModuleBtn";
export default {
  components: {
    ModuleTable,
    AddModuleBtn
  },
  data: () => {
    return {
      modules: []
    };
  },
  methods: {
    //addModule 添加模块成功,更新前端维护的数据
    addModule(module) {
      this.modules.push(module);
    },
  }
};
</script>

