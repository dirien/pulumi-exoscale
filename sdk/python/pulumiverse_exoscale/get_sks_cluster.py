# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetSksClusterResult',
    'AwaitableGetSksClusterResult',
    'get_sks_cluster',
    'get_sks_cluster_output',
]

@pulumi.output_type
class GetSksClusterResult:
    """
    A collection of values returned by getSksCluster.
    """
    def __init__(__self__, addons=None, aggregation_ca=None, auto_upgrade=None, cni=None, control_plane_ca=None, created_at=None, description=None, endpoint=None, exoscale_ccm=None, exoscale_csi=None, id=None, kubelet_ca=None, labels=None, metrics_server=None, name=None, nodepools=None, oidc=None, service_level=None, state=None, version=None, zone=None):
        if addons and not isinstance(addons, list):
            raise TypeError("Expected argument 'addons' to be a list")
        pulumi.set(__self__, "addons", addons)
        if aggregation_ca and not isinstance(aggregation_ca, str):
            raise TypeError("Expected argument 'aggregation_ca' to be a str")
        pulumi.set(__self__, "aggregation_ca", aggregation_ca)
        if auto_upgrade and not isinstance(auto_upgrade, bool):
            raise TypeError("Expected argument 'auto_upgrade' to be a bool")
        pulumi.set(__self__, "auto_upgrade", auto_upgrade)
        if cni and not isinstance(cni, str):
            raise TypeError("Expected argument 'cni' to be a str")
        pulumi.set(__self__, "cni", cni)
        if control_plane_ca and not isinstance(control_plane_ca, str):
            raise TypeError("Expected argument 'control_plane_ca' to be a str")
        pulumi.set(__self__, "control_plane_ca", control_plane_ca)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if exoscale_ccm and not isinstance(exoscale_ccm, bool):
            raise TypeError("Expected argument 'exoscale_ccm' to be a bool")
        pulumi.set(__self__, "exoscale_ccm", exoscale_ccm)
        if exoscale_csi and not isinstance(exoscale_csi, bool):
            raise TypeError("Expected argument 'exoscale_csi' to be a bool")
        pulumi.set(__self__, "exoscale_csi", exoscale_csi)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kubelet_ca and not isinstance(kubelet_ca, str):
            raise TypeError("Expected argument 'kubelet_ca' to be a str")
        pulumi.set(__self__, "kubelet_ca", kubelet_ca)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if metrics_server and not isinstance(metrics_server, bool):
            raise TypeError("Expected argument 'metrics_server' to be a bool")
        pulumi.set(__self__, "metrics_server", metrics_server)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if nodepools and not isinstance(nodepools, list):
            raise TypeError("Expected argument 'nodepools' to be a list")
        pulumi.set(__self__, "nodepools", nodepools)
        if oidc and not isinstance(oidc, dict):
            raise TypeError("Expected argument 'oidc' to be a dict")
        pulumi.set(__self__, "oidc", oidc)
        if service_level and not isinstance(service_level, str):
            raise TypeError("Expected argument 'service_level' to be a str")
        pulumi.set(__self__, "service_level", service_level)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    @_utilities.deprecated("""This attribute has been replaced by `exoscale_ccm`/`metrics_server` attributes, it will be removed in a future release.""")
    def addons(self) -> Sequence[str]:
        return pulumi.get(self, "addons")

    @property
    @pulumi.getter(name="aggregationCa")
    def aggregation_ca(self) -> str:
        """
        The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
        """
        return pulumi.get(self, "aggregation_ca")

    @property
    @pulumi.getter(name="autoUpgrade")
    def auto_upgrade(self) -> Optional[bool]:
        """
        Enable automatic upgrading of the control plane version.
        """
        return pulumi.get(self, "auto_upgrade")

    @property
    @pulumi.getter
    def cni(self) -> Optional[str]:
        """
        The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
        """
        return pulumi.get(self, "cni")

    @property
    @pulumi.getter(name="controlPlaneCa")
    def control_plane_ca(self) -> str:
        """
        The CA certificate (in PEM format) for TLS communications between control plane components.
        """
        return pulumi.get(self, "control_plane_ca")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The cluster creation date.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A free-form text describing the cluster.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        """
        The cluster API endpoint.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="exoscaleCcm")
    def exoscale_ccm(self) -> Optional[bool]:
        """
        Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
        """
        return pulumi.get(self, "exoscale_ccm")

    @property
    @pulumi.getter(name="exoscaleCsi")
    def exoscale_csi(self) -> Optional[bool]:
        """
        Deploy the Exoscale [Container Storage Interface](https://github.com/exoscale/exoscale-csi-driver/) on worker nodes (boolean; default: `false`; requires the CCM to be enabled).
        """
        return pulumi.get(self, "exoscale_csi")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="kubeletCa")
    def kubelet_ca(self) -> str:
        """
        The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
        """
        return pulumi.get(self, "kubelet_ca")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Mapping[str, str]]:
        """
        A map of key/value labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="metricsServer")
    def metrics_server(self) -> Optional[bool]:
        """
        Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
        """
        return pulumi.get(self, "metrics_server")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def nodepools(self) -> Sequence[str]:
        """
        The list of exoscale*sks*nodepool (IDs) attached to the cluster.
        """
        return pulumi.get(self, "nodepools")

    @property
    @pulumi.getter
    def oidc(self) -> 'outputs.GetSksClusterOidcResult':
        """
        An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
        """
        return pulumi.get(self, "oidc")

    @property
    @pulumi.getter(name="serviceLevel")
    def service_level(self) -> Optional[str]:
        """
        The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
        """
        return pulumi.get(self, "service_level")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The cluster state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetSksClusterResult(GetSksClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSksClusterResult(
            addons=self.addons,
            aggregation_ca=self.aggregation_ca,
            auto_upgrade=self.auto_upgrade,
            cni=self.cni,
            control_plane_ca=self.control_plane_ca,
            created_at=self.created_at,
            description=self.description,
            endpoint=self.endpoint,
            exoscale_ccm=self.exoscale_ccm,
            exoscale_csi=self.exoscale_csi,
            id=self.id,
            kubelet_ca=self.kubelet_ca,
            labels=self.labels,
            metrics_server=self.metrics_server,
            name=self.name,
            nodepools=self.nodepools,
            oidc=self.oidc,
            service_level=self.service_level,
            state=self.state,
            version=self.version,
            zone=self.zone)


def get_sks_cluster(addons: Optional[Sequence[str]] = None,
                    aggregation_ca: Optional[str] = None,
                    auto_upgrade: Optional[bool] = None,
                    cni: Optional[str] = None,
                    control_plane_ca: Optional[str] = None,
                    created_at: Optional[str] = None,
                    description: Optional[str] = None,
                    endpoint: Optional[str] = None,
                    exoscale_ccm: Optional[bool] = None,
                    exoscale_csi: Optional[bool] = None,
                    id: Optional[str] = None,
                    kubelet_ca: Optional[str] = None,
                    labels: Optional[Mapping[str, str]] = None,
                    metrics_server: Optional[bool] = None,
                    name: Optional[str] = None,
                    nodepools: Optional[Sequence[str]] = None,
                    oidc: Optional[Union['GetSksClusterOidcArgs', 'GetSksClusterOidcArgsDict']] = None,
                    service_level: Optional[str] = None,
                    state: Optional[str] = None,
                    version: Optional[str] = None,
                    zone: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSksClusterResult:
    """
    Use this data source to access information about an existing resource.

    :param str aggregation_ca: The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
    :param bool auto_upgrade: Enable automatic upgrading of the control plane version.
    :param str cni: The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
    :param str control_plane_ca: The CA certificate (in PEM format) for TLS communications between control plane components.
    :param str created_at: The cluster creation date.
    :param str description: A free-form text describing the cluster.
    :param str endpoint: The cluster API endpoint.
    :param bool exoscale_ccm: Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
    :param bool exoscale_csi: Deploy the Exoscale [Container Storage Interface](https://github.com/exoscale/exoscale-csi-driver/) on worker nodes (boolean; default: `false`; requires the CCM to be enabled).
    :param str id: The ID of this resource.
    :param str kubelet_ca: The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
    :param Mapping[str, str] labels: A map of key/value labels.
    :param bool metrics_server: Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
    :param Sequence[str] nodepools: The list of exoscale*sks*nodepool (IDs) attached to the cluster.
    :param Union['GetSksClusterOidcArgs', 'GetSksClusterOidcArgsDict'] oidc: An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
    :param str service_level: The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
    :param str state: The cluster state.
    :param str version: The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
    """
    __args__ = dict()
    __args__['addons'] = addons
    __args__['aggregationCa'] = aggregation_ca
    __args__['autoUpgrade'] = auto_upgrade
    __args__['cni'] = cni
    __args__['controlPlaneCa'] = control_plane_ca
    __args__['createdAt'] = created_at
    __args__['description'] = description
    __args__['endpoint'] = endpoint
    __args__['exoscaleCcm'] = exoscale_ccm
    __args__['exoscaleCsi'] = exoscale_csi
    __args__['id'] = id
    __args__['kubeletCa'] = kubelet_ca
    __args__['labels'] = labels
    __args__['metricsServer'] = metrics_server
    __args__['name'] = name
    __args__['nodepools'] = nodepools
    __args__['oidc'] = oidc
    __args__['serviceLevel'] = service_level
    __args__['state'] = state
    __args__['version'] = version
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('exoscale:index/getSksCluster:getSksCluster', __args__, opts=opts, typ=GetSksClusterResult).value

    return AwaitableGetSksClusterResult(
        addons=pulumi.get(__ret__, 'addons'),
        aggregation_ca=pulumi.get(__ret__, 'aggregation_ca'),
        auto_upgrade=pulumi.get(__ret__, 'auto_upgrade'),
        cni=pulumi.get(__ret__, 'cni'),
        control_plane_ca=pulumi.get(__ret__, 'control_plane_ca'),
        created_at=pulumi.get(__ret__, 'created_at'),
        description=pulumi.get(__ret__, 'description'),
        endpoint=pulumi.get(__ret__, 'endpoint'),
        exoscale_ccm=pulumi.get(__ret__, 'exoscale_ccm'),
        exoscale_csi=pulumi.get(__ret__, 'exoscale_csi'),
        id=pulumi.get(__ret__, 'id'),
        kubelet_ca=pulumi.get(__ret__, 'kubelet_ca'),
        labels=pulumi.get(__ret__, 'labels'),
        metrics_server=pulumi.get(__ret__, 'metrics_server'),
        name=pulumi.get(__ret__, 'name'),
        nodepools=pulumi.get(__ret__, 'nodepools'),
        oidc=pulumi.get(__ret__, 'oidc'),
        service_level=pulumi.get(__ret__, 'service_level'),
        state=pulumi.get(__ret__, 'state'),
        version=pulumi.get(__ret__, 'version'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_sks_cluster)
def get_sks_cluster_output(addons: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           aggregation_ca: Optional[pulumi.Input[Optional[str]]] = None,
                           auto_upgrade: Optional[pulumi.Input[Optional[bool]]] = None,
                           cni: Optional[pulumi.Input[Optional[str]]] = None,
                           control_plane_ca: Optional[pulumi.Input[Optional[str]]] = None,
                           created_at: Optional[pulumi.Input[Optional[str]]] = None,
                           description: Optional[pulumi.Input[Optional[str]]] = None,
                           endpoint: Optional[pulumi.Input[Optional[str]]] = None,
                           exoscale_ccm: Optional[pulumi.Input[Optional[bool]]] = None,
                           exoscale_csi: Optional[pulumi.Input[Optional[bool]]] = None,
                           id: Optional[pulumi.Input[Optional[str]]] = None,
                           kubelet_ca: Optional[pulumi.Input[Optional[str]]] = None,
                           labels: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                           metrics_server: Optional[pulumi.Input[Optional[bool]]] = None,
                           name: Optional[pulumi.Input[Optional[str]]] = None,
                           nodepools: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           oidc: Optional[pulumi.Input[Optional[Union['GetSksClusterOidcArgs', 'GetSksClusterOidcArgsDict']]]] = None,
                           service_level: Optional[pulumi.Input[Optional[str]]] = None,
                           state: Optional[pulumi.Input[Optional[str]]] = None,
                           version: Optional[pulumi.Input[Optional[str]]] = None,
                           zone: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSksClusterResult]:
    """
    Use this data source to access information about an existing resource.

    :param str aggregation_ca: The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
    :param bool auto_upgrade: Enable automatic upgrading of the control plane version.
    :param str cni: The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
    :param str control_plane_ca: The CA certificate (in PEM format) for TLS communications between control plane components.
    :param str created_at: The cluster creation date.
    :param str description: A free-form text describing the cluster.
    :param str endpoint: The cluster API endpoint.
    :param bool exoscale_ccm: Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
    :param bool exoscale_csi: Deploy the Exoscale [Container Storage Interface](https://github.com/exoscale/exoscale-csi-driver/) on worker nodes (boolean; default: `false`; requires the CCM to be enabled).
    :param str id: The ID of this resource.
    :param str kubelet_ca: The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
    :param Mapping[str, str] labels: A map of key/value labels.
    :param bool metrics_server: Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
    :param Sequence[str] nodepools: The list of exoscale*sks*nodepool (IDs) attached to the cluster.
    :param Union['GetSksClusterOidcArgs', 'GetSksClusterOidcArgsDict'] oidc: An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
    :param str service_level: The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
    :param str state: The cluster state.
    :param str version: The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
    """
    ...
