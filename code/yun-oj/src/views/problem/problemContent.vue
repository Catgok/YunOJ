<template>
  <div>
    <div style="text-align:left; margin:0 15px 15px 5px;font-size: 30px">
      {{ this.$route.params.problemId }}. {{ problemInfo.title }}
    </div>
    <div style="display: flex">
      <div style="flex: 20;border-top: 1px #938c8c solid;">
        <div style="text-align:left;" v-html="renderedMarkdown"></div>
      </div>
      <div style="flex: 8;">
        <div class="problem-attach-info problem-attach-info-0">
          <div>时/空限制 :</div>
          <div>{{ problemInfo.timeLimit }} ms / {{ problemInfo.memoryLimit }} MB</div>
        </div>
        <div class="problem-attach-info">
          <div>总通过/总尝试 :</div>
          <div>{{ problemInfo.passCount }} / {{ problemInfo.tryCount }}</div>
        </div>
        <div class="problem-attach-info problem-attach-info-0">
          <div>难度 :</div>
          <div>{{ problemInfo.hardLevel }}</div>
        </div>
        <div @click="problemSubmission()" class="problem-attach-info problem-submission-button">
          <div>提交记录</div>
          <!--          <el-button link @click="problemSubmission()" style="margin: 0;border: 0">提交记录</el-button>-->
        </div>
      </div>
    </div>
    <div>

      <div style="display: flex;justify-content: space-between;margin: 20px 0 1px 0;background-color: #DCDCDC">
        <div style="display: flex;align-items: center;font-size: 20px;margin-left: 5px">
          <div>在线IDE</div>
        </div>
        <div>
          <el-select v-model="language" placeholder="C++" style="width: 100px">
            <el-option v-for="item in languageOptions" :key="item.value" :label="item.label" :value="item.value"/>
          </el-select>
        </div>
      </div>

      <div class="code-editor">
        <codemirror
            class="code-editor-codemirror"
            style="height: 400px;"
            v-model="inputCode"
            placeholder="打出你的代码吧！"
            :autofocus="true"
            :indent-with-tab="true"
            :tab-size="4"
            :extensions="codemirrorExtensions"
            @change="codeChange"
            @focus="codeFocus"
            @blur="codeBlur"
        />
      </div>

      <div style="display: flex;justify-content: space-between;align-items: center">
        <div></div>
        <div style="display: flex;justify-content: flex-end;margin: 25px 10px 15px 0">
          <el-button @click="runCode">运行</el-button>
          <el-button @click="submitCode">提交</el-button>
        </div>
      </div>

      <div style="display: flex;">
        <div>
          <div class="put-text-tip-line">
            <div class="put-text-tip-line-text"><\>输入</div>
          </div>
          <div class="put-text-show">
            <textarea v-model="inputCase" class="textarea-class input-case-textarea"></textarea>
          </div>
        </div>

        <div>
          <div class="put-text-tip-line">
            <div class="put-text-tip-line-text"><\>输出</div>
          </div>
          <div class="put-text-show">
            <textarea v-model="outputCase" readonly class="textarea-class output-case-textarea"></textarea>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {cpp} from '@codemirror/lang-cpp'
import {autocompletion} from '@codemirror/autocomplete'
import {autoTextarea, getLoginStatus} from "@/utils/utils";

import {basicLight} from '@uiw/codemirror-theme-basic/light'
import {Codemirror} from 'vue-codemirror'
import {ElButton, ElInput, ElOption, ElSelect} from "element-plus"

export default {
  components: {Codemirror, ElInput, ElButton, ElSelect, ElOption},
  data() {
    return {
      codemirrorExtensions: [cpp(), autocompletion(), basicLight],
      inputCode: '',
      inputCase: '',
      outputCase: '运行后会显示输出哦~~',
      problemInfo: {
        problemId: '',
        title: '',
        content: ''
      },
      language: '',
      languageOptions: [
        {label: 'C++', value: 1,},
        // {label: '其他语言', value: 2,},
      ],
      submitId: '',
    }
  },
  created() {
    this.problemInfo.problemId = this.$route.params.problemId
    this.problemInfo.title = sessionStorage.getItem('problemTitle.' + this.problemInfo.problemId)
    this.getProblemInfo(this.problemInfo.problemId)
  },
  mounted() {
    const sourceTextElement = document.querySelector(".input-case-textarea")
    const targetTextElement = document.querySelector(".output-case-textarea")
    autoTextarea(sourceTextElement)
    autoTextarea(targetTextElement)
  },
  computed: {
    renderedMarkdown() {
      // https://github.com/runarberg/markdown-it-math
      const md = require('markdown-it')().use(require('markdown-it-math'),)
      return md.render(this.problemInfo.content);
    },
  },
  methods: {
    codeChange() {
      // console.log('change', this.inputCode)
    },
    codeFocus() {
      // console.log('focus', this.inputCode)
    },
    codeBlur() {
      // console.log('blur', this.inputCode)
    },
    problemSubmission() {
      const path = `/problem/submission/${this.problemInfo.problemId}`
      this.$router.push(path)
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
        this.problemInfo.passCount = resp.data.passCount
        this.problemInfo.tryCount = resp.data.submitCount
        this.problemInfo.hardLevel = resp.data.hardLevel
        sessionStorage.setItem('problemTitle.' + this.problemInfo.problemId, this.problemInfo.title)
      })
    },
    runCode() {
      if (getLoginStatus() === false) {
        this.$router.push('/login')
        return
      }
      const req = {
        code: this.inputCode,
        language: 1,
        input: this.inputCase,
      }
      this.$axios.post('/judge/onlineJudge', req).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          // todo
          return
        }
        this.outputCase = resp.data
      })
    },
    submitCode() {
      if (getLoginStatus() === false) {
        this.$router.push('/login')
        return
      }
      const userInfo = JSON.parse(localStorage.getItem('userInfo'))
      const req = {
        problemId: this.problemInfo.problemId,
        userId: userInfo.userId,
        code: this.inputCode,
        language: 1,
      }
      this.$axios.post('/problem/submit', req).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          // todo
          return
        }
        this.submitId = resp.data
      })
    }
  }
}
</script>


<style scoped>
.code-editor {
  width: 100%;
}

.code-editor-codemirror {
  text-align: left;
}

.textarea-class {
  width: calc(40vw - 30px);
  min-height: 150px;
  _height: 120px;
  margin: 0;
  outline: 0;
  font-family: inherit;
  font-size: 18px;
  line-height: 24px;
  padding: 0;
  word-wrap: break-word;
  overflow-x: hidden;
  overflow-y: auto;
  border: 1px gray solid;
  overflow-wrap: break-word;
  resize: none;
}

.put-text-show {
  flex: 1;
  flex-basis: 50%;
  font-family: inherit;
  font-size: 18px;
}

.problem-attach-info {
  border-left: 1px #938c8c solid;
  display: flex;
  justify-content: space-between;
  padding: 20px 16px 20px 8px;
}

.problem-attach-info-0 {
  background-color: #D3D3D3;
  border-top: 1px #938c8c solid;
  border-bottom: 1px #938c8c solid;
}

.problem-submission-button:hover {
  cursor: pointer;
  color: blue;
}

.put-text-tip-line {
  padding-left: 1px;
  text-align: left;
  border-left: #DCDCDC 1px solid;
  background-color: #DCDCDC;
}

.put-text-tip-line-text {
  border-top: #DCDCDC 1px solid;
  width: 100px;
  background-color: white;
}
</style>