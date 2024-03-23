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
  appName: number;
  appVersion: string;
  appUrl: string;
  mark: string;
  createdAt: string;
  updatedAt: string;
}

export const defaultState: State = {
  id: 0,
  appName: 0,
  appVersion: '',
  appUrl: '',
  mark: '',
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
  appName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入应用类型',
  },
  appVersion: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入应用版本',
  },
  appUrl: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入更新包地址',
  },
  mark: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入备忘',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'appName',
    component: 'NInputNumber',
    label: '应用类型',
    componentProps: {
      placeholder: '请输入应用类型',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
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
]);

export const columns = [
  {
    title: 'id',
    key: 'id',
  },
  {
    title: '应用类型',
    key: 'appName',
  },
  {
    title: '应用版本',
    key: 'appVersion',
  },
  {
    title: '备忘',
    key: 'mark',
  },
  {
    title: '创建时间',
    key: 'createdAt',
  },
  {
    title: '更新时间',
    key: 'updatedAt',
  },
];