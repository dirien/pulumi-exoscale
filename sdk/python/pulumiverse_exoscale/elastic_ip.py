# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ElasticIpArgs', 'ElasticIp']

@pulumi.input_type
class ElasticIpArgs:
    def __init__(__self__, *,
                 zone: pulumi.Input[str],
                 address_family: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 healthcheck: Optional[pulumi.Input['ElasticIpHealthcheckArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 reverse_dns: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ElasticIp resource.
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        :param pulumi.Input[str] address_family: ❗ The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
        :param pulumi.Input[str] description: A free-form text describing the Elastic IP (EIP).
        :param pulumi.Input['ElasticIpHealthcheckArgs'] healthcheck: Healthcheck configuration for *managed* EIPs. It can not be added to an existing *Unmanaged* EIP.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value labels.
        :param pulumi.Input[str] reverse_dns: Domain name for reverse DNS record.
        """
        pulumi.set(__self__, "zone", zone)
        if address_family is not None:
            pulumi.set(__self__, "address_family", address_family)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if healthcheck is not None:
            pulumi.set(__self__, "healthcheck", healthcheck)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if reverse_dns is not None:
            pulumi.set(__self__, "reverse_dns", reverse_dns)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        """
        ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> Optional[pulumi.Input[str]]:
        """
        ❗ The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
        """
        return pulumi.get(self, "address_family")

    @address_family.setter
    def address_family(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_family", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A free-form text describing the Elastic IP (EIP).
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def healthcheck(self) -> Optional[pulumi.Input['ElasticIpHealthcheckArgs']]:
        """
        Healthcheck configuration for *managed* EIPs. It can not be added to an existing *Unmanaged* EIP.
        """
        return pulumi.get(self, "healthcheck")

    @healthcheck.setter
    def healthcheck(self, value: Optional[pulumi.Input['ElasticIpHealthcheckArgs']]):
        pulumi.set(self, "healthcheck", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of key/value labels.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="reverseDns")
    def reverse_dns(self) -> Optional[pulumi.Input[str]]:
        """
        Domain name for reverse DNS record.
        """
        return pulumi.get(self, "reverse_dns")

    @reverse_dns.setter
    def reverse_dns(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reverse_dns", value)


@pulumi.input_type
class _ElasticIpState:
    def __init__(__self__, *,
                 address_family: Optional[pulumi.Input[str]] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 healthcheck: Optional[pulumi.Input['ElasticIpHealthcheckArgs']] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 reverse_dns: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ElasticIp resources.
        :param pulumi.Input[str] address_family: ❗ The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
        :param pulumi.Input[str] cidr: The Elastic IP (EIP) CIDR.
        :param pulumi.Input[str] description: A free-form text describing the Elastic IP (EIP).
        :param pulumi.Input['ElasticIpHealthcheckArgs'] healthcheck: Healthcheck configuration for *managed* EIPs. It can not be added to an existing *Unmanaged* EIP.
        :param pulumi.Input[str] ip_address: The Elastic IP (EIP) IPv4 or IPv6 address.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value labels.
        :param pulumi.Input[str] reverse_dns: Domain name for reverse DNS record.
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        if address_family is not None:
            pulumi.set(__self__, "address_family", address_family)
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if healthcheck is not None:
            pulumi.set(__self__, "healthcheck", healthcheck)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if reverse_dns is not None:
            pulumi.set(__self__, "reverse_dns", reverse_dns)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> Optional[pulumi.Input[str]]:
        """
        ❗ The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
        """
        return pulumi.get(self, "address_family")

    @address_family.setter
    def address_family(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_family", value)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        The Elastic IP (EIP) CIDR.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A free-form text describing the Elastic IP (EIP).
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def healthcheck(self) -> Optional[pulumi.Input['ElasticIpHealthcheckArgs']]:
        """
        Healthcheck configuration for *managed* EIPs. It can not be added to an existing *Unmanaged* EIP.
        """
        return pulumi.get(self, "healthcheck")

    @healthcheck.setter
    def healthcheck(self, value: Optional[pulumi.Input['ElasticIpHealthcheckArgs']]):
        pulumi.set(self, "healthcheck", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The Elastic IP (EIP) IPv4 or IPv6 address.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of key/value labels.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="reverseDns")
    def reverse_dns(self) -> Optional[pulumi.Input[str]]:
        """
        Domain name for reverse DNS record.
        """
        return pulumi.get(self, "reverse_dns")

    @reverse_dns.setter
    def reverse_dns(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reverse_dns", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class ElasticIp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_family: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 healthcheck: Optional[pulumi.Input[Union['ElasticIpHealthcheckArgs', 'ElasticIpHealthcheckArgsDict']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 reverse_dns: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manage Exoscale [Elastic IPs (EIP)](https://community.exoscale.com/documentation/compute/eip/).

        Corresponding data source: exoscale_elastic_ip.

        ## Example Usage

        *Unmanaged* EIPv4:

        ```python
        import pulumi
        import pulumiverse_exoscale as exoscale

        my_elastic_ip = exoscale.ElasticIp("my_elastic_ip", zone="ch-gva-2")
        ```

        *Managed* EIPv6:

        ```python
        import pulumi
        import pulumiverse_exoscale as exoscale

        my_managed_elastic_ip = exoscale.ElasticIp("my_managed_elastic_ip",
            zone="ch-gva-2",
            address_family="inet6",
            reverse_dns="example.net",
            healthcheck={
                "mode": "https",
                "port": 443,
                "uri": "/health",
                "interval": 5,
                "timeout": 3,
                "strikes_ok": 2,
                "strikes_fail": 3,
                "tls_sni": "example.net",
            })
        ```

        Please refer to the examples
        directory for complete configuration examples.

        ## Import

        An existing Elastic IP (EIP) may be imported by `<ID>@<zone>`:

        ```sh
        $ pulumi import exoscale:index/elasticIp:ElasticIp \\
        ```

          exoscale_elastic_ip.my_elastic_ip \\

          f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_family: ❗ The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
        :param pulumi.Input[str] description: A free-form text describing the Elastic IP (EIP).
        :param pulumi.Input[Union['ElasticIpHealthcheckArgs', 'ElasticIpHealthcheckArgsDict']] healthcheck: Healthcheck configuration for *managed* EIPs. It can not be added to an existing *Unmanaged* EIP.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value labels.
        :param pulumi.Input[str] reverse_dns: Domain name for reverse DNS record.
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ElasticIpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage Exoscale [Elastic IPs (EIP)](https://community.exoscale.com/documentation/compute/eip/).

        Corresponding data source: exoscale_elastic_ip.

        ## Example Usage

        *Unmanaged* EIPv4:

        ```python
        import pulumi
        import pulumiverse_exoscale as exoscale

        my_elastic_ip = exoscale.ElasticIp("my_elastic_ip", zone="ch-gva-2")
        ```

        *Managed* EIPv6:

        ```python
        import pulumi
        import pulumiverse_exoscale as exoscale

        my_managed_elastic_ip = exoscale.ElasticIp("my_managed_elastic_ip",
            zone="ch-gva-2",
            address_family="inet6",
            reverse_dns="example.net",
            healthcheck={
                "mode": "https",
                "port": 443,
                "uri": "/health",
                "interval": 5,
                "timeout": 3,
                "strikes_ok": 2,
                "strikes_fail": 3,
                "tls_sni": "example.net",
            })
        ```

        Please refer to the examples
        directory for complete configuration examples.

        ## Import

        An existing Elastic IP (EIP) may be imported by `<ID>@<zone>`:

        ```sh
        $ pulumi import exoscale:index/elasticIp:ElasticIp \\
        ```

          exoscale_elastic_ip.my_elastic_ip \\

          f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2

        :param str resource_name: The name of the resource.
        :param ElasticIpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ElasticIpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_family: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 healthcheck: Optional[pulumi.Input[Union['ElasticIpHealthcheckArgs', 'ElasticIpHealthcheckArgsDict']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 reverse_dns: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ElasticIpArgs.__new__(ElasticIpArgs)

            __props__.__dict__["address_family"] = address_family
            __props__.__dict__["description"] = description
            __props__.__dict__["healthcheck"] = healthcheck
            __props__.__dict__["labels"] = labels
            __props__.__dict__["reverse_dns"] = reverse_dns
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
            __props__.__dict__["cidr"] = None
            __props__.__dict__["ip_address"] = None
        super(ElasticIp, __self__).__init__(
            'exoscale:index/elasticIp:ElasticIp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_family: Optional[pulumi.Input[str]] = None,
            cidr: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            healthcheck: Optional[pulumi.Input[Union['ElasticIpHealthcheckArgs', 'ElasticIpHealthcheckArgsDict']]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            reverse_dns: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'ElasticIp':
        """
        Get an existing ElasticIp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_family: ❗ The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
        :param pulumi.Input[str] cidr: The Elastic IP (EIP) CIDR.
        :param pulumi.Input[str] description: A free-form text describing the Elastic IP (EIP).
        :param pulumi.Input[Union['ElasticIpHealthcheckArgs', 'ElasticIpHealthcheckArgsDict']] healthcheck: Healthcheck configuration for *managed* EIPs. It can not be added to an existing *Unmanaged* EIP.
        :param pulumi.Input[str] ip_address: The Elastic IP (EIP) IPv4 or IPv6 address.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value labels.
        :param pulumi.Input[str] reverse_dns: Domain name for reverse DNS record.
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ElasticIpState.__new__(_ElasticIpState)

        __props__.__dict__["address_family"] = address_family
        __props__.__dict__["cidr"] = cidr
        __props__.__dict__["description"] = description
        __props__.__dict__["healthcheck"] = healthcheck
        __props__.__dict__["ip_address"] = ip_address
        __props__.__dict__["labels"] = labels
        __props__.__dict__["reverse_dns"] = reverse_dns
        __props__.__dict__["zone"] = zone
        return ElasticIp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> pulumi.Output[str]:
        """
        ❗ The Elastic IP (EIP) address family (`inet4` or `inet6`; default: `inet4`).
        """
        return pulumi.get(self, "address_family")

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Output[str]:
        """
        The Elastic IP (EIP) CIDR.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        A free-form text describing the Elastic IP (EIP).
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def healthcheck(self) -> pulumi.Output['outputs.ElasticIpHealthcheck']:
        """
        Healthcheck configuration for *managed* EIPs. It can not be added to an existing *Unmanaged* EIP.
        """
        return pulumi.get(self, "healthcheck")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        The Elastic IP (EIP) IPv4 or IPv6 address.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of key/value labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="reverseDns")
    def reverse_dns(self) -> pulumi.Output[Optional[str]]:
        """
        Domain name for reverse DNS record.
        """
        return pulumi.get(self, "reverse_dns")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        return pulumi.get(self, "zone")

