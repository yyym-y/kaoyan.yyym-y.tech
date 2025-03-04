<template>
  <el-carousel height="24em">
    <div>
      <el-carousel-item v-for="(item, index) in carouselList" :key="item.carouselId + 'carsouel' + index">
        <Carousel_item v-if="item != null" :carousel-item="item"></Carousel_item>
      </el-carousel-item>
    </div>
  </el-carousel>
</template>

<script lang="ts" setup>
import type { carousel } from '@/model/bussModel/carousel';
import Carousel_item from './Carousel_item.vue';

const props = defineProps<{carouselList:carousel[]}>()

</script>

<style>
  .el-carousel__item h3 {
    color: #475669;
    font-size: 18px;
    opacity: 0.75;
    line-height: 300px;
    margin: 0;
  }
  
  .el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
  }
  
  .el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
  }
</style>