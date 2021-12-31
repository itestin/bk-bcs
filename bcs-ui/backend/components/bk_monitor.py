# -*- coding: utf-8 -*-
"""
Tencent is pleased to support the open source community by making 蓝鲸智云PaaS平台社区版 (BlueKing PaaS Community
Edition) available.
Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://opensource.org/licenses/MIT

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.

蓝鲸监控接口封装
"""
import logging
import time

from django.conf import settings

from backend.components.utils import http_get, http_post
from backend.utils.basic import normalize_metric

logger = logging.getLogger(__name__)

BK_MONITOR_QUERY_HOST = settings.BK_MONITOR_QUERY_HOST

# 磁盘统计 允许的文件系统
DISK_FSTYPE = "ext[234]|btrfs|xfs|zfs"
# 磁盘统计 允许的挂载目录
DISK_MOUNTPOINT = "/data"


def query_range(query, start, end, step, project_id=None):
    """范围请求API"""
    url = f'{BK_MONITOR_QUERY_HOST}/query/ts/promql'
    data = {"promql": query, "start": str(int(start)), "end": str(int(end)), "step": f"{step}s"}
    logger.info("prometheus query_range: %s", data)
    bkmonitor_resp = http_post(url, json=data, timeout=120, raise_exception=False)
    prom_resp = bkmonitor_resp2prom(bkmonitor_resp)
    return prom_resp


def query(_query, timestamp=None, project_id=None):
    """查询API"""
    end = time.time()
    # 蓝鲸监控没有实时数据接口, 这里的方案是向前追溯5分钟, 取最新的一个点
    start = end - 300
    url = f'{BK_MONITOR_QUERY_HOST}/query/ts/promql'
    data = {"promql": _query, "start": str(int(start)), "end": str(int(end)), "step": "60s"}
    logger.info("prometheus query: %s", data)
    bkmonitor_resp = http_post(url, json=data, timeout=120, raise_exception=False)
    logger.info("prometheus query_range: %s", bkmonitor_resp)
    prom_resp = bkmonitor_resp2prom(bkmonitor_resp)
    return prom_resp


def bkmonitor_resp2prom(response):
    """蓝鲸监控数据返回转换为prom返回"""
    data = {'resultType': 'matrix', 'result': []}
    series_list = response.get('series') or []
    for series in series_list:
        metric = dict(zip(series['group_keys'], series['group_values']))
        values = []

        # 蓝鲸监控返回的values可能会变化, 通过 columes 字段顺序判断
        if series["columns"][0] == "_value":
            value_index = 0
            timestamp_index = 1
        else:
            value_index = 1
            timestamp_index = 0
        for value in series["values"]:
            values.append((value[timestamp_index], str(value[value_index])))

        data['result'].append({'metric': metric, 'values': values})
    prom_resp = {'data': data}
    return prom_resp


def get_first_value(prom_resp, fill_zero=True):
    """获取返回的第一个值"""
    data = prom_resp.get("data") or {}
    result = data.get("result") or []
    if not result:
        if fill_zero:
            # 返回0字符串, 和promtheus保存一致
            return "0"
        return None

    values = result[0]["values"]
    if not values:
        if fill_zero:
            return "0"
        return None

    # 取最后一个值, 转换为prom字符串格式
    last_value = values[-1]

    return last_value[1]


def get_targets(project_id, cluster_id, dedup=True):
    """获取集群的targets"""
    resp = []
    return resp


def get_cluster_cpu_usage(cluster_id, node_ip_list, bk_biz_id=None):
    """获取集群nodeCPU使用率"""
    node_ip_list = "|".join(node_ip_list)

    cpu_used_prom_query = f"""
        sum(bkmonitor:system:cpu_detail:usage{{bk_biz_id="{bk_biz_id}", ip=~"{node_ip_list}"}}) / 100
    """  # noqa

    cpu_count_prom_query = f"""
        count(bkmonitor:system:cpu_detail:usage{{bk_biz_id="{bk_biz_id}", ip=~"{node_ip_list}"}})
    """  # noqa

    data = {"used": get_first_value(query(cpu_used_prom_query)), "total": get_first_value(query(cpu_count_prom_query))}
    return data


def get_cluster_cpu_usage_range(cluster_id, node_ip_list, bk_biz_id=None):
    """获取集群nodeCPU使用率"""
    end = time.time()
    start = end - 3600
    step = 60

    node_ip_list = "|".join(node_ip_list)
    prom_query = f"""
        sum(bkmonitor:system:cpu_detail:usage{{bk_biz_id="{bk_biz_id}", ip=~"{node_ip_list}"}}) /
        count(bkmonitor:system:cpu_detail:usage{{bk_biz_id="{bk_biz_id}", ip=~"{node_ip_list}"}})"""  # noqa

    resp = query_range(prom_query, start, end, step)
    return resp.get("data") or {}


def get_cluster_memory_usage(cluster_id, node_ip_list, bk_biz_id=None):
    """获取集群nodeCPU使用率"""

    node_ip_list = "|".join(node_ip_list)

    memory_total_prom_query = f"""
        sum(bkmonitor:system:mem:total{{bk_biz_id="{bk_biz_id}", ip=~"{node_ip_list}"}})
    """

    memory_used_prom_query = f"""
        sum(bkmonitor:system:mem:used{{bk_biz_id="{bk_biz_id}", ip=~"{node_ip_list}"}})
    """  # noqa

    data = {
        "used_bytes": get_first_value(query(memory_used_prom_query)),
        "total_bytes": get_first_value(query(memory_total_prom_query)),
    }
    return data


def get_cluster_memory_usage_range(cluster_id, node_ip_list, bk_biz_id=None):
    """获取集群nodeCPU使用率"""
    end = time.time()
    start = end - 3600
    step = 60

    node_ip_list = "|".join(node_ip_list)
    prom_query = f"""
        (sum(bkmonitor:system:mem:used{{bk_biz_id="{bk_biz_id}", ip=~"{node_ip_list}"}}) /
        sum(bkmonitor:system:mem:total{{bk_biz_id="{bk_biz_id}", ip=~"{node_ip_list}"}})) *
        100
    """  # noqa

    resp = query_range(prom_query, start, end, step)
    return resp.get("data") or {}


def get_cluster_disk_usage(cluster_id, node_ip_list, bk_biz_id=None):
    """获取集群nodeCPU使用率"""
    node_ip_list = "|".join(node_ip_list)

    disk_total_prom_query = f"""
        sum(bkmonitor:system:disk:total{{bk_biz_id="{bk_biz_id}", mount_point=~"{ DISK_MOUNTPOINT }", ip=~"{node_ip_list}"}})
    """  # noqa

    disk_used_prom_query = f"""
        sum(bkmonitor:system:disk:used{{bk_biz_id="{bk_biz_id}", mount_point=~"{ DISK_MOUNTPOINT }", ip=~"{node_ip_list}"}})
    """  # noqa

    data = {
        "used_bytes": get_first_value(query(disk_used_prom_query)),
        "total_bytes": get_first_value(query(disk_total_prom_query)),
    }
    return data


def get_cluster_disk_usage_range(cluster_id, node_ip_list, bk_biz_id=None):
    """获取k8s集群磁盘使用率"""
    end = time.time()
    start = end - 3600
    step = 60

    node_ip_list = "|".join(node_ip_list)

    prom_query = f"""
        sum(bkmonitor:system:disk:used{{bk_biz_id="{bk_biz_id}", mount_point=~"{ DISK_MOUNTPOINT }", ip=~"{node_ip_list}"}}) /
        sum(bkmonitor:system:disk:total{{bk_biz_id="{bk_biz_id}", mount_point=~"{ DISK_MOUNTPOINT }", ip=~"{node_ip_list}"}})
    """

    resp = query_range(prom_query, start, end, step)
    return resp.get("data") or {}


def get_node_info(cluster_id, ip, bk_biz_id=None):
    prom_query = f"""
        cadvisor_version_info{{cluster_id="{cluster_id}", instance=~"{ip}:\\\\d+"}} or
        node_uname_info{{cluster_id="{cluster_id}", job="node-exporter", instance=~"{ip}:\\\\d+"}} or
        label_replace(sum by (instance) (count without(cpu, mode) (node_cpu_seconds_total{{cluster_id="{cluster_id}", job="node-exporter", mode="idle", instance=~"{ip}:\\\\d+"}})), "metric_name", "cpu_count", "instance", ".*") or
        label_replace(sum by (instance) (node_memory_MemTotal_bytes{{cluster_id="{cluster_id}", job="node-exporter", instance=~"{ip}:\\\\d+"}}), "metric_name", "memory", "instance", ".*") or
        label_replace(sum by (instance) (node_filesystem_size_bytes{{cluster_id="{cluster_id}", job="node-exporter", instance=~"{ip}:\\\\d+", fstype=~"{ DISK_FSTYPE }", mountpoint=~"{ DISK_MOUNTPOINT }"}}), "metric_name", "disk", "instance", ".*")
    """  # noqa

    resp = query(prom_query)
    return resp.get("data") or {}


def get_container_pod_count(cluster_id, ip, bk_biz_id=None):
    """获取K8S节点容器/Pod数量"""
    prom_query = f"""
        label_replace(sum by (instance) ({{__name__="kubelet_running_container_count", cluster_id="{cluster_id}", instance=~"{ip}:\\\\d+"}}), "metric_name", "container_count", "instance", ".*") or
        label_replace(sum by (instance) ({{__name__="kubelet_running_pod_count", cluster_id="{cluster_id}", instance=~"{ip}:\\\\d+"}}), "metric_name", "pod_count", "instance", ".*")
    """  # noqa
    resp = query(prom_query)
    return resp.get("data") or {}


def get_node_cpu_usage(cluster_id, ip, bk_biz_id=None):
    """获取CPU总使用率"""
    bk_biz_id = "2"
    prom_query = f"""
        sum(bkmonitor:system:cpu_detail:usage{{bk_biz_id="{bk_biz_id}", ip="{ip}"}}) /
        count(bkmonitor:system:cpu_detail:usage{{bk_biz_id="{bk_biz_id}", ip="{ip}"}})"""  # noqa

    resp = query(prom_query)
    value = get_first_value(resp)
    return value


def get_node_cpu_usage_range(cluster_id, ip, start, end, bk_biz_id=None):
    """获取CPU总使用率
    start, end单位为毫秒，和数据平台保持一致
    """
    step = (end - start) // 60

    prom_query = f"""
        sum(bkmonitor:system:cpu_detail:usage{{bk_biz_id="{bk_biz_id}", ip="{ip}"}}) /
        count(bkmonitor:system:cpu_detail:usage{{bk_biz_id="{bk_biz_id}", ip="{ip}"}})"""  # noqa

    resp = query_range(prom_query, start, end, step)
    return resp.get("data") or {}


def get_node_memory_usage(cluster_id, ip, bk_biz_id=None):
    """获取节点内存使用率"""
    prom_query = f"""
        (sum(bkmonitor:system:mem:used{{bk_biz_id="{bk_biz_id}", ip="{ip}"}}) /
        sum(bkmonitor:system:mem:total{{bk_biz_id="{bk_biz_id}", ip="{ip}"}})) *
        100
    """  # noqa

    resp = query(prom_query)
    value = get_first_value(resp)
    return value


def get_node_memory_usage_range(cluster_id, ip, start, end, bk_biz_id=None):
    """获取CPU总使用率
    start, end单位为毫秒，和数据平台保持一致
    """
    step = (end - start) // 60
    prom_query = f"""
        (sum(bkmonitor:system:mem:used{{bk_biz_id="{bk_biz_id}", ip="{ip}"}}) /
        sum(bkmonitor:system:mem:total{{bk_biz_id="{bk_biz_id}", ip="{ip}"}})) *
        100
    """  # noqa

    resp = query_range(prom_query, start, end, step)
    return resp.get("data") or {}


def get_node_disk_usage(cluster_id, ip, bk_biz_id=None):
    prom_query = f"""
        (sum(bkmonitor:system:disk:used{{bk_biz_id="{bk_biz_id}", mount_point=~"{ DISK_MOUNTPOINT }", ip="{ip}"}}) /
        sum(bkmonitor:system:disk:total{{bk_biz_id="{bk_biz_id}", mount_point=~"{ DISK_MOUNTPOINT }", ip="{ip}"}})) *
        100
    """  # noqa

    value = get_first_value(query(prom_query))
    return value


def get_node_network_receive(cluster_id, ip, start, end, bk_biz_id=None):
    """获取网络数据
    start, end单位为毫秒，和数据平台保持一致
    数据单位KB/s
    """
    step = (end - start) // 60
    prom_query = f"""
        max(bkmonitor:system:net:speed_recv{{bk_biz_id="{bk_biz_id}", ip="{ ip }"}})
    """  # noqa
    resp = query_range(prom_query, start, end, step)
    return resp.get("data") or {}


def get_node_network_transmit(cluster_id, ip, start, end, bk_biz_id=None):
    step = (end - start) // 60
    prom_query = f"""
        max(bkmonitor:system:net:speed_send{{bk_biz_id="{bk_biz_id}", ip="{ ip }"}})
        """  # noqa
    resp = query_range(prom_query, start, end, step)
    return resp.get("data") or {}


def get_node_diskio_usage(cluster_id, ip, bk_biz_id=None):
    """获取当前磁盘IO"""
    prom_query = f"""
        max(bkmonitor:system:io:util{{bk_biz_id="{bk_biz_id}", ip="{ip}"}}) * 100
    """  # noqa

    value = get_first_value(query(prom_query))
    return value


def get_node_diskio_usage_range(cluster_id, ip, start, end, bk_biz_id=None):
    """获取磁盘IO数据
    start, end单位为毫秒，和数据平台保持一致
    数据单位KB/s
    """
    step = (end - start) // 60
    prom_query = f"""
        max(bkmonitor:system:io:util{{bk_biz_id="{bk_biz_id}", ip="{ip}"}}) * 100
    """  # noqa

    resp = query_range(prom_query, start, end, step, bk_biz_id=None)
    return resp.get("data") or {}


def get_pod_cpu_usage_range(cluster_id, namespace, pod_name_list, start, end, bk_biz_id=None):
    """获取CPU总使用率
    start, end单位为毫秒，和数据平台保持一致
    """
    step = (end - start) // 60
    pod_name_list = "|".join(pod_name_list)

    porm_query = f"""
        sum by (pod_name) (rate(container_cpu_usage_seconds_total{{cluster_id="{cluster_id}", bk_biz_id="{bk_biz_id}", namespace=~"{ namespace }",
        pod_name=~"{ pod_name_list }", container_name!="", container_name!="POD"}}[2m])) * 100
        """  # noqa
    resp = query_range(porm_query, start, end, step)

    return resp.get("data") or {}


def get_pod_memory_usage_range(cluster_id, namespace, pod_name_list, start, end, bk_biz_id=None):
    """获取CPU总使用率
    start, end单位为毫秒，和数据平台保持一致
    """
    step = (end - start) // 60
    pod_name_list = "|".join(pod_name_list)

    porm_query = f"""
        sum by (pod_name) (container_memory_rss{{cluster_id="{cluster_id}", namespace=~"{ namespace }", pod_name=~"{ pod_name_list }",
        container_name!="", container_name!="POD"}})
        """  # noqa
    resp = query_range(porm_query, start, end, step)

    return resp.get("data") or {}


def get_pod_network_receive(cluster_id, namespace, pod_name_list, start, end, bk_biz_id=None):
    """获取网络数据
    start, end单位为毫秒，和数据平台保持一致
    """
    step = (end - start) // 60
    pod_name_list = "|".join(pod_name_list)

    prom_query = f"""
        sum by(pod_name) (rate(container_network_receive_bytes_total{{cluster_id="{cluster_id}", namespace=~"{ namespace }", pod_name=~"{ pod_name_list }"}}[2m]))
        """  # noqa

    resp = query_range(prom_query, start, end, step)
    return resp.get("data") or {}


def get_pod_network_transmit(cluster_id, namespace, pod_name_list, start, end, bk_biz_id=None):
    step = (end - start) // 60
    pod_name_list = "|".join(pod_name_list)

    prom_query = f"""
        sum by(pod_name) (rate(container_network_transmit_bytes_total{{cluster_id="{cluster_id}",  namespace=~"{ namespace }", pod_name=~"{ pod_name_list }"}}[2m]))
        """  # noqa

    resp = query_range(prom_query, start, end, step)
    return resp.get("data") or {}


def get_container_cpu_usage_range(cluster_id, namespace, pod_name, container_id_list, start, end, bk_biz_id=None):
    """获取CPU总使用率
    start, end单位为毫秒，和数据平台保持一致
    """
    step = (end - start) // 60
    container_id_list = "|".join(f".*{i}.*" for i in container_id_list)

    prom_query = f"""
        sum by(container_name) (rate(container_cpu_usage_seconds_total{{cluster_id="{cluster_id}", namespace=~"{ namespace }", pod_name=~"{pod_name}",
        container_name!="", container_name!="POD", BcsNetworkContainer!="true", id=~"{ container_id_list }"}}[2m])) * 100
        """  # noqa

    resp = query_range(prom_query, start, end, step)
    return resp.get("data") or {}


def get_container_cpu_limit(cluster_id, namespace, pod_name, container_id_list, bk_biz_id=None):
    """获取CPU总使用率
    start, end单位为毫秒，和数据平台保持一致
    """

    container_id_list = "|".join(f".*{i}.*" for i in container_id_list)

    prom_query = f"""
        max by(container_name) (container_spec_cpu_quota{{cluster_id="{cluster_id}", namespace=~"{ namespace }", pod_name=~"{pod_name}",
        container_name!="", container_name!="POD", BcsNetworkContainer!="true", id=~"{ container_id_list }"}})
        """  # noqa

    resp = query(prom_query)
    return resp.get("data") or {}


def get_container_memory_usage_range(cluster_id, namespace, pod_name, container_id_list, start, end, bk_biz_id=None):
    """获取CPU总使用率
    start, end单位为毫秒，和数据平台保持一致
    """
    step = (end - start) // 60
    container_id_list = "|".join(f".*{i}.*" for i in container_id_list)

    prom_query = f"""
        sum by(container_name) (container_memory_rss{{cluster_id="{cluster_id}", namespace=~"{ namespace }",pod_name=~"{pod_name}",
        container_name!="", container_name!="POD", BcsNetworkContainer!="true", id=~"{ container_id_list }"}})
        """  # noqa

    resp = query_range(prom_query, start, end, step)
    return resp.get("data") or {}


def get_container_memory_limit(cluster_id, namespace, pod_name, container_id_list, bk_biz_id=None):
    """获取CPU总使用率
    start, end单位为毫秒，和数据平台保持一致
    """

    container_id_list = "|".join(f".*{i}.*" for i in container_id_list)

    prom_query = f"""
        max by(container_name) (container_spec_memory_limit_bytes{{cluster_id="{cluster_id}", namespace=~"{ namespace }", pod_name=~"{pod_name}",
        container_name!="", container_name!="POD", BcsNetworkContainer!="true", id=~"{ container_id_list }"}}) > 0
        """  # noqa

    resp = query(prom_query)
    return resp.get("data") or {}


def get_container_disk_read(cluster_id, namespace, pod_name, container_id_list, start, end, bk_biz_id=None):
    step = (end - start) // 60
    container_id_list = "|".join(f".*{i}.*" for i in container_id_list)

    prom_query = f"""
        sum by(container_name) (container_fs_reads_bytes_total{{cluster_id="{cluster_id}", namespace=~"{ namespace }", pod_name=~"{pod_name}",
        container_name!="", container_name!="POD", BcsNetworkContainer!="true", id=~"{container_id_list}"}})
        """  # noqa

    resp = query_range(prom_query, start, end, step)
    return resp.get("data") or {}


def get_container_disk_write(cluster_id, namespace, pod_name, container_id_list, start, end, bk_biz_id=None):
    step = (end - start) // 60
    container_id_list = "|".join(f".*{i}.*" for i in container_id_list)

    prom_query = f"""
        sum by(container_name) (container_fs_writes_bytes_total{{cluster_id="{cluster_id}", namespace=~"{ namespace }", pod_name=~"{pod_name}",
        container_name!="", container_name!="POD", BcsNetworkContainer!="true", id=~"{container_id_list}"}})
        """  # noqa

    resp = query_range(prom_query, start, end, step)
    return resp.get("data") or {}


def mesos_agent_memory_usage(cluster_id, ip, bk_biz_id=None):
    """mesos内存使用率"""
    data = {"total": "0", "remain": "0"}
    return data


def mesos_agent_cpu_usage(cluster_id, ip, bk_biz_id=None):
    """mesosCPU使用率"""
    data = {"total": "0", "remain": "0"}
    return data


def mesos_agent_ip_remain_count(cluster_id, ip, bk_biz_id=None):
    """mesos 剩余IP数量"""
    value = 0
    return value


def mesos_cluster_cpu_usage(cluster_id, node_list, bk_biz_id=None):
    """mesos集群CPU使用率"""
    data = {"total": "0", "remain": "0"}
    return data


def mesos_cluster_memory_usage(cluster_id, node_list, bk_biz_id=None):
    """mesos集群mem使用率"""
    data = {"total": "0", "remain": "0"}
    return data


def mesos_cluster_cpu_resource_remain_range(cluster_id, start, end, bk_biz_id=None):
    """mesos集群CPU剩余量, 单位核"""
    data = {}
    return data


def mesos_cluster_cpu_resource_total_range(cluster_id, start, end, bk_biz_id=None):
    """mesos集群CPU总量, 单位核"""
    data = {}
    return data


def mesos_cluster_memory_resource_remain_range(cluster_id, start, end, bk_biz_id=None):
    """mesos集群内存剩余量, 单位MB"""
    data = {}
    return data


def mesos_cluster_memory_resource_total_range(cluster_id, start, end, bk_biz_id=None):
    """mesos集群内存总量, 单位MB"""
    data = {}
    return data


def mesos_cluster_cpu_resource_used_range(cluster_id, start, end, bk_biz_id=None):
    """mesos集群使用的CPU, 单位核"""
    data = {}
    return data


def mesos_cluster_memory_resource_used_range(cluster_id, start, end, bk_biz_id=None):
    """mesos集群使用的内存, 单位MB"""
    data = {}
    return data
