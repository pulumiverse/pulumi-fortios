# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VdomSettingArgs', 'VdomSetting']

@pulumi.input_type
class VdomSettingArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 short_name: Optional[pulumi.Input[str]] = None,
                 temporary: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VdomSetting resource.
        :param pulumi.Input[str] name: VDOM name.
        :param pulumi.Input[str] short_name: VDOM short name.
        :param pulumi.Input[str] temporary: Temporary.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if short_name is not None:
            pulumi.set(__self__, "short_name", short_name)
        if temporary is not None:
            pulumi.set(__self__, "temporary", temporary)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        VDOM name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> Optional[pulumi.Input[str]]:
        """
        VDOM short name.
        """
        return pulumi.get(self, "short_name")

    @short_name.setter
    def short_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "short_name", value)

    @property
    @pulumi.getter
    def temporary(self) -> Optional[pulumi.Input[str]]:
        """
        Temporary.
        """
        return pulumi.get(self, "temporary")

    @temporary.setter
    def temporary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "temporary", value)


@pulumi.input_type
class _VdomSettingState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 short_name: Optional[pulumi.Input[str]] = None,
                 temporary: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VdomSetting resources.
        :param pulumi.Input[str] name: VDOM name.
        :param pulumi.Input[str] short_name: VDOM short name.
        :param pulumi.Input[str] temporary: Temporary.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if short_name is not None:
            pulumi.set(__self__, "short_name", short_name)
        if temporary is not None:
            pulumi.set(__self__, "temporary", temporary)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        VDOM name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> Optional[pulumi.Input[str]]:
        """
        VDOM short name.
        """
        return pulumi.get(self, "short_name")

    @short_name.setter
    def short_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "short_name", value)

    @property
    @pulumi.getter
    def temporary(self) -> Optional[pulumi.Input[str]]:
        """
        Temporary.
        """
        return pulumi.get(self, "temporary")

    @temporary.setter
    def temporary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "temporary", value)


class VdomSetting(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 short_name: Optional[pulumi.Input[str]] = None,
                 temporary: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to configure VDOM of FortiOS. The API user of the token for this feature should have a super admin profile, It can be set in CLI while GUI does not allow.

        !> **Warning:** The resource will be deprecated and replaced by new resource `system.Vdom`, we recommend that you use the new resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        test2 = fortios.system.VdomSetting("test2",
            short_name="aa1122",
            temporary="0")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: VDOM name.
        :param pulumi.Input[str] short_name: VDOM short name.
        :param pulumi.Input[str] temporary: Temporary.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[VdomSettingArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to configure VDOM of FortiOS. The API user of the token for this feature should have a super admin profile, It can be set in CLI while GUI does not allow.

        !> **Warning:** The resource will be deprecated and replaced by new resource `system.Vdom`, we recommend that you use the new resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        test2 = fortios.system.VdomSetting("test2",
            short_name="aa1122",
            temporary="0")
        ```

        :param str resource_name: The name of the resource.
        :param VdomSettingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VdomSettingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 short_name: Optional[pulumi.Input[str]] = None,
                 temporary: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VdomSettingArgs.__new__(VdomSettingArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["short_name"] = short_name
            __props__.__dict__["temporary"] = temporary
        super(VdomSetting, __self__).__init__(
            'fortios:system/vdomSetting:VdomSetting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            short_name: Optional[pulumi.Input[str]] = None,
            temporary: Optional[pulumi.Input[str]] = None) -> 'VdomSetting':
        """
        Get an existing VdomSetting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: VDOM name.
        :param pulumi.Input[str] short_name: VDOM short name.
        :param pulumi.Input[str] temporary: Temporary.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VdomSettingState.__new__(_VdomSettingState)

        __props__.__dict__["name"] = name
        __props__.__dict__["short_name"] = short_name
        __props__.__dict__["temporary"] = temporary
        return VdomSetting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        VDOM name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> pulumi.Output[str]:
        """
        VDOM short name.
        """
        return pulumi.get(self, "short_name")

    @property
    @pulumi.getter
    def temporary(self) -> pulumi.Output[str]:
        """
        Temporary.
        """
        return pulumi.get(self, "temporary")

