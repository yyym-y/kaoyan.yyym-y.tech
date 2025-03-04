<template>
    <div class="icon-outter">
      <el-image
        class="icon-size"
        :src="require('@/assets/img1.gif')" ></el-image>
    </div>
</template>

<script lang="ts" setup>

</script>


<style lang="less">
@import url(@/public.less);
.icon-outter {
    height: @Icon-Outer-Height;
    .icon-size {
        width: @Icon-width;
        height: @Icon-height;
    }
}
</style>