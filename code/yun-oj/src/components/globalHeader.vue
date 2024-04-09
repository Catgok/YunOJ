<template>
  <div class="header-header">
    <div class="logo">
      <el-image style="height: 55px" :src="url" fit="scale-down"/>
    </div>
    <div style="display: flex; justify-content: space-between; width: 100%;">
      <div class="functional-blocks">
        <div v-for="block in functionalBlocks" :key="block.id" class="functional-blocks-block"
             @click="functionalBlocksClick(block.path)">
          {{ block.name }}
        </div>
      </div>
      <div class="user-name-block">
        <div v-if=this.loginStatus style="display: flex">
          <div>{{ this.username }}</div>
          <div style="margin-left: 10px;margin-right: 10px" @click="logout">退出</div>
        </div>
        <div v-else @click="toLogin()"> 登录/注册</div>
      </div>
    </div>
  </div>
</template>

<script>
import {ElImage} from "element-plus";
import {eventBus} from '@/utils/eventBus.js'

export default {
  name: 'GlobalHeader',
  components: {ElImage},
  data() {
    return {
      loginStatus: false,
      username: '',
      url: require('@/assets/logo.png'),
      functionalBlocks: [
        {id: 1, name: '首页', path: '/'},
        {id: 2, name: '题目', path: '/problem'},
        {id: 3, name: '竞赛', path: '/contest'},
        {id: 4, name: '在线IDE', path: '/onlineIDE'},
        // {id: 5, name: '公告', path: '/announcement'},
      ],
    }
  },
  created() {
    eventBus.on('globalHeaderLoadUserInfo', this.loadUserInfo)
  },
  mounted() {
    this.loadUserInfo()
  },
  updated() {
    this.loadUserInfo()
  },
  methods: {
    loadUserInfo() {
      const userInfo = JSON.parse(localStorage.getItem('userInfo'))
      if (userInfo !== null) {
        this.loginStatus = true
        this.username = userInfo.username
      }
    },
    functionalBlocksClick(path) {
      this.$router.push(path)
    },
    toLogin() {
      this.$router.push('/login')
    },
    logout() {
      localStorage.removeItem('userInfo')
      localStorage.removeItem('loginStatus')
      this.loginStatus = false
      this.username = ''
    }
  }
}
</script>

<style scoped>
.header-header {
  padding-left: 8vw;
  padding-right: 8vw;
  display: flex;
}

.logo {
  //border: 1px solid red;
}

.functional-blocks {
  margin-left: 30px;
  display: flex;
  //background-color: #FFA500;
  //border: 1px solid red;
}

.functional-blocks-block {
  background-color: rgb(251, 191, 36, 0.7);
  padding: 0 5px 0 5px;
  width: 90px;
  text-align: center;
  height: 56px;
  line-height: 56px;
}

.functional-blocks-block:hover {
  background-color: rgb(251, 191, 36, 0.9);
  cursor: pointer;
}

.user-name-block {
  text-align: center;
  height: 56px;
  line-height: 56px;
  display: flex;
}

.logout-box {
  background-color: #f0f0f0;
  padding: 10px;
  margin-top: 10px;
}
</style>
