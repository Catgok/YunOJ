<template>
  <div>
    <div style="width: 100%;text-align:center; font-size: 30px;padding: 13px 0 13px 0">
      YunOJ 题库
    </div>
    <div style="height: 40px">
    </div>

    <el-table :data="problemData" stripe style="width: 100%;">
      <el-table-column type="index" label="#" width="70" :index="ProblemListIndexMethod"/>
      <el-table-column prop="problemName" label="题目名称">
        <template #default="scope">
          <div @click="handleProblemClick(scope.row)" class="problem-item">{{ scope.row.problemName }}</div>
        </template>
      </el-table-column>
      <el-table-column prop="hardLevel" label="难度" width="150"/>
      <el-table-column prop="passRate" label="通过率" width="150"/>
    </el-table>

    <div style="display: flex;flex-direction: row-reverse; margin: 30px 0 40px 0 ">
      <el-pagination small background layout="prev, pager, next" :total="total" :current-page="currentPage"
                     :page-size="pageSize" @current-change="handleCurrentChange"></el-pagination>
    </div>
    <div style="text-align: right">
      <el-button type="success" v-show="isCoach()&&total!==0&&total<=currentPage*pageSize" @click="routeToNewProblem">
        新建题目
      </el-button>
    </div>
  </div>
</template>

<script>
import {ElButton, ElPagination, ElTable, ElTableColumn} from "element-plus";
import {hardLevelMap} from "@/utils/globalStaticData";
import {isCoach} from "@/utils/utils";

export default {
  components: {ElTable, ElTableColumn, ElPagination, ElButton},
  data() {
    return {
      total: 0,
      currentPage: 1,
      pageSize: 10,
      problemData: []
    }
  },
  created() {
    this.loadProblemData(1)
  },
  methods: {
    isCoach,
    routeToNewProblem() {
      this.$router.push('/problem/new')
    },
    ProblemListIndexMethod(index) {
      return (this.currentPage - 1) * this.pageSize + index + 1
    },
    handleProblemClick(row) {
      const url = `/problem/${row.id}`;
      sessionStorage.setItem('problemTitle.' + row.id, row.problemName)
      window.open(url);
    },
    handleCurrentChange(val) {
      this.currentPage = val
      this.problemData = []
      this.loadProblemData(val)
    },
    loadProblemData(pageNo) {
      const req = {
        pageNo: pageNo,
        pageSize: this.pageSize
      }
      this.$axios.post('/problem/getByPage', req).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          return
        }
        this.total = resp.total
        resp.data.forEach(problem => {
          this.problemData.push({
            id: problem.problemId,
            problemName: problem.title,
            hardLevel: hardLevelMap[problem.hardLevel],
            passRate: problem.submitCount === 0 ? '0%' : 100 * problem.passCount / problem.submitCount + '%'
          })
        })
      })
    }
  },
}
</script>

<style scoped>

.problem-item:hover {
  cursor: pointer;
  color: blue;
  text-decoration: underline;
}
</style>