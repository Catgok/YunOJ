<template>
  <div>
    <div style="width: 100%;text-align:center; font-size: 30px;padding: 13px 0 13px 0">
      YunOJ 竞赛
    </div>
    <div style="height: 40px">
    </div>

    <el-table :data="contestData" stripe style="width: 100%;">
      <el-table-column type="index" label="#" width="70" :index="ContestListIndexMethod"/>
      <el-table-column prop="name" label="竞赛名称">
        <template #default="scope">
          <div @click="handleContestProblemClick(scope.row)" class="problem-item">{{ scope.row.name }}</div>
        </template>
      </el-table-column>
      <el-table-column prop="startTime" label="开始时间" width="150"/>
      <el-table-column prop="endTime" label="结束时间" width="150"/>
      <el-table-column width="150">
        <template #default="scope">
          <el-button text @click="signupContest(scope.row)">报名</el-button>
        </template>
      </el-table-column>
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
import moment from 'moment';
import {ElButton, ElPagination, ElTable, ElTableColumn} from "element-plus";

export default {
  components: {ElTable, ElTableColumn, ElPagination, ElButton},
  data() {
    return {
      total: 5,
      currentPage: 1,
      pageSize: 5,
      contestData: []
    }
  },
  created() {
    this.loadContestListData(1)
  },
  methods: {
    ContestListIndexMethod(index) {
      return (this.currentPage - 1) * this.pageSize + index + 1
    },
    handleContestProblemClick(row) {
      const url = `/contest/${row.id}`;
      window.open(url);
    },
    handleCurrentChange(val) {
      this.currentPage = val
      this.problemData = []
      this.loadContestListData(val)
    },
    loadContestListData(pageNo) {
      let contestData = this.contestData
      const req = {
        pageNo: pageNo,
        pageSize: this.pageSize
      }
      this.$axios.post('/contest/getContestsByPage', req).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          return
        }
        this.total = resp.total
        for (let i = 0; i < resp.data.length; i++) {
          contestData.push({
            id: resp.data[i].id,
            name: resp.data[i].name,
            startTime: moment(resp.data[i].startTime * 1000).format('YYYY-MM-DD HH:mm'),
            endTime: moment(resp.data[i].endTime * 1000).format('YYYY-MM-DD HH:mm'),
          })
        }
      })
    },
    signupContest(row) {
      const req = {
        contestId: row.id
      }
      this.$axios.post('/contest/signUpContest', req).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          return
        }

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