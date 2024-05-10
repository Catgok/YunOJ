<template>
  <div>
    <div style="width: 100%;text-align:center; font-size: 30px;padding: 13px 0 13px 0">
      YunOJ 竞赛
    </div>
    <div style="height: 40px">
    </div>
    <el-table :data="contestData" stripe style="width: 100%;">
      <el-table-column type="index" label="#" width="70" :index="contestListIndexMethod"/>
      <el-table-column prop="name" label="竞赛名称">
        <template #default="scope">
          <div style="color: #3498db" @click="handleContestClick(scope.row)" class="contest-item">{{ scope.row.name }}</div>
        </template>
      </el-table-column>
      <el-table-column prop="startTime" label="开始时间" width="180"/>
      <el-table-column prop="endTime" label="结束时间" width="180"/>
      <el-table-column width="150">
        <template #default="scope" style="width: 50px ">
          <el-button style="color: #79a978" link v-show="!signUpStatus(scope.row)" @click="signupContest(scope.row)">报名</el-button>
          <div v-show="signUpStatus(scope.row)">已报名</div>
        </template>
      </el-table-column>
    </el-table>

    <div style="display: flex;flex-direction: row-reverse; margin: 30px 0 40px 0 ">
      <el-pagination small background layout="prev, pager, next" :total="total" :current-page="currentPage"
                     :page-size="pageSize" @current-change="handleCurrentChange"></el-pagination>
    </div>
    <div style="text-align: right">
      <el-button type="success" v-if="isCoach()&&total!==0&&total<=currentPage*pageSize" @click="routeToNewContest">
        新建竞赛
      </el-button>
    </div>
  </div>
</template>

<script>
import moment from 'moment';
import {ElButton, ElPagination, ElTable, ElTableColumn} from "element-plus";
import {isCoach} from "@/utils/utils";
import {eventBus} from "@/utils/eventBus";

export default {
  components: {ElTable, ElTableColumn, ElPagination, ElButton},
  data() {
    return {
      total: 0,
      currentPage: 1,
      pageSize: 10,
      contestData: [],
      signedContest: []
    }
  },
  created() {
    this.loadContestListData(1)
    this.getSigned()
  },
  methods: {
    getSigned() {
      this.$axios.post('/contest/getSignUpContests').then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          return
        }
        this.signedContest = resp.data
        // console.log(this.signedContest)
      })
    },
    signUpStatus(row){
      return this.signedContest.includes(row.id)
    },
    isCoach,
    contestListIndexMethod(index) {
      return (this.currentPage - 1) * this.pageSize + index + 1
    },
    routeToNewContest() {
      this.$router.push('/contest/new')
    },
    handleContestClick(row) {
      const url = `/contest/${row.id}`;
      window.open(url);
    },
    handleCurrentChange(val) {
      this.currentPage = val
      this.problemData = []
      this.loadContestListData(val)
    },
    loadContestListData(pageNo) {
      this.contestData = []
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
          this.contestData.push({
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
        if (resp.code === 10301) {
          const noticeData = {type: "error", message: '竞赛未开始或已结束', duration: 3000}
          eventBus.emit('globalNotice', noticeData)
        }
        if (resp.code !== 0) {
          return
        }
        const noticeData = {type: "success", message: '报名竞赛成功', duration: 1300}
        eventBus.emit('globalNotice', noticeData)

      })
    }
  },
}
</script>

<style scoped>
.contest-item:hover {
  cursor: pointer;
  color: blue;
  text-decoration: underline;
}
</style>