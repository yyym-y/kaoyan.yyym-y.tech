<template>
<div>
  <el-tabs v-model="activeName">
      <el-tab-pane label="影视管理" name="vedioContro">
        <Vedio_adder></Vedio_adder>
      </el-tab-pane>
      <el-tab-pane label="轮播图管理" name="carouselContro">
        
      </el-tab-pane>
  </el-tabs>
</div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';

import Vedio_adder from './Vedio_adder.vue';

let activeName = ref("vedioContro")
</script>

<style>

</style>