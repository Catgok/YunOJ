<template>
  <div class="header-header">
    <div class="logo">
      <el-image style="height: 55px" :src="url" fit="scale-down"/>
    </div>
    <div style="display: flex; justify-content: space-between; width: 100%;">
      <div class="functional-blocks">
        <div v-for="block in functionalBlocks" :key="block.id" class="functional-blocks-block" @click="clickk(block.path)">
          {{ block.name }}
        </div>
      </div>
      <div class="hhh123">
        <div>头像</div>
      </div>
    </div>
  </div>
</template>

<script>
import {ElImage} from "element-plus";

export default {
  name: 'GlobalHeader',
  components: {ElImage},
  data() {
    return {
      url: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
      functionalBlocks: [
        {id: 1, name: '首页',path: '/'},
        {id: 2, name: '题目',path: '/problem'},
        {id: 3, name: '竞赛',path: '/contest'},
        {id: 4, name: '公告',path: '/announcement'},
        {id: 5, name: '在线IDE',path: '/onlineIDE'},
      ]
    }
  },
  methods: {
    clickk(path) {
      this.$router.push(path)
    }
  }
}
</script>

<style scoped>
.header-header {
  padding-left: 8vw;
  padding-right: 8vw;
  display: flex;
}

.logo {
  //border: 1px solid red;
}

.functional-blocks {
  margin-left: 30px;
  display: flex;
  //border: 1px solid red;
}

.functional-blocks-block {
  margin-right: 13px;
  margin-left: 13px;
  width: 60px;
  //text-align: center;
  vertical-align: middle;
}
</style>
