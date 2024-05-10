<template>
  <div>
    <div>
      <el-table :data="contestProblemsInfo" stripe style="width: 100%;" table-layout="auto">
        <el-table-column type="index" :index="contestProblemsListIndexMethod"/>
        <el-table-column prop="title" label="题目名称">
          <template #default="scope">
            <div @click="handleContestProblemClick(scope.row)" class="contest-problem-item"> {{ scope.row.title }}</div>
          </template>
        </el-table-column>
        <el-table-column align="right">
          <template #default="scope">
            <div style="color: #3498db;margin-right: 80px" @click="handleContestProblemClick(scope.row)"
                 class="contest-problem-item">查看题目
            </div>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>

</template>

<script>
import {eventBus} from "@/utils/eventBus";
import {ElButton, ElTable, ElTableColumn} from "element-plus";

export default {
  name: 'contestProblemList',
  components: {ElTable, ElTableColumn, ElButton},
  data() {
    return {
      contestId: '',
      contestProblemsInfo: [],
    }
  },
  created() {
    eventBus.on('contestProblemListEvent', this.setContestId)
  },
  methods: {
    setContestId(contestId) {
      this.contestId = contestId
      this.loadContestProblemInfo(this.contestId)
    },
    loadContestProblemInfo(contestId) {
      const req = {
        contestId: parseInt(contestId),
      }
      this.$axios.post('/contest/getContestProblems', req).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          return
        }
        if (resp.data !== null && resp.data !== []) this.contestProblemsInfo = resp.data.sort((a, b) => a.problemId - b.problemId)
      })
    },
    handleContestProblemClick(row) {
      const url = `/contest/${this.contestId}/problem/${row.problemId}`
      window.open(url)
    },
    contestProblemsListIndexMethod(index) {
      return index + 1
    }
  }
}
</script>
<style scoped>
.contest-problem-item:hover {
  cursor: pointer;
  color: blue;
  text-decoration: underline;
}
</style>