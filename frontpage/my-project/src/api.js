import axios from 'axios';

// 配置 axios 实例
const api = axios.create({
  baseURL: 'admin',
  timeout: 10000, // 请求超时设置
  headers: {
    'Content-Type': 'application/json',
  }
});

// 获取节点列表
export const listNodes = () => {
  return api.get('/node/list');
};

// 添加节点
export const addNode = (nodeData) => {
  return api.post('/node/add', nodeData);
};

// 删除节点
export const removeNode = (nodeData) => {
  return api.post('/node/remove', nodeData);
};

// 获取端点列表
export const listEndpoints = () => {
  return api.get('/endpoint/list');
};

// 添加端点
export const addEndpoint = (endpointData) => {
  return api.post('/endpoint/add', endpointData);
};

// 删除端点
export const removeEndpoint = (endpointData) => {
  return api.post('/endpoint/remove', endpointData);
};

export const listIndex = (indexData) => {
  const params = new URLSearchParams();

  // 将 index 对象展开成 key=value 的格式
  for (const key in indexData.index) {
    if (indexData.index[key]) {  // 只添加非空值
      params.append(`index.${key}`, indexData.index[key]);
    }
  }

  return api.get('/index/list', { params });
};
export const createIndex = (indexData) => {
  return api.post('/index/create', indexData);
};

const IPFS_GATEWAYS = [
  'http://ipfs.default.svc.cluster.local:5001',
  // 'http://localhost:5001'
];

// 创建带探活检测的动态网关客户端
const createIpfsClient = async () => {
  for (const gateway of IPFS_GATEWAYS) {
    try {
      // 发送探活请求检查网关可用性
      await axios.post(`${gateway}/api/v0/version`, { timeout: 2000 });
      return axios.create({
        baseURL: gateway,
        timeout: 15000,
        responseType: 'arraybuffer' // 二进制响应处理
      });
    } catch (e) {
      console.warn(`Gateway ${gateway} unavailable, trying next...`);
    }
  }
  throw new Error('All IPFS gateways are down');
};

// 核心下载方法（支持进度回调）
export const downloadFromIPFS = async (cid, filename, onProgress) => {
  const client = await createIpfsClient();
  
  try {
    // 通过 /api/v0/cat 获取原始数据流
    const response = await client.post(`/api/v0/cat?arg=${cid}`, {
      onDownloadProgress: (progressEvent) => {
        if (onProgress && progressEvent.total) {
          const percent = Math.round((progressEvent.loaded * 100) / progressEvent.total);
          onProgress(percent);
        }
      }
    });

    // 处理文件名逻辑
    const finalFilename = filename || `ipfs_${cid.slice(0, 8)}`;

    // 创建 Blob 并触发下载
    const blob = new Blob([response.data]);
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = finalFilename;
    document.body.appendChild(a);
    a.click();
    window.URL.revokeObjectURL(url);
    document.body.removeChild(a);

    return true;
  } catch (error) {
    console.error('IPFS下载失败:', error);
    return false;
  }
};

// 辅助方法：获取文件元数据
export const getFileMetadata = async (cid) => {
  try {
    const client = await createIpfsClient();
    
    // 调用新接口并指定 JSON 格式输出
    const { data } = await client.post('/api/v0/files/stat', null, {
      params: {
        arg: `/ipfs/${cid}`,      // 需要完整 IPFS 路径
        format: 'json',           // 强制返回 JSON 格式
        withLocal: false           // 可选：包含本地存储信息
      }
    });

    // 解析返回数据结构
    return {
      size: data.cumulsize || 0,  // 新接口返回字段名改为小写
      type: data.type, // 类型字段标准化
    };
  } catch (e) {
    console.error('获取元数据失败:', e);
    return { 
      size: 'unknown', 
      type: 'unknown',
    };
  }
};