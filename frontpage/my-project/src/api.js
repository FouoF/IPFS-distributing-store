import axios from 'axios';

// 配置 axios 实例
const api = axios.create({
  // baseURL: 'admin',// 假设 nginx 配置了转发到后端服务的路径
  baseURL: 'http://localhost:8000/admin', //test url
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