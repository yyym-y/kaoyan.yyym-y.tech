<template>
<div>
  <el-tabs v-model="activeName">
      <el-tab-pane label="影视管理" name="vedioContro">
        <Vedio_adder></Vedio_adder>
        <Vedio_table></Vedio_table>
      </el-tab-pane>
      <el-tab-pane label="轮播图管理" name="carouselContro">
        
      </el-tab-pane>
  </el-tabs>
</div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';

import Vedio_adder from './Vedio_adder.vue';
import Vedio_table from './table/Vedio_table.vue';

let activeName = ref("vedioContro")
</script>

<style>

</style>