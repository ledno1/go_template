<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form :label-width="80" :model="formValue" :rules="rules" ref="formRef">
        <n-form-item label="系统开启接单" path="basicSystemOpen">
          <n-switch size="large" v-model:value="formValue.open" @update:value="systemOpenChange"/>
        </n-form-item>
        <div>
          <n-space>
            <n-button type="primary" @click="formSubmit">保存更新</n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {onMounted, ref} from 'vue';
import {useDialog, useMessage} from 'naive-ui';
import {getConfig, updateConfig} from '@/api/sys/config';

const group = ref('businessBasic');
const show = ref(false);
const formRef: any = ref(null);
const message = useMessage();
const dialog = useDialog();

const formValue = ref({
  open: true,
});

const rules = {
  basicName: {
    required: true,
    message: '请输入网站名称',
    trigger: 'blur',
  },
};

function systemOpenChange(value) {
  dialog.warning({
    title: '提示',
    content: `您确定要${value ? "启动" : "关闭"}接单系统吗？该操作修改后立马生效，请慎重操作！`,
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      // message.success('操作成功');
      formSubmit();
    },
    onNegativeClick: () => {
      formValue.value.open = true;
    },
  });
}

function formSubmit() {
  formRef.value.validate((errors) => {
    if (!errors) {
      updateConfig({group: group.value, list: formValue.value}).then((_res) => {
        message.success('更新成功');
        load();
      });
    } else {
      message.error('验证失败，请填写完整信息');
    }
  });
}

onMounted(() => {
  load();
});

function load() {
  show.value = true;
  new Promise((_resolve, _reject) => {
    getConfig({group: group.value})
      .then((res) => {
        formValue.value = res.list;
      })
      .finally(() => {
        show.value = false;
      });
  });
}
</script>
