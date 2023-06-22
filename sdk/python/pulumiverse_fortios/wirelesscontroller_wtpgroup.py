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

__all__ = ['WirelesscontrollerWtpgroupArgs', 'WirelesscontrollerWtpgroup']

@pulumi.input_type
class WirelesscontrollerWtpgroupArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wtps: Optional[pulumi.Input[Sequence[pulumi.Input['WirelesscontrollerWtpgroupWtpArgs']]]] = None):
        """
        The set of arguments for constructing a WirelesscontrollerWtpgroup resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: WTP group name.
        :param pulumi.Input[str] platform_type: FortiAP models to define the WTP group platform type.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[Sequence[pulumi.Input['WirelesscontrollerWtpgroupWtpArgs']]] wtps: WTP list. The structure of `wtps` block is documented below.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if platform_type is not None:
            pulumi.set(__self__, "platform_type", platform_type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if wtps is not None:
            pulumi.set(__self__, "wtps", wtps)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        WTP group name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="platformType")
    def platform_type(self) -> Optional[pulumi.Input[str]]:
        """
        FortiAP models to define the WTP group platform type.
        """
        return pulumi.get(self, "platform_type")

    @platform_type.setter
    def platform_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "platform_type", value)

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

    @property
    @pulumi.getter
    def wtps(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelesscontrollerWtpgroupWtpArgs']]]]:
        """
        WTP list. The structure of `wtps` block is documented below.
        """
        return pulumi.get(self, "wtps")

    @wtps.setter
    def wtps(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelesscontrollerWtpgroupWtpArgs']]]]):
        pulumi.set(self, "wtps", value)


@pulumi.input_type
class _WirelesscontrollerWtpgroupState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wtps: Optional[pulumi.Input[Sequence[pulumi.Input['WirelesscontrollerWtpgroupWtpArgs']]]] = None):
        """
        Input properties used for looking up and filtering WirelesscontrollerWtpgroup resources.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: WTP group name.
        :param pulumi.Input[str] platform_type: FortiAP models to define the WTP group platform type.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[Sequence[pulumi.Input['WirelesscontrollerWtpgroupWtpArgs']]] wtps: WTP list. The structure of `wtps` block is documented below.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if platform_type is not None:
            pulumi.set(__self__, "platform_type", platform_type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if wtps is not None:
            pulumi.set(__self__, "wtps", wtps)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        WTP group name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="platformType")
    def platform_type(self) -> Optional[pulumi.Input[str]]:
        """
        FortiAP models to define the WTP group platform type.
        """
        return pulumi.get(self, "platform_type")

    @platform_type.setter
    def platform_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "platform_type", value)

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

    @property
    @pulumi.getter
    def wtps(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelesscontrollerWtpgroupWtpArgs']]]]:
        """
        WTP list. The structure of `wtps` block is documented below.
        """
        return pulumi.get(self, "wtps")

    @wtps.setter
    def wtps(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelesscontrollerWtpgroupWtpArgs']]]]):
        pulumi.set(self, "wtps", value)


class WirelesscontrollerWtpgroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wtps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelesscontrollerWtpgroupWtpArgs']]]]] = None,
                 __props__=None):
        """
        Configure WTP groups.

        ## Import

        WirelessController WtpGroup can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/wirelesscontrollerWtpgroup:WirelesscontrollerWtpgroup labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/wirelesscontrollerWtpgroup:WirelesscontrollerWtpgroup labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: WTP group name.
        :param pulumi.Input[str] platform_type: FortiAP models to define the WTP group platform type.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelesscontrollerWtpgroupWtpArgs']]]] wtps: WTP list. The structure of `wtps` block is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WirelesscontrollerWtpgroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure WTP groups.

        ## Import

        WirelessController WtpGroup can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/wirelesscontrollerWtpgroup:WirelesscontrollerWtpgroup labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/wirelesscontrollerWtpgroup:WirelesscontrollerWtpgroup labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WirelesscontrollerWtpgroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WirelesscontrollerWtpgroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wtps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelesscontrollerWtpgroupWtpArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WirelesscontrollerWtpgroupArgs.__new__(WirelesscontrollerWtpgroupArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["name"] = name
            __props__.__dict__["platform_type"] = platform_type
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["wtps"] = wtps
        super(WirelesscontrollerWtpgroup, __self__).__init__(
            'fortios:index/wirelesscontrollerWtpgroup:WirelesscontrollerWtpgroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            platform_type: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            wtps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelesscontrollerWtpgroupWtpArgs']]]]] = None) -> 'WirelesscontrollerWtpgroup':
        """
        Get an existing WirelesscontrollerWtpgroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: WTP group name.
        :param pulumi.Input[str] platform_type: FortiAP models to define the WTP group platform type.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelesscontrollerWtpgroupWtpArgs']]]] wtps: WTP list. The structure of `wtps` block is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WirelesscontrollerWtpgroupState.__new__(_WirelesscontrollerWtpgroupState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["name"] = name
        __props__.__dict__["platform_type"] = platform_type
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["wtps"] = wtps
        return WirelesscontrollerWtpgroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        WTP group name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="platformType")
    def platform_type(self) -> pulumi.Output[str]:
        """
        FortiAP models to define the WTP group platform type.
        """
        return pulumi.get(self, "platform_type")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def wtps(self) -> pulumi.Output[Optional[Sequence['outputs.WirelesscontrollerWtpgroupWtp']]]:
        """
        WTP list. The structure of `wtps` block is documented below.
        """
        return pulumi.get(self, "wtps")

