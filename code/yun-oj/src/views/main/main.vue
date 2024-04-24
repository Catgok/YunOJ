<template>
  <div style="height: 70vh;width: 100%">
    <div style="display: flex" class="main-page-item">
      <div style="text-align:left;" v-html="renderedMarkdown"></div>
    </div>
    <div style="display: flex;height: 30vh;margin-bottom: 15px;">
      <div style="flex:1;margin-right: 15px;height: 300px" class="main-page-item">
        <div>
          <div class="main-page-text">最新题目</div>
          <el-table :data="problemData" stripe style="width:calc(40vw - 79.5px)">
            <el-table-column prop="id" label="ID" width="60"></el-table-column>
            <el-table-column prop="problemName" label="题目名称">
              <template #default="scope">
                <div @click="handleProblemClick(scope.row)" class="problem-item">{{ scope.row.problemName }}</div>
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
                <div @click="handleContestClick(scope.row)" class="contest-item">{{ scope.row.name }}</div>
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

export default {
  components: {ElTable, ElTableColumn},
  data() {
    return {
      helloText: `欢迎来到 YUN-OJ
# 11
$$a_b$$
      `,
      problemData: [
        {id: 1001, problemName: 'A+B Problem'},
        {id: 1002, problemName: 'A+B Problem2'},
        {id: 1003, problemName: 'A+B Problem3'},
        {id: 1004, problemName: 'A+B Problem4'},
        {id: 1005, problemName: 'A+B Problem5'},
      ],
      contestData: [
        {id: '1', name: '1'},
        {id: '1', name: '2'},
      ]
    }
  },
  created() {
    this.loadRecentProblemData();
    this.loadRecentContestData();
  },
  computed: {
    renderedMarkdown() {
      // https://github.com/runarberg/markdown-it-math
      const md = require('markdown-it')().use(require('markdown-it-math'),)
      return md.render(this.helloText);
    },
  },
  methods: {
    ContestListIndexMethod(index) {
      return index + 1;
    },
    handleProblemClick(row) {
      const url = `/problem/${row.id}`;
      sessionStorage.setItem('problemTitle.' + row.id, row.problemName)
      window.open(url)
    },
    handleContestClick(row) {
      const url = `/contest/${row.id}`;
      window.open(url);
    },
    loadRecentProblemData() {
      // todo
    },
    loadRecentContestData() {
      // todo
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