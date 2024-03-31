<template>
  <div>
    <div style="text-align:left; margin:0 15px 15px 5px;font-size: 30px;display: flex;justify-content: space-between">
      <div>{{ problemId }}. A+B Problem</div>
      <div style="font-size: 20px;  display: flex;justify-content: center;align-items: flex-end;"
           @click="returnProblem()">
        返回题目
      </div>
    </div>
    <div style="border-top: 1px #938c8c solid;">
      <div style="text-align:left;" v-html="renderedMarkdown"> </div>
      <div class="code-editor">
        <codemirror
            class="code-editor-codemirror"
            style="height: 400px;"
            v-model="submissionContent"
            placeholder="打出你的代码吧！"
            :autofocus="true"
            :indent-with-tab="true"
            :tab-size="4"
            :extensions="codemirrorExtensions"
            :options="{readOnly: true}"
            @change="codeChange"
            @focus="codeFocus"
            @blur="codeBlur"
        />
      </div>
    </div>

  </div>
</template>

<script>
import {ElTable, ElTableColumn, ElButton} from "element-plus";

import {cpp} from '@codemirror/lang-cpp'
import {autocompletion} from '@codemirror/autocomplete'

import {basicLight} from '@uiw/codemirror-theme-basic/light'
import {Codemirror} from 'vue-codemirror'

export default {
  components: {ElTable, ElTableColumn, ElButton, Codemirror},
  data() {
    return {
      problemId: this.$route.params.problemId,
      codemirrorExtensions: [cpp(), autocompletion(), basicLight],
      submissionContent: '```cpp \n#include <iostream>\nusing namespace std;\nint main() {\n    int a, b;\n    cin >> a >> b;\n    cout << a + b << endl;\n    return 0;\n}\n```\n',
    }
  },
  computed: {
    renderedMarkdown() {
      const md = require('markdown-it')()
      return md.render(this.submissionContent);
    },
  },
  methods: {
    returnProblem() {
      const path = `/problem/${this.problemId}`;
      this.$router.push(path);
    }
  },
}
</script>

<style scoped>

</style>