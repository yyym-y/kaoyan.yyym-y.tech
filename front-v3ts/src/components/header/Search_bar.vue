<template>
	<div class="search-bar">
		<el-row :gutter="10">
			<el-col :span="16">
				<el-input
						v-model="searchInput"
						@keyup.enter="toSearch"
						placeholder="请输入要搜索的内容, 宁可少字也不要错字~~"
						class="search-input">
				</el-input>
			</el-col>
			<el-col :span="8">
				<el-button plain class="search-but" @click="toSearch">
					<el-icon style="margin-right : 4px"><Search /></el-icon>
					搜索
				</el-button>
			</el-col>
		</el-row>
	</div>
</template>

<script lang="ts" setup>
	import { useRouter } from 'vue-router';

	import { useSearch } from '@/store/search';
	import { storeToRefs } from 'pinia';

	const router = useRouter();
	const searchStore = useSearch()
	let { searchInput } = storeToRefs(searchStore)

	const toSearch = () => {
		console.log("SerchInput : " + searchInput.value)
		router.push({
			name:"search",
			params:{ input: searchInput.value }
		})
	}
</script>

<style lang="less">
</style>