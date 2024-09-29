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

__all__ = ['BlockStorageVolumeArgs', 'BlockStorageVolume']

@pulumi.input_type
class BlockStorageVolumeArgs:
    def __init__(__self__, *,
                 zone: pulumi.Input[str],
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 snapshot_target: Optional[pulumi.Input['BlockStorageVolumeSnapshotTargetArgs']] = None,
                 timeouts: Optional[pulumi.Input['BlockStorageVolumeTimeoutsArgs']] = None):
        """
        The set of arguments for constructing a BlockStorageVolume resource.
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels.
        :param pulumi.Input[str] name: Volume name.
        :param pulumi.Input[int] size: Volume size in GB (default 10). If volume is attached, instance must be stopped to update this value. Volume can only grow, cannot be shrunk.
        :param pulumi.Input['BlockStorageVolumeSnapshotTargetArgs'] snapshot_target: Block storage snapshot to use when creating a volume. Read-only after creation.
        """
        pulumi.set(__self__, "zone", zone)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if snapshot_target is not None:
            pulumi.set(__self__, "snapshot_target", snapshot_target)
        if timeouts is not None:
            pulumi.set(__self__, "timeouts", timeouts)

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
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Resource labels.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Volume name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        Volume size in GB (default 10). If volume is attached, instance must be stopped to update this value. Volume can only grow, cannot be shrunk.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="snapshotTarget")
    def snapshot_target(self) -> Optional[pulumi.Input['BlockStorageVolumeSnapshotTargetArgs']]:
        """
        Block storage snapshot to use when creating a volume. Read-only after creation.
        """
        return pulumi.get(self, "snapshot_target")

    @snapshot_target.setter
    def snapshot_target(self, value: Optional[pulumi.Input['BlockStorageVolumeSnapshotTargetArgs']]):
        pulumi.set(self, "snapshot_target", value)

    @property
    @pulumi.getter
    def timeouts(self) -> Optional[pulumi.Input['BlockStorageVolumeTimeoutsArgs']]:
        return pulumi.get(self, "timeouts")

    @timeouts.setter
    def timeouts(self, value: Optional[pulumi.Input['BlockStorageVolumeTimeoutsArgs']]):
        pulumi.set(self, "timeouts", value)


@pulumi.input_type
class _BlockStorageVolumeState:
    def __init__(__self__, *,
                 blocksize: Optional[pulumi.Input[int]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 snapshot_target: Optional[pulumi.Input['BlockStorageVolumeSnapshotTargetArgs']] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 timeouts: Optional[pulumi.Input['BlockStorageVolumeTimeoutsArgs']] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BlockStorageVolume resources.
        :param pulumi.Input[int] blocksize: Volume block size.
        :param pulumi.Input[str] created_at: Volume creation date.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels.
        :param pulumi.Input[str] name: Volume name.
        :param pulumi.Input[int] size: Volume size in GB (default 10). If volume is attached, instance must be stopped to update this value. Volume can only grow, cannot be shrunk.
        :param pulumi.Input['BlockStorageVolumeSnapshotTargetArgs'] snapshot_target: Block storage snapshot to use when creating a volume. Read-only after creation.
        :param pulumi.Input[str] state: Volume state.
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        if blocksize is not None:
            pulumi.set(__self__, "blocksize", blocksize)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if snapshot_target is not None:
            pulumi.set(__self__, "snapshot_target", snapshot_target)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if timeouts is not None:
            pulumi.set(__self__, "timeouts", timeouts)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def blocksize(self) -> Optional[pulumi.Input[int]]:
        """
        Volume block size.
        """
        return pulumi.get(self, "blocksize")

    @blocksize.setter
    def blocksize(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "blocksize", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Volume creation date.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Resource labels.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Volume name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        Volume size in GB (default 10). If volume is attached, instance must be stopped to update this value. Volume can only grow, cannot be shrunk.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="snapshotTarget")
    def snapshot_target(self) -> Optional[pulumi.Input['BlockStorageVolumeSnapshotTargetArgs']]:
        """
        Block storage snapshot to use when creating a volume. Read-only after creation.
        """
        return pulumi.get(self, "snapshot_target")

    @snapshot_target.setter
    def snapshot_target(self, value: Optional[pulumi.Input['BlockStorageVolumeSnapshotTargetArgs']]):
        pulumi.set(self, "snapshot_target", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Volume state.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def timeouts(self) -> Optional[pulumi.Input['BlockStorageVolumeTimeoutsArgs']]:
        return pulumi.get(self, "timeouts")

    @timeouts.setter
    def timeouts(self, value: Optional[pulumi.Input['BlockStorageVolumeTimeoutsArgs']]):
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


class BlockStorageVolume(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 snapshot_target: Optional[pulumi.Input[Union['BlockStorageVolumeSnapshotTargetArgs', 'BlockStorageVolumeSnapshotTargetArgsDict']]] = None,
                 timeouts: Optional[pulumi.Input[Union['BlockStorageVolumeTimeoutsArgs', 'BlockStorageVolumeTimeoutsArgsDict']]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manage [Exoscale Block Storage](https://community.exoscale.com/documentation/block-storage/) Volume.

        Block Storage offers persistent externally attached volumes for your workloads.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels.
        :param pulumi.Input[str] name: Volume name.
        :param pulumi.Input[int] size: Volume size in GB (default 10). If volume is attached, instance must be stopped to update this value. Volume can only grow, cannot be shrunk.
        :param pulumi.Input[Union['BlockStorageVolumeSnapshotTargetArgs', 'BlockStorageVolumeSnapshotTargetArgsDict']] snapshot_target: Block storage snapshot to use when creating a volume. Read-only after creation.
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BlockStorageVolumeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage [Exoscale Block Storage](https://community.exoscale.com/documentation/block-storage/) Volume.

        Block Storage offers persistent externally attached volumes for your workloads.

        :param str resource_name: The name of the resource.
        :param BlockStorageVolumeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BlockStorageVolumeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 snapshot_target: Optional[pulumi.Input[Union['BlockStorageVolumeSnapshotTargetArgs', 'BlockStorageVolumeSnapshotTargetArgsDict']]] = None,
                 timeouts: Optional[pulumi.Input[Union['BlockStorageVolumeTimeoutsArgs', 'BlockStorageVolumeTimeoutsArgsDict']]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BlockStorageVolumeArgs.__new__(BlockStorageVolumeArgs)

            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["size"] = size
            __props__.__dict__["snapshot_target"] = snapshot_target
            __props__.__dict__["timeouts"] = timeouts
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
            __props__.__dict__["blocksize"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["state"] = None
        super(BlockStorageVolume, __self__).__init__(
            'exoscale:index/blockStorageVolume:BlockStorageVolume',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            blocksize: Optional[pulumi.Input[int]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[int]] = None,
            snapshot_target: Optional[pulumi.Input[Union['BlockStorageVolumeSnapshotTargetArgs', 'BlockStorageVolumeSnapshotTargetArgsDict']]] = None,
            state: Optional[pulumi.Input[str]] = None,
            timeouts: Optional[pulumi.Input[Union['BlockStorageVolumeTimeoutsArgs', 'BlockStorageVolumeTimeoutsArgsDict']]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'BlockStorageVolume':
        """
        Get an existing BlockStorageVolume resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] blocksize: Volume block size.
        :param pulumi.Input[str] created_at: Volume creation date.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels.
        :param pulumi.Input[str] name: Volume name.
        :param pulumi.Input[int] size: Volume size in GB (default 10). If volume is attached, instance must be stopped to update this value. Volume can only grow, cannot be shrunk.
        :param pulumi.Input[Union['BlockStorageVolumeSnapshotTargetArgs', 'BlockStorageVolumeSnapshotTargetArgsDict']] snapshot_target: Block storage snapshot to use when creating a volume. Read-only after creation.
        :param pulumi.Input[str] state: Volume state.
        :param pulumi.Input[str] zone: ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BlockStorageVolumeState.__new__(_BlockStorageVolumeState)

        __props__.__dict__["blocksize"] = blocksize
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["size"] = size
        __props__.__dict__["snapshot_target"] = snapshot_target
        __props__.__dict__["state"] = state
        __props__.__dict__["timeouts"] = timeouts
        __props__.__dict__["zone"] = zone
        return BlockStorageVolume(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def blocksize(self) -> pulumi.Output[int]:
        """
        Volume block size.
        """
        return pulumi.get(self, "blocksize")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Volume creation date.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Volume name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[Optional[int]]:
        """
        Volume size in GB (default 10). If volume is attached, instance must be stopped to update this value. Volume can only grow, cannot be shrunk.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="snapshotTarget")
    def snapshot_target(self) -> pulumi.Output[Optional['outputs.BlockStorageVolumeSnapshotTarget']]:
        """
        Block storage snapshot to use when creating a volume. Read-only after creation.
        """
        return pulumi.get(self, "snapshot_target")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Volume state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def timeouts(self) -> pulumi.Output[Optional['outputs.BlockStorageVolumeTimeouts']]:
        return pulumi.get(self, "timeouts")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        return pulumi.get(self, "zone")

