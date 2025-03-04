<template>
  <div>
    <h2 class="title">开发日志</h2>
    <Markdown_edit 
        :handle-content="handleLogContent" :handle-save="handleLogSave"></Markdown_edit>
    <el-backtop :right="100" :bottom="100" />
  </div>
</template>

<script lang="ts" setup>
import Markdown_edit from '@/components/util/Markdown_edit.vue';
import { getDevelopContent, saveDevelopContent } from '@/network/admin';
import { ElMessage } from 'element-plus';

const handleLogContent = () => {
  return getDevelopContent().then(res => {
    if(res.code == 0) {
      ElMessage.error("获取失败"); return "";
    }
    return res.data
  }).catch(err => {ElMessage.error("获取失败 " + err);})
}
const handleLogSave = (content:string, mode:string) => {
  saveDevelopContent({content}).then(res => {
    if(res.code == 0) {
      ElMessage.error("保存失败"); return "";
    }
    if(mode == "auto") return
    ElMessage.success("保存成功")
  }).catch(err => {ElMessage.error("保存失败 " + err);})
}
</script>

<style lang="less">
@import url("@/public.less");
.title {
    .subtitle()
}
</style>