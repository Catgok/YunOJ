<template>
  <div>
    <div style="text-align:left; margin:0 15px 15px 5px;font-size: 30px">
      <div>{{ contestInfo.name }}</div>
      <div style="display: flex;font-size: 16px">
        <div>{{ contestInfo.startTime }}</div>
        --
        <div>{{ contestInfo.endTime }}</div>
      </div>
    </div>
    <div style="display: flex">
      <div v-for="item in funcList" :key="item.id" class="contest-detail-box" @click="clickFunc(item.id)">
        {{ item.name }}
      </div>
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
import moment from "moment/moment";

export default {
  name: 'contestContent',
  components: {ContestProblemList, ContestRankInfo},
  data() {
    return {
      contestInfo: {
        contestId: 1,
        contestName: '',
        description: '',
      },
      showFunc: '',
      funcList: [
        {name: '比赛信息'},
        {name: '题目列表'},
        {name: '竞赛榜单'},
      ],
      eventsMap: {
        2: 'contestProblemListEvent',
        3: 'contestRankInfoEvent'
      },
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
      const md = require('markdown-it')().use(require('markdown-it-math'))
      return md.render(this.contestInfo.description);
    },
  },
  methods: {
    clickFunc(id) {
      this.showFunc = id

      if (id > 1) eventBus.emit(this.eventsMap[id], this.contestInfo.id)
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
        this.contestInfo.startTime = moment(resp.data.startTime * 1000).format('YYYY-MM-DD HH:mm')
        this.contestInfo.endTime = moment(resp.data.endTime * 1000).format('YYYY-MM-DD HH:mm')
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