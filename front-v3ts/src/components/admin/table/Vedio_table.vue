<template>
    <el-table :data="vedioForm" style="width: 100%">
      <el-table-column label="影视 ID" prop="vedioId" width="80"/>
      <el-table-column label="影视名称" prop="vedioName" width="300"/>
      <el-table-column label="影视类别" prop="typeId">
        <template #default="scope">
            <el-tag>{{ typeList[scope.row.typeId - 1].name }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="总共集数" prop="episode"></el-table-column>
      <el-table-column label="前台可见性" prop="visable">
        <template #default="scope">
            <el-tag v-if="scope.row.visable">可见</el-tag>
            <el-tag v-else >不可见</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="right">
        <template #header>
          <el-input v-model="vedioSerch" size="small" placeholder="Type to search" />
        </template>
        <template #default="scope">
          <el-button size="small" @click="console.log(scope.$index, scope.row)">
            Edit
          </el-button>
          <el-button size="small" type="danger" @click="console.log(scope.$index, scope.row)">
            Delete
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </template>

<script lang="ts" setup>
import { vedioType } from '@/model/bussModel/vedioType';
import { vedioTableResp } from '@/model/respModel/vedioTableResp';
import { getVedioTable } from '@/network/admin';
import { getVedioType } from '@/network/pub';
import { ElMessage } from 'element-plus';
import { ref } from 'vue';

let vedioForm = ref([] as vedioTableResp[])
getVedioTable().then(res => {
    if(res.code == 0) {
        ElMessage.error("影视信息获取失败... " + res.msg); return;
    }
    vedioForm.value = res.data
}).catch(err => {ElMessage.error("影视信息获取失败... " + err);})

let typeList = ref([] as vedioType[])
getVedioType().then(res => {
    if(res.code == 0) {
        console.log("影视类别获取失败 " + res.msg); return;
    }
    typeList.value = res.data
}).catch(err => { console.log("影视类别获取失败 " + err);})

let vedioSerch = ref("")

</script>