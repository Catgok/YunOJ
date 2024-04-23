<template>
  <div>
    <div style="display: flex;justify-content: space-between;margin: 20px 0 1px 0;background-color: #DCDCDC">
      <div style="display: flex;align-items: center;font-size: 20px;margin-left: 5px">
        <div>在线IDE</div>
      </div>
      <div style="margin-right: 15px;display: flex">
        <el-select v-model="language" placeholder="C++" style="width: 100px">
          <el-option v-for="item in languageOptions" :key="item.value" :label="item.label" :value="item.value"/>
        </el-select>
        <div style="margin-left: 10px">
          <el-button style="color:green" type="text" @click="onlineJudge"> > 运行</el-button>
        </div>
      </div>
    </div>

    <div class="code-editor">
      <codemirror class="code-editor-codemirror" style="height: 400px;" v-model="inputCode" :autofocus="true"
                  :indent-with-tab="true" :tab-size="4" :extensions="codemirrorExtensions"/>
    </div>

    <!--    <div> {{ runResult }}</div>-->

    <div style="display: flex;margin-top: 20px">
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
</template>

<script>
// import {defineComponent, shallowRef} from 'vue'
import {cpp} from '@codemirror/lang-cpp'
import {autocompletion} from '@codemirror/autocomplete'
import {autoTextarea} from "@/utils/utils";

import {basicLight} from '@uiw/codemirror-theme-basic/light'
import {Codemirror} from 'vue-codemirror'
import {ElButton, ElInput, ElOption, ElSelect} from "element-plus";

export default {
  name: 'OnlineIDE',
  components: {Codemirror, ElInput, ElButton, ElSelect, ElOption},
  data() {
    return {
      inputCode: `#include <iostream>
using namespace std;
int main() {
    cout << "Hello, World!" << endl;
    return 0;
}`,
      codemirrorExtensions: [cpp(), autocompletion(), basicLight],
      inputCase: '',
      outputCase: '',
      language: '',
      languageOptions: [
        {label: 'C++', value: 1,},
        // {label: '其他语言', value: 2,},
      ],
      runResult: ''
    }
  },
  mounted() {
    const sourceTextElement = document.querySelector(".input-case-textarea");
    autoTextarea(sourceTextElement)
    const targetTextElement = document.querySelector(".output-case-textarea");
    autoTextarea(targetTextElement)
  },
  methods: {
    onlineJudge() {
      this.outputCase = '运行中.....'
      const req = {
        code: this.inputCode,
        language: 1,
        input: this.inputCase,
      }

      this.$axios.post('/judge/onlineJudge', req).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          this.outputCase = resp.message
          return
        }
        this.outputCase = resp.data
        // this.runResult = "运行成功"
      })
    },
  },
}
</script>

<style>
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
