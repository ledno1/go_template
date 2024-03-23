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

export const options = ref<Options>({
  app_type: [],
});

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
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'appName',
    component: 'NSelect',
    label: '应用类型',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择应用类型',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'appVersion',
    component: 'NInput',
    label: '应用版本',
    componentProps: {
      placeholder: '请输入应用版本',
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
    render(row) {
      if (isNullObject(row.appName)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.app_type, row.appName),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.app_type, row.appName),
        }
      );
    },
  },
  {
    title: '应用版本',
    key: 'appVersion',
  },
  {
    title: '更新包地址',
    key: 'appUrl',
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

async function loadOptions() {
  options.value = await Dicts({
    types: [
      'app_type',
   ],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'appName':
        item.componentProps.options = options.value.app_type;
        break;
     }
  }
}

await loadOptions();