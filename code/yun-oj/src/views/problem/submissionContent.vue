<template>
  <div>
    <div style="text-align:left; margin:0 15px 15px 5px;font-size: 30px;display: flex;justify-content: space-between">
      <div>{{ problemId }}. {{ problemTitle }}</div>
      <div style="font-size: 20px;  display: flex;justify-content: center;align-items: flex-end;"
           @click="returnProblem()">
        返回题目
      </div>
    </div>
    <!--    <div style="margin: 5px ">提交记录</div>-->
    <div style="border-top: 1px #938c8c solid;">
      <div class="code-editor" style="pointer-events: none;">
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
        />
      </div>
    </div>

  </div>
</template>

<script>
import {ElButton, ElTable, ElTableColumn} from "element-plus";

import {cpp} from '@codemirror/lang-cpp'
import {basicLight} from '@uiw/codemirror-theme-basic/light'
import {Codemirror} from 'vue-codemirror'

export default {
  components: {ElTable, ElTableColumn, ElButton, Codemirror},
  data() {
    return {
      problemId: '',
      problemTitle: '',
      codemirrorExtensions: [cpp(), basicLight],
      submissionContent: "",
    }
  },
  created() {
    this.problemId = this.$route.params.problemId
    this.problemTitle = sessionStorage.getItem('problemTitle.' + this.problemId)
    this.loadSubmissionContent(this.$route.params.submissionId)
  },
  methods: {
    returnProblem() {
      const path = `/problem/${this.problemId}`;
      this.$router.push(path);
    },
    loadSubmissionContent(submissionId) {
      const req = {
        submitId: parseInt(submissionId)
      }
      this.$axios.post('/problem/getSubmit', req).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          // todo
          return
        }
        this.submissionContent = resp.data.code

      })
    },
  },
}
</script>

<style scoped>

</style>