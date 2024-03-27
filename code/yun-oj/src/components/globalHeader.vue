<template>
  <div class="header-header">
    <div class="logo">
      <el-image style="height: 55px" :src="url" fit="scale-down"/>
    </div>
    <div style="display: flex; justify-content: space-between; width: 100%;">
      <div class="functional-blocks">
        <div class="functional-blocks-block">Function 1</div>
        <div class="functional-blocks-block">Function 2</div>
        <div class="functional-blocks-block">Function 3</div>
        <div class="functional-blocks-block">Function 4</div>
      </div>
      <div class="hhh123">头像</div>
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
      url: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
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
  display: flex;
  //border: 1px solid red;
}

.functional-blocks-block {
  margin-right: 13px;
  margin-left: 13px;
  //text-align: center;
  vertical-align: middle;
}
</style>
