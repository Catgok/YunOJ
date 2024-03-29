<template>
  <div>
    <div style="width: 100%;text-align:center; font-size: 30px;padding: 13px 0 13px 0">
      YunOJ 题库
    </div>
    <div style="height: 40px">
    </div>

    <el-table :data="problemData" stripe style="width: 100%;">
      <el-table-column prop="id" label="ID" width="70"/>
      <el-table-column prop="problemName" label="题目名称">
        <template #default="scope">
          <div @click="handleProblemClick(scope.row.problemId)" class="problem-item">{{ scope.row.problemName }}</div>
        </template>
      </el-table-column>
      <el-table-column prop="hardLevel" label="难度" width="150"/>
      <el-table-column prop="passRate" label="通过率" width="150"/>
    </el-table>

    <div style="display: flex;flex-direction: row-reverse; margin: 30px 0 40px 0 ">
      <el-pagination
          small background layout="prev, pager, next"
          :total="total"
          :current-page="currentPage"
          :page-size="pageSize"
          @current-change="handleCurrentChange"
      ></el-pagination>
    </div>
  </div>
</template>

<script>
import {ElPagination, ElTable, ElTableColumn} from "element-plus";

export default {
  components: {ElTable, ElTableColumn, ElPagination},
  data() {
    return {
      total: 10000,
      currentPage: 1,
      pageSize: 50,
      problemData: []
    }
  },
  methods: {
    handleProblemClick(problemId) {
      const url = `/problem/${problemId}`;
      window.open(url);
    },
    handleCurrentChange(val) {
      this.currentPage = val
    },
    loadProblemData(pageNo) {
      // this.$axios.get('/problem').then((res) => {
      //   this.problemData = res.data
      // })
      for (let i = 1; i <= 50; i++) {
        let item = {
          problemId: 1000 + i,
          problemName: '题目' + i,
          hardLevel: 1,
          passRate: 100 - i,
        }
        this.problemData.push(item)
      }
    }
  },
  created() {
    this.loadProblemData(1)
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