import { h, ref } from 'vue';
import { NAvatar, NImage, NTag, NSwitch, NRate } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';

import { isArray, isNullObject } from '@/utils/is';
import { getFileExt } from '@/utils/urlUtils';
import { defRangeShortcuts, defShortcuts, formatToDate } from '@/utils/dateUtil';
import { validate } from '@/utils/validateUtil';
import { getOptionLabel, getOptionTag, Options, errorImg } from '@/utils/hotgo';


export interface State {
  id: number;
  uid: number;
  haxKey: string;
  authCount: number;
  authSuccessCount: number;
  createdAt: string;
  updatedAt: string;
}

export const defaultState: State = {
  id: 0,
  uid: 0,
  haxKey: '',
  authCount: 0,
  authSuccessCount: 0,
  createdAt: '',
  updatedAt: '',
};

export function newState(state: State | null): State {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(defaultState);
}


export const rules = {
  uid: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入管理员ID',
  },
  haxKey: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入密钥',
  },
  authCount: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入授权总数',
  },
  authSuccessCount: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入授权成功',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'adminMemberRealName',
    component: 'NInput',
    label: '真实姓名',
    componentProps: {
      placeholder: '请输入真实姓名',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'adminMemberUsername',
    component: 'NInput',
    label: '帐号',
    componentProps: {
      placeholder: '请输入帐号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

export const columns = [
  {
    title: 'id',
    key: 'id',
  },
  {
    title: '管理员ID',
    key: 'uid',
  },
  {
    title: '密钥',
    key: 'haxKey',
  },
  {
    title: '授权总数',
    key: 'authCount',
  },
  {
    title: '授权成功',
    key: 'authSuccessCount',
  },
  {
    title: '创建时间',
    key: 'createdAt',
  },
  {
    title: '更新时间',
    key: 'updatedAt',
  },
  {
    title: '真实姓名',
    key: 'adminMemberRealName',
  },
  {
    title: '帐号',
    key: 'adminMemberUsername',
  },
];