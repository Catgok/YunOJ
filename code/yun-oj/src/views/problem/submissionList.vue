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
import {ElTable, ElTableColumn, ElButton} from "element-plus";

export default {
  components: {ElTable, ElTableColumn, ElButton},
  data() {
    return {
      problemId: this.$route.params.problemId,
      submissionData: []
    }
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
      for (let i = 1; i <= 10; i++) {
        let item = {
          submissionId: 1000000 + i,
          problemId: 1000 + i,
          problemName: '题目' + i,
          submitTime: '2021-10-10 10:10:10',
          status: 'Accepted',
          language: 'C++',
          runtime: '100ms',
          memory: '100MB',
          codeLength: '1000B',
        }
        this.submissionData.push(item)
      }
    }
  },
  created() {
    this.loadSubmissionData()
  }
}
</script>

<style scoped>

</style>