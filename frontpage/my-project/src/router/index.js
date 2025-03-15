import { createRouter, createWebHistory } from 'vue-router';
import EndpointView from '@/components/EndpointView.vue';
import NodeView from '@/components/NodeView.vue';
import QueryView from '@/components/QueryView.vue';

const routes = [
  {
    path: '/endpoint',
    name: 'Endpoint',
    component: EndpointView
  },
  {
    path: '/node',
    name: 'Node',
    component: NodeView
  },
  {
    path: '/query',
    name: 'Query',
    component: QueryView
  },
  {
    path: '/',
    redirect: '/endpoint' // 默认跳转到 Endpoint 视图
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
