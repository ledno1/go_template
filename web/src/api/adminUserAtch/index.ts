import { http, jumpExport } from '@/utils/http/axios';

// 获取授权/密钥列表
export function List(params) {
  return http.request({
    url: '/adminUserAtch/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除授权/密钥
export function Delete(params) {
  return http.request({
    url: '/adminUserAtch/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑授权/密钥
export function Edit(params) {
  return http.request({
    url: '/adminUserAtch/edit',
    method: 'POST',
    params,
  });
}




// 获取授权/密钥指定详情
export function View(params) {
  return http.request({
    url: '/adminUserAtch/view',
    method: 'GET',
    params,
  });
}