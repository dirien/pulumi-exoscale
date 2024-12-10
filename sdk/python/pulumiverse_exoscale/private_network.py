# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PrivateNetworkArgs', 'PrivateNetwork']

@pulumi.input_type
class PrivateNetworkArgs:
    def __init__(__self__, *,
                 zone: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PrivateNetwork resource.
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        :param pulumi.Input[str] description: A free-form text describing the network.
        :param pulumi.Input[str] end_ip: (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value labels.
        :param pulumi.Input[str] name: The private network name.
        :param pulumi.Input[str] netmask: (For managed Privnets) The network mask defining the IPv4 network allowed for static leases.
        :param pulumi.Input[str] start_ip: (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        """
        pulumi.set(__self__, "zone", zone)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if end_ip is not None:
            pulumi.set(__self__, "end_ip", end_ip)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if netmask is not None:
            pulumi.set(__self__, "netmask", netmask)
        if start_ip is not None:
            pulumi.set(__self__, "start_ip", start_ip)

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
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A free-form text describing the network.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> Optional[pulumi.Input[str]]:
        """
        (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        """
        return pulumi.get(self, "end_ip")

    @end_ip.setter
    def end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_ip", value)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The private network name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def netmask(self) -> Optional[pulumi.Input[str]]:
        """
        (For managed Privnets) The network mask defining the IPv4 network allowed for static leases.
        """
        return pulumi.get(self, "netmask")

    @netmask.setter
    def netmask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netmask", value)

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> Optional[pulumi.Input[str]]:
        """
        (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        """
        return pulumi.get(self, "start_ip")

    @start_ip.setter
    def start_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_ip", value)


@pulumi.input_type
class _PrivateNetworkState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PrivateNetwork resources.
        :param pulumi.Input[str] description: A free-form text describing the network.
        :param pulumi.Input[str] end_ip: (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value labels.
        :param pulumi.Input[str] name: The private network name.
        :param pulumi.Input[str] netmask: (For managed Privnets) The network mask defining the IPv4 network allowed for static leases.
        :param pulumi.Input[str] start_ip: (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if end_ip is not None:
            pulumi.set(__self__, "end_ip", end_ip)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if netmask is not None:
            pulumi.set(__self__, "netmask", netmask)
        if start_ip is not None:
            pulumi.set(__self__, "start_ip", start_ip)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A free-form text describing the network.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> Optional[pulumi.Input[str]]:
        """
        (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        """
        return pulumi.get(self, "end_ip")

    @end_ip.setter
    def end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_ip", value)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The private network name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def netmask(self) -> Optional[pulumi.Input[str]]:
        """
        (For managed Privnets) The network mask defining the IPv4 network allowed for static leases.
        """
        return pulumi.get(self, "netmask")

    @netmask.setter
    def netmask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netmask", value)

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> Optional[pulumi.Input[str]]:
        """
        (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        """
        return pulumi.get(self, "start_ip")

    @start_ip.setter
    def start_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_ip", value)

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


class PrivateNetwork(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manage Exoscale [Private Networks](https://community.exoscale.com/documentation/compute/private-networks/).

        Corresponding data source: exoscale_private_network.

        ## Example Usage

        *Unmanaged* private network:

        ```python
        import pulumi
        import pulumiverse_exoscale as exoscale

        my_private_network = exoscale.PrivateNetwork("my_private_network",
            zone="ch-gva-2",
            name="my-private-network")
        ```

        *Managed* private network:

        ```python
        import pulumi
        import pulumiverse_exoscale as exoscale

        my_managed_private_network = exoscale.PrivateNetwork("my_managed_private_network",
            zone="ch-gva-2",
            name="my-managed-private-network",
            netmask="255.255.255.0",
            start_ip="10.0.0.20",
            end_ip="10.0.0.253")
        ```

        Please refer to the examples
        directory for complete configuration examples.

        ## Import

        An existing private network may be imported by `<ID>@<zone>`:

        ```sh
        $ pulumi import exoscale:index/privateNetwork:PrivateNetwork \\
        ```

          exoscale_private_network.my_private_network \\

          f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A free-form text describing the network.
        :param pulumi.Input[str] end_ip: (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value labels.
        :param pulumi.Input[str] name: The private network name.
        :param pulumi.Input[str] netmask: (For managed Privnets) The network mask defining the IPv4 network allowed for static leases.
        :param pulumi.Input[str] start_ip: (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateNetworkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage Exoscale [Private Networks](https://community.exoscale.com/documentation/compute/private-networks/).

        Corresponding data source: exoscale_private_network.

        ## Example Usage

        *Unmanaged* private network:

        ```python
        import pulumi
        import pulumiverse_exoscale as exoscale

        my_private_network = exoscale.PrivateNetwork("my_private_network",
            zone="ch-gva-2",
            name="my-private-network")
        ```

        *Managed* private network:

        ```python
        import pulumi
        import pulumiverse_exoscale as exoscale

        my_managed_private_network = exoscale.PrivateNetwork("my_managed_private_network",
            zone="ch-gva-2",
            name="my-managed-private-network",
            netmask="255.255.255.0",
            start_ip="10.0.0.20",
            end_ip="10.0.0.253")
        ```

        Please refer to the examples
        directory for complete configuration examples.

        ## Import

        An existing private network may be imported by `<ID>@<zone>`:

        ```sh
        $ pulumi import exoscale:index/privateNetwork:PrivateNetwork \\
        ```

          exoscale_private_network.my_private_network \\

          f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2

        :param str resource_name: The name of the resource.
        :param PrivateNetworkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateNetworkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateNetworkArgs.__new__(PrivateNetworkArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["end_ip"] = end_ip
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["netmask"] = netmask
            __props__.__dict__["start_ip"] = start_ip
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
        super(PrivateNetwork, __self__).__init__(
            'exoscale:index/privateNetwork:PrivateNetwork',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            end_ip: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            netmask: Optional[pulumi.Input[str]] = None,
            start_ip: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'PrivateNetwork':
        """
        Get an existing PrivateNetwork resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A free-form text describing the network.
        :param pulumi.Input[str] end_ip: (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value labels.
        :param pulumi.Input[str] name: The private network name.
        :param pulumi.Input[str] netmask: (For managed Privnets) The network mask defining the IPv4 network allowed for static leases.
        :param pulumi.Input[str] start_ip: (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrivateNetworkState.__new__(_PrivateNetworkState)

        __props__.__dict__["description"] = description
        __props__.__dict__["end_ip"] = end_ip
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["netmask"] = netmask
        __props__.__dict__["start_ip"] = start_ip
        __props__.__dict__["zone"] = zone
        return PrivateNetwork(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A free-form text describing the network.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> pulumi.Output[Optional[str]]:
        """
        (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        """
        return pulumi.get(self, "end_ip")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of key/value labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The private network name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def netmask(self) -> pulumi.Output[Optional[str]]:
        """
        (For managed Privnets) The network mask defining the IPv4 network allowed for static leases.
        """
        return pulumi.get(self, "netmask")

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> pulumi.Output[Optional[str]]:
        """
        (For managed Privnets) The first/last IPv4 addresses used by the DHCP service for dynamic leases.
        """
        return pulumi.get(self, "start_ip")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        return pulumi.get(self, "zone")

