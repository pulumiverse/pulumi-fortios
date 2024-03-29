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

__all__ = ['ServergroupArgs', 'Servergroup']

@pulumi.input_type
class ServergroupArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ldb_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_lists: Optional[pulumi.Input[Sequence[pulumi.Input['ServergroupServerListArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Servergroup resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] ldb_method: Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
        :param pulumi.Input[str] name: Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
        :param pulumi.Input[Sequence[pulumi.Input['ServergroupServerListArgs']]] server_lists: Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `server_list` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if ldb_method is not None:
            pulumi.set(__self__, "ldb_method", ldb_method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if server_lists is not None:
            pulumi.set(__self__, "server_lists", server_lists)
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
    @pulumi.getter(name="ldbMethod")
    def ldb_method(self) -> Optional[pulumi.Input[str]]:
        """
        Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
        """
        return pulumi.get(self, "ldb_method")

    @ldb_method.setter
    def ldb_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldb_method", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serverLists")
    def server_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServergroupServerListArgs']]]]:
        """
        Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `server_list` block is documented below.
        """
        return pulumi.get(self, "server_lists")

    @server_lists.setter
    def server_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServergroupServerListArgs']]]]):
        pulumi.set(self, "server_lists", value)

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
class _ServergroupState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ldb_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_lists: Optional[pulumi.Input[Sequence[pulumi.Input['ServergroupServerListArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Servergroup resources.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] ldb_method: Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
        :param pulumi.Input[str] name: Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
        :param pulumi.Input[Sequence[pulumi.Input['ServergroupServerListArgs']]] server_lists: Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `server_list` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if ldb_method is not None:
            pulumi.set(__self__, "ldb_method", ldb_method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if server_lists is not None:
            pulumi.set(__self__, "server_lists", server_lists)
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
    @pulumi.getter(name="ldbMethod")
    def ldb_method(self) -> Optional[pulumi.Input[str]]:
        """
        Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
        """
        return pulumi.get(self, "ldb_method")

    @ldb_method.setter
    def ldb_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldb_method", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serverLists")
    def server_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServergroupServerListArgs']]]]:
        """
        Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `server_list` block is documented below.
        """
        return pulumi.get(self, "server_lists")

    @server_lists.setter
    def server_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServergroupServerListArgs']]]]):
        pulumi.set(self, "server_lists", value)

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


class Servergroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ldb_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServergroupServerListArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure an ICAP server group consisting of multiple forward servers. Supports failover and load balancing. Applies to FortiOS Version `>= 7.2.0`.

        ## Import

        Icap ServerGroup can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:icap/servergroup:Servergroup labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:icap/servergroup:Servergroup labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] ldb_method: Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
        :param pulumi.Input[str] name: Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServergroupServerListArgs']]]] server_lists: Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `server_list` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ServergroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure an ICAP server group consisting of multiple forward servers. Supports failover and load balancing. Applies to FortiOS Version `>= 7.2.0`.

        ## Import

        Icap ServerGroup can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:icap/servergroup:Servergroup labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:icap/servergroup:Servergroup labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param ServergroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServergroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ldb_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServergroupServerListArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServergroupArgs.__new__(ServergroupArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["ldb_method"] = ldb_method
            __props__.__dict__["name"] = name
            __props__.__dict__["server_lists"] = server_lists
            __props__.__dict__["vdomparam"] = vdomparam
        super(Servergroup, __self__).__init__(
            'fortios:icap/servergroup:Servergroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            ldb_method: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            server_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServergroupServerListArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Servergroup':
        """
        Get an existing Servergroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] ldb_method: Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
        :param pulumi.Input[str] name: Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServergroupServerListArgs']]]] server_lists: Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `server_list` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServergroupState.__new__(_ServergroupState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["ldb_method"] = ldb_method
        __props__.__dict__["name"] = name
        __props__.__dict__["server_lists"] = server_lists
        __props__.__dict__["vdomparam"] = vdomparam
        return Servergroup(resource_name, opts=opts, __props__=__props__)

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
        Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        """
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter(name="ldbMethod")
    def ldb_method(self) -> pulumi.Output[str]:
        """
        Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
        """
        return pulumi.get(self, "ldb_method")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serverLists")
    def server_lists(self) -> pulumi.Output[Optional[Sequence['outputs.ServergroupServerList']]]:
        """
        Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `server_list` block is documented below.
        """
        return pulumi.get(self, "server_lists")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

