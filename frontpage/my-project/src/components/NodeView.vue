<template>
  <div>
    <div class="header-container">
      <h1>节点管理</h1>
      <el-button type="primary" @click="openAddNodeDialog">接入节点</el-button>
    </div>

    <!-- 表格部分（绑定获取的节点数据） -->
    <el-table :data="nodes" style="width: 100%">
      <el-table-column prop="id" label="编号"></el-table-column>
      <el-table-column prop="name" label="名称"></el-table-column>
      <el-table-column prop="status" label="状态"></el-table-column>
      <el-table-column prop="addr" label="地址"></el-table-column>
      <el-table-column label="Actions">
        <template #default="{ row }">
          <el-button size="small" @click="removeNode(row)">Delete</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 添加节点的弹出表单对话框 -->
    <el-dialog
      title="Add New Node"
      v-model="dialogVisible"
      width="400px"
      @close="resetForm"
    >
      <el-form :model="newNode" ref="form">
        <el-form-item label="ID" :rules="[{ required: true, message: 'Please input ID', trigger: 'blur' }]">
          <el-input v-model="newNode.node.id"></el-input>
        </el-form-item>
        <el-form-item label="Name" :rules="[{ required: true, message: 'Please input name', trigger: 'blur' }]">
          <el-input v-model="newNode.node.name"></el-input>
        </el-form-item>
        <el-form-item label="Address" :rules="[{ required: true, message: 'Please input address', trigger: 'blur' }]">
          <el-input v-model="newNode.node.addr"></el-input>
        </el-form-item>
      </el-form>

      <template #footer>
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="addNewNode">Add</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { listNodes, addNode, removeNode } from '@/api'; // 导入 API

export default {
  name: 'NodeView',
  data() {
    return {
      nodes: [],
      dialogVisible: false,
      newNode: {node: {id: '', name: '', addr: '', status: '' }},
    };
  },
  created() {
    this.fetchNodes();
  },
  methods: {
    async fetchNodes() {
      try {
        const response = await listNodes();
        this.nodes = response.data.node;
      } catch (error) {
        this.$message.error('Failed to load nodes');
      }
    },

    openAddNodeDialog() {
      this.dialogVisible = true;
    },

    async addNewNode() {
      try {
        await addNode(this.newNode);
        this.dialogVisible = false;
        this.resetForm();
        this.fetchNodes(); // 重新加载节点数据
      } catch (error) {
        this.$message.error('Failed to add node');
      }
    },

    async removeNode(node) {
      try {
        await removeNode(node.addr);
        this.fetchNodes(); // 重新加载节点数据
      } catch (error) {
        this.$message.error('Failed to remove node');
      }
    },

    resetForm() {
      this.newNode ={node: { id: '', name: '', addr: '' }};
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

