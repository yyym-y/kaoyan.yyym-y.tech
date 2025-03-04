<template>
  <div class="carousel_item-outter">
    <img class="carousel-img-item"
      :src="absUrl"
      @click="jumpTo()">
    <Vedio_item :if-show-name="false" :vedio-info="prop.carouselItem.vedio"
      v-if="carouselItem.ifShowVedio" ref="carouselVedio"
      class="carousel-video-item"
      ></Vedio_item>
  </div>
</template>

<script lang="ts" setup>
import { type carousel } from '@/model/bussModel/carousel';
import { useServeURL } from '@/store/network';
import { storeToRefs } from 'pinia';
import { ref } from 'vue';
import Vedio_item from '../util/Vedio_item.vue';

const prop = defineProps<{ carouselItem:carousel }>()
const serveStore = useServeURL()
const { mergeURL } = storeToRefs(serveStore)
const carouselVedio = ref()

let absUrl = ref("")
absUrl.value = mergeURL.value(prop.carouselItem.cover, "image")

const jumpTo = () => {
  if(prop.carouselItem.ifShowVedio == false)
    return
  carouselVedio.value.jumpTo()
}

</script>

<style lang="less">
.carousel-img-item {
  object-fit: scale-down !important;
  height: 24em;
  width: 100%;
}
.carousel-video-item {
    position: absolute;
    // margin-right: 5vw;
    right: 10%;
    top: 30%;
}
.carousel_item-outter {
  width: 100%;
  height: 100%;
}
</style>