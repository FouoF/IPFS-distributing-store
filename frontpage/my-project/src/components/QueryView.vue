<template>
  <div>
    <div class="header-container">
      <h1>数据查询</h1>
      <!-- <el-button type="primary" @click="">添加索引</el-button> -->
      <!-- <el-button @click="testdownloadRecord" type="primary" size="small" >testDownload</el-button> -->
    </div>

    <el-row :gutter="20">
      <el-col :span="8">
        <el-select v-model="selectedOption1" placeholder="请选择数据板块" @change="fetchOptions2">
          <el-option v-for="option in options1" :key="option" :label="option" :value="option"></el-option>
        </el-select>
      </el-col>

      <el-col :span="8">
        <el-select v-model="selectedOption2" placeholder="请选择索引" @change="fetchOptions3">
          <el-option v-for="option in options2" :key="option" :label="option" :value="option"></el-option>
        </el-select>
      </el-col>

      <el-col :span="8">
        <el-select v-model="selectedOption3" placeholder="请选择索引" @change="fetchData">
          <el-option v-for="option in options3" :key="option" :label="option" :value="option"></el-option>
        </el-select>
      </el-col>
    </el-row>

    <el-table v-if="records.length" :data="records" style="width: 100%">
      <el-table-column prop="leafname" label="Name"></el-table-column>
      <el-table-column prop="cid" label="CID"></el-table-column>
      <el-table-column label="Action" width="180">
        <template #default="{ row }">
          <el-button @click="downloadRecord(row)" type="primary" size="small" :loading="row.downloading">Download</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { ref, onMounted, watch } from 'vue';
import { listIndex, downloadFromIPFS, getFileMetadata, listLeaf } from '@/api'; // 引入 API 接口
import { ElMessage, ElMessageBox } from 'element-plus';

export default {
  setup() {
    const selectedOption1 = ref('');
    const selectedOption2 = ref('');
    const selectedOption3 = ref('');
    const options1 = ref([]);
    const options2 = ref([]);
    const options3 = ref([]);
    const records = ref([]);

    // 获取选项1的接口
    const fetchOptions1 = async () => {
      try {
        const index = { index: { name: '', L1: '', L2: '', leafname: '' } };
        console.log(index);
        const response = await listIndex(index); // 假设你有这个接口
        options1.value = response.data.index;
        if (options1.value.length > 0) {
          selectedOption1.value = options1.value[0]; // 默认选择第一个选项
          fetchOptions2(); // 获取第二个下拉框的选项
        }
      } catch (error) {
        console.error('Failed to fetch options1', error);
      }
    };

    // 获取选项2的接口（根据选项1的值获取）
    const fetchOptions2 = async () => {
      if (!selectedOption1.value) return;
      try {
        const index = { index: { name: selectedOption1.value, L1: '', L2: '', leafname: '' } };
        console.log(index);
        const response = await listIndex(index); // 假设你有这个接口
        options2.value = response.data.index;
        if (options2.value.length > 0) {
          selectedOption2.value = options2.value[0]; // 默认选择第一个选项
          fetchOptions3(); // 获取第三个下拉框的选项
        }
      } catch (error) {
        console.error('Failed to fetch options2', error);
      }
    };

    // 获取选项3的接口（根据选项1和选项2的值获取）
    const fetchOptions3 = async () => {
      if (!selectedOption1.value || !selectedOption2.value) return;
      try {
        const index = { index: { name: selectedOption1.value, L1: selectedOption2.value, L2: '', leafname: '' } };
        console.log(index);
        const response = await listIndex(index); // 假设你有这个接口
        options3.value = response.data.index;
        if (options3.value.length > 0) {
          selectedOption3.value = options3.value[0]; // 默认选择第一个选项
          fetchData(); // 获取数据
        }
      } catch (error) {
        console.error('Failed to fetch options3', error);
      }
    };

    // 获取记录的接口
    const fetchData = async () => {
      if (selectedOption1.value && selectedOption2.value && selectedOption3.value) {
        try {
          const index = { index: { name: selectedOption1.value, L1: selectedOption2.value, L2: selectedOption3.value, leafname: '' } };
          console.log(index);
          const response = await listLeaf(index);
          records.value = response.data.leaf;
        } catch (error) {
          console.error('Failed to fetch records', error);
        }
      } else {
        records.value = [];
      }
    };

    // 下载记录的逻辑
    const downloadRecord = async (row) => {
      try {
        // 获取文件元数据（大小、类型）
        const { size, type } = await getFileMetadata(row.cid);
        
        // 类型检查
        if (type === 'directory') {
          ElMessage.error('不支持下载目录');
          return;
        }

        // 大文件确认（超过50MB）
        if (size > 50 * 1024 * 1024) {
          await ElMessageBox.confirm(
            `文件较大（${(size / 1024 / 1024).toFixed(1)}MB），确认下载吗？`,
            '下载确认',
            { type: 'warning' }
          );
        }

        // 显示进度提示
        const message = ElMessage({
          message: `下载准备中...`,
          type: 'info',
          duration: 0,
          showClose: true
        });

        // 执行下载
        const success = await downloadFromIPFS(
          row.cid,
          row.leafname || `file_${row.cid.slice(0, 8)}`, // 使用 leafname 或生成默认名
          (p) => {
            message.message = `下载进度: ${p}%`;
          }
        );

        // 关闭进度提示
        message.close();

        if (success) {
          ElMessage.success('下载已开始');
        } else {
          ElMessage.warning('已切换备用网关下载');
        }
      } catch (error) {
        if (error !== 'cancel') { // 忽略用户取消的情况
          ElMessage.error(`下载失败: ${error.message}`);
        }
      }
    };

    // const testdownloadRecord = async () => {
    //   try {
    //     // 获取文件元数据（大小、类型）
    //     const row = {leafname: "流汗黄豆.png", cid: "QmUDjVT54GiUFkW8Nc4Q2nTVKxEDmirKtK9oRFRKmhZLzL"}
    //     const { size, type } = await getFileMetadata(row.cid);
        
    //     // 类型检查
    //     if (type === 'directory') {
    //       ElMessage.error('不支持下载目录');
    //       return;
    //     }

    //     // 大文件确认（超过50MB）
    //     if (size > 50 * 1024 * 1024) {
    //       await ElMessageBox.confirm(
    //         `文件较大（${(size / 1024 / 1024).toFixed(1)}MB），确认下载吗？`,
    //         '下载确认',
    //         { type: 'warning' }
    //       );
    //     }

    //     // 显示进度提示
    //     const message = ElMessage({
    //       message: `下载准备中...`,
    //       type: 'info',
    //       duration: 0,
    //       showClose: true
    //     });

    //     // 执行下载
    //     const success = await downloadFromIPFS(
    //       row.cid,
    //       row.leafname || `file_${row.cid.slice(0, 8)}`, // 使用 leafname 或生成默认名
    //       (p) => {
    //         message.message = `下载进度: ${p}%`;
    //       }
    //     );

    //     // 关闭进度提示
    //     message.close();

    //     if (success) {
    //       ElMessage.success('下载已开始');
    //     } else {
    //       ElMessage.warning('已切换备用网关下载');
    //     }
    //   } catch (error) {
    //     if (error !== 'cancel') { // 忽略用户取消的情况
    //       ElMessage.error(`下载失败: ${error.message}`);
    //     }
    //   }
    // };

    // 页面加载时获取选项
    onMounted(() => {
      fetchOptions1(); // 获取第一个下拉框的选项
    });

    // 监听 selectedOption1 的变化，变化时调用 fetchOptions2
    watch(selectedOption1, () => {
      fetchOptions2();
    });

    // 监听 selectedOption2 的变化，变化时调用 fetchOptions3
    watch(selectedOption2, () => {
      fetchOptions3();
    });

    return {
      selectedOption1,
      selectedOption2,
      selectedOption3,
      options1,
      options2,
      options3,
      records,
      fetchData,
      downloadRecord,
      // testdownloadRecord
    };
  }
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
</style>
