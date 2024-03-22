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

__all__ = ['SnmpArgs', 'Snmp']

@pulumi.input_type
class SnmpArgs:
    def __init__(__self__, *,
                 communities: Optional[pulumi.Input[Sequence[pulumi.Input['SnmpCommunityArgs']]]] = None,
                 contact_info: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 engine_id: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 trap_high_cpu_threshold: Optional[pulumi.Input[int]] = None,
                 trap_high_mem_threshold: Optional[pulumi.Input[int]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input['SnmpUserArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Snmp resource.
        :param pulumi.Input[Sequence[pulumi.Input['SnmpCommunityArgs']]] communities: SNMP Community Configuration. The structure of `community` block is documented below.
        :param pulumi.Input[str] contact_info: Contact Information.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] engine_id: AC SNMP engineId string (maximum 24 characters).
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[int] trap_high_cpu_threshold: CPU usage when trap is sent.
        :param pulumi.Input[int] trap_high_mem_threshold: Memory usage when trap is sent.
        :param pulumi.Input[Sequence[pulumi.Input['SnmpUserArgs']]] users: SNMP User Configuration. The structure of `user` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if communities is not None:
            pulumi.set(__self__, "communities", communities)
        if contact_info is not None:
            pulumi.set(__self__, "contact_info", contact_info)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if engine_id is not None:
            pulumi.set(__self__, "engine_id", engine_id)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if trap_high_cpu_threshold is not None:
            pulumi.set(__self__, "trap_high_cpu_threshold", trap_high_cpu_threshold)
        if trap_high_mem_threshold is not None:
            pulumi.set(__self__, "trap_high_mem_threshold", trap_high_mem_threshold)
        if users is not None:
            pulumi.set(__self__, "users", users)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def communities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SnmpCommunityArgs']]]]:
        """
        SNMP Community Configuration. The structure of `community` block is documented below.
        """
        return pulumi.get(self, "communities")

    @communities.setter
    def communities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SnmpCommunityArgs']]]]):
        pulumi.set(self, "communities", value)

    @property
    @pulumi.getter(name="contactInfo")
    def contact_info(self) -> Optional[pulumi.Input[str]]:
        """
        Contact Information.
        """
        return pulumi.get(self, "contact_info")

    @contact_info.setter
    def contact_info(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "contact_info", value)

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
    @pulumi.getter(name="engineId")
    def engine_id(self) -> Optional[pulumi.Input[str]]:
        """
        AC SNMP engineId string (maximum 24 characters).
        """
        return pulumi.get(self, "engine_id")

    @engine_id.setter
    def engine_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_id", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        """
        Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        """
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter(name="trapHighCpuThreshold")
    def trap_high_cpu_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        CPU usage when trap is sent.
        """
        return pulumi.get(self, "trap_high_cpu_threshold")

    @trap_high_cpu_threshold.setter
    def trap_high_cpu_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "trap_high_cpu_threshold", value)

    @property
    @pulumi.getter(name="trapHighMemThreshold")
    def trap_high_mem_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Memory usage when trap is sent.
        """
        return pulumi.get(self, "trap_high_mem_threshold")

    @trap_high_mem_threshold.setter
    def trap_high_mem_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "trap_high_mem_threshold", value)

    @property
    @pulumi.getter
    def users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SnmpUserArgs']]]]:
        """
        SNMP User Configuration. The structure of `user` block is documented below.
        """
        return pulumi.get(self, "users")

    @users.setter
    def users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SnmpUserArgs']]]]):
        pulumi.set(self, "users", value)

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
class _SnmpState:
    def __init__(__self__, *,
                 communities: Optional[pulumi.Input[Sequence[pulumi.Input['SnmpCommunityArgs']]]] = None,
                 contact_info: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 engine_id: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 trap_high_cpu_threshold: Optional[pulumi.Input[int]] = None,
                 trap_high_mem_threshold: Optional[pulumi.Input[int]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input['SnmpUserArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Snmp resources.
        :param pulumi.Input[Sequence[pulumi.Input['SnmpCommunityArgs']]] communities: SNMP Community Configuration. The structure of `community` block is documented below.
        :param pulumi.Input[str] contact_info: Contact Information.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] engine_id: AC SNMP engineId string (maximum 24 characters).
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[int] trap_high_cpu_threshold: CPU usage when trap is sent.
        :param pulumi.Input[int] trap_high_mem_threshold: Memory usage when trap is sent.
        :param pulumi.Input[Sequence[pulumi.Input['SnmpUserArgs']]] users: SNMP User Configuration. The structure of `user` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if communities is not None:
            pulumi.set(__self__, "communities", communities)
        if contact_info is not None:
            pulumi.set(__self__, "contact_info", contact_info)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if engine_id is not None:
            pulumi.set(__self__, "engine_id", engine_id)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if trap_high_cpu_threshold is not None:
            pulumi.set(__self__, "trap_high_cpu_threshold", trap_high_cpu_threshold)
        if trap_high_mem_threshold is not None:
            pulumi.set(__self__, "trap_high_mem_threshold", trap_high_mem_threshold)
        if users is not None:
            pulumi.set(__self__, "users", users)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def communities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SnmpCommunityArgs']]]]:
        """
        SNMP Community Configuration. The structure of `community` block is documented below.
        """
        return pulumi.get(self, "communities")

    @communities.setter
    def communities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SnmpCommunityArgs']]]]):
        pulumi.set(self, "communities", value)

    @property
    @pulumi.getter(name="contactInfo")
    def contact_info(self) -> Optional[pulumi.Input[str]]:
        """
        Contact Information.
        """
        return pulumi.get(self, "contact_info")

    @contact_info.setter
    def contact_info(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "contact_info", value)

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
    @pulumi.getter(name="engineId")
    def engine_id(self) -> Optional[pulumi.Input[str]]:
        """
        AC SNMP engineId string (maximum 24 characters).
        """
        return pulumi.get(self, "engine_id")

    @engine_id.setter
    def engine_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_id", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        """
        Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        """
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter(name="trapHighCpuThreshold")
    def trap_high_cpu_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        CPU usage when trap is sent.
        """
        return pulumi.get(self, "trap_high_cpu_threshold")

    @trap_high_cpu_threshold.setter
    def trap_high_cpu_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "trap_high_cpu_threshold", value)

    @property
    @pulumi.getter(name="trapHighMemThreshold")
    def trap_high_mem_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Memory usage when trap is sent.
        """
        return pulumi.get(self, "trap_high_mem_threshold")

    @trap_high_mem_threshold.setter
    def trap_high_mem_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "trap_high_mem_threshold", value)

    @property
    @pulumi.getter
    def users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SnmpUserArgs']]]]:
        """
        SNMP User Configuration. The structure of `user` block is documented below.
        """
        return pulumi.get(self, "users")

    @users.setter
    def users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SnmpUserArgs']]]]):
        pulumi.set(self, "users", value)

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


class Snmp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 communities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnmpCommunityArgs']]]]] = None,
                 contact_info: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 engine_id: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 trap_high_cpu_threshold: Optional[pulumi.Input[int]] = None,
                 trap_high_mem_threshold: Optional[pulumi.Input[int]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnmpUserArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure SNMP. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        WirelessController Snmp can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:wirelesscontroller/snmp:Snmp labelname WirelessControllerSnmp
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:wirelesscontroller/snmp:Snmp labelname WirelessControllerSnmp
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnmpCommunityArgs']]]] communities: SNMP Community Configuration. The structure of `community` block is documented below.
        :param pulumi.Input[str] contact_info: Contact Information.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] engine_id: AC SNMP engineId string (maximum 24 characters).
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[int] trap_high_cpu_threshold: CPU usage when trap is sent.
        :param pulumi.Input[int] trap_high_mem_threshold: Memory usage when trap is sent.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnmpUserArgs']]]] users: SNMP User Configuration. The structure of `user` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SnmpArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure SNMP. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        WirelessController Snmp can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:wirelesscontroller/snmp:Snmp labelname WirelessControllerSnmp
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:wirelesscontroller/snmp:Snmp labelname WirelessControllerSnmp
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SnmpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnmpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 communities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnmpCommunityArgs']]]]] = None,
                 contact_info: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 engine_id: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 trap_high_cpu_threshold: Optional[pulumi.Input[int]] = None,
                 trap_high_mem_threshold: Optional[pulumi.Input[int]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnmpUserArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SnmpArgs.__new__(SnmpArgs)

            __props__.__dict__["communities"] = communities
            __props__.__dict__["contact_info"] = contact_info
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["engine_id"] = engine_id
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["trap_high_cpu_threshold"] = trap_high_cpu_threshold
            __props__.__dict__["trap_high_mem_threshold"] = trap_high_mem_threshold
            __props__.__dict__["users"] = users
            __props__.__dict__["vdomparam"] = vdomparam
        super(Snmp, __self__).__init__(
            'fortios:wirelesscontroller/snmp:Snmp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            communities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnmpCommunityArgs']]]]] = None,
            contact_info: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            engine_id: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            trap_high_cpu_threshold: Optional[pulumi.Input[int]] = None,
            trap_high_mem_threshold: Optional[pulumi.Input[int]] = None,
            users: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnmpUserArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Snmp':
        """
        Get an existing Snmp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnmpCommunityArgs']]]] communities: SNMP Community Configuration. The structure of `community` block is documented below.
        :param pulumi.Input[str] contact_info: Contact Information.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] engine_id: AC SNMP engineId string (maximum 24 characters).
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[int] trap_high_cpu_threshold: CPU usage when trap is sent.
        :param pulumi.Input[int] trap_high_mem_threshold: Memory usage when trap is sent.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnmpUserArgs']]]] users: SNMP User Configuration. The structure of `user` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnmpState.__new__(_SnmpState)

        __props__.__dict__["communities"] = communities
        __props__.__dict__["contact_info"] = contact_info
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["engine_id"] = engine_id
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["trap_high_cpu_threshold"] = trap_high_cpu_threshold
        __props__.__dict__["trap_high_mem_threshold"] = trap_high_mem_threshold
        __props__.__dict__["users"] = users
        __props__.__dict__["vdomparam"] = vdomparam
        return Snmp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def communities(self) -> pulumi.Output[Optional[Sequence['outputs.SnmpCommunity']]]:
        """
        SNMP Community Configuration. The structure of `community` block is documented below.
        """
        return pulumi.get(self, "communities")

    @property
    @pulumi.getter(name="contactInfo")
    def contact_info(self) -> pulumi.Output[str]:
        """
        Contact Information.
        """
        return pulumi.get(self, "contact_info")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="engineId")
    def engine_id(self) -> pulumi.Output[str]:
        """
        AC SNMP engineId string (maximum 24 characters).
        """
        return pulumi.get(self, "engine_id")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        """
        Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        """
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter(name="trapHighCpuThreshold")
    def trap_high_cpu_threshold(self) -> pulumi.Output[int]:
        """
        CPU usage when trap is sent.
        """
        return pulumi.get(self, "trap_high_cpu_threshold")

    @property
    @pulumi.getter(name="trapHighMemThreshold")
    def trap_high_mem_threshold(self) -> pulumi.Output[int]:
        """
        Memory usage when trap is sent.
        """
        return pulumi.get(self, "trap_high_mem_threshold")

    @property
    @pulumi.getter
    def users(self) -> pulumi.Output[Optional[Sequence['outputs.SnmpUser']]]:
        """
        SNMP User Configuration. The structure of `user` block is documented below.
        """
        return pulumi.get(self, "users")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

