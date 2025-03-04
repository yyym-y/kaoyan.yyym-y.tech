<template>
<div>
    <el-upload :action="uploadUrl"
        :headers="header" :limit="1"
        :on-success="handleSuccess" :on-remove="handleRemove"
        v-model:file-list="fileList" list-type="picture">
        <el-button type="primary">点击上传</el-button>
        <template #tip>
            <div class="el-upload__tip">
                只能上传jpg/png文件，且不超过500kb
            </div>
        </template>
    </el-upload>
</div>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue";

import { useServeURL } from "@/store/network";

import { UploadProps, UploadUserFile } from "element-plus";
import { storeToRefs } from "pinia";

let serveStore = useServeURL()
let { mergeURL } = storeToRefs(serveStore)

let fileName = ref("")
let tempUrl = ref("")
let fileList = ref<UploadUserFile[]>([])
let uploadUrl = ref(mergeURL.value("uploadImg", "/root"))
let header:Record<string, any> = reactive({
    'Authorization' : localStorage.getItem("token")
})

const handleRemove: UploadProps['onRemove'] = (file: any, fileList: any) => {
    console.log(file, fileList);
}
const handleSuccess : UploadProps['onSuccess'] = (response: any, file: any, fileList: any) => {
    fileName.value = file.name
    tempUrl.value = file.url
    console.log(fileName.value, tempUrl.value)
}
const getFileName = () => {
    return fileName.value
}
const getTempUrl= () => {
    return tempUrl.value
}

defineExpose({ getFileName, getTempUrl })
</script>

<style>

</style>