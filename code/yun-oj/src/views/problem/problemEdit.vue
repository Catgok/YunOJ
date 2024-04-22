<template>
  <div>
    <div>编辑题目</div>
    <div style="width: 100%">
      <el-form :model="problemInfo" label-position="left" label-width="auto" style="max-width: 1000px">
        <el-form-item label="题目名称">
          <el-input v-model="problemInfo.title"/>
        </el-form-item>
        <el-form-item label="时/空限制">
          <div style="display: flex">
            <div>时间限制:</div>
            <el-input-number v-model="problemInfo.timeLimit" :step="1" :min="100" :max="10000" step-strictly/>
            ms
          </div>
          <div style="display: flex">
            <div>空间限制:</div>
            <el-input-number v-model="problemInfo.memoryLimit" :step="1" :min="16" :max="2048" step-strictly/>
            MB
          </div>
        </el-form-item>
        <el-form-item label="题目描述">
          <el-input type="textarea" :autosize="{ minRows: 10, maxRows: 100 }" v-model="problemInfo.content"/>
        </el-form-item>
        <el-form-item label="题目预览">
          <div style="width:100%;height:100%;border: 1px black solid;min-height: 20vh">
            <div style="" v-html="renderedMarkdown"></div>
          </div>
        </el-form-item>
        <el-form-item label="上传样例">
          <el-upload
              v-model:file-list="fileList"
              class="upload-demo"
              multiple
              :auto-upload="false"
              :limit="20"
              :on-change="handleChange"
              :on-exceed="handleExceed"
          >
            <el-button type="primary">点击上传样例</el-button>
          </el-upload>
        </el-form-item>
      </el-form>

    </div>
    <div>
      <el-button type="primary" @click="onSubmit">Create</el-button>
      <el-button>Cancel</el-button>
    </div>
  </div>
</template>

<script>
import {ElButton, ElForm, ElFormItem, ElInput, ElInputNumber, ElUpload,} from "element-plus";

export default {
  components: {ElForm, ElFormItem, ElButton, ElInput, ElInputNumber, ElUpload},
  data() {
    return {
      isNewProblem: false,
      problemInfo: {
        problemId: '',
        title: '',
        content: '',
        timeLimit: 1000,
        memoryLimit: 128,
      },
      judgeCases: [],
      fileList: []
    }
  },
  created() {
    this.isNewProblem = this.$route.name === 'problem-new'
  },
  computed: {
    renderedMarkdown() {
      // https://github.com/runarberg/markdown-it-math
      const md = require('markdown-it')().use(require('markdown-it-math'),)
      return md.render(this.problemInfo.content);
    },
  },
  methods: {
    onSubmit() {
      console.log('submit!')
    },
    handleExceed() {
      console.log("最多10个样例")
    },
    handleChange(file) {
      let reader = new FileReader()
      if (typeof FileReader === "undefined") return
      reader.readAsArrayBuffer(file.raw);
      reader.onload = function (e) {
        const fileIntS = new Uint8Array(e.target.result)
        const snippets = new TextDecoder('utf-8').decode(fileIntS)
        console.log(snippets)
      };
    },
  }
}
</script>
<style scoped>

</style>