<template>
  <div>
    <div style="font-size: 28px; border-bottom: 1px #938c8c solid;margin-bottom: 30px">{{ titleInfo }}</div>
    <div style="width: 100%">
      <el-form :model="problemInfo" label-position="left" label-width="auto" style="max-width: 1000px">
        <el-form-item label="题目名称">
          <el-input v-model="problemInfo.title"/>
        </el-form-item>
        <el-form-item label="时/空限制">
          <div style="display: flex;width: 100%">
            <div style="display: flex;flex: 1">
              <div style="margin-right: 20px">时间限制(ms) ：</div>
              <el-input-number v-model="problemInfo.timeLimit" :step="1" :min="100" :max="10000" step-strictly/>
            </div>
            <div style="display: flex;flex: 1">
              <div style="margin-right: 20px">空间限制(MB) ：</div>
              <el-input-number v-model="problemInfo.memoryLimit" :step="1" :min="16" :max="2048" step-strictly/>
            </div>
          </div>
        </el-form-item>
        <el-form-item label="题目难度">
          <el-select v-model="problemInfo.hardLevel" style="width: 180px">
            <el-option v-for="item in hardLevelOption" :key="item.value" :label="item.label" :value="item.value"/>
          </el-select>
        </el-form-item>
        <el-form-item label="题目描述">
          <el-input type="textarea" :autosize="{ minRows: 10, maxRows: 100 }" v-model="problemInfo.content"/>
        </el-form-item>
        <el-form-item label="描述预览">
          <div style="width:100%;height:100%;border: 1px black solid;min-height: 20vh;padding-left:10px">
            <div style="" v-html="renderedMarkdown"></div>
          </div>
        </el-form-item>
        <el-form-item label="上传样例">
          <div style="min-height: 10vh;display: flex;width: 100%">
            <el-upload v-model:file-list="inputFileList" class="case-upload" multiple :auto-upload="false" :limit="20"
                       :on-change="handleInputChange" :on-exceed="handleExceed">
              <el-button>上传输入样例</el-button>
            </el-upload>
            <el-upload v-model:file-list="outputFileList" class="case-upload" multiple :auto-upload="false" :limit="20"
                       :on-change="handleOutputChange" :on-exceed="handleExceed">
              <el-button>上传输出样例</el-button>
            </el-upload>
          </div>
        </el-form-item>
      </el-form>
    </div>
    <div style="text-align: right">
      <el-button type="primary" @click="onSubmitProblemInfo">{{ submitProblemInfo }}</el-button>
      <el-button type="primary" @click="onSubmitProblemCase">{{ submitCaseInfo }}</el-button>
    </div>
  </div>
</template>

<script>
import {ElButton, ElForm, ElFormItem, ElInput, ElInputNumber, ElOption, ElSelect, ElUpload,} from "element-plus";
import {hardLevelMap} from "@/utils/globalStaticData";

export default {
  components: {ElForm, ElFormItem, ElButton, ElInput, ElInputNumber, ElUpload, ElSelect, ElOption},
  data() {
    return {
      isNewProblem: false,
      titleInfo: '',
      submitProblemInfo: '',
      submitCaseInfo: '更新样例',
      problemInfo: {problemId: '', title: '', content: '', timeLimit: 1000, memoryLimit: 128, hardLevel: 1,},
      judgeCases: {input: [], output: []},
      inputFileList: [],
      outputFileList: [],
      hardLevelOption: [
        {value: 1, label: hardLevelMap[1]},
        {value: 2, label: hardLevelMap[2]},
        {value: 3, label: hardLevelMap[3]},
      ]
    }
  },
  created() {
    this.isNewProblem = this.$route.name === 'problemNew'
    if (this.isNewProblem) {
      this.titleInfo = '新建题目'
      this.submitProblemInfo = '创建题目'
    } else {
      this.titleInfo = '编辑题目'
      this.submitProblemInfo = '提交修改'
      this.getProblemInfo(this.$route.params.problemId)
    }
  },
  computed: {
    renderedMarkdown() {
      // https://github.com/runarberg/markdown-it-math
      const md = require('markdown-it')().use(require('markdown-it-math'),)
      return md.render(this.problemInfo.content);
    },
  },
  methods: {
    onSubmitProblemInfo() {
      // todo
      console.log(this.problemInfo);
    },
    onSubmitProblemCase() {
      // todo
      console.log(this.judgeCases);
    },
    handleExceed() {
      console.log("最多10个样例")
    },
    handleInputChange(file) {
      this.getFileContent(file).then(data => {
        this.judgeCases.input.push(data)
      })
    },
    handleOutputChange(file) {
      this.getFileContent(file).then(data => {
        this.judgeCases.output.push(data)
      })
    },
    handleOutputRemove(file) {
      console.log(file)
    },
    getFileContent(file) {
      return new Promise((resolve, reject) => {
        let reader = new FileReader();
        if (typeof FileReader === "undefined") {
          reject("FileReader is not supported in this browser.");
          return;
        }
        reader.readAsArrayBuffer(file.raw);
        reader.onload = function (e) {
          const fileIntS = new Uint8Array(e.target.result);
          const snippets = new TextDecoder('utf-8').decode(fileIntS);
          resolve(snippets);
        };
        reader.onerror = function (error) {
          reject(error);
        };
      });
    },
    getProblemInfo(problemId) {
      const req = {
        problemId: parseInt(problemId),
      }
      this.$axios.post('/problem/get', req).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          return
        }
        this.problemInfo.problemId = resp.data.problemId
        this.problemInfo.title = resp.data.title
        this.problemInfo.content = resp.data.description
        this.problemInfo.timeLimit = resp.data.timeLimit
        this.problemInfo.memoryLimit = resp.data.memoryLimit
        this.problemInfo.hardLevel = hardLevelMap[resp.data.hardLevel]
      })
    },
  }
}
</script>
<style scoped>
.case-upload {
  flex: 1;
}
</style>