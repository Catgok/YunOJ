<template>
  <div>
    <div style="text-align:left; margin:0 15px 15px 5px">
      <div>
        <div style="font-size: 30px">{{ contestInfo.name }}</div>
        <div style="display: flex;font-size: 16px;justify-content: space-between;">
          <div style="display: flex">
            <div>{{ contestInfo.startTime }}</div>
            &nbsp; -- &nbsp;
            <div>{{ contestInfo.endTime }}</div>
          </div>
        </div>
      </div>
    </div>
    <div style="display: flex;justify-content: space-between;    align-items: flex-end;">
      <el-radio-group v-model="this.showFunc">
        <el-radio-button v-for="item in funcList" :key="item.id" :value="item.id" @click="clickFunc(item.id)">
          {{ item.name }}
        </el-radio-button>
      </el-radio-group>
      <el-button @click="routeToEditContest()">编辑竞赛</el-button>
    </div>
    <div style="display: flex">
      <div style="flex: 20;border-top: 1px #938c8c solid;">
        <div v-show="showFunc===1">
          <div style="text-align:left;" v-html="renderedMarkdown"></div>
        </div>
        <contest-problem-list v-show="showFunc===2"></contest-problem-list>
        <contest-rank-info v-show="showFunc===3"></contest-rank-info>
      </div>
    </div>
  </div>
</template>

<script>
import ContestProblemList from "@/views/contest/detail/contestProblemList.vue";
import ContestRankInfo from "@/views/contest/detail/contestRankInfo.vue";
import {eventBus} from "@/utils/eventBus";
import {formDate} from "@/utils/utils";
import {ElButton, ElRadioButton, ElRadioGroup} from "element-plus";
import texzilla from "texzilla";

export default {
  name: 'contestContent',
  components: {ContestProblemList, ContestRankInfo, ElRadioButton, ElRadioGroup, ElButton},
  data() {
    return {
      contestInfo: {contestId: 1, contestName: '', description: '',},
      showFunc: 1,
      funcList: [{name: '比赛信息'}, {name: '题目列表'}, {name: '竞赛榜单'},],
      eventsMap: {2: 'contestProblemListEvent', 3: 'contestRankInfoEvent'},
    }
  },
  created() {
    this.showFunc = 1
    this.contestInfo.id = this.$route.params.contestId
    for (let i = 0; i < this.funcList.length; i++) {
      this.funcList[i].id = i + 1
    }
    this.loadContestInfo(this.contestInfo.id)
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
    clickFunc(id) {
      this.showFunc = id
      if (id > 1) eventBus.emit(this.eventsMap[id], this.contestInfo.id)
    },
    routeToEditContest() {
      this.$router.push(`/contest/edit/${this.contestInfo.id}`)
    },
    loadContestInfo(contestId) {
      const req = {
        contestId: parseInt(contestId),
      }
      this.$axios.post('/contest/getContestById', req).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          return
        }
        this.contestInfo = resp.data
        this.contestInfo.startTime = formDate(resp.data.startTime)
        this.contestInfo.endTime = formDate(resp.data.endTime)
      })
    }
  }
}
</script>


<style scoped>
.contest-detail-box {
  flex: 1;
  border: 1px #938c8c solid;
  margin: 0 5px 0 0;
  padding: 5px;
  text-align: center;
}
</style>