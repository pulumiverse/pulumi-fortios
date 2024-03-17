# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['Ippool6Args', 'Ippool6']

@pulumi.input_type
class Ippool6Args:
    def __init__(__self__, *,
                 endip: pulumi.Input[str],
                 startip: pulumi.Input[str],
                 add_nat46_route: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat46: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Ippool6 resource.
        :param pulumi.Input[str] endip: Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        :param pulumi.Input[str] startip: First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        :param pulumi.Input[str] add_nat46_route: Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] name: IPv6 IP pool name.
        :param pulumi.Input[str] nat46: Enable/disable NAT46. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "endip", endip)
        pulumi.set(__self__, "startip", startip)
        if add_nat46_route is not None:
            pulumi.set(__self__, "add_nat46_route", add_nat46_route)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nat46 is not None:
            pulumi.set(__self__, "nat46", nat46)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def endip(self) -> pulumi.Input[str]:
        """
        Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        """
        return pulumi.get(self, "endip")

    @endip.setter
    def endip(self, value: pulumi.Input[str]):
        pulumi.set(self, "endip", value)

    @property
    @pulumi.getter
    def startip(self) -> pulumi.Input[str]:
        """
        First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        """
        return pulumi.get(self, "startip")

    @startip.setter
    def startip(self, value: pulumi.Input[str]):
        pulumi.set(self, "startip", value)

    @property
    @pulumi.getter(name="addNat46Route")
    def add_nat46_route(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "add_nat46_route")

    @add_nat46_route.setter
    def add_nat46_route(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "add_nat46_route", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 IP pool name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def nat46(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable NAT46. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "nat46")

    @nat46.setter
    def nat46(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat46", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _Ippool6State:
    def __init__(__self__, *,
                 add_nat46_route: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 endip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat46: Optional[pulumi.Input[str]] = None,
                 startip: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Ippool6 resources.
        :param pulumi.Input[str] add_nat46_route: Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] endip: Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        :param pulumi.Input[str] name: IPv6 IP pool name.
        :param pulumi.Input[str] nat46: Enable/disable NAT46. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] startip: First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if add_nat46_route is not None:
            pulumi.set(__self__, "add_nat46_route", add_nat46_route)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if endip is not None:
            pulumi.set(__self__, "endip", endip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nat46 is not None:
            pulumi.set(__self__, "nat46", nat46)
        if startip is not None:
            pulumi.set(__self__, "startip", startip)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="addNat46Route")
    def add_nat46_route(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "add_nat46_route")

    @add_nat46_route.setter
    def add_nat46_route(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "add_nat46_route", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter
    def endip(self) -> Optional[pulumi.Input[str]]:
        """
        Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        """
        return pulumi.get(self, "endip")

    @endip.setter
    def endip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 IP pool name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def nat46(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable NAT46. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "nat46")

    @nat46.setter
    def nat46(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat46", value)

    @property
    @pulumi.getter
    def startip(self) -> Optional[pulumi.Input[str]]:
        """
        First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        """
        return pulumi.get(self, "startip")

    @startip.setter
    def startip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "startip", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class Ippool6(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_nat46_route: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 endip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat46: Optional[pulumi.Input[str]] = None,
                 startip: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure IPv6 IP pools.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.firewall.Ippool6("trname",
            endip="2001:3ca1:10f:1a:121b::19",
            startip="2001:3ca1:10f:1a:121b::10")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Firewall Ippool6 can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:firewall/ippool6:Ippool6 labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:firewall/ippool6:Ippool6 labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] add_nat46_route: Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] endip: Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        :param pulumi.Input[str] name: IPv6 IP pool name.
        :param pulumi.Input[str] nat46: Enable/disable NAT46. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] startip: First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Ippool6Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure IPv6 IP pools.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.firewall.Ippool6("trname",
            endip="2001:3ca1:10f:1a:121b::19",
            startip="2001:3ca1:10f:1a:121b::10")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Firewall Ippool6 can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:firewall/ippool6:Ippool6 labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:firewall/ippool6:Ippool6 labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param Ippool6Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Ippool6Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_nat46_route: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 endip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat46: Optional[pulumi.Input[str]] = None,
                 startip: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = Ippool6Args.__new__(Ippool6Args)

            __props__.__dict__["add_nat46_route"] = add_nat46_route
            __props__.__dict__["comments"] = comments
            if endip is None and not opts.urn:
                raise TypeError("Missing required property 'endip'")
            __props__.__dict__["endip"] = endip
            __props__.__dict__["name"] = name
            __props__.__dict__["nat46"] = nat46
            if startip is None and not opts.urn:
                raise TypeError("Missing required property 'startip'")
            __props__.__dict__["startip"] = startip
            __props__.__dict__["vdomparam"] = vdomparam
        super(Ippool6, __self__).__init__(
            'fortios:firewall/ippool6:Ippool6',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            add_nat46_route: Optional[pulumi.Input[str]] = None,
            comments: Optional[pulumi.Input[str]] = None,
            endip: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            nat46: Optional[pulumi.Input[str]] = None,
            startip: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Ippool6':
        """
        Get an existing Ippool6 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] add_nat46_route: Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] endip: Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        :param pulumi.Input[str] name: IPv6 IP pool name.
        :param pulumi.Input[str] nat46: Enable/disable NAT46. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] startip: First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _Ippool6State.__new__(_Ippool6State)

        __props__.__dict__["add_nat46_route"] = add_nat46_route
        __props__.__dict__["comments"] = comments
        __props__.__dict__["endip"] = endip
        __props__.__dict__["name"] = name
        __props__.__dict__["nat46"] = nat46
        __props__.__dict__["startip"] = startip
        __props__.__dict__["vdomparam"] = vdomparam
        return Ippool6(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addNat46Route")
    def add_nat46_route(self) -> pulumi.Output[str]:
        """
        Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "add_nat46_route")

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def endip(self) -> pulumi.Output[str]:
        """
        Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        """
        return pulumi.get(self, "endip")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        IPv6 IP pool name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def nat46(self) -> pulumi.Output[str]:
        """
        Enable/disable NAT46. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "nat46")

    @property
    @pulumi.getter
    def startip(self) -> pulumi.Output[str]:
        """
        First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        """
        return pulumi.get(self, "startip")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

