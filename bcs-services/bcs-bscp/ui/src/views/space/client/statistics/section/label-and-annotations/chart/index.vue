<template>
  <Teleport :disabled="!isOpenFullScreen" to="body">
    <div ref="containerRef" :class="{ fullscreen: isOpenFullScreen }">
      <Card :title="$t(`按 {n} 统计`, { n: label })" :height="368">
        <template #operation>
          <OperationBtn
            :is-open-full-screen="isOpenFullScreen"
            @refresh="emits('refresh')"
            @toggle-full-screen="isOpenFullScreen = !isOpenFullScreen" />
        </template>
        <template #head-suffix>
          <bk-tag theme="info" type="stroke" style="margin-left: 8px"> {{ $t('标签') }} </bk-tag>
          <TriggerBtn v-model:currentType="currentType" style="margin-left: 8px" />
        </template>
        <bk-loading class="loading-wrap" :loading="loading">
          <component
            :bk-biz-id="bkBizId"
            :app-id="appId"
            :is="currentComponent"
            :data="data"
            :label="label"
            @jump="jumpToSearch($event as string)" />
        </bk-loading>
      </Card>
    </div>
  </Teleport>
</template>

<script lang="ts" setup>
  import { ref, onMounted, computed, watch } from 'vue';
  import Card from '../../../components/card.vue';
  import TriggerBtn from '../../../components/trigger-btn.vue';
  import Pie from './pie.vue';
  import Column from './column.vue';
  import Table from './table.vue';
  import OperationBtn from '../../../components/operation-btn.vue';
  import { IClientLabelItem } from '../../../../../../../../types/client';
  import { useRouter } from 'vue-router';

  const router = useRouter();

  const emits = defineEmits(['refresh']);

  const props = defineProps<{
    bkBizId: string;
    appId: number;
    label: string;
    data: IClientLabelItem[];
    loading: boolean;
  }>();

  const currentType = ref('column');
  const componentMap = {
    pie: Pie,
    column: Column,
    table: Table,
  };
  const isOpenFullScreen = ref(false);
  const containerRef = ref();
  const initialWidth = ref(0);

  const currentComponent = computed(() => componentMap[currentType.value as keyof typeof componentMap]);

  onMounted(() => {
    initialWidth.value = containerRef.value.offsetWidth;
  });

  watch(
    () => isOpenFullScreen.value,
    (val) => {
      containerRef.value!.style.width = val ? '100%' : `${initialWidth.value}px`;
    },
  );

  const jumpToSearch = (value: string) => {
    const routeData = router.resolve({
      name: 'client-search',
      params: { appId: props.appId, bizId: props.bkBizId },
      query: { label: `${props.label}=${value}` },
    });
    window.open(routeData.href, '_blank');
  };
</script>

<style scoped lang="scss">
  .fullscreen {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    z-index: 5000;
    .card {
      width: 100%;
      height: 100vh !important;
      :deep(.operation-btn) {
        top: 0 !important;
      }
    }
  }
  .loading-wrap {
    height: 100%;
  }
</style>
