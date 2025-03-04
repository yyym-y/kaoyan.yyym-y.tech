<template>
    <div class="Nav-menu">
      <el-menu
        default-active="1" >
        <el-menu-item index="1" @click="ToHome()">
          <el-icon><House /></el-icon>
          <span>首页</span>
        </el-menu-item>
        <el-menu-item index="2" @click="ToNote()">
          <el-icon><Notebook /></el-icon>
          <span>共享笔记</span>
        </el-menu-item>
      </el-menu>
    </div>
</template>

<script lang="ts" setup>
  import { useRouter } from 'vue-router';
  const router = useRouter();

  const ToNote = () => {
    router.push({ name: "note" }).catch(err => {
      console.error(err);
    });
  };
  const ToHome = () => {
    router.push({ name : "home"}).catch(err => {
      console.error(err)
    })
  }
</script>

<style lang="less">
.Nav-menu {
  margin-left: 40px;
}
</style>
