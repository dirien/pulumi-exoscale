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

__all__ = [
    'GetElasticIPResult',
    'AwaitableGetElasticIPResult',
    'get_elastic_ip',
    'get_elastic_ip_output',
]

@pulumi.output_type
class GetElasticIPResult:
    """
    A collection of values returned by getElasticIP.
    """
    def __init__(__self__, address_family=None, cidr=None, description=None, healthchecks=None, id=None, ip_address=None, labels=None, reverse_dns=None, zone=None):
        if address_family and not isinstance(address_family, str):
            raise TypeError("Expected argument 'address_family' to be a str")
        pulumi.set(__self__, "address_family", address_family)
        if cidr and not isinstance(cidr, str):
            raise TypeError("Expected argument 'cidr' to be a str")
        pulumi.set(__self__, "cidr", cidr)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if healthchecks and not isinstance(healthchecks, list):
            raise TypeError("Expected argument 'healthchecks' to be a list")
        pulumi.set(__self__, "healthchecks", healthchecks)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_address and not isinstance(ip_address, str):
            raise TypeError("Expected argument 'ip_address' to be a str")
        pulumi.set(__self__, "ip_address", ip_address)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if reverse_dns and not isinstance(reverse_dns, str):
            raise TypeError("Expected argument 'reverse_dns' to be a str")
        pulumi.set(__self__, "reverse_dns", reverse_dns)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> str:
        """
        The Elastic IP (EIP) address family (`inet4` or `inet6`).
        """
        return pulumi.get(self, "address_family")

    @property
    @pulumi.getter
    def cidr(self) -> str:
        """
        The Elastic IP (EIP) CIDR.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The Elastic IP (EIP) description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def healthchecks(self) -> Sequence['outputs.GetElasticIPHealthcheckResult']:
        """
        The *managed* EIP healthcheck configuration.
        """
        return pulumi.get(self, "healthchecks")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The Elastic IP (EIP) ID to match (conflicts with `ip_address` and `labels`).
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[str]:
        """
        The EIP IPv4 or IPv6 address to match (conflicts with `id` and `labels`).
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Mapping[str, str]]:
        """
        The EIP labels to match (conflicts with `ip_address` and `id`).
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="reverseDns")
    def reverse_dns(self) -> str:
        """
        Domain name for reverse DNS record.
        """
        return pulumi.get(self, "reverse_dns")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        The Exocale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        return pulumi.get(self, "zone")


class AwaitableGetElasticIPResult(GetElasticIPResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetElasticIPResult(
            address_family=self.address_family,
            cidr=self.cidr,
            description=self.description,
            healthchecks=self.healthchecks,
            id=self.id,
            ip_address=self.ip_address,
            labels=self.labels,
            reverse_dns=self.reverse_dns,
            zone=self.zone)


def get_elastic_ip(id: Optional[str] = None,
                   ip_address: Optional[str] = None,
                   labels: Optional[Mapping[str, str]] = None,
                   zone: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetElasticIPResult:
    """
    Use this data source to access information about an existing resource.

    :param str id: The Elastic IP (EIP) ID to match (conflicts with `ip_address` and `labels`).
    :param str ip_address: The EIP IPv4 or IPv6 address to match (conflicts with `id` and `labels`).
    :param Mapping[str, str] labels: The EIP labels to match (conflicts with `ip_address` and `id`).
    :param str zone: The Exocale [Zone](https://www.exoscale.com/datacenters/) name.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['ipAddress'] = ip_address
    __args__['labels'] = labels
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('exoscale:index/getElasticIP:getElasticIP', __args__, opts=opts, typ=GetElasticIPResult).value

    return AwaitableGetElasticIPResult(
        address_family=pulumi.get(__ret__, 'address_family'),
        cidr=pulumi.get(__ret__, 'cidr'),
        description=pulumi.get(__ret__, 'description'),
        healthchecks=pulumi.get(__ret__, 'healthchecks'),
        id=pulumi.get(__ret__, 'id'),
        ip_address=pulumi.get(__ret__, 'ip_address'),
        labels=pulumi.get(__ret__, 'labels'),
        reverse_dns=pulumi.get(__ret__, 'reverse_dns'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_elastic_ip)
def get_elastic_ip_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                          ip_address: Optional[pulumi.Input[Optional[str]]] = None,
                          labels: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                          zone: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetElasticIPResult]:
    """
    Use this data source to access information about an existing resource.

    :param str id: The Elastic IP (EIP) ID to match (conflicts with `ip_address` and `labels`).
    :param str ip_address: The EIP IPv4 or IPv6 address to match (conflicts with `id` and `labels`).
    :param Mapping[str, str] labels: The EIP labels to match (conflicts with `ip_address` and `id`).
    :param str zone: The Exocale [Zone](https://www.exoscale.com/datacenters/) name.
    """
    ...
