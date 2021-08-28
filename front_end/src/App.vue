<template>
  <div style="height:100%;">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: "App",
  components: {}
};
</script>

<style>
.content-wrapper {
  padding: 10px;
}
html,
body {
  height: 100%;
  background: black;
  color: green;
  margin: 0px;
  padding: 5px;
  box-sizing: border-box;
}
.container {
  width: 100%;
  height: 100%;
}
.border {
  border: 1px solid green;
}
.m-2 {
  margin: 20px;
}
.m-rl-auto {
  margin-right: auto;
  margin-left: auto;
}
.m-tb-2 {
  margin-bottom: 20px;
  margin-top: 20px;
}
.login {
  width: 40%;
  max-width: 300px;
  min-width: 250px;
}
.p-2 {
  padding: 20px;
}
.flex {
  display: flex;
}
.flex-dir-col {
  flex-direction: column;
}
.flex-center {
  justify-content: center;
}
.flex-grow {
  flex-grow: 1;
}
.flex-space-between {
  justify-content: space-around;
}
.btn-green {
  border: 1px dotted green;
  width: 100%;
  margin: 10px;
  height: 35px;
  background: black;
  color: green;
  cursor: pointer;
}
.btn-green:hover {
  background: green;
  color: black;
}
.input-yellow {
  background: transparent;
  color: yellow;
  border: 1px hidden green;
  border-bottom: 1px dashed green;
  outline: none;
}
.alert {
  text-align: center;
  color: yellow;
  font-size: 14px;
}
.el-aside {
  width: 200px !important;
}
.el-container {
  flex-grow: 1;
  overflow: scroll;
}
.el-footer {
  height: 60px;
  line-height: 60px;
  text-align: center;
}
</style>