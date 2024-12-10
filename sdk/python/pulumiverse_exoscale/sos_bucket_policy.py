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

__all__ = ['SosBucketPolicyArgs', 'SosBucketPolicy']

@pulumi.input_type
class SosBucketPolicyArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 policy: pulumi.Input[str],
                 zone: pulumi.Input[str],
                 timeouts: Optional[pulumi.Input['SosBucketPolicyTimeoutsArgs']] = None):
        """
        The set of arguments for constructing a SosBucketPolicy resource.
        :param pulumi.Input[str] bucket: ❗ The name of the bucket to which the policy is to be applied.
        :param pulumi.Input[str] policy: The content of the policy
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "policy", policy)
        pulumi.set(__self__, "zone", zone)
        if timeouts is not None:
            pulumi.set(__self__, "timeouts", timeouts)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        """
        ❗ The name of the bucket to which the policy is to be applied.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Input[str]:
        """
        The content of the policy
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy", value)

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
    def timeouts(self) -> Optional[pulumi.Input['SosBucketPolicyTimeoutsArgs']]:
        return pulumi.get(self, "timeouts")

    @timeouts.setter
    def timeouts(self, value: Optional[pulumi.Input['SosBucketPolicyTimeoutsArgs']]):
        pulumi.set(self, "timeouts", value)


@pulumi.input_type
class _SosBucketPolicyState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 timeouts: Optional[pulumi.Input['SosBucketPolicyTimeoutsArgs']] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SosBucketPolicy resources.
        :param pulumi.Input[str] bucket: ❗ The name of the bucket to which the policy is to be applied.
        :param pulumi.Input[str] policy: The content of the policy
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if timeouts is not None:
            pulumi.set(__self__, "timeouts", timeouts)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        """
        ❗ The name of the bucket to which the policy is to be applied.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        The content of the policy
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter
    def timeouts(self) -> Optional[pulumi.Input['SosBucketPolicyTimeoutsArgs']]:
        return pulumi.get(self, "timeouts")

    @timeouts.setter
    def timeouts(self, value: Optional[pulumi.Input['SosBucketPolicyTimeoutsArgs']]):
        pulumi.set(self, "timeouts", value)

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


class SosBucketPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 timeouts: Optional[pulumi.Input[Union['SosBucketPolicyTimeoutsArgs', 'SosBucketPolicyTimeoutsArgsDict']]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manage Exoscale [SOS Bucket Policies](https://community.exoscale.com/documentation/storage/bucketpolicy/).

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: ❗ The name of the bucket to which the policy is to be applied.
        :param pulumi.Input[str] policy: The content of the policy
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SosBucketPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage Exoscale [SOS Bucket Policies](https://community.exoscale.com/documentation/storage/bucketpolicy/).

        :param str resource_name: The name of the resource.
        :param SosBucketPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SosBucketPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 timeouts: Optional[pulumi.Input[Union['SosBucketPolicyTimeoutsArgs', 'SosBucketPolicyTimeoutsArgsDict']]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SosBucketPolicyArgs.__new__(SosBucketPolicyArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__.__dict__["policy"] = policy
            __props__.__dict__["timeouts"] = timeouts
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
        super(SosBucketPolicy, __self__).__init__(
            'exoscale:index/sosBucketPolicy:SosBucketPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[str]] = None,
            timeouts: Optional[pulumi.Input[Union['SosBucketPolicyTimeoutsArgs', 'SosBucketPolicyTimeoutsArgsDict']]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'SosBucketPolicy':
        """
        Get an existing SosBucketPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: ❗ The name of the bucket to which the policy is to be applied.
        :param pulumi.Input[str] policy: The content of the policy
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SosBucketPolicyState.__new__(_SosBucketPolicyState)

        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["policy"] = policy
        __props__.__dict__["timeouts"] = timeouts
        __props__.__dict__["zone"] = zone
        return SosBucketPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        ❗ The name of the bucket to which the policy is to be applied.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        The content of the policy
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter
    def timeouts(self) -> pulumi.Output[Optional['outputs.SosBucketPolicyTimeouts']]:
        return pulumi.get(self, "timeouts")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        return pulumi.get(self, "zone")

