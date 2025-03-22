<template>
  <div>
    <h3 class="title" >欢迎使用 ORC 光学识别功能 :)</h3>
    <el-tabs v-model="activeName">
      <el-tab-pane class="tab" label="图片转 Latex" name="first">
        <Img2LatexVue></Img2LatexVue>
      </el-tab-pane>
      <el-tab-pane class="tab" label="图片转文字" name="second">
        <Build_EmptyVue></Build_EmptyVue>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script lang="ts" setup>
import Build_EmptyVue from "@/components/util/Build_Empty.vue";
import { ref } from "vue";
import Img2LatexVue from "@/components/orc/Img2Latex.vue";

let activeName = ref("first")
let uploadURL = ""
</script>

<style lang="less">
@import url("@/public.less");
.title {
  .subtitle()
}
</style>