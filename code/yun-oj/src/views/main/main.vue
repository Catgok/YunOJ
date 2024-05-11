<template>
  <div style="height: 70vh;width: 100%">
    <div style="display: flex;margin-top: -20px" class="main-page-item">
      <div style="text-align:left;" v-html="renderedMarkdown"></div>
    </div>
    <div style="display: flex;height: 30vh;margin-bottom: 15px;">
      <div style="flex:1;margin-right: 15px;height: 300px" class="main-page-item">
        <div>
          <div class="main-page-text">最新题目</div>
          <el-table :data="problemData" stripe style="width:calc(40vw - 79.5px)">
            <el-table-column prop="problemId" label="ID" width="60"></el-table-column>
            <el-table-column prop="title" label="题目名称">
              <template #default="scope">
                <div style="color: #3498db" @click="handleProblemClick(scope.row)" class="problem-item">{{ scope.row.title }}</div>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>
      <div style="flex:1;height: 300px" class="main-page-item">
        <div>
          <div class="main-page-text">最新竞赛</div>
          <el-table :data="contestData" stripe style="width:calc(40vw - 79.5px)">
            <el-table-column type="index" label="#" width="70" :index="ContestListIndexMethod"/>
            <el-table-column prop="name" label="竞赛名称">
              <template #default="scope">
                <div style="color: #3498db" @click="handleContestClick(scope.row)" class="contest-item">{{ scope.row.name }}</div>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {ElTable, ElTableColumn} from "element-plus";
import texzilla from "texzilla";

export default {
  components: {ElTable, ElTableColumn},
  data() {
    return {
      helloText: `
# 欢迎来到 YUN-OJ
### 这里是属于YNU的ACM竞赛基地
### 点击题目开始你的ACM之旅吧！
      `,
      problemData: [],
      contestData: []
    }
  },
  created() {
    this.loadRecentProblemData();
    this.loadRecentContestData();
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
      return md.render(this.helloText);
    },
  },
  methods: {
    ContestListIndexMethod(index) {
      return index + 1;
    },
    handleProblemClick(row) {
      const url = `/problem/${row.problemId}`;
      sessionStorage.setItem('problemTitle.' + row.problemId, row.problemName)
      window.open(url)
    },
    handleContestClick(row) {
      const url = `/contest/${row.id}`;
      window.open(url);
    },
    loadRecentProblemData() {
      this.$axios.post('/problem/getRecentProblems').then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          return
        }
        const problemTitleList = resp.data
        this.problemData = problemTitleList.sort(() => Math.random() - 0.5).slice(0, 5).sort((a, b) => a.problemId - b.problemId);
      })
    },
    loadRecentContestData() {
      this.$axios.post('/contest/getRecentContests').then((res) => {
        const resp = res.data
        if (resp.code !== 0) {
          return
        }
        const ContestBaseInfoList = resp.data
        this.contestData = ContestBaseInfoList.sort(() => Math.random() - 0.5).slice(0, 5).sort((a, b) => a.id - b.id);
        ;
      })
    }
  }
}
</script>

<style scoped>
.main-page-item {
  border: 1px solid #ccc;
  border-radius: 10px; /* 设置圆角大小 */
  padding: 20px;
  background-color: #f9f9f9;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); /* 添加阴影效果 */
  margin-bottom: 15px;
}

.problem-item:hover {
  cursor: pointer;
  color: blue;
  text-decoration: underline;
}

.contest-item:hover {
  cursor: pointer;
  color: blue;
  text-decoration: underline;
}

.main-page-text {
  margin-bottom: 10px;
}
</style>