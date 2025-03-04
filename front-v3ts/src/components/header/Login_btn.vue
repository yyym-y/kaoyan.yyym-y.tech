<template>
  <div class="login-outer">
    <el-button type="info"
      @click="LogInFormVis=true"
      class="login-btn"
      round>登录账户</el-button>

    <el-dialog title="账户登录" v-model="LogInFormVis" width="27em">
      <Login_form></Login_form>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';

  import Login_form from './form/Login_form.vue';

	const LogInFormVis = ref(false)
</script>

<style lang="less">
.login-outer {
  float: right;
}
</style>
