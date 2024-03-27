<template>
  <div class="main">
    <div style="width:33vw">
      <div class="header">登录一下</div>
      <div class="text">用户名</div>
      <el-input v-model="username" class="input-box"></el-input>
      <div class="text">密码</div>
      <el-input v-model="password" class="input-box" type="password"></el-input>
      <el-button class="login-button" type="success" @click="login">登录</el-button>
      <div class="help-line">
        <div>找回密码</div>
        <div>注册账号</div>
      </div>
    </div>
  </div>
</template>
<script>
import {ElButton, ElHeader, ElInput} from "element-plus";

export default {
  name: "login",
  components: {ElInput, ElButton, ElHeader},
  data() {
    return {
      username: 'public',
      password: 'public',
    }
  },
  methods: {
    login() {
      const req = {
        username: this.username,
        password: this.password
      }
      let that = this
      const redirect = this.$route.query.redirect

      this.$axios.post('/login', req).then((res) => {
        const data = res.data
        if (data.code === 200 && data.result !== "") {
          const token = data.result
          localStorage.setItem('jwtToken', token)
          that.$axios.defaults.headers.common['Authorization'] = `Bearer ${token}`
          if (redirect) that.$router.push(redirect)
          else that.$router.push('/')
        }
      }).catch((error) => {
      });
    },
  }
}
</script>

<style scoped>
.main {
  display: flex;
  flex-direction: column;
  padding-top: 18vh;
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