<template>
  <div>
    <el-button type="primary" @click="dialogVisible = true"
      style="margin-bottom: 8px;">
      <el-icon style="margin-right: 8px;"><Edit /></el-icon>添加影视</el-button>

    <el-dialog title="添加影视" v-model="dialogVisible" width="44%" append-to-body>
      <el-form :label-position="'left'" label-width="100px" v-model="createFromData">
        <el-form-item label="影视名称">
          <el-input v-model="createFromData.vedioName"></el-input>
        </el-form-item>
        <el-form-item label="上传封面">
            <Image_uploader ref="createImgUpload"></Image_uploader>
        </el-form-item>
        <el-form-item label="影视类别">
          <el-select v-model="createFromData.typeId" placeholder="请选择类别">
            <el-option v-for="item in typeList" :key="item.typeId + 'type'"
              :label="item.name" :value="item.typeId"></el-option>
          </el-select>              
        </el-form-item>
        <el-form-item label="影视描述">
          <el-input type="textarea" :autosize="{ minRows: 2, maxRows: 4}"
              placeholder="请输入内容" v-model="createFromData.description"> </el-input>
        </el-form-item>
        <el-form-item label="前台可见性">
          <el-switch v-model="createFromData.visable" active-text="可见" inactive-text="不可见"></el-switch>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm()">创建影视</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { createVedioReq } from '@/model/reqModel/createVedioReq';
import { reactive, ref } from 'vue';

import { getVedioType } from '@/network/pub';
import { vedioType } from '@/model/bussModel/vedioType';

import Image_uploader from '../util/Image_uploader.vue';
import { createVedio } from '@/network/admin';
import { ElMessage } from 'element-plus/es';

let dialogVisible = ref(false)
let createFromData:createVedioReq = reactive({} as createVedioReq)
let typeList = ref([] as vedioType[])
getVedioType().then(res => {
    if(res.code == 0) return
    typeList.value = res.data
})

let createImgUpload = ref()

const submitForm = () => {
    createFromData.vedioCover = createImgUpload.value.getFileName()
    createVedio(createFromData).then(res => {
      if(res.code == 0) {
        ElMessage.error("创建失败 " + res.msg); return;
      }
      ElMessage.success("创建成功")
      setTimeout(() => location.reload(), 700);
    })
}

</script>

<style>

</style>