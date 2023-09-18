# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DomainArgs', 'Domain']

@pulumi.input_type
class DomainArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Domain resource.
        :param pulumi.Input[str] name: ❗ The DNS domain name.
        """
        DomainArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        ❗ The DNS domain name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _DomainState:
    def __init__(__self__, *,
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 expires_on: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Domain resources.
        :param pulumi.Input[bool] auto_renew: Whether the DNS domain has automatic renewal enabled (boolean).
        :param pulumi.Input[str] expires_on: The domain expiration date, if known.
        :param pulumi.Input[str] name: ❗ The DNS domain name.
        :param pulumi.Input[str] state: The domain state.
        :param pulumi.Input[str] token: A security token that can be used as an alternative way to manage DNS domains via the Exoscale API.
        """
        _DomainState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            auto_renew=auto_renew,
            expires_on=expires_on,
            name=name,
            state=state,
            token=token,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             auto_renew: Optional[pulumi.Input[bool]] = None,
             expires_on: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             state: Optional[pulumi.Input[str]] = None,
             token: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if auto_renew is not None:
            warnings.warn("""Not used, will be removed in the future""", DeprecationWarning)
            pulumi.log.warn("""auto_renew is deprecated: Not used, will be removed in the future""")
        if auto_renew is not None:
            _setter("auto_renew", auto_renew)
        if expires_on is not None:
            warnings.warn("""Not used, will be removed in the future""", DeprecationWarning)
            pulumi.log.warn("""expires_on is deprecated: Not used, will be removed in the future""")
        if expires_on is not None:
            _setter("expires_on", expires_on)
        if name is not None:
            _setter("name", name)
        if state is not None:
            warnings.warn("""Not used, will be removed in the future""", DeprecationWarning)
            pulumi.log.warn("""state is deprecated: Not used, will be removed in the future""")
        if state is not None:
            _setter("state", state)
        if token is not None:
            warnings.warn("""Not used, will be removed in the future""", DeprecationWarning)
            pulumi.log.warn("""token is deprecated: Not used, will be removed in the future""")
        if token is not None:
            _setter("token", token)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the DNS domain has automatic renewal enabled (boolean).
        """
        warnings.warn("""Not used, will be removed in the future""", DeprecationWarning)
        pulumi.log.warn("""auto_renew is deprecated: Not used, will be removed in the future""")

        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter(name="expiresOn")
    def expires_on(self) -> Optional[pulumi.Input[str]]:
        """
        The domain expiration date, if known.
        """
        warnings.warn("""Not used, will be removed in the future""", DeprecationWarning)
        pulumi.log.warn("""expires_on is deprecated: Not used, will be removed in the future""")

        return pulumi.get(self, "expires_on")

    @expires_on.setter
    def expires_on(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_on", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        ❗ The DNS domain name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The domain state.
        """
        warnings.warn("""Not used, will be removed in the future""", DeprecationWarning)
        pulumi.log.warn("""state is deprecated: Not used, will be removed in the future""")

        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        A security token that can be used as an alternative way to manage DNS domains via the Exoscale API.
        """
        warnings.warn("""Not used, will be removed in the future""", DeprecationWarning)
        pulumi.log.warn("""token is deprecated: Not used, will be removed in the future""")

        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)


class Domain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        An existing DNS domain may be imported by `ID`

        ```sh
         $ pulumi import exoscale:index/domain:Domain \\
        ```

         exoscale_domain.my_domain \\

         89083a5c-b648-474a-0000-0000000f67bd

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: ❗ The DNS domain name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DomainArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        An existing DNS domain may be imported by `ID`

        ```sh
         $ pulumi import exoscale:index/domain:Domain \\
        ```

         exoscale_domain.my_domain \\

         89083a5c-b648-474a-0000-0000000f67bd

        :param str resource_name: The name of the resource.
        :param DomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            DomainArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DomainArgs.__new__(DomainArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["auto_renew"] = None
            __props__.__dict__["expires_on"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["token"] = None
        super(Domain, __self__).__init__(
            'exoscale:index/domain:Domain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_renew: Optional[pulumi.Input[bool]] = None,
            expires_on: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None) -> 'Domain':
        """
        Get an existing Domain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_renew: Whether the DNS domain has automatic renewal enabled (boolean).
        :param pulumi.Input[str] expires_on: The domain expiration date, if known.
        :param pulumi.Input[str] name: ❗ The DNS domain name.
        :param pulumi.Input[str] state: The domain state.
        :param pulumi.Input[str] token: A security token that can be used as an alternative way to manage DNS domains via the Exoscale API.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DomainState.__new__(_DomainState)

        __props__.__dict__["auto_renew"] = auto_renew
        __props__.__dict__["expires_on"] = expires_on
        __props__.__dict__["name"] = name
        __props__.__dict__["state"] = state
        __props__.__dict__["token"] = token
        return Domain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> pulumi.Output[bool]:
        """
        Whether the DNS domain has automatic renewal enabled (boolean).
        """
        warnings.warn("""Not used, will be removed in the future""", DeprecationWarning)
        pulumi.log.warn("""auto_renew is deprecated: Not used, will be removed in the future""")

        return pulumi.get(self, "auto_renew")

    @property
    @pulumi.getter(name="expiresOn")
    def expires_on(self) -> pulumi.Output[str]:
        """
        The domain expiration date, if known.
        """
        warnings.warn("""Not used, will be removed in the future""", DeprecationWarning)
        pulumi.log.warn("""expires_on is deprecated: Not used, will be removed in the future""")

        return pulumi.get(self, "expires_on")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        ❗ The DNS domain name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The domain state.
        """
        warnings.warn("""Not used, will be removed in the future""", DeprecationWarning)
        pulumi.log.warn("""state is deprecated: Not used, will be removed in the future""")

        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        A security token that can be used as an alternative way to manage DNS domains via the Exoscale API.
        """
        warnings.warn("""Not used, will be removed in the future""", DeprecationWarning)
        pulumi.log.warn("""token is deprecated: Not used, will be removed in the future""")

        return pulumi.get(self, "token")

