<template>
  <div>
    搜索的内容是 ：{{ input }}
    <buildEmpty></buildEmpty>
  </div>
</template>

<script lang="ts" setup>

import buildEmpty from '@/components/util/Build_Empty.vue';
import { useSearch } from '@/store/search';

const searchStore = useSearch()
const input = searchStore.searchInput
</script>

<style>

</style>