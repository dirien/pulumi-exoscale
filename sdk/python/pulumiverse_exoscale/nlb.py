# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['NLBArgs', 'NLB']

@pulumi.input_type
class NLBArgs:
    def __init__(__self__, *,
                 zone: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NLB resource.
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        :param pulumi.Input[str] description: A free-form text describing the NLB.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value labels.
        :param pulumi.Input[str] name: The network load balancer (NLB) name.
        """
        pulumi.set(__self__, "zone", zone)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)

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
        A free-form text describing the NLB.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

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
        The network load balancer (NLB) name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _NLBState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NLB resources.
        :param pulumi.Input[str] created_at: The NLB creation date.
        :param pulumi.Input[str] description: A free-form text describing the NLB.
        :param pulumi.Input[str] ip_address: The NLB IPv4 address.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value labels.
        :param pulumi.Input[str] name: The network load balancer (NLB) name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: The list of the exoscale*nlb*service (names).
        :param pulumi.Input[str] state: The current NLB state.
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if services is not None:
            pulumi.set(__self__, "services", services)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The NLB creation date.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A free-form text describing the NLB.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The NLB IPv4 address.
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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The network load balancer (NLB) name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of the exoscale*nlb*service (names).
        """
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "services", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The current NLB state.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

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


class NLB(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        An existing network load balancer (NLB) may be imported by `<ID>@<zone>`console

        ```sh
         $ pulumi import exoscale:index/nLB:NLB \\
        ```

         exoscale_nlb.my_nlb \\

         f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A free-form text describing the NLB.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value labels.
        :param pulumi.Input[str] name: The network load balancer (NLB) name.
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NLBArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        An existing network load balancer (NLB) may be imported by `<ID>@<zone>`console

        ```sh
         $ pulumi import exoscale:index/nLB:NLB \\
        ```

         exoscale_nlb.my_nlb \\

         f81d4fae-7dec-11d0-a765-00a0c91e6bf6@ch-gva-2

        :param str resource_name: The name of the resource.
        :param NLBArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NLBArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NLBArgs.__new__(NLBArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
            __props__.__dict__["created_at"] = None
            __props__.__dict__["ip_address"] = None
            __props__.__dict__["services"] = None
            __props__.__dict__["state"] = None
        super(NLB, __self__).__init__(
            'exoscale:index/nLB:NLB',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            state: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'NLB':
        """
        Get an existing NLB resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The NLB creation date.
        :param pulumi.Input[str] description: A free-form text describing the NLB.
        :param pulumi.Input[str] ip_address: The NLB IPv4 address.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value labels.
        :param pulumi.Input[str] name: The network load balancer (NLB) name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: The list of the exoscale*nlb*service (names).
        :param pulumi.Input[str] state: The current NLB state.
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NLBState.__new__(_NLBState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["ip_address"] = ip_address
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["services"] = services
        __props__.__dict__["state"] = state
        __props__.__dict__["zone"] = zone
        return NLB(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The NLB creation date.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A free-form text describing the NLB.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        The NLB IPv4 address.
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
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The network load balancer (NLB) name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def services(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of the exoscale*nlb*service (names).
        """
        return pulumi.get(self, "services")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current NLB state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        return pulumi.get(self, "zone")

