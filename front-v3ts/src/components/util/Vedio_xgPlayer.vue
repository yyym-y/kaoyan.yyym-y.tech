<template>
  <div>
    <div id="XGplayer" v-touch:hold="handleHold" v-touch:release="handleRelease"></div>
    <transition name="el-fade-in-linear">
      <div class="notice" v-show="showNotce">已为你跳转到上次观看时间</div>
    </transition>
  </div>
</template>

<script lang="ts" setup>

import { onBeforeMount, onBeforeUnmount, onMounted, reactive, ref } from 'vue';
import Player from 'xgplayer'
import HlsPlugin from 'xgplayer-hls'
import "xgplayer/dist/index.min.css";
import TextTrack from 'xgplayer/es/plugins/track'
import 'xgplayer/es/plugins/track/index.css'

const props = defineProps<{vedioUrl : string, vedioId?:number, epicNum?:number}>()
let option = {
  id: "XGplayer",
  fluid : true,
  url: props.vedioUrl,
  plugins: [HlsPlugin, TextTrack],
  commonStyle : {
    progressColor: '#8d77b8',  // 进度条底色
    playedColor: 'gray', // 播放完成部分进度条底色
    cachedColor: '#FFFFFF', // 缓存部分进度条底色
    sliderBtnStyle: {}, // 进度条滑块样式
    volumeColor: '#40a1d9' // 音量颜色
  },
  miniprogress : true,
  closeVideoClick : true,
  playbackRate: [3, 2.5, 2, 1.75, 1.5, 1.25, 1, 0.5],
  hls: {
    fetchOptions: {
      headers: {
        "Authorization": localStorage.getItem("token"),
      },
    }
  },
  startTime : localStorage.getItem(props.vedioId + "-" + props.epicNum) != null ?
    parseInt(localStorage.getItem(props.vedioId + "-" + props.epicNum) as string) : 0
}

let player : any = null
onMounted(() => {
  player = new Player(option);
})
onBeforeUnmount(() => {
  player.destroy()
})

// 切换视频 Url
const changeUrl = (vedioUrl:string) => {
  option.url = vedioUrl
  player = new Player(option);
}
defineExpose({
  changeUrl
})

// 记录历史播放时长
let interval = setInterval(() => {
  if(player == null || props.vedioId == null || props.epicNum == null) return
  localStorage.setItem(props.vedioId + "-" + props.epicNum, player.currentTime)
}, 3000)
onBeforeUnmount(() => {
  clearInterval(interval)
})
// 加载弹出提示
let showNotce = ref(false)
onMounted(() => {
  setTimeout(() => {
    showNotce.value = true
  },700)
  setTimeout(() => {
    showNotce.value = false
  }, 4000)
})

// 设置长按增加倍速 
let isHold = false
let oriPlayRate = 1
let firstTime = true
const handleHold = () => {
  if(firstTime) {
    firstTime = false; return;
  }
  if(player.paused) return
  isHold = true
  oriPlayRate = player.playbackRate
  player.playbackRate = 3
  console.log("hold")
}
const handleRelease = () => {
  if(isHold == false) return;
  isHold = false
  player.playbackRate = oriPlayRate
  console.log("release")
}

</script>

<style lang="less" scoped>
.notice {
  position: absolute;
  height: 20px;
  background-color: rgb(39, 46, 59);
  color: white;
  width: 175px;
  transform: translateY(-32px);
  margin-left: 12px;
  padding-left: 8px;
  padding-top: 4px ;
}
</style>