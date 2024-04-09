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

import {ElContainer, ElFooter, ElHeader, ElMain, ElScrollbar} from "element-plus";
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
  background-image: url('@/assets/bg.png');
  background-size: cover;
}

.global-header {
  background-color: rgb(255, 241, 118, 0.75);
  height: 56px;
}

.global-main {
  margin: 2vw 20px 10vw 20px;
  width: 80vw;
  border-radius: 10px;
  padding: 30px;
  background-color: rgba(255, 255, 255, 0.75);
  backdrop-filter: blur(5px);
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.25); /* 设置阴影效果 */
}

.main-content {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0;
}

.global-footer {
  border: 1px red solid;
  background-color: #f0f0f0;
  height: 40px;
}
</style>