<template>
  <div>
    <div style="text-align:left; margin:0 15px 15px 5px;font-size: 30px;display: flex;justify-content: space-between">
      <div>{{ problemId }}. {{ problemTitle }}</div>
      <div style="font-size: 20px;  display: flex;justify-content: center;align-items: flex-end;"
           @click="returnProblem()"> 返回题目
      </div>
    </div>
    <div style="border-top: 1px #938c8c solid;">
      <div>
        <el-table :data="submissionData" stripe style="width: 100%;">
          <el-table-column prop="submitTime" label="提交时间" width="250"/>
          <el-table-column prop="status" label="评测状态"></el-table-column>
          <el-table-column prop="runtime" label="运行时间" width="120"/>
          <el-table-column prop="memory" label="运行内存" width="120"/>
          <el-table-column prop="language" label="语言" width="100"/>
          <el-table-column width="100">
            <template #default="scope">
              <el-button type="text" @click="showCode(scope.row.submissionId)">查看代码</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>

  </div>
</template>

<script>
import {ElButton, ElTable, ElTableColumn} from "element-plus";

export default {
  components: {ElTable, ElTableColumn, ElButton},
  data() {
    return {
      problemId: this.$route.params.problemId,
      problemTitle: '',
      submissionData: []
    }
  },
  created() {
    // console.log(this.$store.state)
    this.problemId = this.$route.params.problemId
    this.problemTitle = sessionStorage.getItem('problemTitle.' + this.problemId)

    this.loadSubmissionData()
  },
  methods: {
    returnProblem() {
      const path = `/problem/${this.problemId}`;
      this.$router.push(path);
    },
    showCode(submissionId) {
      const path = `/problem/submission/${this.problemId}/${submissionId}`;
      this.$router.push(path);
    },

    loadSubmissionData() {
      const userInfo = JSON.parse(localStorage.getItem('userInfo'))
      const req = {
        userId: userInfo.userId,
        problemId: parseInt(this.problemId)
      }

      this.$axios.post('/problem/getSubmissionByProblemId', req).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          // todo
          return
        }
        if (resp.data === null || resp.data.length === 0) {
          return
        }
        resp.data.forEach((item) => {
          this.submissionData.push({
            submissionId: item.submitId,
            problemId: item.problemId,
            problemName: item.problemName,
            submitTime: '2021-10-10 10:10:10', // todo
            status: item.status,
            language: item.language,
            runtime: (item.time / 1000000).toFixed(3) + 'ms',
            memory: (item.memory / 1024 / 1024).toFixed(3) + 'MB',
            code: item.code,
          })
        })
      })
    }
  },
}
</script>

<style scoped>

</style>