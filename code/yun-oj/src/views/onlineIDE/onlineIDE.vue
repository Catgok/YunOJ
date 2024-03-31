<template>
  <div>
<!--    todo 宽度-->
    <div style="display: flex;justify-content: space-between;margin: 20px 0 1px 0;background-color: #acacac">
      <div>在线IDE</div>
      <div style="margin-right: 15px;display: flex">
        <div>C++</div>
        <div>
          <el-button type="text"> 运行</el-button>
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

    <!--    <div style="display: flex;justify-content: flex-end;margin: 25px 10px 15px 0">-->
    <!--      <el-button>运行</el-button>-->
    <!--      <el-button>提交</el-button>-->
    <!--    </div>-->
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

.put-text-tip-line {
  //padding-left: 5px;
  text-align: left;
  border-left: red 1px solid;
  background-color: #acacac;
}

.put-text-tip-line-text {
  border-top: #938c8c 1px solid;
  width: 100px;
  background-color: white;
}
</style>
