<template>
  <div class="nameOuter">
    <el-dropdown  @command="handleCommand" class="setupStyle">
      <span class="el-dropdown-link">
        {{ username }}
        <el-icon class="el-icon--right"><arrow-down /></el-icon>
      </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item v-if="type == '0'" command="admin">控制台</el-dropdown-item>
        <el-dropdown-item v-if="type == '0'" command="developLog">开发日志</el-dropdown-item>
        <el-dropdown-item command="changepwd">修改密码</el-dropdown-item>
        <el-dropdown-item command="exit">退出登录</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
    <el-dialog title="修改密码" v-model="changepwd" width="31em">
        <ChangePwd_form></ChangePwd_form>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { useRouter } from 'vue-router';

  import { ArrowDown } from '@element-plus/icons-vue'

  import ChangePwd_form from './form/ChangePwd_form.vue';

  let username = ref(localStorage.getItem("username"))
  let type = ref(localStorage.getItem("type"))
  let changepwd = ref(false)
  const router = useRouter()

  const exitLog = () => {
    localStorage.setItem("username", "");
    localStorage.setItem("token", "");
    localStorage.setItem("type", "-1");
    setTimeout(() => location.reload(), 700);
  } 

  const jumpToAdmin = () => {
    router.push({ name: "admin" }).catch(err => {
      console.error(err);
    });
  }

  const jumpToLog = () => {
    router.push({ name: "log" }).catch(err => {
      console.error(err);
    });
  }

  const handleCommand = (command:string) => {
    switch (command) {
        case "exit": { exitLog(); break;}
        case "changepwd": { changepwd.value = true; break;}
        case "admin": { jumpToAdmin(); break;}
        case "developLog" : {jumpToLog(); break;}
    }
  }
</script>

<style>
.example-showcase .el-dropdown-link {
  cursor: pointer;
  color: var(--el-color-primary);
  display: flex;
  align-items: center;
}
.setupStyle .el-dropdown-link:focus {
  outline: none;
}
.nameOuter {
  margin-top: 16px;
  margin-right: 20px;
  float: right;
}
.el-icon--right {
  transform: translateY(3px);
}
</style>