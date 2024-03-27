<template>
  <div class="common-layout">
    <el-container>
      <el-header class="global-header">
        <GlobalHeader/>
      </el-header>
      <el-main class="global-main">
        <div class="main-content">
          Main
        </div>
      </el-main>
<!--      <el-footer class="global-footer">-->
<!--        Footer-->
<!--      </el-footer>-->
    </el-container>
  </div>
</template>

<script>

import {ElContainer, ElHeader, ElMain, ElFooter} from "element-plus";
import GlobalHeader from "@/components/globalHeader.vue";

export default {
  name: "AppIndex",
  components: {ElContainer, ElHeader, ElMain, ElFooter, GlobalHeader},
}
</script>

<style scoped>
.common-layout {
  height: 100vh;
  display: flex;
}

.global-header {
  background-color: #f0f0f0;
  //border: 1px red solid;
  height: 56px;
  margin-bottom: 20px;
}

.global-main {
  flex: 1;
  border: 1px red solid;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0;
  background-color: #ffffff;
}

.main-content {
  border: 1px red solid;
  text-align: center;
  width: 80vw;
  height: 100%;
}

.global-footer {
  border: 1px red solid;
  background-color: #f0f0f0;
  height: 40px;
}
</style>