<template>

</template>

<script>
export default {
  components: {},
  data() {
    return {
      submissionData: []
    }
  },
  methods: {
    loadSubmissionData() {
      for (let i = 1; i <= 50; i++) {
        let item = {
          submissionId: 1000 + i,
          problemId: 1000 + i,
          problemName: '题目' + i,
          submitTime: '2021-10-10 10:10:10',
          status: 'Accepted',
          language: 'C++',
          runtime: '100ms',
          memory: '100MB',
          codeLength: '1000B',
        }
        this.submissionData.push(item)
      }
    }
  },
  created() {
    this.loadSubmissionData()
  }
}
</script>

<style scoped>

</style>