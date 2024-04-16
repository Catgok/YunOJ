<template>
  <div>
    <div>contestProblemList</div>
    <div>
      <el-table :data="contestProblemsInfo" stripe style="width: 100%;" table-layout="auto">
        <el-table-column type="index" :index="contestProblemsListIndexMethod"/>
        <el-table-column prop="title" label="题目名称">
          <template #default="scope">
            <div @click="handleContestProblemClick(scope.row)" class="problem-item">
              {{ scope.row.title }}
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
        this.contestProblemsInfo = resp.data
        for (let i = 0; i < this.contestProblemsInfo.length; i++) {
          this.contestProblemsInfo[i].id = i + 1
        }
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
.problem-item:hover {
  cursor: pointer;
  color: blue;
  text-decoration: underline;
}
</style>