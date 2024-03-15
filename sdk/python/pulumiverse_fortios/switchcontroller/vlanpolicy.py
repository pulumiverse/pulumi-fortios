# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['VlanpolicyArgs', 'Vlanpolicy']

@pulumi.input_type
class VlanpolicyArgs:
    def __init__(__self__, *,
                 allowed_vlans: Optional[pulumi.Input[Sequence[pulumi.Input['VlanpolicyAllowedVlanArgs']]]] = None,
                 allowed_vlans_all: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 discard_mode: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fortilink: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 untagged_vlans: Optional[pulumi.Input[Sequence[pulumi.Input['VlanpolicyUntaggedVlanArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Vlanpolicy resource.
        :param pulumi.Input[Sequence[pulumi.Input['VlanpolicyAllowedVlanArgs']]] allowed_vlans: Allowed VLANs to be applied when using this VLAN policy. The structure of `allowed_vlans` block is documented below.
        :param pulumi.Input[str] allowed_vlans_all: Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] description: Description for the VLAN policy.
        :param pulumi.Input[str] discard_mode: Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] fortilink: FortiLink interface for which this VLAN policy belongs to.
        :param pulumi.Input[str] name: VLAN policy name.
        :param pulumi.Input[Sequence[pulumi.Input['VlanpolicyUntaggedVlanArgs']]] untagged_vlans: Untagged VLANs to be applied when using this VLAN policy. The structure of `untagged_vlans` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vlan: Native VLAN to be applied when using this VLAN policy.
        """
        if allowed_vlans is not None:
            pulumi.set(__self__, "allowed_vlans", allowed_vlans)
        if allowed_vlans_all is not None:
            pulumi.set(__self__, "allowed_vlans_all", allowed_vlans_all)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if discard_mode is not None:
            pulumi.set(__self__, "discard_mode", discard_mode)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if fortilink is not None:
            pulumi.set(__self__, "fortilink", fortilink)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if untagged_vlans is not None:
            pulumi.set(__self__, "untagged_vlans", untagged_vlans)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vlan is not None:
            pulumi.set(__self__, "vlan", vlan)

    @property
    @pulumi.getter(name="allowedVlans")
    def allowed_vlans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VlanpolicyAllowedVlanArgs']]]]:
        """
        Allowed VLANs to be applied when using this VLAN policy. The structure of `allowed_vlans` block is documented below.
        """
        return pulumi.get(self, "allowed_vlans")

    @allowed_vlans.setter
    def allowed_vlans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VlanpolicyAllowedVlanArgs']]]]):
        pulumi.set(self, "allowed_vlans", value)

    @property
    @pulumi.getter(name="allowedVlansAll")
    def allowed_vlans_all(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "allowed_vlans_all")

    @allowed_vlans_all.setter
    def allowed_vlans_all(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allowed_vlans_all", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description for the VLAN policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="discardMode")
    def discard_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
        """
        return pulumi.get(self, "discard_mode")

    @discard_mode.setter
    def discard_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "discard_mode", value)

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
    def fortilink(self) -> Optional[pulumi.Input[str]]:
        """
        FortiLink interface for which this VLAN policy belongs to.
        """
        return pulumi.get(self, "fortilink")

    @fortilink.setter
    def fortilink(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fortilink", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        VLAN policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="untaggedVlans")
    def untagged_vlans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VlanpolicyUntaggedVlanArgs']]]]:
        """
        Untagged VLANs to be applied when using this VLAN policy. The structure of `untagged_vlans` block is documented below.
        """
        return pulumi.get(self, "untagged_vlans")

    @untagged_vlans.setter
    def untagged_vlans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VlanpolicyUntaggedVlanArgs']]]]):
        pulumi.set(self, "untagged_vlans", value)

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
    def vlan(self) -> Optional[pulumi.Input[str]]:
        """
        Native VLAN to be applied when using this VLAN policy.
        """
        return pulumi.get(self, "vlan")

    @vlan.setter
    def vlan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vlan", value)


@pulumi.input_type
class _VlanpolicyState:
    def __init__(__self__, *,
                 allowed_vlans: Optional[pulumi.Input[Sequence[pulumi.Input['VlanpolicyAllowedVlanArgs']]]] = None,
                 allowed_vlans_all: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 discard_mode: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fortilink: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 untagged_vlans: Optional[pulumi.Input[Sequence[pulumi.Input['VlanpolicyUntaggedVlanArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Vlanpolicy resources.
        :param pulumi.Input[Sequence[pulumi.Input['VlanpolicyAllowedVlanArgs']]] allowed_vlans: Allowed VLANs to be applied when using this VLAN policy. The structure of `allowed_vlans` block is documented below.
        :param pulumi.Input[str] allowed_vlans_all: Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] description: Description for the VLAN policy.
        :param pulumi.Input[str] discard_mode: Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] fortilink: FortiLink interface for which this VLAN policy belongs to.
        :param pulumi.Input[str] name: VLAN policy name.
        :param pulumi.Input[Sequence[pulumi.Input['VlanpolicyUntaggedVlanArgs']]] untagged_vlans: Untagged VLANs to be applied when using this VLAN policy. The structure of `untagged_vlans` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vlan: Native VLAN to be applied when using this VLAN policy.
        """
        if allowed_vlans is not None:
            pulumi.set(__self__, "allowed_vlans", allowed_vlans)
        if allowed_vlans_all is not None:
            pulumi.set(__self__, "allowed_vlans_all", allowed_vlans_all)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if discard_mode is not None:
            pulumi.set(__self__, "discard_mode", discard_mode)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if fortilink is not None:
            pulumi.set(__self__, "fortilink", fortilink)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if untagged_vlans is not None:
            pulumi.set(__self__, "untagged_vlans", untagged_vlans)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vlan is not None:
            pulumi.set(__self__, "vlan", vlan)

    @property
    @pulumi.getter(name="allowedVlans")
    def allowed_vlans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VlanpolicyAllowedVlanArgs']]]]:
        """
        Allowed VLANs to be applied when using this VLAN policy. The structure of `allowed_vlans` block is documented below.
        """
        return pulumi.get(self, "allowed_vlans")

    @allowed_vlans.setter
    def allowed_vlans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VlanpolicyAllowedVlanArgs']]]]):
        pulumi.set(self, "allowed_vlans", value)

    @property
    @pulumi.getter(name="allowedVlansAll")
    def allowed_vlans_all(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "allowed_vlans_all")

    @allowed_vlans_all.setter
    def allowed_vlans_all(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allowed_vlans_all", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description for the VLAN policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="discardMode")
    def discard_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
        """
        return pulumi.get(self, "discard_mode")

    @discard_mode.setter
    def discard_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "discard_mode", value)

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
    def fortilink(self) -> Optional[pulumi.Input[str]]:
        """
        FortiLink interface for which this VLAN policy belongs to.
        """
        return pulumi.get(self, "fortilink")

    @fortilink.setter
    def fortilink(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fortilink", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        VLAN policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="untaggedVlans")
    def untagged_vlans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VlanpolicyUntaggedVlanArgs']]]]:
        """
        Untagged VLANs to be applied when using this VLAN policy. The structure of `untagged_vlans` block is documented below.
        """
        return pulumi.get(self, "untagged_vlans")

    @untagged_vlans.setter
    def untagged_vlans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VlanpolicyUntaggedVlanArgs']]]]):
        pulumi.set(self, "untagged_vlans", value)

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
    def vlan(self) -> Optional[pulumi.Input[str]]:
        """
        Native VLAN to be applied when using this VLAN policy.
        """
        return pulumi.get(self, "vlan")

    @vlan.setter
    def vlan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vlan", value)


class Vlanpolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_vlans: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VlanpolicyAllowedVlanArgs']]]]] = None,
                 allowed_vlans_all: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 discard_mode: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fortilink: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 untagged_vlans: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VlanpolicyUntaggedVlanArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure VLAN policy to be applied on the managed FortiSwitch ports through port-policy. Applies to FortiOS Version `>= 6.4.0`.

        ## Import

        SwitchController VlanPolicy can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:switchcontroller/vlanpolicy:Vlanpolicy labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:switchcontroller/vlanpolicy:Vlanpolicy labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VlanpolicyAllowedVlanArgs']]]] allowed_vlans: Allowed VLANs to be applied when using this VLAN policy. The structure of `allowed_vlans` block is documented below.
        :param pulumi.Input[str] allowed_vlans_all: Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] description: Description for the VLAN policy.
        :param pulumi.Input[str] discard_mode: Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] fortilink: FortiLink interface for which this VLAN policy belongs to.
        :param pulumi.Input[str] name: VLAN policy name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VlanpolicyUntaggedVlanArgs']]]] untagged_vlans: Untagged VLANs to be applied when using this VLAN policy. The structure of `untagged_vlans` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vlan: Native VLAN to be applied when using this VLAN policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[VlanpolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure VLAN policy to be applied on the managed FortiSwitch ports through port-policy. Applies to FortiOS Version `>= 6.4.0`.

        ## Import

        SwitchController VlanPolicy can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:switchcontroller/vlanpolicy:Vlanpolicy labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:switchcontroller/vlanpolicy:Vlanpolicy labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param VlanpolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VlanpolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_vlans: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VlanpolicyAllowedVlanArgs']]]]] = None,
                 allowed_vlans_all: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 discard_mode: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fortilink: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 untagged_vlans: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VlanpolicyUntaggedVlanArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VlanpolicyArgs.__new__(VlanpolicyArgs)

            __props__.__dict__["allowed_vlans"] = allowed_vlans
            __props__.__dict__["allowed_vlans_all"] = allowed_vlans_all
            __props__.__dict__["description"] = description
            __props__.__dict__["discard_mode"] = discard_mode
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["fortilink"] = fortilink
            __props__.__dict__["name"] = name
            __props__.__dict__["untagged_vlans"] = untagged_vlans
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["vlan"] = vlan
        super(Vlanpolicy, __self__).__init__(
            'fortios:switchcontroller/vlanpolicy:Vlanpolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allowed_vlans: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VlanpolicyAllowedVlanArgs']]]]] = None,
            allowed_vlans_all: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            discard_mode: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            fortilink: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            untagged_vlans: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VlanpolicyUntaggedVlanArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            vlan: Optional[pulumi.Input[str]] = None) -> 'Vlanpolicy':
        """
        Get an existing Vlanpolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VlanpolicyAllowedVlanArgs']]]] allowed_vlans: Allowed VLANs to be applied when using this VLAN policy. The structure of `allowed_vlans` block is documented below.
        :param pulumi.Input[str] allowed_vlans_all: Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] description: Description for the VLAN policy.
        :param pulumi.Input[str] discard_mode: Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] fortilink: FortiLink interface for which this VLAN policy belongs to.
        :param pulumi.Input[str] name: VLAN policy name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VlanpolicyUntaggedVlanArgs']]]] untagged_vlans: Untagged VLANs to be applied when using this VLAN policy. The structure of `untagged_vlans` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vlan: Native VLAN to be applied when using this VLAN policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VlanpolicyState.__new__(_VlanpolicyState)

        __props__.__dict__["allowed_vlans"] = allowed_vlans
        __props__.__dict__["allowed_vlans_all"] = allowed_vlans_all
        __props__.__dict__["description"] = description
        __props__.__dict__["discard_mode"] = discard_mode
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["fortilink"] = fortilink
        __props__.__dict__["name"] = name
        __props__.__dict__["untagged_vlans"] = untagged_vlans
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["vlan"] = vlan
        return Vlanpolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedVlans")
    def allowed_vlans(self) -> pulumi.Output[Optional[Sequence['outputs.VlanpolicyAllowedVlan']]]:
        """
        Allowed VLANs to be applied when using this VLAN policy. The structure of `allowed_vlans` block is documented below.
        """
        return pulumi.get(self, "allowed_vlans")

    @property
    @pulumi.getter(name="allowedVlansAll")
    def allowed_vlans_all(self) -> pulumi.Output[str]:
        """
        Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "allowed_vlans_all")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description for the VLAN policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="discardMode")
    def discard_mode(self) -> pulumi.Output[str]:
        """
        Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
        """
        return pulumi.get(self, "discard_mode")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def fortilink(self) -> pulumi.Output[str]:
        """
        FortiLink interface for which this VLAN policy belongs to.
        """
        return pulumi.get(self, "fortilink")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        VLAN policy name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="untaggedVlans")
    def untagged_vlans(self) -> pulumi.Output[Optional[Sequence['outputs.VlanpolicyUntaggedVlan']]]:
        """
        Untagged VLANs to be applied when using this VLAN policy. The structure of `untagged_vlans` block is documented below.
        """
        return pulumi.get(self, "untagged_vlans")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vlan(self) -> pulumi.Output[str]:
        """
        Native VLAN to be applied when using this VLAN policy.
        """
        return pulumi.get(self, "vlan")

