<template>
  <div class="tree">
    <VTreeSearch @select="handleNodeSelect" :data="data" :expandOnFilter="false"
          selectable animation 
          ref="filetree"
          :showLine="{ type : 'soild', polyline : false }">
      <template #node="{ node }">
        <a-dropdown trigger="contextMenu" alignPoint
          :style="{display:'block'}" ref="dropdown">
          
          <div>
            <span v-if="node.fileType == 'folder'">📂 </span>
            <span v-else-if="node.fileType == 'md'">📄 </span>
            <span v-if="node.fileType == 'pdf'">📙 </span>
            <span>
              {{ node.title }}
            </span>            
          </div>

          <!-- 右键选择 -->
          <template #content>
            <div v-if="type == '0'">
              <!-- 文件夹专属操作 -->
              <a-doption v-if="node.fileType == 'folder'" @click="handleSelect(node, 1)">
                <template #icon><icon-plus /></template>
                <template #default>新建文件</template>
                </a-doption>
              <a-doption v-if="node.fileType == 'folder'" @click="handleSelect(node, 2)">
                <template #icon><icon-plus /></template>
                <template #default>新建文件夹</template></a-doption>
              <a-doption v-if="node.fileType == 'folder'" @click="handleSelect(node, 3)">
                <template #icon><icon-upload /></template>
                <template #default>在此处上传文件</template></a-doption>   
              <!-- 通用操作 -->
              <a-doption @click="handleSelect(node, 4)">
                <template #icon><icon-edit /></template>
                <template #default>重命名</template></a-doption>
              <a-doption @click="handleSelect(node, 5)">
                <template #icon><icon-minus /></template>
                <template #default>删除</template></a-doption>              
            </div>
          </template>
        </a-dropdown>
      </template>
    </VTreeSearch> 
    <!-- 上传文件 -->
    <el-dialog v-model="uploadVis" title="上传" width="500px" class="uploadDialog">
      <File_uploader :order="uploadOrder" :actionUrl="arctionUrl" ref="fileUploader"></File_uploader>
      <div class="buttonOutter">
        <el-button class="uploadButton" @click="uploadSubmit">确认添加</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { VTreeSearch } from '@wsfe/vue-tree'
import { ElMessage, ElMessageBox } from 'element-plus';
import File_uploader from '../util/File_uploader.vue';

import { reactive, ref } from 'vue';
import { useServeURL } from '@/store/network';
import { storeToRefs } from 'pinia';
import { notePath } from '@/model/bussModel/notePath';
import { fileChange, getNotePath } from '@/network/note';

let uploadOrder = ['application/pdf', 'text/markdown', 'text/plain']
const serveStore = useServeURL()
const { mergeURL } = storeToRefs(serveStore)
let arctionUrl = mergeURL.value("note-admin/uploadNote")

// 获取节点数据
let data = ref([] as notePath[])
getNotePath().then(res => {
  data.value = res.data
})

// 传递数据到 useNote
import { useNote } from '@/store/note';
const noteStore = useNote()
let { selectNode } = storeToRefs(noteStore)
// 处理选中的节点信息并存储
const handleNodeSelect = (node:{fileType:string, id: string, title:string}) => {
  if(node.fileType == "folder") return
  selectNode.value.fileType = node.fileType
  selectNode.value.id = node.id
  selectNode.value.title = node.title
}


let filetree = ref()
let type = localStorage.getItem("type")
let uploadVis = ref(false)


const handleAppend = (node:any, title:string, id:string, fileType:string) => {
  filetree.value.append({ title: title, id: id, fileType:fileType}, node.id)
}
const handleRemove = (node:any) => {
  filetree.value.remove(node.id)
}
const handleUpdateSingleName = (node:any, newName:string, id:string) => {
  filetree.value.updateNode(node.id, { title: newName, id:id })
}

const toServe = (action : number, origin: string, final:string) => {
  fileChange({action, origin, final}).then(res => {
    if(res.code == 0) {
      ElMessage.error("操作未生效" + res.msg); return
    }
  }).catch(res => {
    ElMessage.error("操作未生效" + res.msg);
  })
}

let fileUploader = ref()
let uploadNode = ref()
const handleSelect = (node:any, action:number) => {
  switch(action) {
    case 1: {
      toServe(1, node.id, node.id + "\\新建文件.md")
      handleAppend(node, "新建文件", node.id + "\\新建文件.md", "md");
      break;
    }
    case 2: {
      toServe(2, node.id, node.id + "\\新建文件夹")
      handleAppend(node, "新建文件夹", node.id + "\\新建文件夹", "folder");
      break;
    }
    case 3: {
      ElMessage.info("未建设好"); break
      uploadVis.value = true
      uploadNode.value = node
      break;
    }
    case 4: {
      ElMessageBox.prompt('请输入新的文件名', 'Tip', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
      }).then(({ value }) => {
        const lastIndex = node.id.lastIndexOf('.');
        let fixx = lastIndex == -1 ? "" : node.id.substring(lastIndex + 1);
        if(node.fileType != "folder") fixx = "." + fixx
        const lastSplit = node.id.lastIndexOf('\\');
        console.log(lastSplit)
        let before = lastSplit == -1 ? "" : node.id.substring(0, lastSplit);
        before = before == "" ? before : before + "\\"
        let newId = before + value + fixx
        console.log(newId)
        handleUpdateSingleName(node, value, newId)
        toServe(4, node.id, newId)
        ElMessage({ type: 'success', message: `修改为 ${value}`})
      }).catch(() => {
          ElMessage({ type: 'info', message: '取消', })
      })
      break;
    }
    case 5: {
      ElMessageBox.confirm('确定要删除吗？可能无法恢复', 'Warning', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning',
      }).then(() => {
        handleRemove(node)
        toServe(5, node.id, node.id)
        ElMessage({ type: 'success', message: '删除成功', })
      }).catch(() => {
        ElMessage({type: 'info', message: '取消删除', })
      })
      break;
    }
  }
};

const uploadSubmit = () => {
  let fileName = fileUploader.value.getFileName()
  if(fileName == "") return
  const lastIndex = fileName.lastIndexOf('.');
  let type = lastIndex == -1 ? "" : fileName.substring(lastIndex + 1);
  toServe(3, fileName, uploadNode.value.id + "\\" + fileName)
  handleAppend(uploadNode.value, uploadNode.value.id + "\\" + fileName, fileName, type)
  uploadVis.value = false
}

</script>

<style lang="less">
@import '~@wsfe/vue-tree/style.css';
@import '~@wsfe/vue-tree/src/styles/index.less';
</style>

<style lang="less">
.buttonOutter {
  width: 100%;
  display: flex;
  justify-content: center;
}
</style>