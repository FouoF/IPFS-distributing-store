<template>
  <div id="app">
    <!-- 这里是应用的导航栏 -->
    <el-container>
      <el-header>
        <el-menu mode="horizontal" :default-active="activeMenu">
          <el-menu-item index="1" @click="navigateTo('endpoint')">Endpoint</el-menu-item>
          <el-menu-item index="2" @click="navigateTo('node')">Node</el-menu-item>
          <el-menu-item index="3" @click="navigateTo('query')">Query</el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <!-- 这里是根据路由显示的内容 -->
        <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import { useRouter } from 'vue-router';

export default defineComponent({
  name: 'App',
  setup() {
    const router = useRouter();
    const activeMenu = '1'; // 默认选中的菜单项

    const navigateTo = (view) => {
      // 跳转到不同的视图
      router.push(`/${view}`);
    };

    return {
      activeMenu,
      navigateTo
    };
  }
});
</script>

<style>
/* 可以自定义样式，这里只提供基础布局 */
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.el-header {
  background-color: #20a0ff;
  color: #fff;
  text-align: center;
}

.el-main {
  margin-top: 20px;
}

.el-menu {
  background-color: transparent;
}
</style>
