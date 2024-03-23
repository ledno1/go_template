import { http, jumpExport } from '@/utils/http/axios';

// 获取插件客户端更新列表
export function List(params) {
  return http.request({
    url: '/clientUpdateJ/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除插件客户端更新
export function Delete(params) {
  return http.request({
    url: '/clientUpdateJ/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑插件客户端更新
export function Edit(params) {
  return http.request({
    url: '/clientUpdateJ/edit',
    method: 'POST',
    params,
  });
}




// 获取插件客户端更新指定详情
export function View(params) {
  return http.request({
    url: '/clientUpdateJ/view',
    method: 'GET',
    params,
  });
}



// 导出插件客户端更新
export function Export(params) {
  jumpExport('/clientUpdateJ/export', params);
}