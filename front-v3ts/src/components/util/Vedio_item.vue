<template>
    <div :class="choiceClass">
      <img class="imgItem" :src="absUrl" @click="jumpTo()">
      <el-text truncated size="large" class="text"
        @click="jumpTo()" v-if="ifShowName">{{ props.vedioInfo.vedioName }}</el-text>
    </div>
</template>
  
<script lang="ts" setup>
import { type vedio } from '@/model/bussModel/vedio';
import { ref } from 'vue';
import { useServeURL } from '../../store/network';
import { storeToRefs } from 'pinia';
import { useVedio } from '@/store/vedio';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';

const props = withDefaults(defineProps<{
  vedioInfo : vedio,
  ifShowName? : boolean,
}>(), {
  ifShowName : () => false
})

let absUrl = ref("")
const serveStore = useServeURL()
const { mergeURL } = storeToRefs(serveStore)
absUrl.value = mergeURL.value(props.vedioInfo.vedioCover, "image")

let choiceClass = ref("normal-video-item")
if (props.vedioInfo.typeId == 2) {
  choiceClass.value = "lecture-vedio-item"
}

const veidoStore = useVedio()
const router = useRouter()
const jumpTo = () => {
  veidoStore.vedioInfo = props.vedioInfo
  let episoNum = 1
  let lastEpiso = localStorage.getItem("lastEpiso" + String(props.vedioInfo.vedioId))
  if(lastEpiso != null) {
    episoNum = parseInt(lastEpiso)
    ElMessage.info("已跳转到上次观看的集数 : 第" + lastEpiso + "集")
  }
  router.push({name : "vedio", params : {
    vedioId : props.vedioInfo.vedioId,
    episoNum : episoNum
  }}).catch(err => {
    console.error(err)
  })
}

defineExpose({ jumpTo })
</script>
  
<style scoped lang="less">
@import url("@/public.less");
  .outer-center() {
    display: flex;
    justify-content: center;
    flex-wrap: wrap
  }
  .normal-video-item {
    .normal-video-item();
    .outer-center();
  }
  .lecture-vedio-item {
    .lecture-vedio-item();
    .outer-center();
  }
  .name {
      width: 100%;
      text-align: center;
  }
  .imgItem {
      width: 100%; border-radius: 4px;
  }
  .text {
    margin-top: 4px;
    margin-bottom: 4px;
    user-select: none;
  }
</style>