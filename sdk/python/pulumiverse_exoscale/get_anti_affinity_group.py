# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetAntiAffinityGroupResult',
    'AwaitableGetAntiAffinityGroupResult',
    'get_anti_affinity_group',
    'get_anti_affinity_group_output',
]

@pulumi.output_type
class GetAntiAffinityGroupResult:
    """
    A collection of values returned by getAntiAffinityGroup.
    """
    def __init__(__self__, id=None, instances=None, name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The anti-affinity group ID to match (conflicts with `name`).
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def instances(self) -> Sequence[str]:
        """
        The list of attached exoscale*compute*instance (IDs).
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The group name to match (conflicts with `id`).
        """
        return pulumi.get(self, "name")


class AwaitableGetAntiAffinityGroupResult(GetAntiAffinityGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAntiAffinityGroupResult(
            id=self.id,
            instances=self.instances,
            name=self.name)


def get_anti_affinity_group(id: Optional[str] = None,
                            name: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAntiAffinityGroupResult:
    """
    Fetch Exoscale [Anti-Affinity Groups](https://community.exoscale.com/documentation/compute/anti-affinity-groups/) data.

    Corresponding resource: exoscale_anti_affinity_group.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_exoscale as exoscale

    my_anti_affinity_group = exoscale.get_anti_affinity_group(name="my-anti-affinity-group")
    pulumi.export("myAntiAffinityGroupId", my_anti_affinity_group.id)
    ```
    <!--End PulumiCodeChooser -->

    Please refer to the examples
    directory for complete configuration examples.


    :param str id: The anti-affinity group ID to match (conflicts with `name`).
    :param str name: The group name to match (conflicts with `id`).
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('exoscale:index/getAntiAffinityGroup:getAntiAffinityGroup', __args__, opts=opts, typ=GetAntiAffinityGroupResult).value

    return AwaitableGetAntiAffinityGroupResult(
        id=pulumi.get(__ret__, 'id'),
        instances=pulumi.get(__ret__, 'instances'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_anti_affinity_group)
def get_anti_affinity_group_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                                   name: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAntiAffinityGroupResult]:
    """
    Fetch Exoscale [Anti-Affinity Groups](https://community.exoscale.com/documentation/compute/anti-affinity-groups/) data.

    Corresponding resource: exoscale_anti_affinity_group.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_exoscale as exoscale

    my_anti_affinity_group = exoscale.get_anti_affinity_group(name="my-anti-affinity-group")
    pulumi.export("myAntiAffinityGroupId", my_anti_affinity_group.id)
    ```
    <!--End PulumiCodeChooser -->

    Please refer to the examples
    directory for complete configuration examples.


    :param str id: The anti-affinity group ID to match (conflicts with `name`).
    :param str name: The group name to match (conflicts with `id`).
    """
    ...
