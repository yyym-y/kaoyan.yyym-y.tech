<template>
  <div class="noteOutter" ref="outtter">
    <el-slider v-model="splitSize" style="width: 100%" :format-tooltip="formatter"
      :min="0" :max="1" :step="0.001" ref="silder" size="small"/>
    <a-split :directions="['right']" style="width:100%; height: 100%; margin-top: 10px;"
      v-model:size="splitSize">
      <template #first>
        <div>
          <h3 class="subtitle">笔记目录</h3>
          <File_tree></File_tree>         
        </div>
      </template>
      <template #second>
        <div style="margin-left: 16px;">
        <h3 class="subtitle">{{ selectNode.title }}</h3>
        <Markdown_edit v-if="selectNode.fileType == 'md'" ref="mdEditer"
          :handle-content="handleNoteContent" :handle-save="handleNoteSave"></Markdown_edit>
        <Pdf_viewer v-if="selectNode.fileType == 'pdf'" :note-id="selectNode.id"></Pdf_viewer>
        <el-backtop :right="100" :bottom="100" />
        </div>
      </template>
    </a-split>
  </div>
</template>

<script lang="ts" setup>
import File_tree from '@/components/note/File_tree.vue';
import Markdown_edit from '@/components/util/Markdown_edit.vue';
import Pdf_viewer from '@/components/note/Pdf_viewer.vue';
import { onMounted, ref, watch, watchEffect } from 'vue';

// 调整 split 的大小
let splitSize = ref(0)
const formatter = (value: number) => { 
  return `${(value * 100).toFixed(2)}%`
};
let fullWidth = ref(0)
let outtter = ref()
onMounted(() => {
  fullWidth.value = outtter.value.clientWidth
  splitSize.value = 230 / fullWidth.value
})

// 获取存储的数据
import { useNote } from '@/store/note';
import { storeToRefs } from 'pinia';
const noteStore = useNote()
let { selectNode } = storeToRefs(noteStore)

// 监听ID的变化
let mdEditer = ref()
watch(selectNode, (newV, oldV) => {
  if(mdEditer.value == null || mdEditer.value.refresh == null) return
  mdEditer.value.refresh()
}, {deep : true})

// 获取内容和保存的函数
import { ElMessage } from 'element-plus';
import { getNoteContent, saveContent } from '@/network/note';

const handleNoteContent = async ():Promise<string> => {
  return getNoteContent({ notePath: selectNode.value.id }).then(res => {
    if (res.code === 0) {
        ElMessage.error("获取失败"); return "";
    }
    return res.data
  })
}

const handleNoteSave = (content:string, mode:string) => {
  console.log(content)
  saveContent({
    notePath : selectNode.value.id,
    content : content
  }).then(res => {
    if (res.code === 0) {
        ElMessage.error("保存失败"); return;
    }
    if(mode == "auto") return
    ElMessage.success("保存成功");
  })
}


</script>

<style lang="less" scoped>
@import url("@/public.less");
.subtitle {
  .subtitle();
  margin-right: 12px;
  width: auto;
}
.noteOutter {
  width: 100%;
  height: 100%;
}
</style>