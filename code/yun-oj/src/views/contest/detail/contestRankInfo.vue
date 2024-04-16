<template>
  <div>
    <div>contestRankInfo</div>
    <div></div>
  </div>
</template>

<script>
import {eventBus} from "@/utils/eventBus";

export default {
  name: 'contestRankInfo',
  data() {
    return {
      contestId: '',
      contestRankInfo: [],
    }
  },
  created() {
    eventBus.on('contestRankInfoEvent', this.setContestId)
  },
  methods: {
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
        this.contestRankInfo = resp.data
        console.log(this.contestRankInfo)
        // todo rank处理
      })
    },
  }
}
</script>
<style scoped>
</style>