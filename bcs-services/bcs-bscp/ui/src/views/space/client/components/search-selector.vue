<template>
  <section class="section">
    <bk-popover
      trigger="manual"
      ext-cls="search-selector"
      :is-show="isShowPopover"
      :arrow="false"
      placement="bottom-start"
      theme="light"
      @after-show="handleGetSearchList('recent')">
      <div
        class="search-wrap"
        :data-placeholder="inputPlacehoder"
        v-bk-tooltips="{ content: inputPlacehoder, disabled: locale === 'zh-cn' || !inputPlacehoder }"
        @click="isShowPopover = !isShowPopover">
        <div class="search-condition-list">
          <bk-tag
            v-for="(condition, index) in searchConditionList"
            :key="condition.key"
            style="margin-right: 6px"
            closable
            @close="handleConditionClose(index)">
            {{ condition.content }}
          </bk-tag>
        </div>
        <div class="search-container-input">
          <bk-input
            v-model="searchStr"
            ref="inputRef"
            class="input"
            placeholder=" "
            :readonly="!searchStr"
            @focus="inputFocus = true"
            @blur="handleConfirmConditionItem"
            @enter="handleConfirmConditionItem" />
        </div>
        <div
          v-if="searchConditionList.length && isClientSearch"
          :class="['set-used', { light: isCommonlyUsedBtnLight }]"
          v-bk-tooltips="{ content: t('设为常用') }"
          @click.stop="handleOpenSetCommonlyDialg(true)">
          <span class="bk-bscp-icon icon-star-fill"></span>
        </div>
      </div>
      <template #content>
        <div v-if="!showChildSelector" v-click-outside="() => (isShowPopover = false)" class="menu-wrap">
          <div class="search-condition">
            <div class="title">{{ t('查询条件') }}</div>
            <div v-for="item in selectorData" :key="item.value" class="search-item" @click="handleSelectParent(item)">
              {{ item.name }}
            </div>
          </div>
          <div class="resent-search">
            <div class="title">{{ t('最近查询') }}</div>
            <bk-loading :loading="resentSearchListLoading">
              <div
                v-for="item in recentSearchList"
                :key="item.id"
                class="search-item"
                @click="handleSelectRecentSearch(item)">
                <bk-overflow-title type="tips">{{ item.spec.search_name }}</bk-overflow-title>
              </div>
            </bk-loading>
          </div>
        </div>
        <div v-else class="children-menu-wrap" v-click-outside="() => (isShowPopover = false)">
          <div v-for="item in childSelectorData" :key="item.value" class="search-item" @click="handleSelectChild(item)">
            {{ item.name }}
          </div>
        </div>
      </template>
    </bk-popover>
    <div v-if="isClientSearch" class="commonly-wrap">
      <template v-for="(item, index) in commonlySearchList" :key="item.id">
        <CommonlyUsedTag
          v-if="index < 5"
          :commonly-search-item="item"
          @update="handleOpenSetCommonlyDialg(false, item)"
          @click="searchConditionList = cloneDeep(item.search_condition)"
          @delete="handleOpenDeleteCommonlyDialog(item)" />
      </template>
      <bk-popover
        trigger="manual"
        ext-cls="all-commonly-search-popover"
        placement="bottom-start"
        theme="light"
        :is-show="isShowAllCommonSearchPopover"
        :arrow="false">
        <bk-button theme="primary" text @click="isShowAllCommonSearchPopover = !isShowAllCommonSearchPopover">
          {{ t('全部常用查询') }}
        </bk-button>
        <template #content>
          <div
            v-for="item in commonlySearchList"
            :key="item.id"
            class="search-item"
            v-click-outside="() => (isShowAllCommonSearchPopover = false)"
            @click="handleSelectCommonSearch(item)">
            <div class="name">
              <bk-overflow-title>{{ item.spec.search_name }}</bk-overflow-title>
            </div>
            <div class="action-icon" v-if="item.spec.creator !== 'system'">
              <EditLine class="icon edit" @click.stop="handleOpenSetCommonlyDialg(false, item)" />
              <Error class="icon close" @click.stop="handleOpenDeleteCommonlyDialog(item)" />
            </div>
          </div>
        </template>
      </bk-popover>
    </div>
    <SetCommonlyDialog
      :is-show="isShowSetCommonlyDialog"
      :is-create="isCreateCommonlyUsed"
      :name="selectedCommomlyItem?.spec.search_name"
      @create="handleConfirmCreateCommonlyUsed"
      @update="handleConfirmUpdateCommonlyUsed"
      @close="isShowSetCommonlyDialog = false" />
    <bk-dialog
      :is-show="isShowDeleteCommonlyDialog"
      :ext-cls="'delete-commonly-dialog'"
      :width="400"
      @closed="isShowDeleteCommonlyDialog = false">
      <div class="head">{{ t('确认删除该常用查询?') }}</div>
      <div class="body">
        <span class="label">{{ t('名称') }} : </span>
        <span class="name">{{ selectedDeleteCommonlyItem?.spec.search_name }}</span>
      </div>
      <div class="footer">
        <div class="btns">
          <bk-button theme="danger" @click="handleConfirmDeleteCommonlyUsed">{{ t('删除') }}</bk-button>
          <bk-button @click="isShowDeleteCommonlyDialog = false">{{ t('取消') }}</bk-button>
        </div>
      </div>
    </bk-dialog>
  </section>
</template>

<script lang="ts" setup>
  import { nextTick, ref, computed, watch, onMounted, onBeforeUnmount } from 'vue';
  import { storeToRefs } from 'pinia';
  import { EditLine, Error } from 'bkui-vue/lib/icon';
  import { CLIENT_SEARCH_DATA, CLIENT_STATISTICS_SEARCH_DATA, CLIENT_STATUS_MAP } from '../../../../constants/client';
  import { ISelectorItem, ISearchCondition, ICommonlyUsedItem, IClinetCommonQuery } from '../../../../../types/client';
  import {
    getClientSearchRecord,
    createClientSearchRecord,
    updateClientSearchRecord,
    deleteClientSearchRecord,
  } from '../../../../api/client';
  import useClientStore from '../../../../store/client';
  import SetCommonlyDialog from './set-commonly-dialog.vue';
  import CommonlyUsedTag from './commonly-used-tag.vue';
  import { Message } from 'bkui-vue';
  import { cloneDeep } from 'lodash';
  import { useRoute } from 'vue-router';
  import { useI18n } from 'vue-i18n';

  const { t, locale } = useI18n();

  const clientStore = useClientStore();
  const { searchQuery } = storeToRefs(clientStore);

  const route = useRoute();

  const props = defineProps<{
    bkBizId: string;
    appId: number;
  }>();

  const isShowPopover = ref(false);
  const searchConditionList = ref<ISearchCondition[]>([]);
  const showChildSelector = ref(false);
  const childSelectorData = ref<ISelectorItem[]>();
  const searchStr = ref('');
  const inputRef = ref();
  const parentSelecte = ref<ISelectorItem>();
  const recentSearchList = ref<ICommonlyUsedItem[]>([]);
  const resentSearchListLoading = ref(false);
  const commonlySearchList = ref<ICommonlyUsedItem[]>([]);
  const inputFocus = ref(false);
  const isShowSetCommonlyDialog = ref(false);
  const isCreateCommonlyUsed = ref(true);
  const selectedCommomlyItem = ref<ICommonlyUsedItem>();
  const isShowSetCommonlyDropdown = ref(false);
  const isShowDeleteCommonlyDialog = ref(false);
  const selectedDeleteCommonlyItem = ref<ICommonlyUsedItem>();
  const isShowAllCommonSearchPopover = ref(false);

  const inputPlacehoder = computed(() => {
    if (searchConditionList.value.length || searchStr.value || inputFocus.value) return '';
    return t('UID/IP/标签/当前配置版本/目标配置版本/最近一次拉取配置状态/在线状态/客户端组件版本');
  });

  const isClientSearch = computed(() => route.name === 'client-search');

  const selectorData = computed(() => (isClientSearch.value ? CLIENT_SEARCH_DATA : CLIENT_STATISTICS_SEARCH_DATA));

  const isCommonlyUsedBtnLight = computed(() => {
    return commonlySearchList.value.some((commonlySearchItem) => {
      if (commonlySearchItem.search_condition.length !== searchConditionList.value.length) return false;
      return commonlySearchItem.search_condition.every((commonlySearchConditionList) => {
        const { key, value } = commonlySearchConditionList;
        return searchConditionList.value.findIndex((item) => item.key === key && item.value === value) > -1;
      });
    });
  });

  watch(
    () => searchConditionList.value,
    () => {
      // 搜索框和查询条件都为空时不需要转换查询参数
      if (searchConditionList.value.length === 0 && Object.keys(searchQuery.value.search!).length === 0) return;
      handleSearchConditionChangeQuery();
    },
    { deep: true },
  );

  watch(
    () => props.appId,
    () => {
      handleGetSearchList('common');
    },
  );

  watch(
    () => searchQuery.value.search,
    (val) => {
      if (Object.keys(val!).length === 0) {
        searchConditionList.value = [];
      } else {
        handleAddRecentSearch();
      }
    },
  );

  watch(
    () => isShowPopover.value,
    (val) => {
      if (val && !searchStr.value) {
        showChildSelector.value = false;
        parentSelecte.value = undefined;
      }
    },
  );

  watch(
    () => isShowPopover.value,
    (val) => {
      if (val && !searchStr.value) {
        showChildSelector.value = false;
        parentSelecte.value = undefined;
      }
    },
  );

  onMounted(() => {
    handleGetSearchList('common');
    const entries = Object.entries(route.query);
    if (entries.length === 0) return;
    const { name, value } = CLIENT_SEARCH_DATA.find((item) => item.value === entries[0][0])!;
    searchConditionList.value.push({
      content: `${name} : ${entries[0][1]}`,
      value: entries[0][1] as string,
      key: value,
    });
  });

  onBeforeUnmount(() => {
    clientStore.$patch((state) => {
      state.searchQuery.search = {};
    });
  });

  // 选择父选择器
  const handleSelectParent = (parentSelectorItem: ISelectorItem) => {
    parentSelecte.value = parentSelectorItem;
    // 如果有子选择项就展示 没有就用户手动输入
    if (parentSelectorItem?.children) {
      childSelectorData.value = parentSelectorItem.children;
      showChildSelector.value = true;
    } else {
      isShowPopover.value = false;
      nextTick(() => inputRef.value.focus());
    }
    searchStr.value = `${parentSelectorItem?.name} : `;
  };

  // 选择子选择器
  const handleSelectChild = (childrenSelectoreItem: ISelectorItem) => {
    showChildSelector.value = false;
    isShowPopover.value = false;
    // 重复的查询项去重
    const index = searchConditionList.value.findIndex(
      (item) => item.key === parentSelecte.value?.value && item.key !== 'label',
    );
    if (index > -1) handleConditionClose(index);
    searchConditionList.value.push({
      key: parentSelecte.value!.value,
      value: childrenSelectoreItem.value,
      content: `${parentSelecte.value?.name} : ${childrenSelectoreItem.name}`,
    });
    searchStr.value = '';
  };

  // 手动输入确认搜索项
  const handleConfirmConditionItem = () => {
    const conditionValue = searchStr.value.split(' : ', 2);
    inputFocus.value = false;
    if (!conditionValue[1]) {
      searchStr.value = '';
      return;
    }
    // 重复的查询项去重
    const index = searchConditionList.value.findIndex(
      (item) => item.key === parentSelecte.value?.value && item.key !== 'label',
    );
    if (index > -1) handleConditionClose(index);
    searchConditionList.value.push({
      key: parentSelecte.value!.value,
      value: conditionValue[1],
      content: `${parentSelecte.value?.name} : ${conditionValue[1]}`,
    });
    searchStr.value = '';
  };

  // 获取最近搜索记录和常用搜索记录
  const handleGetSearchList = async (search_type: string) => {
    if (!props.appId) return;
    try {
      resentSearchListLoading.value = search_type === 'recent';
      const params: IClinetCommonQuery = {
        start: 0,
        limit: 10,
        search_type,
      };
      if (search_type === 'common') {
        params.all = true;
      } else {
        isClientSearch.value ? (params.search_type = 'query') : (params.search_type = 'statistic');
      }
      const res = await getClientSearchRecord(props.bkBizId, props.appId, params);
      const searchList = res.data.details;
      searchList.forEach((item: ICommonlyUsedItem) => handleQueryChangeSearchCondition(item));
      if (search_type === 'recent') {
        recentSearchList.value = searchList;
      } else {
        commonlySearchList.value = searchList;
      }
    } catch (error) {
      console.error(error);
    } finally {
      resentSearchListLoading.value = false;
    }
  };

  // 删除查询条件
  const handleConditionClose = (index: number) => {
    searchConditionList.value.splice(index, 1);
  };

  // 添加最近查询
  const handleAddRecentSearch = async () => {
    await createClientSearchRecord(props.bkBizId, props.appId, {
      search_type: isClientSearch.value ? 'query' : 'statistic',
      search_condition: searchQuery.value.search!,
    });
  };

  // 设置常用查询
  const handleConfirmCreateCommonlyUsed = async (search_name: string) => {
    try {
      await createClientSearchRecord(props.bkBizId, props.appId, {
        search_condition: searchQuery.value.search!,
        search_type: 'common',
        search_name,
      });
      isShowSetCommonlyDialog.value = false;
      handleGetSearchList('common');
      Message({
        theme: 'success',
        message: t('常用查询添加成功'),
      });
    } catch (error) {
      console.error(error);
    }
  };

  // 更新常用查询
  const handleConfirmUpdateCommonlyUsed = async (search_name: string) => {
    try {
      await updateClientSearchRecord(props.bkBizId, props.appId, selectedCommomlyItem.value!.id, {
        search_condition: selectedCommomlyItem.value!.spec.search_condition,
        search_type: 'common',
        search_name,
      });
      isShowSetCommonlyDialog.value = false;
      handleGetSearchList('common');
      Message({
        theme: 'success',
        message: t('常用查询修改成功'),
      });
    } catch (error) {
      console.error(error);
    }
  };

  // 删除常用查询
  const handleOpenDeleteCommonlyDialog = (item: ICommonlyUsedItem) => {
    selectedDeleteCommonlyItem.value = item;
    isShowDeleteCommonlyDialog.value = true;
  };

  const handleConfirmDeleteCommonlyUsed = async () => {
    try {
      await deleteClientSearchRecord(props.bkBizId, props.appId, selectedDeleteCommonlyItem.value!.id);
      isShowDeleteCommonlyDialog.value = false;
      handleGetSearchList('common');
      Message({
        theme: 'success',
        message: t('常用查询删除成功'),
      });
    } catch (error) {
      console.error(error);
    }
  };

  const handleOpenSetCommonlyDialg = (isCreate: boolean, item?: ICommonlyUsedItem) => {
    if (isCreate) {
      if (isCommonlyUsedBtnLight.value) return;
      isCreateCommonlyUsed.value = true;
    } else {
      isCreateCommonlyUsed.value = false;
      selectedCommomlyItem.value = item;
    }
    isShowSetCommonlyDialog.value = true;
    isShowSetCommonlyDropdown.value = false;
  };

  // 查询条件转换为查询参数
  const handleSearchConditionChangeQuery = () => {
    const query: { [key: string]: any } = {};
    const label: { [key: string]: any } = {};
    searchConditionList.value.forEach((item) => {
      if (item.key === 'label') {
        const labelValue = item.value.split('=', 2);
        label[labelValue[0]] = labelValue[1] || '';
        query[item.key] = label;
      } else if (item.key === 'online_status' || item.key === 'release_change_status') {
        if (query[item.key]) {
          query[item.key].push(item.value);
        } else {
          query[item.key] = [item.value];
        }
      } else {
        query[item.key] = item.value.trim();
      }
    });
    clientStore.$patch((state) => {
      state.searchQuery.search = query;
    });
  };

  // 查询参数转换为查询条件并获取查询名
  const handleQueryChangeSearchCondition = (item: ICommonlyUsedItem) => {
    const searchList: ISearchCondition[] = [];
    const searchName: string[] = [];
    const query: { [key: string]: any } = item.spec.search_condition;
    Object.keys(query).forEach((key) => {
      if (key === 'label') {
        const labelValue = query[key];
        Object.keys(labelValue).forEach((label) => {
          const value = labelValue[label] || '';
          const content = value ? `${t('标签')}:${label}=${labelValue[label]}` : `${t('标签')}:${label}`;
          searchList.push({
            key,
            value: `${label}=${value}`,
            content,
          });
          searchName.push(content);
        });
      } else if (key === 'online_status' || key === 'release_change_status') {
        query[key].forEach((value: string) => {
          const content = `${selectorData.value.find((item) => item.value === key)?.name} : ${
            CLIENT_STATUS_MAP[value as keyof typeof CLIENT_STATUS_MAP]
          }`;
          searchList.push({
            key,
            value,
            content,
          });
          searchName.push(content);
        });
      } else {
        const content = `${selectorData.value.find((item) => item.value === key)?.name} : ${query[key]}`;
        searchList.push({
          key,
          value: query[key],
          content,
        });
        searchName.push(content);
      }
    });
    item.search_condition = searchList;
    item.spec.search_name = item.spec.search_name || searchName.join(';');
  };

  const handleSelectRecentSearch = (item: ICommonlyUsedItem) => {
    searchConditionList.value = cloneDeep(item.search_condition);
    isShowPopover.value = false;
  };

  const handleSelectCommonSearch = (item: ICommonlyUsedItem) => {
    searchConditionList.value = cloneDeep(item.search_condition);
    isShowAllCommonSearchPopover.value = false;
  };
</script>

<style scoped lang="scss">
  .section {
    position: relative;
  }
  .search-wrap {
    position: relative;
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    padding-left: 8px;
    width: 670px;
    min-height: 32px;
    background: #fff;
    border: 1px solid #c4c6cc;
    &::after {
      position: absolute;
      width: calc(100% - 16px);
      content: attr(data-placeholder);
      color: #c4c6cc;
      overflow: hidden;
      white-space: nowrap;
      text-overflow: ellipsis;
    }
    .search-container-input {
      min-width: 40px;
      .input {
        border: none;
        height: 100%;
        outline: none;
        box-shadow: none;
      }
    }
    .search-condition-list {
      display: flex;
      align-items: center;
      flex-wrap: wrap;
    }
    .set-used {
      position: absolute;
      right: 8px;
      display: flex;
      align-items: center;
      justify-content: center;
      width: 24px;
      height: 24px;
      background: #f0f1f5;
      border-radius: 2px;
      color: #c4c6cc;
      font-size: 14px;
      &.light span {
        color: #ff9c01;
      }
    }
  }
  .menu-wrap {
    display: flex;
    justify-content: space-between;
    width: calc(670px - 16px);
    padding: 8px;
    .title {
      width: 319px;
      height: 24px;
      background: #eaebf0;
      border-radius: 2px;
      padding-left: 8px;
      color: #313238;
      line-height: 24px;
      margin-bottom: 8px;
    }
    .search-item {
      width: 319px;
      height: 32px;
      padding-left: 12px;
      line-height: 32px;
      &:hover {
        background: #f5f7fa;
      }
    }
  }
  .children-menu-wrap {
    min-width: 200px;
    padding: 4px 0;
    div {
      display: flex;
      align-items: center;
      height: 32px;
      padding: 0 8px;
      color: #63656e;
      cursor: pointer;
      &:hover {
        background: #f5f7fa;
      }
    }
  }
  .commonly-wrap {
    position: absolute;
    height: 26px;
    display: flex;
    align-items: center;
    margin-top: 6px;
  }
  .action-item {
    width: 58px;
    height: 32px;
    line-height: 32px;
    text-align: center;
    color: #63656e;
    cursor: pointer;
    &:hover {
      background: #f5f7fa;
    }
  }
</style>

<style lang="scss">
  .bk-popover.bk-pop2-content.search-selector {
    padding: 0;
  }
  .commonly-search-item-popover.bk-popover.bk-pop2-content {
    padding: 4px 0;
  }
  .all-commonly-search-popover.bk-popover.bk-pop2-content {
    padding: 4px 0;
    max-height: 168px;
    overflow-y: auto;
    .search-item {
      display: flex;
      align-items: center;
      justify-content: space-between;
      padding: 0 8px 0 12px;
      width: 238px;
      height: 32px;
      cursor: pointer;
      color: #63656e;
      &:hover {
        background-color: #f5f7fa;
      }
      .name {
        max-width: 120px;
      }
      .action-icon {
        display: flex;
        align-items: center;
        .icon:hover {
          color: #3a84ff;
        }
        .edit {
          margin-right: 17px;
        }
        .close {
          font-size: 14px;
        }
      }
    }
  }

  .delete-commonly-dialog .bk-modal-body {
    .bk-modal-header {
      display: none;
    }
    .bk-modal-footer {
      display: none;
    }
    .bk-modal-content {
      padding: 48px 24px 0 24px;
      .head {
        font-size: 20px;
        color: #313238;
        text-align: center;
      }
      .body {
        min-height: 32px;
        line-height: 32px;
        background: #f5f7fa;
        padding-left: 16px;
        margin-top: 16px;
        .label {
          color: #63656e;
        }
        .name {
          color: #313238;
        }
      }
      .footer {
        display: flex;
        justify-content: space-around;
        .btns {
          margin-top: 24px;
          .bk-button {
            width: 88px;
          }
          .bk-button:nth-child(1) {
            margin-right: 8px;
          }
        }
      }
    }
  }
</style>
