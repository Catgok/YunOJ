<template>
  <div class="main">
    <div style="width:33vw">
      <div class="header">{{ headText }}</div>
      <div class="text">用户名</div>
      <el-input v-model="username" class="input-box"
                :placeholder="funcType==='login'?'用户名,手机号,邮箱':''"></el-input>
      <div v-show="funcType==='register'" class="text">手机号</div>
      <el-input v-show="funcType==='register'" v-model="phone" class="input-box"></el-input>
      <div class="text">密码</div>
      <el-input v-model="password" class="input-box" type="password"></el-input>
      <el-button class="login-button" :type="funcType==='login'?'success':'primary'" @click="doFunc">
        {{ funcName }}
      </el-button>
      <div class="help-line">
        <div @click="forgetPass">找回密码</div>
        <div @click="changeLoginAndRegister">{{ switchText }}</div>
      </div>
    </div>
  </div>
</template>

<script>
import {ElButton, ElHeader, ElInput} from "element-plus";
import {eventBus} from '@/utils/eventBus.js'

export default {
  name: "login",
  components: {ElInput, ElButton, ElHeader},
  data() {
    return {
      funcType: 'login',
      username: 'u1se',
      password: 'pass',
      headText: '登录一下',
      funcName: '登录',
      switchText: '注册账号',
      phone: '',
    }
  },
  methods: {
    changeLoginAndRegister() {
      if (this.funcType === 'login') {
        this.funcType = 'register'
        this.headText = '注册账号'
        this.funcName = '注册'
        this.switchText = '已有账号？点此登录'
      } else {
        this.funcType = 'login'
        this.headText = '登录一下'
        this.funcName = '登录'
        this.switchText = '注册账号'
      }
    },
    doFunc() {
      if (this.funcType === 'login') this.login()
      else this.register()
    },
    login() {
      const req = {
        userKey: this.username,
        password: this.password
      }
      let noticeData = {type: "success", message: "登录成功", duration: 1300}
      this.$axios.post('/user/login', req).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          noticeData = {type: "error", message: "用户名或密码错误", duration: 2500}
          eventBus.emit('globalNotice', noticeData)
          return
        }
        localStorage.setItem('userInfo', JSON.stringify(resp.data))
        localStorage.setItem('U-Token', resp.utoken);
        // localStorage.setItem('loginStatus', 'true')
        eventBus.emit('globalHeaderLoadUserInfoEvent')
        eventBus.emit('globalNotice', noticeData)
        this.$router.go(-1)
      })
    },
    register() {
      const req = {
        username: this.username,
        phone: this.phone,
        password: this.password
      }
      this.$axios.post('/user/register', req).then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          let noticeData = {type: "error", message: "用户名或手机号已使用，请重试", duration: 2500}
          eventBus.emit('globalNotice', noticeData)
          return
        }
        if (resp.data) {
          this.login()
        }
      })
    },
    forgetPass() {
      const noticeData = {type: "info", message: "请联系管理员", duration: 1300}
      eventBus.emit('globalNotice', noticeData)
    }
  }
}
</script>

<style scoped>
.main {
  display: flex;
  flex-direction: column;
  padding-top: 4vh;
  padding-bottom: 8vh;
  justify-content: center;
  align-items: center
}

.header {
  font-size: 27px;
  text-align: center;
  margin-bottom: 40px;
}

.text {
  user-select: none;
  font-size: 18px;
}

.input-box {
  height: 40px;
  margin: 7px 0 15px 0;
}

.login-button {
  margin-top: 10px;
  height: 35px;
  width: 100%;
}

.help-line {
  display: flex;
  justify-content: space-between;
  margin-top: 15px;
  font-size: 14px;
  cursor: pointer;
  user-select: none;
}
</style>