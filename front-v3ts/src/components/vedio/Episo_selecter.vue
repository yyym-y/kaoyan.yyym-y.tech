<template>
  <div>
    <a-scrollbar :style="scrollStyle" ref="scrollbarRef" class="scoller">
      <p v-for="(item, index) in playList" :key="item.name + index + 'episo'"
        :ref="setEpisoItemRef"
        class="episo-item" @click="jumpEpiso(index)">{{ item.name }}</p>
    </a-scrollbar>
  </div>
</template>

<script lang="ts" setup>
import { playInfo } from '@/model/bussModel/playInfo';
import { getPlayList } from '@/network/vedio';
import { useServeURL } from '@/store/network';
import { ElMessage } from 'element-plus';
import { storeToRefs } from 'pinia/dist/pinia';
import { useRoute, useRouter } from 'vue-router/dist/vue-router';
import { ref, watch } from 'vue/dist/vue';


const props = defineProps<{
  vedioId : number,
  vedioUrl_b : string,
  EpisoName_b : string
}>()

let vedioUrl = ref(props.vedioUrl_b)
let EpisoName = ref(props.EpisoName_b)

const route = useRoute()
const router = useRouter()
const ServeStore = useServeURL()
let { mergeURL } = storeToRefs(ServeStore)

let pEpisoNum_s = route.params.episoNum as string
let pEpisoNum = parseInt(pEpisoNum_s)
let playList = ref([] as playInfo[])
getPlayList({vedioId : props.vedioId}).then(res => {
  if(res.code == 0) { console.error(res.msg); return; }
  playList.value = res.data
  if(pEpisoNum > playList.value.length || pEpisoNum <= 0) {
    ElMessage.error("集数错误, 不存在该集");
    return
  }
  let base = "vedio/" + String(props.vedioId) + "/" + pEpisoNum_s
  vedioUrl.value = mergeURL.value(playList.value[pEpisoNum - 1].url, base)

  // 缓存看过的集数
  localStorage.setItem("lastEpiso" + props.vedioId, pEpisoNum_s)
  EpisoName.value = playList.value[pEpisoNum - 1].name
  // 滚动条跳转到指定地点并高亮
  setTimeout(() => {
    let height = 50 * (pEpisoNum - 1)
    // jumpScroll(height)
    // heightlight(pEpisoNum)
  }, 700);
}).catch(err => {
  ElMessage.error("集数错误, 不存在该集");
  console.error(err)
})


watch(() => props.vedioUrl_b, (newValue) => {
  vedioUrl.value = newValue;
});
const updateValue = () => {
  emit('update:modelValue', localValue.value);
};
</script>

<style>

</style>