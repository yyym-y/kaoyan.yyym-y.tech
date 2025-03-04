<template>
  <div>
    <h3 class="title" >欢迎来到管理员控制面板 :)</h3>
    <Dashboard></Dashboard>
  </div>
</template>

<script lang="ts" setup>
import Dashboard from '@/components/admin/Dashboard.vue';

</script>

<style lang="less">
@import url("@/public.less");
.title {
  .subtitle()
}
</style>