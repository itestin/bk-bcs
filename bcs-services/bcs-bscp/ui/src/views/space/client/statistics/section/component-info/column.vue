<template>
  <div ref="canvasRef" class="canvas-wrap">
    <Tooltip ref="tooltipRef" />
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, ref, watch } from 'vue';
  import { Column } from '@antv/g2plot';
  import Tooltip from '../../components/tooltip.vue';

  const props = defineProps<{
    data: any;
  }>();
  const canvasRef = ref<HTMLElement>();
  const tooltipRef = ref();
  let columnPlot: Column;
  const data = ref(props.data || []);

  watch(
    () => props.data,
    () => {
      data.value = props.data;
      columnPlot.changeData(data.value);
    },
  );

  onMounted(() => {
    initChart();
  });

  const initChart = () => {
    columnPlot = new Column(canvasRef.value!, {
      data: props.data,
      isStack: true,
      color: ['#3E96C2', '#61B2C2', '#85CCA8', '#B5E0AB'],
      xField: 'client_type',
      yField: 'value',
      yAxis: {
        grid: {
          line: {
            style: {
              stroke: '#979BA5',
              lineDash: [4, 5],
            },
          },
        },
      },
      seriesField: 'client_version',
      maxColumnWidth: 80,
      label: {
        // 可手动配置 label 数据标签位置
        position: 'middle', // 'top', 'bottom', 'middle'
      },
      legend: {
        position: 'bottom',
      },
    });
    columnPlot.render();
  };
</script>

<style scoped lang="scss"></style>
