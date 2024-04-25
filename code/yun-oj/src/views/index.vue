<template>
  <div class="common-layout">
    <global-notice></global-notice>
    <el-container>
      <el-header class="global-header">
        <global-header/>
      </el-header>
      <el-scrollbar class="main-content">
        <el-main :style="computedMainStyle" class="global-main">
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
import GlobalNotice from "@/components/globalNotice.vue";

export default {
  name: "AppIndex",
  components: {ElContainer, ElHeader, ElMain, ElFooter, ElScrollbar, GlobalHeader, GlobalNotice},
  data() {
    return {
      isMainRoute: true,
    }
  },
  computed: {
    computedMainStyle() {
      if (!this.isMainRoute) {
        return {
          backgroundColor: `rgba(255, 255, 255, 0.75)`,
          backdropFilter: `blur(0.5px)`,
          marginBottom: `10vw`,
          boxShadow: `0 0 10px rgba(0, 0, 0, 0.25)`
        };
      }
    }
  },
  watch: {
    $route(to, from) {
      this.isMainRoute = to.name === 'main';
    }
  },
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
  margin: 2vw 20px 0 20px;
  width: 80vw;
  border-radius: 10px;
  padding: 30px;
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