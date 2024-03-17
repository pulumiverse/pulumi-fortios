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

__all__ = ['SwitchinterfaceArgs', 'Switchinterface']

@pulumi.input_type
class SwitchinterfaceArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 intra_switch_policy: Optional[pulumi.Input[str]] = None,
                 mac_ttl: Optional[pulumi.Input[int]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchinterfaceMemberArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 span: Optional[pulumi.Input[str]] = None,
                 span_dest_port: Optional[pulumi.Input[str]] = None,
                 span_direction: Optional[pulumi.Input[str]] = None,
                 span_source_ports: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchinterfaceSpanSourcePortArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Switchinterface resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] intra_switch_policy: Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
        :param pulumi.Input[int] mac_ttl: Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
        :param pulumi.Input[Sequence[pulumi.Input['SwitchinterfaceMemberArgs']]] members: Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
        :param pulumi.Input[str] name: Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
        :param pulumi.Input[str] span: Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] span_dest_port: SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
        :param pulumi.Input[str] span_direction: The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
        :param pulumi.Input[Sequence[pulumi.Input['SwitchinterfaceSpanSourcePortArgs']]] span_source_ports: Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `span_source_port` block is documented below.
        :param pulumi.Input[str] type: Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
        :param pulumi.Input[str] vdom: VDOM that the software switch belongs to.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if intra_switch_policy is not None:
            pulumi.set(__self__, "intra_switch_policy", intra_switch_policy)
        if mac_ttl is not None:
            pulumi.set(__self__, "mac_ttl", mac_ttl)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if span is not None:
            pulumi.set(__self__, "span", span)
        if span_dest_port is not None:
            pulumi.set(__self__, "span_dest_port", span_dest_port)
        if span_direction is not None:
            pulumi.set(__self__, "span_direction", span_direction)
        if span_source_ports is not None:
            pulumi.set(__self__, "span_source_ports", span_source_ports)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vdom is not None:
            pulumi.set(__self__, "vdom", vdom)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

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
    @pulumi.getter(name="intraSwitchPolicy")
    def intra_switch_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
        """
        return pulumi.get(self, "intra_switch_policy")

    @intra_switch_policy.setter
    def intra_switch_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "intra_switch_policy", value)

    @property
    @pulumi.getter(name="macTtl")
    def mac_ttl(self) -> Optional[pulumi.Input[int]]:
        """
        Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
        """
        return pulumi.get(self, "mac_ttl")

    @mac_ttl.setter
    def mac_ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mac_ttl", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchinterfaceMemberArgs']]]]:
        """
        Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchinterfaceMemberArgs']]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def span(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "span")

    @span.setter
    def span(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "span", value)

    @property
    @pulumi.getter(name="spanDestPort")
    def span_dest_port(self) -> Optional[pulumi.Input[str]]:
        """
        SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
        """
        return pulumi.get(self, "span_dest_port")

    @span_dest_port.setter
    def span_dest_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "span_dest_port", value)

    @property
    @pulumi.getter(name="spanDirection")
    def span_direction(self) -> Optional[pulumi.Input[str]]:
        """
        The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
        """
        return pulumi.get(self, "span_direction")

    @span_direction.setter
    def span_direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "span_direction", value)

    @property
    @pulumi.getter(name="spanSourcePorts")
    def span_source_ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchinterfaceSpanSourcePortArgs']]]]:
        """
        Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `span_source_port` block is documented below.
        """
        return pulumi.get(self, "span_source_ports")

    @span_source_ports.setter
    def span_source_ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchinterfaceSpanSourcePortArgs']]]]):
        pulumi.set(self, "span_source_ports", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def vdom(self) -> Optional[pulumi.Input[str]]:
        """
        VDOM that the software switch belongs to.
        """
        return pulumi.get(self, "vdom")

    @vdom.setter
    def vdom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdom", value)

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
class _SwitchinterfaceState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 intra_switch_policy: Optional[pulumi.Input[str]] = None,
                 mac_ttl: Optional[pulumi.Input[int]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchinterfaceMemberArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 span: Optional[pulumi.Input[str]] = None,
                 span_dest_port: Optional[pulumi.Input[str]] = None,
                 span_direction: Optional[pulumi.Input[str]] = None,
                 span_source_ports: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchinterfaceSpanSourcePortArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Switchinterface resources.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] intra_switch_policy: Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
        :param pulumi.Input[int] mac_ttl: Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
        :param pulumi.Input[Sequence[pulumi.Input['SwitchinterfaceMemberArgs']]] members: Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
        :param pulumi.Input[str] name: Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
        :param pulumi.Input[str] span: Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] span_dest_port: SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
        :param pulumi.Input[str] span_direction: The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
        :param pulumi.Input[Sequence[pulumi.Input['SwitchinterfaceSpanSourcePortArgs']]] span_source_ports: Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `span_source_port` block is documented below.
        :param pulumi.Input[str] type: Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
        :param pulumi.Input[str] vdom: VDOM that the software switch belongs to.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if intra_switch_policy is not None:
            pulumi.set(__self__, "intra_switch_policy", intra_switch_policy)
        if mac_ttl is not None:
            pulumi.set(__self__, "mac_ttl", mac_ttl)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if span is not None:
            pulumi.set(__self__, "span", span)
        if span_dest_port is not None:
            pulumi.set(__self__, "span_dest_port", span_dest_port)
        if span_direction is not None:
            pulumi.set(__self__, "span_direction", span_direction)
        if span_source_ports is not None:
            pulumi.set(__self__, "span_source_ports", span_source_ports)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vdom is not None:
            pulumi.set(__self__, "vdom", vdom)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

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
    @pulumi.getter(name="intraSwitchPolicy")
    def intra_switch_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
        """
        return pulumi.get(self, "intra_switch_policy")

    @intra_switch_policy.setter
    def intra_switch_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "intra_switch_policy", value)

    @property
    @pulumi.getter(name="macTtl")
    def mac_ttl(self) -> Optional[pulumi.Input[int]]:
        """
        Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
        """
        return pulumi.get(self, "mac_ttl")

    @mac_ttl.setter
    def mac_ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mac_ttl", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchinterfaceMemberArgs']]]]:
        """
        Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchinterfaceMemberArgs']]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def span(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "span")

    @span.setter
    def span(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "span", value)

    @property
    @pulumi.getter(name="spanDestPort")
    def span_dest_port(self) -> Optional[pulumi.Input[str]]:
        """
        SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
        """
        return pulumi.get(self, "span_dest_port")

    @span_dest_port.setter
    def span_dest_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "span_dest_port", value)

    @property
    @pulumi.getter(name="spanDirection")
    def span_direction(self) -> Optional[pulumi.Input[str]]:
        """
        The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
        """
        return pulumi.get(self, "span_direction")

    @span_direction.setter
    def span_direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "span_direction", value)

    @property
    @pulumi.getter(name="spanSourcePorts")
    def span_source_ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchinterfaceSpanSourcePortArgs']]]]:
        """
        Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `span_source_port` block is documented below.
        """
        return pulumi.get(self, "span_source_ports")

    @span_source_ports.setter
    def span_source_ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchinterfaceSpanSourcePortArgs']]]]):
        pulumi.set(self, "span_source_ports", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def vdom(self) -> Optional[pulumi.Input[str]]:
        """
        VDOM that the software switch belongs to.
        """
        return pulumi.get(self, "vdom")

    @vdom.setter
    def vdom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdom", value)

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


class Switchinterface(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 intra_switch_policy: Optional[pulumi.Input[str]] = None,
                 mac_ttl: Optional[pulumi.Input[int]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchinterfaceMemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 span: Optional[pulumi.Input[str]] = None,
                 span_dest_port: Optional[pulumi.Input[str]] = None,
                 span_direction: Optional[pulumi.Input[str]] = None,
                 span_source_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchinterfaceSpanSourcePortArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure software switch interfaces by grouping physical and WiFi interfaces.

        ## Import

        System SwitchInterface can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/switchinterface:Switchinterface labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/switchinterface:Switchinterface labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] intra_switch_policy: Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
        :param pulumi.Input[int] mac_ttl: Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchinterfaceMemberArgs']]]] members: Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
        :param pulumi.Input[str] name: Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
        :param pulumi.Input[str] span: Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] span_dest_port: SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
        :param pulumi.Input[str] span_direction: The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchinterfaceSpanSourcePortArgs']]]] span_source_ports: Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `span_source_port` block is documented below.
        :param pulumi.Input[str] type: Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
        :param pulumi.Input[str] vdom: VDOM that the software switch belongs to.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SwitchinterfaceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure software switch interfaces by grouping physical and WiFi interfaces.

        ## Import

        System SwitchInterface can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/switchinterface:Switchinterface labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/switchinterface:Switchinterface labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SwitchinterfaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SwitchinterfaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 intra_switch_policy: Optional[pulumi.Input[str]] = None,
                 mac_ttl: Optional[pulumi.Input[int]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchinterfaceMemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 span: Optional[pulumi.Input[str]] = None,
                 span_dest_port: Optional[pulumi.Input[str]] = None,
                 span_direction: Optional[pulumi.Input[str]] = None,
                 span_source_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchinterfaceSpanSourcePortArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SwitchinterfaceArgs.__new__(SwitchinterfaceArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["intra_switch_policy"] = intra_switch_policy
            __props__.__dict__["mac_ttl"] = mac_ttl
            __props__.__dict__["members"] = members
            __props__.__dict__["name"] = name
            __props__.__dict__["span"] = span
            __props__.__dict__["span_dest_port"] = span_dest_port
            __props__.__dict__["span_direction"] = span_direction
            __props__.__dict__["span_source_ports"] = span_source_ports
            __props__.__dict__["type"] = type
            __props__.__dict__["vdom"] = vdom
            __props__.__dict__["vdomparam"] = vdomparam
        super(Switchinterface, __self__).__init__(
            'fortios:system/switchinterface:Switchinterface',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            intra_switch_policy: Optional[pulumi.Input[str]] = None,
            mac_ttl: Optional[pulumi.Input[int]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchinterfaceMemberArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            span: Optional[pulumi.Input[str]] = None,
            span_dest_port: Optional[pulumi.Input[str]] = None,
            span_direction: Optional[pulumi.Input[str]] = None,
            span_source_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchinterfaceSpanSourcePortArgs']]]]] = None,
            type: Optional[pulumi.Input[str]] = None,
            vdom: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Switchinterface':
        """
        Get an existing Switchinterface resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] intra_switch_policy: Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
        :param pulumi.Input[int] mac_ttl: Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchinterfaceMemberArgs']]]] members: Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
        :param pulumi.Input[str] name: Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
        :param pulumi.Input[str] span: Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] span_dest_port: SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
        :param pulumi.Input[str] span_direction: The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchinterfaceSpanSourcePortArgs']]]] span_source_ports: Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `span_source_port` block is documented below.
        :param pulumi.Input[str] type: Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
        :param pulumi.Input[str] vdom: VDOM that the software switch belongs to.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SwitchinterfaceState.__new__(_SwitchinterfaceState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["intra_switch_policy"] = intra_switch_policy
        __props__.__dict__["mac_ttl"] = mac_ttl
        __props__.__dict__["members"] = members
        __props__.__dict__["name"] = name
        __props__.__dict__["span"] = span
        __props__.__dict__["span_dest_port"] = span_dest_port
        __props__.__dict__["span_direction"] = span_direction
        __props__.__dict__["span_source_ports"] = span_source_ports
        __props__.__dict__["type"] = type
        __props__.__dict__["vdom"] = vdom
        __props__.__dict__["vdomparam"] = vdomparam
        return Switchinterface(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="intraSwitchPolicy")
    def intra_switch_policy(self) -> pulumi.Output[str]:
        """
        Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
        """
        return pulumi.get(self, "intra_switch_policy")

    @property
    @pulumi.getter(name="macTtl")
    def mac_ttl(self) -> pulumi.Output[int]:
        """
        Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
        """
        return pulumi.get(self, "mac_ttl")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Optional[Sequence['outputs.SwitchinterfaceMember']]]:
        """
        Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def span(self) -> pulumi.Output[str]:
        """
        Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "span")

    @property
    @pulumi.getter(name="spanDestPort")
    def span_dest_port(self) -> pulumi.Output[str]:
        """
        SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
        """
        return pulumi.get(self, "span_dest_port")

    @property
    @pulumi.getter(name="spanDirection")
    def span_direction(self) -> pulumi.Output[str]:
        """
        The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
        """
        return pulumi.get(self, "span_direction")

    @property
    @pulumi.getter(name="spanSourcePorts")
    def span_source_ports(self) -> pulumi.Output[Optional[Sequence['outputs.SwitchinterfaceSpanSourcePort']]]:
        """
        Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `span_source_port` block is documented below.
        """
        return pulumi.get(self, "span_source_ports")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def vdom(self) -> pulumi.Output[str]:
        """
        VDOM that the software switch belongs to.
        """
        return pulumi.get(self, "vdom")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

