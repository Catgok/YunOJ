<template>
  <div>
    <div style="font-size: 28px; border-bottom: 1px #938c8c solid;margin-bottom: 30px">{{ titleInfo }}</div>
    <div style="width: 100%">
      <el-form :model="contestInfo" label-position="left" label-width="auto" style="max-width: 1300px">
        <el-form-item label="竞赛名称">
          <el-input v-model="contestInfo.name"/>
        </el-form-item>
        <el-form-item label="竞赛时间">
          <div class="block">
            <el-date-picker v-model="contestTimeInfo" type="datetimerange" start-placeholder="开始时间"
                            end-placeholder="结束时间" format="YYYY-M-D HH:mm:ss" date-format="YYYY/M/D ddd"
                            time-format="HH:mm:ss"/>
          </div>
        </el-form-item>
        <el-form-item label="竞赛描述">
          <el-input type="textarea" :autosize="{ minRows: 10, maxRows: 100 }" v-model="contestInfo.description"/>
        </el-form-item>
        <el-form-item label="描述预览">
          <div style="width:100%;height:100%;border: 1px black solid;min-height: 20vh;padding-left:10px">
            <div style="" v-html="renderedMarkdown"></div>
          </div>
        </el-form-item>
        <el-form-item label="竞赛题目">
          <el-table :data="contestProblemInfo" style="width: 100%">
            <el-table-column label="题目ID" prop="id" width="180"/>
            <el-table-column label="题目名称" prop="name"/>
            <el-table-column align="right">
              <template #header>
                <div style="display: flex;justify-content: flex-end;">
                  <el-input v-model="problemSearchInput" size="small" placeholder="输入题目ID" style="width: 180px"/>
                  <el-button @click="handleAddProblem">添加</el-button>
                </div>
              </template>
              <template #default="scope">
                <el-button size="small" type="danger" @click="handleDeleteProblem(scope.row)">
                  删除
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-form-item>
      </el-form>
    </div>
    <div style="text-align: right">
      <el-button type="primary" @click="onSubmitContestInfo">{{ submitContestInfo }}</el-button>
    </div>
  </div>
</template>

<script>
import {
  ElButton,
  ElCol,
  ElDatePicker,
  ElForm,
  ElFormItem,
  ElInput,
  ElTable,
  ElTableColumn,
  ElTimePicker
} from "element-plus";
import {eventBus} from "@/utils/eventBus";
import texzilla from "texzilla";

export default {
  components: {ElForm, ElFormItem, ElButton, ElCol, ElTimePicker, ElInput, ElDatePicker, ElTable, ElTableColumn},
  data() {
    return {
      titleInfo: '',
      isNewContest: false,
      submitContestInfo: '',
      contestInfo: {id: '', name: '', startTime: '', endTime: '', description: ''},
      contestTimeInfo: [],
      problemSearchInput: '',
      contestProblemInfo: []
    }
  },
  created() {
    this.isNewContest = this.$route.name === 'contestNew'
    if (this.isNewContest) {
      this.titleInfo = '新建竞赛'
      this.submitContestInfo = '创建竞赛'
    } else {
      this.titleInfo = '编辑竞赛'
      this.submitContestInfo = '提交修改'
      this.contestInfo.id = this.$route.params.contestId
      this.getContestInfo(this.contestInfo.id)
    }
  },
  computed: {
    renderedMarkdown() {
      // https://github.com/runarberg/markdown-it-math
      const texzilla = require('texzilla');
      const md = require('markdown-it')()
          .use(require('markdown-it-math'), {
            inlineOpen: '$',
            inlineClose: '$',
            blockOpen: '$$',
            blockClose: '$$',
            inlineRenderer: function(str) {
              return texzilla.toMathMLString(str);
            },
            blockRenderer: function(str) {
              return texzilla.toMathMLString(str, true);
            }
          });
      return md.render(this.contestInfo.description);
    },
  },
  methods: {
    getContestInfo(contestId) {
      const req = {
        contestId: parseInt(contestId),
      }
      this.$axios.post('/contest/getContestById', req).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          return
        }
        this.contestInfo.name = resp.data.name
        this.contestInfo.description = resp.data.description
        this.contestInfo.startTime = resp.data.startTime
        this.contestInfo.endTime = resp.data.endTime
        this.contestTimeInfo.push(new Date(this.contestInfo.startTime * 1000), new Date(this.contestInfo.endTime * 1000))
      })
      const getContestProblemsReq = {
        contestId: parseInt(contestId),
      }
      this.$axios.post('/contest/getContestProblems', getContestProblemsReq).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          return
        }
        if (resp.data && resp.data.length > 0)
          this.contestProblemInfo = resp.data.map(item => ({id: item.problemId, name: item.title}));
        else this.contestProblemInfo = []
      })
    },
    onSubmitContestInfo() {
      this.contestInfo.startTime = this.contestTimeInfo[0].getTime() / 1000
      this.contestInfo.endTime = this.contestTimeInfo[1].getTime() / 1000
      if (this.isNewContest) {
        let addContestReq = {
          name: this.contestInfo.name,
          description: this.contestInfo.description,
          startTime: this.contestInfo.startTime,
          endTime: this.contestInfo.endTime
        }
        this.$axios.post('/contest/createContest', addContestReq).then((res) => {
          const resp = res.data
          if (resp.code !== 0) {
            return
          }
          this.contestInfo.id = resp.data
          if (this.contestProblemInfo.length === 0) {
            let noticeData = {type: "success", message: "新建竞赛成功", duration: 1300}
            eventBus.emit('globalNotice', noticeData)
            this.$router.go(-1)
            return;
          }
          let updateContestProblemReq = {
            contestId: parseInt(this.contestInfo.id),
            problemIds: []
          }
          updateContestProblemReq.problemIds = this.contestProblemInfo.map(item => item.id)
          this.$axios.post('/contest/addProblem', updateContestProblemReq).then((res) => {
            const resp = res.data
            if (resp.code !== 0) {
              return
            }
            let noticeData = {type: "success", message: "新建竞赛成功", duration: 1300}
            eventBus.emit('globalNotice', noticeData)
            this.$router.go(-1)
          })
        })
      } else {
        let updateContestProblemReq = {
          contestId: parseInt(this.contestInfo.id),
          problemIds: []
        }
        updateContestProblemReq.problemIds = this.contestProblemInfo.map(item => item.id)
        this.$axios.post('/contest/addProblem', updateContestProblemReq).then((res) => {
          const resp = res.data
          if (resp.code !== 0) {
            return
          }
        })
        let updateContestReq = {
          contest: this.contestInfo
        }
        updateContestReq.contest.id = parseInt(updateContestReq.contest.id)
        this.$axios.post('/contest/updateContest', updateContestReq).then((res) => {
          const resp = res.data
          if (resp.code !== 0) {
            return
          }
          let noticeData = {type: "success", message: "更新竞赛成功", duration: 1300}
          eventBus.emit('globalNotice', noticeData)
          this.$router.go(-1)
        })
      }
    },
    handleAddProblem() {
      const req = {
        problemId: parseInt(this.problemSearchInput),
      }
      this.$axios.post('/problem/get', req).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          let noticeData = {type: "error", message: "题目ID不存在", duration: 3000}
          eventBus.emit('globalNotice', noticeData)
          return
        }
        const problemInfo = {id: parseInt(this.problemSearchInput), name: resp.data.title}
        this.contestProblemInfo.push(problemInfo)
      })
    },
    handleDeleteProblem(row) {
      const index = this.contestProblemInfo.findIndex(problem => problem.id === row.id);
      if (index !== -1) this.contestProblemInfo.splice(index, 1);
    }
  }

}
</script>

<style scoped>
</style>