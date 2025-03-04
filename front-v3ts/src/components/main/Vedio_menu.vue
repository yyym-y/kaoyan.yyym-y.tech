<template>
    <div class="vedio-outter">
        <div v-for="(item, index) in vedioList" :key="item + 'vedio' + index">
        <Vedio_item :if-show-name="true" :vedio-info="item"></Vedio_item>
        </div>
        <i v-for="i in 6" :key="i + '<i>'" :class="emptyClass"></i>
    </div>
</template>

<script lang="ts" setup>
import { vedio } from '@/model/bussModel/vedio';
import { ref } from 'vue';

import Vedio_item from '../util/Vedio_item.vue';

const props = defineProps<{vedioList:vedio[]}>()
let emptyClass = ref("normal-video-item")
if(props.vedioList != null && props.vedioList.length != 0 && props.vedioList[0].typeId == 2) {
    emptyClass.value = "lecture-vedio-item"
}

</script>

<style lang="less">
@import url(@/public.less);
.normal-video-item {
  .normal-video-item()
}
.lecture-vedio-item {
  .lecture-vedio-item()
}
</style>