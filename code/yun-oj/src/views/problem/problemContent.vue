<template>
  <div>
    <div style="text-align:left; margin:0 15px 15px 5px;font-size: 30px">
      {{ this.$route.params.problemId }}. {{ problemInfo.title }}
    </div>
    <div style="display: flex">
      <div style="flex: 10;border-top: 1px #938c8c solid;">
        <div style="text-align:left;" v-html="renderedMarkdown"></div>
      </div>
      <div style="flex: 4;">
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

      <div style="display: flex;justify-content: space-between;margin: 20px 0 1px 0;background-color: #acacac">
        <div>在线IDE</div>
        <div style="margin-right: 15px">
          <div>C++</div>
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

      <div style="display: flex;justify-content: flex-end;margin: 25px 10px 15px 0">
        <el-button>运行</el-button>
        <el-button>提交</el-button>
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
import {autoTextarea} from "@/utils/utils";

import {basicLight} from '@uiw/codemirror-theme-basic/light'
import {Codemirror} from 'vue-codemirror'
import {ElButton, ElInput} from "element-plus"

export default {
  components: {Codemirror, ElInput, ElButton},
  data() {
    return {
      codemirrorExtensions: [cpp(), autocompletion(), basicLight],
      inputCode: '',
      inputCase: '',
      outputCase: '运行后会显示输出哦~~',
      problemInfo: {
        problemId: '',
        title: 'A+B Problem',
        content: '输入两个整数，求这两个整数的和是多少。\n### 输入格式\n输入两个整数$$A$$, $$B$$, 用空格隔开\n\n### 输出格式\n输出一个整数，表示这两个数的和\n\n### 数据范围\n$$0 <= A,B <= 10^8$$\n### 输入样例\n``` text\n3 4\n```\n### 输出样例\n ```text\n7\n```',
        timeLimit: 1000,
        memoryLimit: 256,
        passCount: 100,
        tryCount: 120,
        hardLevel: 1,
      },
    }
  },
  mounted() {
    this.problemInfo.problemId = this.$route.params.problemId
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
      console.log('change', this.inputCode)
    },
    codeFocus() {
      console.log('focus', this.inputCode)
    },
    codeBlur() {
      console.log('blur', this.inputCode)
    },
    problemSubmission() {
      const path = `/problem/submission/${this.problemInfo.problemId}`
      this.$router.push(path)
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
  width: calc(40vw - 20px);
  min-height: 148px;
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
  //margin: 10px 1px 10px 1px;
}

.problem-attach-info-0 {
  background-color: #f0f0f0;
  border-top: 1px #938c8c solid;
  border-bottom: 1px #938c8c solid;
}

.problem-submission-button:hover {
  cursor: pointer;
  color: blue;
  //background-color: #938c8c;
}

.put-text-tip-line {
  //padding-left: 5px;
  text-align: left;
  border-left: red 1px solid;
  background-color: #acacac;
}

.put-text-tip-line-text {
  width: 100px;
  background-color: white;
}
</style>