<template>
  <div class="header">
		<el-row :gutter="10">
			<el-col :span="16">
				<Search_Bar></Search_Bar>
			</el-col>
			<el-col :span="8" class="righter">
				<Login_Btn v-if="username == '' || username == null"></Login_Btn>
				<Username_Link v-else :username="username"></Username_Link>
			</el-col>
		</el-row>
  </div>
</template>

<script lang="ts" setup>
	import { ref } from 'vue';

	import Search_Bar from './Search_bar.vue'
	import Login_Btn from './Login_btn.vue'
	import Username_Link from './Username_link.vue'

	const username = ref(localStorage.getItem("username"))
</script>

<style lang="less">
@import url('@/public.less');
.header {
  height: 97%;
  .Border-Bottom(#C0C4CC);
}
</style>