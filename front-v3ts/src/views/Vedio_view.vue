<template>
    <div>
      <h1 class="subtitle1">{{ vedioInfo.vedioName }}</h1>
      <div>
        <el-row ref="vedioOuter">
          <el-col :span="17">
            <h3 class="subtitle3">{{ EpisoName }}</h3>
            <Vedio_xgPlayer :vedio-url="vedioUrl" ref="xgVedio" :vedio-id="pVedioId" :epic-num="pEpisoNum"></Vedio_xgPlayer>
          </el-col>
          <el-col :span="7">
            <h3 class="subtitle3" style="margin-left: 12px;">选集</h3>
            <a-scrollbar :style="scrollStyle" ref="scrollbarRef" class="scoller">
              <p v-for="(item, index) in playList" :key="item.name + index + 'episo'"
                :ref="setEpisoItemRef"
                class="episo-item" @click="jumpEpiso(index)">{{ item.name }}</p>
            </a-scrollbar>
          </el-col>
        </el-row>
      </div>
      <h3 class="subtitle2">视频简介</h3>
      <div>
        <el-empty v-if="vedioInfo.description == ''" description="暂时没有简介哦"></el-empty>
        <el-text v-else size="large">{{ vedioInfo.description }}</el-text>
      </div>
    </div>
</template>
  
<script lang="ts" setup>
import Vedio_xgPlayer from '@/components/util/Vedio_xgPlayer.vue';
import { playInfo } from '@/model/bussModel/playInfo';
import { getPlayList, getVedioById } from '@/network/vedio';
import { useServeURL } from '@/store/network';
import { useVedio } from '@/store/vedio';
import { ElMessage } from 'element-plus';
import { storeToRefs } from 'pinia';
import {  onBeforeUpdate , onUpdated, reactive, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useResizeObserver } from '@vueuse/core';

let endErr:any = null

// 校验是否登录
let username = localStorage.getItem("username")
let type = localStorage.getItem("type")
if(username == null || type == null || username == "" || type == "-1") {
  ElMessage.error("你还未登录， 请登录后访问...")
  endErr = true
}

// 获取 Vedio 的高度
const xgVedio = ref()
let scrollStyle = reactive({
  height : "0px"
})
const resizeObserver = useResizeObserver(xgVedio, (entries) => {
  const entry = entries[0]
  const { width, height } = entry.contentRect
  scrollStyle.height = String(height) + "px"
});

// 通过 VedioId 获取 Vedio 的信息
const route = useRoute()
const router = useRouter()
let pVedioId_s = route.params.vedioId as string
let pVedioId = parseInt(pVedioId_s)
const vedioStore = useVedio()
let { vedioInfo } = storeToRefs(vedioStore)
if(pVedioId != vedioInfo.value.vedioId) {
  getVedioById({ vedioId : pVedioId}).then(res => {
    if(res.code == 0) {
      if(!endErr) {
        ElMessage.error("URL 错误, 不存在该视频")
        endErr = true        
      }
      return
    }
    vedioInfo.value = res.data
  }).catch(err => {
    if(!endErr) {
        ElMessage.error("URL 错误, 不存在该视频")
        endErr = true        
    }
    console.error(err)
  })
}

// 校验播放的集数
let vedioUrl = ref("")
const serveStore = useServeURL()
const { mergeURL } = storeToRefs(serveStore)
let EpisoName = ref("")
let scrollbarRef = ref()


const jumpScroll = (num:number) => { // 跳转
  scrollbarRef.value.scrollTop(num)
}

const episoItemRef = ref([] as HTMLElement[]);
const setEpisoItemRef = (el: any) => {
  if (el) { episoItemRef.value.push(el) }
}
onBeforeUpdate(() => {
  episoItemRef.value = []
})
onUpdated(() => {
  console.log(episoItemRef)
})

let lastSelect = -1
const heightlight = (num:number) => {
  if (num < 1 || num > episoItemRef.length) {
    console.warn('Invalid index:', num);
    return;
  }
  if(lastSelect != -1) {
    console.log(episoItemRef.value)
    console.log(lastSelect)
    console.log(episoItemRef.value[lastSelect])
    episoItemRef.value[lastSelect].style.backgroundColor = '#e7f1ff';
    episoItemRef.value[lastSelect].style.color = '#409eff';
  }
  episoItemRef.value[num - 1].style.backgroundColor = 'lightblue';
  episoItemRef.value[num - 1].style.color = '#165DFF';
  lastSelect = num - 1
}

let pEpisoNum_s = route.params.episoNum as string
let pEpisoNum = parseInt(pEpisoNum_s)
let playList = ref([] as playInfo[])
getPlayList({vedioId : pVedioId}).then(res => {
  if(res.code == 0) { console.error(res.msg); return; }
  playList.value = res.data
  if(pEpisoNum > playList.value.length || pEpisoNum <= 0) {
    if(! endErr) {
      ElMessage.error("集数错误, 不存在该集"); endErr = true
    }
    return
  }
  let base = "vedio/vedioPlay/" + pVedioId_s + "/" + pEpisoNum_s
  vedioUrl.value = mergeURL.value(playList.value[pEpisoNum - 1].url, base)
  xgVedio.value.changeUrl(vedioUrl.value)
  // 缓存看过的集数
  localStorage.setItem("lastEpiso" + pVedioId_s, pEpisoNum_s)
  EpisoName.value = playList.value[pEpisoNum - 1].name
  // 滚动条跳转到指定地点并高亮
  setTimeout(() => {
    let height = 50 * (pEpisoNum - 1)
    jumpScroll(height)
    heightlight(pEpisoNum)
  }, 700);
}).catch(err => {
  if(! endErr) {
      ElMessage.error("集数错误, 不存在该集"); endErr = true
  }
  console.error(err)
})

const jumpEpiso = (num:number) => {
  heightlight(num + 1)
  let base = "vedio/vedioPlay/" + pVedioId_s + "/" + String(num + 1)
  vedioUrl.value = mergeURL.value(playList.value[num].url, base)
  console.log(vedioUrl.value)
  xgVedio.value.changeUrl(vedioUrl.value)
  // 缓存并跳转
  router.replace({
    params: {
      ...route.params,
      episoNum : String(num + 1)
    }
  })
  localStorage.setItem("lastEpiso" + pVedioId_s, String(num + 1))
  EpisoName.value = playList.value[num].name
}

</script>

<style lang="less">
@import url("@/public.less");
.subtitle1 {
  .subtitle();
}
.subtitle2 {
  .subtitle();
}
.subtitle3 {
  .subtitle();
}
.desciption {
  margin-top: 24px;
}
.episo-item {
  display: flex;
  align-items: center;
  height: 40px;
  margin: 10px;
  text-align: center;
  border-radius: 4px;
  background: #e7f1ff;
  padding-left: 12px;
  color: #409eff;
  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
}
.episoOutter {
  height: 100%;
}
.scoller {
  overflow: auto;
}
</style>