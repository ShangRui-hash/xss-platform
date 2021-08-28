<template>
  <el-menu default-active="0" style="height:100%" router>
    <el-menu-item :key="item.index" v-for="(item) in navbar" :index="item.to">
      <i :class="item.icon"></i>
      {{item.title}}
    </el-menu-item>
  </el-menu>
</template>

<script>
export default {
  props: ["navbar"]
};
</script>