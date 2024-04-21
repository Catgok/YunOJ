<template>
</template>

<script>
import {eventBus} from '@/utils/eventBus'
import {ElNotification} from 'element-plus'

export default {
  mounted() {
    eventBus.on('globalNotice', this.showMessage);
  },
  methods: {
    showMessage(msgData) {
      let noticeOptions = {}
      if (msgData.title !== undefined) noticeOptions.title = msgData.title
      if (msgData.message !== undefined) noticeOptions.message = msgData.message
      if (msgData.type !== undefined) noticeOptions.type = msgData.type
      if (msgData.duration !== undefined) noticeOptions.duration = msgData.duration
      ElNotification(noticeOptions)
    }
  }
}
</script>
<style scoped>

</style>