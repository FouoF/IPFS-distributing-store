<template>
  <div>
    <div class="header-container">
      <h1>数据源视图</h1>
      <el-button type="primary" @click="openAddEndpointDialog" class="add-endpoint-btn">接入数据源</el-button>
    </div>

    <!-- 表格部分（绑定获取的端点数据） -->
    <el-table :data="endpoints" style="width: 100%">
      <el-table-column prop="id" label="编号"></el-table-column>
      <el-table-column prop="description" label="描述"></el-table-column>
      <el-table-column prop="addr" label="地址"></el-table-column>
      <el-table-column label="Actions">
        <template #default="{ row }">
          <el-button size="small" @click="removeEndpoint(row)">Delete</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 添加端点的弹出表单对话框 -->
    <el-dialog
      title="添加数据源"
      v-model="dialogVisible"
      width="400px"
      @close="resetForm"
    >
      <el-form :model="newEndpoint" ref="form">
        <el-form-item label="编号" :rules="[{ required: true, message: 'Please input ID', trigger: 'blur' }]">
          <el-input v-model="newEndpoint.id"></el-input>
        </el-form-item>
        <el-form-item label="描述" :rules="[{ required: true, message: 'Please input description', trigger: 'blur' }]">
          <el-input v-model="newEndpoint.description"></el-input>
        </el-form-item>
        <el-form-item label="地址" :rules="[{ required: true, message: 'Please input address', trigger: 'blur' }]">
          <el-input v-model="newEndpoint.addr"></el-input>
        </el-form-item>
        <el-form-item label="版块" :rules="[{ required: true, message: 'Please input address', trigger: 'blur' }]">
          <el-input v-model="newEndpoint.index.name"></el-input>
        </el-form-item>
        <el-form-item label="一级索引" :rules="[{ required: true, message: 'Please input address', trigger: 'blur' }]">
          <el-input v-model="newEndpoint.index.l1"></el-input>
        </el-form-item>
        <el-form-item label="二级索引" :rules="[{ required: true, message: 'Please input address', trigger: 'blur' }]">
          <el-input v-model="newEndpoint.index.l2"></el-input>
        </el-form-item>
      </el-form>

      <template #footer>
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="addNewEndpoint">Add</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { listEndpoints, addEndpoint, removeEndpoint } from '@/api'; // 导入 API

export default {
  name: 'EndpointView',
  data() {
    return {
      endpoints: [], // 存储从后端获取的端点数据
      dialogVisible: false,
      newEndpoint: {
      id: '',
      description: '',
      address: '',
      index: { name: '', l1: '', l2: '', leafname: '' } // 直接定义新对象
    }
    };
  },
  created() {
    this.fetchEndpoints(); // 页面加载时获取端点数据
  },
  methods: {
    // 获取端点列表
    async fetchEndpoints() {
      try {
        const response = await listEndpoints();
        this.endpoints = response.data.endpoint; // 根据接口文档，数据在 `endpoint` 字段中
      } catch (error) {
        this.$message.error('Failed to load endpoints');
      }
    },

    // 打开添加端点的弹出框
    openAddEndpointDialog() {
      this.dialogVisible = true;
    },

    // 添加新的端点
    async addNewEndpoint() {
      try {
        const success = await addEndpoint(this.newEndpoint);
        if (success) {
          this.dialogVisible = false;  // 关闭对话框
          this.resetForm();  // 重置表单
          this.fetchEndpoints();  // 重新加载端点数据
        } else {
          this.$message.error('Failed to add endpoint');
        }
      } catch (error) {
        this.$message.error('Failed to add endpoint');
      }
    },

    // 删除端点
    async removeEndpoint(endpoint) {
      try {
        await removeEndpoint(endpoint.index);  // 使用 `index` 来删除端点
        this.fetchEndpoints();  // 重新加载端点数据
      } catch (error) {
        this.$message.error('Failed to remove endpoint');
      }
    },

    // 重置表单
    resetForm() {
      this.newEndpoint = { id: '', description: '', addr: '' };
    },
  },
};
</script>

<style scoped>
.header-container {
  display: flex;
  justify-content: space-between; /* 将标题和按钮放到两侧 */
  align-items: center; /* 垂直居中对齐 */
  height: 50px; /* 设置容器固定高度，避免内容增加导致高度过高 */
  padding: 0 20px; /* 为左右两侧添加内边距，避免紧贴边缘 */
  background-color: #f4fcf7; /* 背景色 */
  margin-bottom: 20px; /* 给下面的内容添加一些间距 */
}
/* 页面容器样式 */
div {
  padding: 20px;
  background-color: #f4f7fc;
  font-family: 'Arial', sans-serif;
}

/* 标题样式 */
h1 {
  font-size: 2rem;
  color: #333;
  margin-bottom: 20px;
}

/* 布局容器样式 */
.el-row {
  margin-bottom: 30px;
}

/* 每一列的间隔 */
.el-col {
  padding: 10px;
}

/* 下拉框的样式 */
.el-select {
  width: 100%;
  border-radius: 4px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  font-size: 14px;
  background-color: #fff;
  transition: border 0.3s ease;
}

/* 聚焦时下拉框样式 */
.el-select:focus {
  border-color: #409eff;
}

/* 表格样式 */
.el-table {
  margin-top: 20px;
  border-radius: 8px;
  overflow: hidden;
  background-color: #fff;
}

/* 表格头部样式 */
.el-table th {
  background-color: #f2f6fc;
  color: #333;
  font-weight: bold;
}

/* 表格内容行样式 */
.el-table td {
  padding: 15px;
  color: #666;
}

/* 操作按钮样式 */
.el-button {
  background-color: #409eff;
  color: white;
  border-radius: 4px;
  transition: background-color 0.3s ease;
}

/* 鼠标悬停按钮样式 */
.el-button:hover {
  background-color: #66b1ff;
}

/* 弹出框样式 */
.el-dialog {
  border-radius: 8px;
  overflow: hidden;
  background-color: #fff;
}

/* 表单项样式 */
.el-form-item {
  margin-bottom: 20px;
}

/* 表单输入框样式 */
.el-input {
  border-radius: 4px;
  background-color: #f7f7f7;
  border: 1px solid #dcdfe6;
  transition: border-color 0.3s ease;
}

/* 聚焦时输入框样式 */
.el-input:focus {
  border-color: #409eff;
}

/* 按钮内边距与间距 */
.el-button + .el-button {
  margin-left: 10px;
}

/* 当没有数据时的提示文字 */
.el-table--empty span {
  color: #b5b5b5;
  font-size: 14px;
  text-align: center;
}

/* 设置选择框的间隔 */
.el-row .el-col {
  margin-bottom: 20px;
}
</style>
