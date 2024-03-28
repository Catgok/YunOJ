<template>
  <div>
    <div> title {{ this.$route.params.problemId }}</div>
    <div style="display: flex">
      <div style="flex: 8;border: 1px red solid">题目描述</div>
      <div style="flex: 3;border: 1px red solid">
        <div>时/空限制:</div>
        <div>通过/尝试:</div>
        <div>难度:</div>
        <div>提交记录</div>
      </div>
    </div>
    <div>
      <div>代码框</div>
      <div>运行&提交</div>
      <div style="display: flex;flex-direction: row">
        <div>输入</div>
        <div>输出</div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  components: {},
  data() {
    return {}
  },
  methods: {},
  created() {
  }
}
</script>

<style scoped>

</style>