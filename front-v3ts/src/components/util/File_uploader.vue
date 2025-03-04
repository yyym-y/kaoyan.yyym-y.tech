<template>
  <div>
    <el-upload style="margin-top: 16px" drag :headers="header" :action="actionUrl"
        :show-file-list="true" :limit="1"
        :disabled="finished"
        :on-success="handleVideoSuccess"
        :before-upload="beforeUploadVideo"
        :file-list="fileList"
        :on-progress="uploadVideoProcess">
        <i v-if="!finished"  class="el-icon-upload"></i>
        <div v-if="!finished" class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
        <el-result v-if="finished" icon="success" style="margin-top: 20px;n"></el-result>
    </el-upload>
  </div>
</template>

<script lang="ts" setup>
import { ElMessage } from 'element-plus';
import { reactive, ref } from 'vue';


const props = defineProps<{ actionUrl:string, order:string[] }>()

const beforeUploadVideo = (file: { type: string; }) => {
    if (props.order.indexOf(file.type) === -1) {
        ElMessage.error("请上传正确的格式")
        return false
    }
    return true;
}

const fileList = ref([])
const header = reactive({
    'Authorization' : localStorage.getItem("token")
})

let videoUploadPercent = ref(0)
const uploadVideoProcess = (event: any, file: any, fileList: any) => {
    videoUploadPercent.value = Number(file.percentage)
}

let fileName = ref("")
let finished = ref(false)
const handleVideoSuccess = (res: any, file: { name: string; }) => {
    finished.value = true
    fileName.value = file.name
}

const getFileName = () => {
    return fileName.value
}

defineExpose( { getFileName } )
</script>

<style>

</style>