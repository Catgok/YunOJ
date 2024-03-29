<template>
  <div>
    <div style="display: flex;justify-content: space-between">
      <div>在线IDE</div>
      <div style="display:flex;">
        <div>C++</div>
        <div>
          <el-button>运行</el-button>
        </div>
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

    <div style="display: flex;">
      <div>
        <div style="text-align: left">输入</div>
        <div class="put-text-show">
          <textarea v-model="inputCase" class="textarea-class input-case-textarea"></textarea>
        </div>
      </div>

      <div>
        <div style="text-align: left">输出</div>
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
import {ElButton, ElInput} from "element-plus";

export default {
  name: 'OnlineIDE',
  components: {Codemirror, ElInput, ElButton},
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
      outputCase: '2323',
    }
  },
  mounted() {
    const sourceTextElement = document.querySelector(".input-case-textarea");
    autoTextarea(sourceTextElement)
    const targetTextElement = document.querySelector(".output-case-textarea");
    autoTextarea(targetTextElement)
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
    }
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
  width: 40vw;
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
  border: 1px red solid;
  overflow-wrap: break-word;
  resize: none;
}

.put-text-show {
  flex: 1;
  flex-basis: 50%;
  font-family: inherit;
  font-size: 18px;
}
</style>
