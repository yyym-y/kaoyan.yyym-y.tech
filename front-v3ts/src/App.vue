<template>
  <div id="app">
    <el-container>
        <el-aside width="200px" class="ell-aside">
          <Aside></Aside>
        </el-aside>
        <el-container>
            <el-header>
              <Headers></Headers>
            </el-header>
            <el-main>
              <router-view></router-view>
            </el-main>
        </el-container>
    </el-container>
  </div>
</template>

<script lang="ts">

import Aside from "./components/aside/Aside.vue";
import Headers from './components/header/Header.vue'
export default {
  components : {
    Aside,
    Headers
  }
}

document.addEventListener('contextmenu', function (e) {
e.preventDefault();
})
</script>


<style lang="less">
  .el-header {
    color: #333;
    line-height: 60px;
    background-color: #f7f8f9;
    min-width: 700px;
  }
  .el-aside {
    color: #333;
    text-align: center;
    line-height: 200px;
    height: 100vh;
    border-right-style: solid;
    border-right-color: #C0C4CC;
    border-right-width: 1px;

  }
  .el-main {
    color: #333;
    background-color: #f7f8f9;
  }
  .ell-aside {
    overflow-y: hidden !important;
  }
</style>
