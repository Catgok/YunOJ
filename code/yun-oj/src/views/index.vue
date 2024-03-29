<template>
  <div class="common-layout">
    <el-container>
      <el-header class="global-header">
        <GlobalHeader/>
      </el-header>
      <el-scrollbar class="main-content">
        <el-main class="global-main">
          <router-view></router-view>
        </el-main>
      </el-scrollbar>
      <!--      <el-footer class="global-footer">-->
      <!--        Footer-->
      <!--      </el-footer>-->
    </el-container>
  </div>
</template>

<script>

import {ElContainer, ElHeader, ElMain, ElFooter, ElScrollbar} from "element-plus";
import GlobalHeader from "@/components/globalHeader.vue";

export default {
  name: "AppIndex",
  components: {ElContainer, ElHeader, ElMain, ElFooter, ElScrollbar, GlobalHeader},
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
}

.global-main {
  border: 1px red solid;
  background-color: #ffffff;
  margin: 0 20px 0 20px;
  width: 80vw;
}

.main-content {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0;
  border: 1px red solid;
}

.global-footer {
  border: 1px red solid;
  background-color: #f0f0f0;
  height: 40px;
}
</style>