<template>
  <div>
    <el-form :label-position="'left'" label-width="100px" :model="formInfo">
			<el-form-item label="用户名/邮箱">
				<el-input v-model="formInfo.key"></el-input>
			</el-form-item>
			<el-form-item label="密码">
				<div>
					<el-input v-model="formInfo.password" show-password></el-input>
					<el-link style="float: right;" :underline="false"
						@click="registerVisible = true">注册账号</el-link>
					<i style="float: right; margin-left: 5px; margin-right: 5px">|</i>
					<el-link style="float: right;" :underline="false"
						@click="refindVisible = true">找回密码</el-link>					
				</div>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="submitForm()">登录</el-button>
			</el-form-item>
    </el-form>

    <el-dialog title="找回密码" v-model="refindVisible" width="31em" append-to-body>
			<Refind_form></Refind_form>
    </el-dialog>
    <el-dialog title="注册账户" v-model="registerVisible" width="31em" append-to-body>
			<Register_form></Register_form>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue';

import { ElMessage } from 'element-plus'
import { login } from '@/network/user';
import { loginReq } from '@/model/reqModel/userReq';

import Refind_form from './Refind_form.vue';
import Register_form from './Register_form.vue';
	// data
	const refindVisible = ref(false)
	const registerVisible = ref(false)
	const formInfo:loginReq = reactive({
		key : "",
		password : "",
	})

	// method
	const submitForm = () => {
		console.log("submit", formInfo)
		if(formInfo.key == "" || formInfo.password == "") {
			ElMessage.error({ message: '账号或密码不能为空' });
			return
		}
		login(formInfo).then((res) => {
			console.log(res)
			if(res.code == 0) {
				ElMessage.error({ message: '账号或密码错误' });
				return
			}
			ElMessage({ message: '登录成功', type: 'success', })
			let data = res.data
			localStorage.setItem('token', data.token)
			localStorage.setItem('username', data.username)
			localStorage.setItem('type', String(data.type))
			setTimeout(() => location.reload(), 700);
		}).catch((err) => {
			console.log(err)
			ElMessage.error({ message: '账号或密码错误' })
		})
	}
</script>

<style>

</style>
