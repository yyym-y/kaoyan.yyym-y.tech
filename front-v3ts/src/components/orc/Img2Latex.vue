<template>
  <div>
    <el-upload class="upload-demo" drag :action="uploadUrl" v-show="ifSuccess"
      :on-success="handleSuccess" :on-remove="handleRemove" :headers="header">
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text"> Drop file here or <em>click to upload</em></div>
      <template #tip>
        <div class="el-upload__tip"> jpg/png files with a size less than 500kb </div>
      </template>
    </el-upload>
    <el-divider />
    <h3 class="result"> 结果展示 </h3>
    <el-input v-model="result" style="width:100%" 
      :autosize="{ minRows: 2, maxRows: 4 }" type="textarea"/>
    <el-divider />
    <el-form label-width="80px">
      <el-form-item label="选择模型">
        <el-select v-model="modelValue" placeholder="Select" style="width: 240px">
          <el-option key="model1" label="轻量公式识别模型" :value="1"/>
          <el-option key="model2" label="标准公式识别模型" :value="2"/>
        </el-select>        
      </el-form-item>
      <el-form-item label="">
        <el-button type="primary" @click="begin">开始识别</el-button>    
      </el-form-item>
    </el-form>
  </div>
</template>

<script lang="ts" setup>
import { useServeURL } from "@/store/network";
import { ElMessage, UploadProps } from "element-plus";
import { storeToRefs } from "pinia";
import { reactive, ref } from "vue";
import { latexOrc } from '../../network/orc';

let serveStore = useServeURL()
let { mergeURL } = storeToRefs(serveStore)
let uploadUrl = ref(mergeURL.value("uploadImg", "/orc"))

let fileName = ref("")
let temFileUrl = ref("")
let ifSuccess = ref(true)
let header:Record<string, any> = reactive({
    'Authorization' : localStorage.getItem("token")
})
const handleRemove: UploadProps['onRemove'] = (file: any, fileList: any) => {
    console.log(file, fileList);
}
const handleSuccess : UploadProps['onSuccess'] = (response: any, file: any, fileList: any) => {
    console.log(response)
    fileName.value = response.data.fileName
    temFileUrl.value = file.url
}

let modelValue = ref(2)

let result = ref("")

const begin = () => {
  latexOrc({
    fileName: fileName.value,
    model : modelValue.value
  }).then(res => {
    console.log(res)
    if(res.code == 0) {
      ElMessage.error("识别失败[本服务器异常] " + res.msg)
      return
    }
    if(res.data.res == null) {
      ElMessage.error("识别失败[api服务器异常] " + res.data.err_info.err_msg + "/" + res.data.err_info.err_type)
      return
    }
    result.value = res.data.res.latex
  })
}
</script>

<style lang="less" scoped>
@import url("@/public.less");
.result {
  font-weight: 500;
}
</style>