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

__all__ = ['StandaloneclusterArgs', 'Standalonecluster']

@pulumi.input_type
class StandaloneclusterArgs:
    def __init__(__self__, *,
                 asymmetric_traffic_control: Optional[pulumi.Input[str]] = None,
                 cluster_peers: Optional[pulumi.Input[Sequence[pulumi.Input['StandaloneclusterClusterPeerArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 encryption: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 group_member_id: Optional[pulumi.Input[int]] = None,
                 layer2_connection: Optional[pulumi.Input[str]] = None,
                 psksecret: Optional[pulumi.Input[str]] = None,
                 session_sync_dev: Optional[pulumi.Input[str]] = None,
                 standalone_group_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Standalonecluster resource.
        :param pulumi.Input[str] asymmetric_traffic_control: Asymmetric traffic control mode. Valid values: `cps-preferred`, `strict-anti-replay`.
        :param pulumi.Input[Sequence[pulumi.Input['StandaloneclusterClusterPeerArgs']]] cluster_peers: Configure FortiGate Session Life Support Protocol (FGSP) session synchronization. The structure of `cluster_peer` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] encryption: Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[int] group_member_id: Cluster member ID. On FortiOS versions 6.4.0: 0 - 3. On FortiOS versions >= 6.4.1: 0 - 15.
        :param pulumi.Input[str] layer2_connection: Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
        :param pulumi.Input[str] psksecret: Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
        :param pulumi.Input[str] session_sync_dev: Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
        :param pulumi.Input[int] standalone_group_id: Cluster group ID (0 - 255). Must be the same for all members.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if asymmetric_traffic_control is not None:
            pulumi.set(__self__, "asymmetric_traffic_control", asymmetric_traffic_control)
        if cluster_peers is not None:
            pulumi.set(__self__, "cluster_peers", cluster_peers)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if encryption is not None:
            pulumi.set(__self__, "encryption", encryption)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if group_member_id is not None:
            pulumi.set(__self__, "group_member_id", group_member_id)
        if layer2_connection is not None:
            pulumi.set(__self__, "layer2_connection", layer2_connection)
        if psksecret is not None:
            pulumi.set(__self__, "psksecret", psksecret)
        if session_sync_dev is not None:
            pulumi.set(__self__, "session_sync_dev", session_sync_dev)
        if standalone_group_id is not None:
            pulumi.set(__self__, "standalone_group_id", standalone_group_id)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="asymmetricTrafficControl")
    def asymmetric_traffic_control(self) -> Optional[pulumi.Input[str]]:
        """
        Asymmetric traffic control mode. Valid values: `cps-preferred`, `strict-anti-replay`.
        """
        return pulumi.get(self, "asymmetric_traffic_control")

    @asymmetric_traffic_control.setter
    def asymmetric_traffic_control(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "asymmetric_traffic_control", value)

    @property
    @pulumi.getter(name="clusterPeers")
    def cluster_peers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StandaloneclusterClusterPeerArgs']]]]:
        """
        Configure FortiGate Session Life Support Protocol (FGSP) session synchronization. The structure of `cluster_peer` block is documented below.
        """
        return pulumi.get(self, "cluster_peers")

    @cluster_peers.setter
    def cluster_peers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StandaloneclusterClusterPeerArgs']]]]):
        pulumi.set(self, "cluster_peers", value)

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
    def encryption(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "encryption")

    @encryption.setter
    def encryption(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption", value)

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
    @pulumi.getter(name="groupMemberId")
    def group_member_id(self) -> Optional[pulumi.Input[int]]:
        """
        Cluster member ID. On FortiOS versions 6.4.0: 0 - 3. On FortiOS versions >= 6.4.1: 0 - 15.
        """
        return pulumi.get(self, "group_member_id")

    @group_member_id.setter
    def group_member_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "group_member_id", value)

    @property
    @pulumi.getter(name="layer2Connection")
    def layer2_connection(self) -> Optional[pulumi.Input[str]]:
        """
        Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
        """
        return pulumi.get(self, "layer2_connection")

    @layer2_connection.setter
    def layer2_connection(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "layer2_connection", value)

    @property
    @pulumi.getter
    def psksecret(self) -> Optional[pulumi.Input[str]]:
        """
        Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
        """
        return pulumi.get(self, "psksecret")

    @psksecret.setter
    def psksecret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "psksecret", value)

    @property
    @pulumi.getter(name="sessionSyncDev")
    def session_sync_dev(self) -> Optional[pulumi.Input[str]]:
        """
        Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
        """
        return pulumi.get(self, "session_sync_dev")

    @session_sync_dev.setter
    def session_sync_dev(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session_sync_dev", value)

    @property
    @pulumi.getter(name="standaloneGroupId")
    def standalone_group_id(self) -> Optional[pulumi.Input[int]]:
        """
        Cluster group ID (0 - 255). Must be the same for all members.
        """
        return pulumi.get(self, "standalone_group_id")

    @standalone_group_id.setter
    def standalone_group_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "standalone_group_id", value)

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
class _StandaloneclusterState:
    def __init__(__self__, *,
                 asymmetric_traffic_control: Optional[pulumi.Input[str]] = None,
                 cluster_peers: Optional[pulumi.Input[Sequence[pulumi.Input['StandaloneclusterClusterPeerArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 encryption: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 group_member_id: Optional[pulumi.Input[int]] = None,
                 layer2_connection: Optional[pulumi.Input[str]] = None,
                 psksecret: Optional[pulumi.Input[str]] = None,
                 session_sync_dev: Optional[pulumi.Input[str]] = None,
                 standalone_group_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Standalonecluster resources.
        :param pulumi.Input[str] asymmetric_traffic_control: Asymmetric traffic control mode. Valid values: `cps-preferred`, `strict-anti-replay`.
        :param pulumi.Input[Sequence[pulumi.Input['StandaloneclusterClusterPeerArgs']]] cluster_peers: Configure FortiGate Session Life Support Protocol (FGSP) session synchronization. The structure of `cluster_peer` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] encryption: Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[int] group_member_id: Cluster member ID. On FortiOS versions 6.4.0: 0 - 3. On FortiOS versions >= 6.4.1: 0 - 15.
        :param pulumi.Input[str] layer2_connection: Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
        :param pulumi.Input[str] psksecret: Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
        :param pulumi.Input[str] session_sync_dev: Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
        :param pulumi.Input[int] standalone_group_id: Cluster group ID (0 - 255). Must be the same for all members.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if asymmetric_traffic_control is not None:
            pulumi.set(__self__, "asymmetric_traffic_control", asymmetric_traffic_control)
        if cluster_peers is not None:
            pulumi.set(__self__, "cluster_peers", cluster_peers)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if encryption is not None:
            pulumi.set(__self__, "encryption", encryption)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if group_member_id is not None:
            pulumi.set(__self__, "group_member_id", group_member_id)
        if layer2_connection is not None:
            pulumi.set(__self__, "layer2_connection", layer2_connection)
        if psksecret is not None:
            pulumi.set(__self__, "psksecret", psksecret)
        if session_sync_dev is not None:
            pulumi.set(__self__, "session_sync_dev", session_sync_dev)
        if standalone_group_id is not None:
            pulumi.set(__self__, "standalone_group_id", standalone_group_id)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="asymmetricTrafficControl")
    def asymmetric_traffic_control(self) -> Optional[pulumi.Input[str]]:
        """
        Asymmetric traffic control mode. Valid values: `cps-preferred`, `strict-anti-replay`.
        """
        return pulumi.get(self, "asymmetric_traffic_control")

    @asymmetric_traffic_control.setter
    def asymmetric_traffic_control(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "asymmetric_traffic_control", value)

    @property
    @pulumi.getter(name="clusterPeers")
    def cluster_peers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StandaloneclusterClusterPeerArgs']]]]:
        """
        Configure FortiGate Session Life Support Protocol (FGSP) session synchronization. The structure of `cluster_peer` block is documented below.
        """
        return pulumi.get(self, "cluster_peers")

    @cluster_peers.setter
    def cluster_peers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StandaloneclusterClusterPeerArgs']]]]):
        pulumi.set(self, "cluster_peers", value)

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
    def encryption(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "encryption")

    @encryption.setter
    def encryption(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption", value)

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
    @pulumi.getter(name="groupMemberId")
    def group_member_id(self) -> Optional[pulumi.Input[int]]:
        """
        Cluster member ID. On FortiOS versions 6.4.0: 0 - 3. On FortiOS versions >= 6.4.1: 0 - 15.
        """
        return pulumi.get(self, "group_member_id")

    @group_member_id.setter
    def group_member_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "group_member_id", value)

    @property
    @pulumi.getter(name="layer2Connection")
    def layer2_connection(self) -> Optional[pulumi.Input[str]]:
        """
        Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
        """
        return pulumi.get(self, "layer2_connection")

    @layer2_connection.setter
    def layer2_connection(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "layer2_connection", value)

    @property
    @pulumi.getter
    def psksecret(self) -> Optional[pulumi.Input[str]]:
        """
        Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
        """
        return pulumi.get(self, "psksecret")

    @psksecret.setter
    def psksecret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "psksecret", value)

    @property
    @pulumi.getter(name="sessionSyncDev")
    def session_sync_dev(self) -> Optional[pulumi.Input[str]]:
        """
        Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
        """
        return pulumi.get(self, "session_sync_dev")

    @session_sync_dev.setter
    def session_sync_dev(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session_sync_dev", value)

    @property
    @pulumi.getter(name="standaloneGroupId")
    def standalone_group_id(self) -> Optional[pulumi.Input[int]]:
        """
        Cluster group ID (0 - 255). Must be the same for all members.
        """
        return pulumi.get(self, "standalone_group_id")

    @standalone_group_id.setter
    def standalone_group_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "standalone_group_id", value)

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


class Standalonecluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asymmetric_traffic_control: Optional[pulumi.Input[str]] = None,
                 cluster_peers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StandaloneclusterClusterPeerArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 encryption: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 group_member_id: Optional[pulumi.Input[int]] = None,
                 layer2_connection: Optional[pulumi.Input[str]] = None,
                 psksecret: Optional[pulumi.Input[str]] = None,
                 session_sync_dev: Optional[pulumi.Input[str]] = None,
                 standalone_group_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes. Applies to FortiOS Version `>= 6.4.0`.

        ## Import

        System StandaloneCluster can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/standalonecluster:Standalonecluster labelname SystemStandaloneCluster
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/standalonecluster:Standalonecluster labelname SystemStandaloneCluster
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] asymmetric_traffic_control: Asymmetric traffic control mode. Valid values: `cps-preferred`, `strict-anti-replay`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StandaloneclusterClusterPeerArgs']]]] cluster_peers: Configure FortiGate Session Life Support Protocol (FGSP) session synchronization. The structure of `cluster_peer` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] encryption: Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[int] group_member_id: Cluster member ID. On FortiOS versions 6.4.0: 0 - 3. On FortiOS versions >= 6.4.1: 0 - 15.
        :param pulumi.Input[str] layer2_connection: Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
        :param pulumi.Input[str] psksecret: Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
        :param pulumi.Input[str] session_sync_dev: Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
        :param pulumi.Input[int] standalone_group_id: Cluster group ID (0 - 255). Must be the same for all members.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[StandaloneclusterArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes. Applies to FortiOS Version `>= 6.4.0`.

        ## Import

        System StandaloneCluster can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/standalonecluster:Standalonecluster labelname SystemStandaloneCluster
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/standalonecluster:Standalonecluster labelname SystemStandaloneCluster
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param StandaloneclusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StandaloneclusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asymmetric_traffic_control: Optional[pulumi.Input[str]] = None,
                 cluster_peers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StandaloneclusterClusterPeerArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 encryption: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 group_member_id: Optional[pulumi.Input[int]] = None,
                 layer2_connection: Optional[pulumi.Input[str]] = None,
                 psksecret: Optional[pulumi.Input[str]] = None,
                 session_sync_dev: Optional[pulumi.Input[str]] = None,
                 standalone_group_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StandaloneclusterArgs.__new__(StandaloneclusterArgs)

            __props__.__dict__["asymmetric_traffic_control"] = asymmetric_traffic_control
            __props__.__dict__["cluster_peers"] = cluster_peers
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["encryption"] = encryption
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["group_member_id"] = group_member_id
            __props__.__dict__["layer2_connection"] = layer2_connection
            __props__.__dict__["psksecret"] = None if psksecret is None else pulumi.Output.secret(psksecret)
            __props__.__dict__["session_sync_dev"] = session_sync_dev
            __props__.__dict__["standalone_group_id"] = standalone_group_id
            __props__.__dict__["vdomparam"] = vdomparam
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["psksecret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Standalonecluster, __self__).__init__(
            'fortios:system/standalonecluster:Standalonecluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            asymmetric_traffic_control: Optional[pulumi.Input[str]] = None,
            cluster_peers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StandaloneclusterClusterPeerArgs']]]]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            encryption: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            group_member_id: Optional[pulumi.Input[int]] = None,
            layer2_connection: Optional[pulumi.Input[str]] = None,
            psksecret: Optional[pulumi.Input[str]] = None,
            session_sync_dev: Optional[pulumi.Input[str]] = None,
            standalone_group_id: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Standalonecluster':
        """
        Get an existing Standalonecluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] asymmetric_traffic_control: Asymmetric traffic control mode. Valid values: `cps-preferred`, `strict-anti-replay`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StandaloneclusterClusterPeerArgs']]]] cluster_peers: Configure FortiGate Session Life Support Protocol (FGSP) session synchronization. The structure of `cluster_peer` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] encryption: Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[int] group_member_id: Cluster member ID. On FortiOS versions 6.4.0: 0 - 3. On FortiOS versions >= 6.4.1: 0 - 15.
        :param pulumi.Input[str] layer2_connection: Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
        :param pulumi.Input[str] psksecret: Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
        :param pulumi.Input[str] session_sync_dev: Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
        :param pulumi.Input[int] standalone_group_id: Cluster group ID (0 - 255). Must be the same for all members.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StandaloneclusterState.__new__(_StandaloneclusterState)

        __props__.__dict__["asymmetric_traffic_control"] = asymmetric_traffic_control
        __props__.__dict__["cluster_peers"] = cluster_peers
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["encryption"] = encryption
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["group_member_id"] = group_member_id
        __props__.__dict__["layer2_connection"] = layer2_connection
        __props__.__dict__["psksecret"] = psksecret
        __props__.__dict__["session_sync_dev"] = session_sync_dev
        __props__.__dict__["standalone_group_id"] = standalone_group_id
        __props__.__dict__["vdomparam"] = vdomparam
        return Standalonecluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="asymmetricTrafficControl")
    def asymmetric_traffic_control(self) -> pulumi.Output[str]:
        """
        Asymmetric traffic control mode. Valid values: `cps-preferred`, `strict-anti-replay`.
        """
        return pulumi.get(self, "asymmetric_traffic_control")

    @property
    @pulumi.getter(name="clusterPeers")
    def cluster_peers(self) -> pulumi.Output[Optional[Sequence['outputs.StandaloneclusterClusterPeer']]]:
        """
        Configure FortiGate Session Life Support Protocol (FGSP) session synchronization. The structure of `cluster_peer` block is documented below.
        """
        return pulumi.get(self, "cluster_peers")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def encryption(self) -> pulumi.Output[str]:
        """
        Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        """
        Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        """
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter(name="groupMemberId")
    def group_member_id(self) -> pulumi.Output[int]:
        """
        Cluster member ID. On FortiOS versions 6.4.0: 0 - 3. On FortiOS versions >= 6.4.1: 0 - 15.
        """
        return pulumi.get(self, "group_member_id")

    @property
    @pulumi.getter(name="layer2Connection")
    def layer2_connection(self) -> pulumi.Output[str]:
        """
        Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
        """
        return pulumi.get(self, "layer2_connection")

    @property
    @pulumi.getter
    def psksecret(self) -> pulumi.Output[Optional[str]]:
        """
        Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
        """
        return pulumi.get(self, "psksecret")

    @property
    @pulumi.getter(name="sessionSyncDev")
    def session_sync_dev(self) -> pulumi.Output[str]:
        """
        Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
        """
        return pulumi.get(self, "session_sync_dev")

    @property
    @pulumi.getter(name="standaloneGroupId")
    def standalone_group_id(self) -> pulumi.Output[int]:
        """
        Cluster group ID (0 - 255). Must be the same for all members.
        """
        return pulumi.get(self, "standalone_group_id")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[str]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

