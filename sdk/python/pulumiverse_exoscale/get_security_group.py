# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSecurityGroupResult',
    'AwaitableGetSecurityGroupResult',
    'get_security_group',
    'get_security_group_output',
]

@pulumi.output_type
class GetSecurityGroupResult:
    """
    A collection of values returned by getSecurityGroup.
    """
    def __init__(__self__, external_sources=None, id=None, name=None):
        if external_sources and not isinstance(external_sources, list):
            raise TypeError("Expected argument 'external_sources' to be a list")
        pulumi.set(__self__, "external_sources", external_sources)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="externalSources")
    def external_sources(self) -> Sequence[str]:
        """
        The list of external network sources, in [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notatio) notation.
        """
        return pulumi.get(self, "external_sources")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The security group ID to match (conflicts with `name`)
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name to match (conflicts with `id`)
        """
        return pulumi.get(self, "name")


class AwaitableGetSecurityGroupResult(GetSecurityGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecurityGroupResult(
            external_sources=self.external_sources,
            id=self.id,
            name=self.name)


def get_security_group(id: Optional[str] = None,
                       name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecurityGroupResult:
    """
    Use this data source to access information about an existing resource.

    :param str id: The security group ID to match (conflicts with `name`)
    :param str name: The name to match (conflicts with `id`)
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('exoscale:index/getSecurityGroup:getSecurityGroup', __args__, opts=opts, typ=GetSecurityGroupResult).value

    return AwaitableGetSecurityGroupResult(
        external_sources=pulumi.get(__ret__, 'external_sources'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_security_group)
def get_security_group_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                              name: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecurityGroupResult]:
    """
    Use this data source to access information about an existing resource.

    :param str id: The security group ID to match (conflicts with `name`)
    :param str name: The name to match (conflicts with `id`)
    """
    ...
