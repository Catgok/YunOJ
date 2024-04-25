<template>
  <div style="margin-top: 20px">
    <el-table :data="contestRankTableInfo" stripe style="width: 100% " table-layout="auto">
      <el-table-column type="index" label="排名" width="70" :index="ContestRankIndexMethod"/>
      <el-table-column prop="username" label="用户名"/>
      <el-table-column prop="solved" label="通过" width="70"/>
      <el-table-column prop="penalty" label="罚时" width="120"/>

      <el-table-column v-for="item in contestRankTitleInfo" :key=item.idx :prop=item.problemId
                       :label="item.problemId" width="100" style="text-align: center">
        <template #header>
          <div style="text-align:center;">
            <el-popover placement="top-start" :width="200" trigger="hover" :content="item.problemName">
              <template #reference>
                {{ item.idx }}
              </template>

            </el-popover>
          </div>
        </template>

        <template #default="scope">
          <div v-show="calcShowStatus(scope.row,item.problemId)===1" class="rank-ceil-1">
            <div>+</div>
            <div>
              {{ getSuccessTimes(scope.row, item.problemId) }} / {{ getFirstPassTime(scope.row, item.problemId) }}
            </div>
          </div>
          <div v-show="calcShowStatus(scope.row,item.problemId)===2" class="rank-ceil-2">
            <div>+</div>
            <div>
              {{ getSuccessTimes(scope.row, item.problemId) }} / {{ getFirstPassTime(scope.row, item.problemId) }}
            </div>
          </div>
          <div v-show="calcShowStatus(scope.row,item.problemId)===3" class="rank-ceil-3">
            <div>-</div>
            <div>{{ getSubmitTimes(scope.row, item.problemId) }}</div>
          </div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import {eventBus} from "@/utils/eventBus";
import {ElPopover, ElTable, ElTableColumn} from "element-plus";

export default {
  components: {ElTable, ElTableColumn, ElPopover},
  data() {
    return {
      contestId: '',
      contestRankTitleInfo: [],
      contestRankTableInfo: [],
    }
  },
  created() {
    eventBus.on('contestRankInfoEvent', this.setContestId)
  },
  methods: {
    calcShowStatus(row, problemId) {
      // 1: 一血 2: 通过 3: 未通过 4: 未提交
      if (row[problemId] === undefined) return 4
      if (!row[problemId].isPass) return 3
      const passTime = row[problemId].firstPassTime
      for (let i = 0; i < this.contestRankTableInfo.length; i++) {
        if (this.contestRankTableInfo[i].username !== row.username && this.contestRankTableInfo[i][problemId] !== undefined
            && this.contestRankTableInfo[i][problemId].isPass && this.contestRankTableInfo[i][problemId].firstPassTime < passTime) {
          return 2
        }
      }
      return 1;
    },
    getSubmitTimes(row, problemId) {
      if (row[problemId] === undefined) return 0
      return row[problemId].submitTimes
    },
    getSuccessTimes(row, problemId) {
      if (row[problemId] === undefined) return 0
      return row[problemId].tryTimes + 1
    },
    getFirstPassTime(row, problemId) {
      if (row[problemId] === undefined) return ''
      return Math.round(row[problemId].firstPassTime / 60)
    },
    ContestRankIndexMethod(index) {
      return index + 1
    },
    setContestId(contestId) {
      this.contestId = contestId
      this.loadContestRankInfo(this.contestId)
    },
    loadContestRankInfo(contestId) {
      const req = {
        contestId: parseInt(contestId),
      }
      this.$axios.post('/contest/getContestRank', req).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          return
        }
        this.doRank(resp.data)
      })
    },
    doRank(rankInfo) {
      this.contestRankTableInfo = []
      this.contestRankTitleInfo = []
      rankInfo.forEach(data => {
        let existingUser = this.contestRankTableInfo.find(user => user.username === data.userName);
        if (!existingUser) {
          existingUser = {username: data.userName, solved: 0, penalty: 0};
          this.contestRankTableInfo.push(existingUser);
        }
        existingUser[data.problemId.toString()] = {
          isPass: data.isPass,
          submitTimes: data.submitTimes,
          tryTimes: data.tryTimes,
          firstPassTime: data.firstPassTime
        };
        if (data.isPass) {
          existingUser.solved++;
          existingUser.penalty += data.firstPassTime + data.tryTimes * 20 * 60;
        }
      });
      this.contestRankTableInfo.sort((a, b) => {
        if (a.solved !== b.solved) return b.solved - a.solved;
        return a.penalty - b.penalty;
      });
      this.contestRankTableInfo.forEach((user, index) => {
        user.penalty = Math.round(user.penalty / 60);
      });
      console.log(this.contestRankTableInfo)

      const req = {
        contestId: parseInt(this.contestId),
      }
      this.$axios.post('/contest/getContestProblems', req).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          return
        }
        const contestProblemsInfo = resp.data.sort((a, b) => a.problemId - b.problemId)
        for (let i = 0; i < contestProblemsInfo.length; i++) {
          this.contestRankTitleInfo.push({
            idx: String.fromCharCode(65 + i),
            problemId: contestProblemsInfo[i].problemId.toString(),
            problemName: contestProblemsInfo[i].title,
          })
        }
        // console.log(this.contestRankTitleInfo)
      })

    }
  }
}
</script>
<style>
.el-table .cell {
  padding: 0 3px 0 3px

}

.rank-ceil-1 {
  padding: 0;
  font-weight: bold;
  text-align: center;
  background-color: limegreen;
}

.rank-ceil-2 {
  padding: 0;
  font-weight: bold;
  text-align: center;
  background: rgb(231, 254, 189);
}

.rank-ceil-3 {
  padding: 0;
  font-weight: bold;
  text-align: center;
  background: pink;
}
</style>