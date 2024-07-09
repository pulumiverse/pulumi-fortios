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

__all__ = ['OndemandsnifferArgs', 'Ondemandsniffer']

@pulumi.input_type
class OndemandsnifferArgs:
    def __init__(__self__, *,
                 advanced_filter: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 hosts: Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferHostArgs']]]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 max_packet_count: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 non_ip_packet: Optional[pulumi.Input[str]] = None,
                 ports: Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferPortArgs']]]] = None,
                 protocols: Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferProtocolArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Ondemandsniffer resource.
        :param pulumi.Input[str] advanced_filter: Advanced freeform filter that will be used over existing filter settings if set. Can only be used by super admin.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[Sequence[pulumi.Input['OndemandsnifferHostArgs']]] hosts: IPv4 or IPv6 hosts to filter in this traffic sniffer. The structure of `hosts` block is documented below.
        :param pulumi.Input[str] interface: Interface name that on-demand packet sniffer will take place.
        :param pulumi.Input[int] max_packet_count: Maximum number of packets to capture per on-demand packet sniffer.
        :param pulumi.Input[str] name: On-demand packet sniffer name.
        :param pulumi.Input[str] non_ip_packet: Include non-IP packets. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['OndemandsnifferPortArgs']]] ports: Ports to filter for in this traffic sniffer. The structure of `ports` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['OndemandsnifferProtocolArgs']]] protocols: Protocols to filter in this traffic sniffer. The structure of `protocols` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if advanced_filter is not None:
            pulumi.set(__self__, "advanced_filter", advanced_filter)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if hosts is not None:
            pulumi.set(__self__, "hosts", hosts)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if max_packet_count is not None:
            pulumi.set(__self__, "max_packet_count", max_packet_count)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if non_ip_packet is not None:
            pulumi.set(__self__, "non_ip_packet", non_ip_packet)
        if ports is not None:
            pulumi.set(__self__, "ports", ports)
        if protocols is not None:
            pulumi.set(__self__, "protocols", protocols)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="advancedFilter")
    def advanced_filter(self) -> Optional[pulumi.Input[str]]:
        """
        Advanced freeform filter that will be used over existing filter settings if set. Can only be used by super admin.
        """
        return pulumi.get(self, "advanced_filter")

    @advanced_filter.setter
    def advanced_filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "advanced_filter", value)

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
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        """
        Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        """
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter
    def hosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferHostArgs']]]]:
        """
        IPv4 or IPv6 hosts to filter in this traffic sniffer. The structure of `hosts` block is documented below.
        """
        return pulumi.get(self, "hosts")

    @hosts.setter
    def hosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferHostArgs']]]]):
        pulumi.set(self, "hosts", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        Interface name that on-demand packet sniffer will take place.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="maxPacketCount")
    def max_packet_count(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of packets to capture per on-demand packet sniffer.
        """
        return pulumi.get(self, "max_packet_count")

    @max_packet_count.setter
    def max_packet_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_packet_count", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        On-demand packet sniffer name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nonIpPacket")
    def non_ip_packet(self) -> Optional[pulumi.Input[str]]:
        """
        Include non-IP packets. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "non_ip_packet")

    @non_ip_packet.setter
    def non_ip_packet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "non_ip_packet", value)

    @property
    @pulumi.getter
    def ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferPortArgs']]]]:
        """
        Ports to filter for in this traffic sniffer. The structure of `ports` block is documented below.
        """
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferPortArgs']]]]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter
    def protocols(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferProtocolArgs']]]]:
        """
        Protocols to filter in this traffic sniffer. The structure of `protocols` block is documented below.
        """
        return pulumi.get(self, "protocols")

    @protocols.setter
    def protocols(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferProtocolArgs']]]]):
        pulumi.set(self, "protocols", value)

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
class _OndemandsnifferState:
    def __init__(__self__, *,
                 advanced_filter: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 hosts: Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferHostArgs']]]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 max_packet_count: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 non_ip_packet: Optional[pulumi.Input[str]] = None,
                 ports: Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferPortArgs']]]] = None,
                 protocols: Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferProtocolArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Ondemandsniffer resources.
        :param pulumi.Input[str] advanced_filter: Advanced freeform filter that will be used over existing filter settings if set. Can only be used by super admin.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[Sequence[pulumi.Input['OndemandsnifferHostArgs']]] hosts: IPv4 or IPv6 hosts to filter in this traffic sniffer. The structure of `hosts` block is documented below.
        :param pulumi.Input[str] interface: Interface name that on-demand packet sniffer will take place.
        :param pulumi.Input[int] max_packet_count: Maximum number of packets to capture per on-demand packet sniffer.
        :param pulumi.Input[str] name: On-demand packet sniffer name.
        :param pulumi.Input[str] non_ip_packet: Include non-IP packets. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['OndemandsnifferPortArgs']]] ports: Ports to filter for in this traffic sniffer. The structure of `ports` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['OndemandsnifferProtocolArgs']]] protocols: Protocols to filter in this traffic sniffer. The structure of `protocols` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if advanced_filter is not None:
            pulumi.set(__self__, "advanced_filter", advanced_filter)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if hosts is not None:
            pulumi.set(__self__, "hosts", hosts)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if max_packet_count is not None:
            pulumi.set(__self__, "max_packet_count", max_packet_count)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if non_ip_packet is not None:
            pulumi.set(__self__, "non_ip_packet", non_ip_packet)
        if ports is not None:
            pulumi.set(__self__, "ports", ports)
        if protocols is not None:
            pulumi.set(__self__, "protocols", protocols)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="advancedFilter")
    def advanced_filter(self) -> Optional[pulumi.Input[str]]:
        """
        Advanced freeform filter that will be used over existing filter settings if set. Can only be used by super admin.
        """
        return pulumi.get(self, "advanced_filter")

    @advanced_filter.setter
    def advanced_filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "advanced_filter", value)

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
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        """
        Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        """
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter
    def hosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferHostArgs']]]]:
        """
        IPv4 or IPv6 hosts to filter in this traffic sniffer. The structure of `hosts` block is documented below.
        """
        return pulumi.get(self, "hosts")

    @hosts.setter
    def hosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferHostArgs']]]]):
        pulumi.set(self, "hosts", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        Interface name that on-demand packet sniffer will take place.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="maxPacketCount")
    def max_packet_count(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of packets to capture per on-demand packet sniffer.
        """
        return pulumi.get(self, "max_packet_count")

    @max_packet_count.setter
    def max_packet_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_packet_count", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        On-demand packet sniffer name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nonIpPacket")
    def non_ip_packet(self) -> Optional[pulumi.Input[str]]:
        """
        Include non-IP packets. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "non_ip_packet")

    @non_ip_packet.setter
    def non_ip_packet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "non_ip_packet", value)

    @property
    @pulumi.getter
    def ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferPortArgs']]]]:
        """
        Ports to filter for in this traffic sniffer. The structure of `ports` block is documented below.
        """
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferPortArgs']]]]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter
    def protocols(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferProtocolArgs']]]]:
        """
        Protocols to filter in this traffic sniffer. The structure of `protocols` block is documented below.
        """
        return pulumi.get(self, "protocols")

    @protocols.setter
    def protocols(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OndemandsnifferProtocolArgs']]]]):
        pulumi.set(self, "protocols", value)

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


class Ondemandsniffer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 advanced_filter: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 hosts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OndemandsnifferHostArgs']]]]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 max_packet_count: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 non_ip_packet: Optional[pulumi.Input[str]] = None,
                 ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OndemandsnifferPortArgs']]]]] = None,
                 protocols: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OndemandsnifferProtocolArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure on-demand packet sniffer. Applies to FortiOS Version `>= 7.4.4`.

        ## Import

        Firewall OnDemandSniffer can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:firewall/ondemandsniffer:Ondemandsniffer labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:firewall/ondemandsniffer:Ondemandsniffer labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] advanced_filter: Advanced freeform filter that will be used over existing filter settings if set. Can only be used by super admin.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OndemandsnifferHostArgs']]]] hosts: IPv4 or IPv6 hosts to filter in this traffic sniffer. The structure of `hosts` block is documented below.
        :param pulumi.Input[str] interface: Interface name that on-demand packet sniffer will take place.
        :param pulumi.Input[int] max_packet_count: Maximum number of packets to capture per on-demand packet sniffer.
        :param pulumi.Input[str] name: On-demand packet sniffer name.
        :param pulumi.Input[str] non_ip_packet: Include non-IP packets. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OndemandsnifferPortArgs']]]] ports: Ports to filter for in this traffic sniffer. The structure of `ports` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OndemandsnifferProtocolArgs']]]] protocols: Protocols to filter in this traffic sniffer. The structure of `protocols` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[OndemandsnifferArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure on-demand packet sniffer. Applies to FortiOS Version `>= 7.4.4`.

        ## Import

        Firewall OnDemandSniffer can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:firewall/ondemandsniffer:Ondemandsniffer labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:firewall/ondemandsniffer:Ondemandsniffer labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param OndemandsnifferArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OndemandsnifferArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 advanced_filter: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 hosts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OndemandsnifferHostArgs']]]]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 max_packet_count: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 non_ip_packet: Optional[pulumi.Input[str]] = None,
                 ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OndemandsnifferPortArgs']]]]] = None,
                 protocols: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OndemandsnifferProtocolArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OndemandsnifferArgs.__new__(OndemandsnifferArgs)

            __props__.__dict__["advanced_filter"] = advanced_filter
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["hosts"] = hosts
            __props__.__dict__["interface"] = interface
            __props__.__dict__["max_packet_count"] = max_packet_count
            __props__.__dict__["name"] = name
            __props__.__dict__["non_ip_packet"] = non_ip_packet
            __props__.__dict__["ports"] = ports
            __props__.__dict__["protocols"] = protocols
            __props__.__dict__["vdomparam"] = vdomparam
        super(Ondemandsniffer, __self__).__init__(
            'fortios:firewall/ondemandsniffer:Ondemandsniffer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            advanced_filter: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            hosts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OndemandsnifferHostArgs']]]]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            max_packet_count: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            non_ip_packet: Optional[pulumi.Input[str]] = None,
            ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OndemandsnifferPortArgs']]]]] = None,
            protocols: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OndemandsnifferProtocolArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Ondemandsniffer':
        """
        Get an existing Ondemandsniffer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] advanced_filter: Advanced freeform filter that will be used over existing filter settings if set. Can only be used by super admin.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OndemandsnifferHostArgs']]]] hosts: IPv4 or IPv6 hosts to filter in this traffic sniffer. The structure of `hosts` block is documented below.
        :param pulumi.Input[str] interface: Interface name that on-demand packet sniffer will take place.
        :param pulumi.Input[int] max_packet_count: Maximum number of packets to capture per on-demand packet sniffer.
        :param pulumi.Input[str] name: On-demand packet sniffer name.
        :param pulumi.Input[str] non_ip_packet: Include non-IP packets. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OndemandsnifferPortArgs']]]] ports: Ports to filter for in this traffic sniffer. The structure of `ports` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OndemandsnifferProtocolArgs']]]] protocols: Protocols to filter in this traffic sniffer. The structure of `protocols` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OndemandsnifferState.__new__(_OndemandsnifferState)

        __props__.__dict__["advanced_filter"] = advanced_filter
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["hosts"] = hosts
        __props__.__dict__["interface"] = interface
        __props__.__dict__["max_packet_count"] = max_packet_count
        __props__.__dict__["name"] = name
        __props__.__dict__["non_ip_packet"] = non_ip_packet
        __props__.__dict__["ports"] = ports
        __props__.__dict__["protocols"] = protocols
        __props__.__dict__["vdomparam"] = vdomparam
        return Ondemandsniffer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="advancedFilter")
    def advanced_filter(self) -> pulumi.Output[Optional[str]]:
        """
        Advanced freeform filter that will be used over existing filter settings if set. Can only be used by super admin.
        """
        return pulumi.get(self, "advanced_filter")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        """
        Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        """
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter
    def hosts(self) -> pulumi.Output[Optional[Sequence['outputs.OndemandsnifferHost']]]:
        """
        IPv4 or IPv6 hosts to filter in this traffic sniffer. The structure of `hosts` block is documented below.
        """
        return pulumi.get(self, "hosts")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[str]:
        """
        Interface name that on-demand packet sniffer will take place.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="maxPacketCount")
    def max_packet_count(self) -> pulumi.Output[int]:
        """
        Maximum number of packets to capture per on-demand packet sniffer.
        """
        return pulumi.get(self, "max_packet_count")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        On-demand packet sniffer name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nonIpPacket")
    def non_ip_packet(self) -> pulumi.Output[str]:
        """
        Include non-IP packets. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "non_ip_packet")

    @property
    @pulumi.getter
    def ports(self) -> pulumi.Output[Optional[Sequence['outputs.OndemandsnifferPort']]]:
        """
        Ports to filter for in this traffic sniffer. The structure of `ports` block is documented below.
        """
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter
    def protocols(self) -> pulumi.Output[Optional[Sequence['outputs.OndemandsnifferProtocol']]]:
        """
        Protocols to filter in this traffic sniffer. The structure of `protocols` block is documented below.
        """
        return pulumi.get(self, "protocols")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[str]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
