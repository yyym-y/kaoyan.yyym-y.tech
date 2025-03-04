<template>
  <div>
    <Carousel v-if="carouselList.length != 0" :carousel-list="carouselList"></Carousel>
    <h3 class="subtitle">全部影视目录</h3>
    <el-tabs v-model="activeName">
      <el-tab-pane v-for="(pr, index) in typeList" :key="pr + 'type' + index"
        class="tab" :label="pr.name" :name="pr.name">
        <Vedio_menu v-if="indexVedio != null && indexVedio[pr.typeId] != null"
          :vedio-list="indexVedio[pr.typeId]"></Vedio_menu>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue';

import Carousel from '@/components/main/Carousel.vue';
import Vedio_menu from '@/components/main/Vedio_menu.vue';
import { type carousel } from '@/model/bussModel/carousel';
import { type vedioType } from '@/model/bussModel/vedioType';
import { getCarousel } from '@/network/pub';
import { getVedioType } from '@/network/pub';
import { getIndexVedio } from '@/network/pub';
import { vedio } from '@/model/bussModel/vedio';

let carouselList = ref([] as carousel[])
getCarousel().then(res => {
  if(res.code == 0)
    return
  carouselList.value = res.data
}).catch(err => (console.error(err)))

let typeList = ref([] as vedioType[])
let activeName = ref("")
getVedioType().then(res => {
  if(res.code == 0)
    return
  typeList.value = res.data
  if(typeList.value.length != 0)
    activeName.value = typeList.value[0].name
}).catch(err => (console.error(err)))

let indexVedio = reactive({} as Record<number, vedio[]>)
getIndexVedio().then(res => {
  if(res.code == 0)
    return
  Object.assign(indexVedio, res.data)
  console.log(indexVedio)
}).catch(err => (console.error(err)))

</script>

<style lang="less">
@import url(@/public.less);
.subtitle {
  .subtitle()
}
.tab {
  min-width: 50px;
}
.vedio-outter {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
}
</style>