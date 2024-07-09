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

__all__ = ['OtvpArgs', 'Otvp']

@pulumi.input_type
class OtvpArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 application: Optional[pulumi.Input[str]] = None,
                 date: Optional[pulumi.Input[int]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 log_packet: Optional[pulumi.Input[str]] = None,
                 metadatas: Optional[pulumi.Input[Sequence[pulumi.Input['OtvpMetadataArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 os: Optional[pulumi.Input[str]] = None,
                 rev: Optional[pulumi.Input[int]] = None,
                 rule_id: Optional[pulumi.Input[int]] = None,
                 service: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Otvp resource.
        :param pulumi.Input[str] action: Action. Valid values: `pass`, `block`.
        :param pulumi.Input[str] application: Vulnerable applications.
        :param pulumi.Input[int] date: Date.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] group: Group.
        :param pulumi.Input[str] location: Vulnerable location.
        :param pulumi.Input[str] log: Enable/disable logging. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] log_packet: Enable/disable packet logging. Valid values: `disable`, `enable`.
        :param pulumi.Input[Sequence[pulumi.Input['OtvpMetadataArgs']]] metadatas: Meta data. The structure of `metadata` block is documented below.
        :param pulumi.Input[str] name: Rule name.
        :param pulumi.Input[str] os: Vulnerable operation systems.
        :param pulumi.Input[int] rev: Revision.
        :param pulumi.Input[int] rule_id: Rule ID.
        :param pulumi.Input[str] service: Vulnerable service.
        :param pulumi.Input[str] severity: Severity.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if application is not None:
            pulumi.set(__self__, "application", application)
        if date is not None:
            pulumi.set(__self__, "date", date)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if group is not None:
            pulumi.set(__self__, "group", group)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if log is not None:
            pulumi.set(__self__, "log", log)
        if log_packet is not None:
            pulumi.set(__self__, "log_packet", log_packet)
        if metadatas is not None:
            pulumi.set(__self__, "metadatas", metadatas)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if os is not None:
            pulumi.set(__self__, "os", os)
        if rev is not None:
            pulumi.set(__self__, "rev", rev)
        if rule_id is not None:
            pulumi.set(__self__, "rule_id", rule_id)
        if service is not None:
            pulumi.set(__self__, "service", service)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Action. Valid values: `pass`, `block`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def application(self) -> Optional[pulumi.Input[str]]:
        """
        Vulnerable applications.
        """
        return pulumi.get(self, "application")

    @application.setter
    def application(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application", value)

    @property
    @pulumi.getter
    def date(self) -> Optional[pulumi.Input[int]]:
        """
        Date.
        """
        return pulumi.get(self, "date")

    @date.setter
    def date(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "date", value)

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
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        Group.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Vulnerable location.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def log(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable logging. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "log")

    @log.setter
    def log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log", value)

    @property
    @pulumi.getter(name="logPacket")
    def log_packet(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable packet logging. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "log_packet")

    @log_packet.setter
    def log_packet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_packet", value)

    @property
    @pulumi.getter
    def metadatas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OtvpMetadataArgs']]]]:
        """
        Meta data. The structure of `metadata` block is documented below.
        """
        return pulumi.get(self, "metadatas")

    @metadatas.setter
    def metadatas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OtvpMetadataArgs']]]]):
        pulumi.set(self, "metadatas", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Rule name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def os(self) -> Optional[pulumi.Input[str]]:
        """
        Vulnerable operation systems.
        """
        return pulumi.get(self, "os")

    @os.setter
    def os(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "os", value)

    @property
    @pulumi.getter
    def rev(self) -> Optional[pulumi.Input[int]]:
        """
        Revision.
        """
        return pulumi.get(self, "rev")

    @rev.setter
    def rev(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rev", value)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> Optional[pulumi.Input[int]]:
        """
        Rule ID.
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rule_id", value)

    @property
    @pulumi.getter
    def service(self) -> Optional[pulumi.Input[str]]:
        """
        Vulnerable service.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[str]]:
        """
        Severity.
        """
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "severity", value)

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
class _OtvpState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 application: Optional[pulumi.Input[str]] = None,
                 date: Optional[pulumi.Input[int]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 log_packet: Optional[pulumi.Input[str]] = None,
                 metadatas: Optional[pulumi.Input[Sequence[pulumi.Input['OtvpMetadataArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 os: Optional[pulumi.Input[str]] = None,
                 rev: Optional[pulumi.Input[int]] = None,
                 rule_id: Optional[pulumi.Input[int]] = None,
                 service: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Otvp resources.
        :param pulumi.Input[str] action: Action. Valid values: `pass`, `block`.
        :param pulumi.Input[str] application: Vulnerable applications.
        :param pulumi.Input[int] date: Date.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] group: Group.
        :param pulumi.Input[str] location: Vulnerable location.
        :param pulumi.Input[str] log: Enable/disable logging. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] log_packet: Enable/disable packet logging. Valid values: `disable`, `enable`.
        :param pulumi.Input[Sequence[pulumi.Input['OtvpMetadataArgs']]] metadatas: Meta data. The structure of `metadata` block is documented below.
        :param pulumi.Input[str] name: Rule name.
        :param pulumi.Input[str] os: Vulnerable operation systems.
        :param pulumi.Input[int] rev: Revision.
        :param pulumi.Input[int] rule_id: Rule ID.
        :param pulumi.Input[str] service: Vulnerable service.
        :param pulumi.Input[str] severity: Severity.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if application is not None:
            pulumi.set(__self__, "application", application)
        if date is not None:
            pulumi.set(__self__, "date", date)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if group is not None:
            pulumi.set(__self__, "group", group)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if log is not None:
            pulumi.set(__self__, "log", log)
        if log_packet is not None:
            pulumi.set(__self__, "log_packet", log_packet)
        if metadatas is not None:
            pulumi.set(__self__, "metadatas", metadatas)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if os is not None:
            pulumi.set(__self__, "os", os)
        if rev is not None:
            pulumi.set(__self__, "rev", rev)
        if rule_id is not None:
            pulumi.set(__self__, "rule_id", rule_id)
        if service is not None:
            pulumi.set(__self__, "service", service)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Action. Valid values: `pass`, `block`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def application(self) -> Optional[pulumi.Input[str]]:
        """
        Vulnerable applications.
        """
        return pulumi.get(self, "application")

    @application.setter
    def application(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application", value)

    @property
    @pulumi.getter
    def date(self) -> Optional[pulumi.Input[int]]:
        """
        Date.
        """
        return pulumi.get(self, "date")

    @date.setter
    def date(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "date", value)

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
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        Group.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Vulnerable location.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def log(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable logging. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "log")

    @log.setter
    def log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log", value)

    @property
    @pulumi.getter(name="logPacket")
    def log_packet(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable packet logging. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "log_packet")

    @log_packet.setter
    def log_packet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_packet", value)

    @property
    @pulumi.getter
    def metadatas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OtvpMetadataArgs']]]]:
        """
        Meta data. The structure of `metadata` block is documented below.
        """
        return pulumi.get(self, "metadatas")

    @metadatas.setter
    def metadatas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OtvpMetadataArgs']]]]):
        pulumi.set(self, "metadatas", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Rule name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def os(self) -> Optional[pulumi.Input[str]]:
        """
        Vulnerable operation systems.
        """
        return pulumi.get(self, "os")

    @os.setter
    def os(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "os", value)

    @property
    @pulumi.getter
    def rev(self) -> Optional[pulumi.Input[int]]:
        """
        Revision.
        """
        return pulumi.get(self, "rev")

    @rev.setter
    def rev(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rev", value)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> Optional[pulumi.Input[int]]:
        """
        Rule ID.
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rule_id", value)

    @property
    @pulumi.getter
    def service(self) -> Optional[pulumi.Input[str]]:
        """
        Vulnerable service.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[str]]:
        """
        Severity.
        """
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "severity", value)

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


class Otvp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 application: Optional[pulumi.Input[str]] = None,
                 date: Optional[pulumi.Input[int]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 log_packet: Optional[pulumi.Input[str]] = None,
                 metadatas: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OtvpMetadataArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 os: Optional[pulumi.Input[str]] = None,
                 rev: Optional[pulumi.Input[int]] = None,
                 rule_id: Optional[pulumi.Input[int]] = None,
                 service: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Show OT patch signatures. Applies to FortiOS Version `>= 7.4.1`.

        ## Import

        Rule Otvp can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:rule/otvp:Otvp labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:rule/otvp:Otvp labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Action. Valid values: `pass`, `block`.
        :param pulumi.Input[str] application: Vulnerable applications.
        :param pulumi.Input[int] date: Date.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] group: Group.
        :param pulumi.Input[str] location: Vulnerable location.
        :param pulumi.Input[str] log: Enable/disable logging. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] log_packet: Enable/disable packet logging. Valid values: `disable`, `enable`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OtvpMetadataArgs']]]] metadatas: Meta data. The structure of `metadata` block is documented below.
        :param pulumi.Input[str] name: Rule name.
        :param pulumi.Input[str] os: Vulnerable operation systems.
        :param pulumi.Input[int] rev: Revision.
        :param pulumi.Input[int] rule_id: Rule ID.
        :param pulumi.Input[str] service: Vulnerable service.
        :param pulumi.Input[str] severity: Severity.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[OtvpArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Show OT patch signatures. Applies to FortiOS Version `>= 7.4.1`.

        ## Import

        Rule Otvp can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:rule/otvp:Otvp labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:rule/otvp:Otvp labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param OtvpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OtvpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 application: Optional[pulumi.Input[str]] = None,
                 date: Optional[pulumi.Input[int]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 log_packet: Optional[pulumi.Input[str]] = None,
                 metadatas: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OtvpMetadataArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 os: Optional[pulumi.Input[str]] = None,
                 rev: Optional[pulumi.Input[int]] = None,
                 rule_id: Optional[pulumi.Input[int]] = None,
                 service: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OtvpArgs.__new__(OtvpArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["application"] = application
            __props__.__dict__["date"] = date
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["group"] = group
            __props__.__dict__["location"] = location
            __props__.__dict__["log"] = log
            __props__.__dict__["log_packet"] = log_packet
            __props__.__dict__["metadatas"] = metadatas
            __props__.__dict__["name"] = name
            __props__.__dict__["os"] = os
            __props__.__dict__["rev"] = rev
            __props__.__dict__["rule_id"] = rule_id
            __props__.__dict__["service"] = service
            __props__.__dict__["severity"] = severity
            __props__.__dict__["vdomparam"] = vdomparam
        super(Otvp, __self__).__init__(
            'fortios:rule/otvp:Otvp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            application: Optional[pulumi.Input[str]] = None,
            date: Optional[pulumi.Input[int]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            group: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            log: Optional[pulumi.Input[str]] = None,
            log_packet: Optional[pulumi.Input[str]] = None,
            metadatas: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OtvpMetadataArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            os: Optional[pulumi.Input[str]] = None,
            rev: Optional[pulumi.Input[int]] = None,
            rule_id: Optional[pulumi.Input[int]] = None,
            service: Optional[pulumi.Input[str]] = None,
            severity: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Otvp':
        """
        Get an existing Otvp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Action. Valid values: `pass`, `block`.
        :param pulumi.Input[str] application: Vulnerable applications.
        :param pulumi.Input[int] date: Date.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] group: Group.
        :param pulumi.Input[str] location: Vulnerable location.
        :param pulumi.Input[str] log: Enable/disable logging. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] log_packet: Enable/disable packet logging. Valid values: `disable`, `enable`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OtvpMetadataArgs']]]] metadatas: Meta data. The structure of `metadata` block is documented below.
        :param pulumi.Input[str] name: Rule name.
        :param pulumi.Input[str] os: Vulnerable operation systems.
        :param pulumi.Input[int] rev: Revision.
        :param pulumi.Input[int] rule_id: Rule ID.
        :param pulumi.Input[str] service: Vulnerable service.
        :param pulumi.Input[str] severity: Severity.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OtvpState.__new__(_OtvpState)

        __props__.__dict__["action"] = action
        __props__.__dict__["application"] = application
        __props__.__dict__["date"] = date
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["group"] = group
        __props__.__dict__["location"] = location
        __props__.__dict__["log"] = log
        __props__.__dict__["log_packet"] = log_packet
        __props__.__dict__["metadatas"] = metadatas
        __props__.__dict__["name"] = name
        __props__.__dict__["os"] = os
        __props__.__dict__["rev"] = rev
        __props__.__dict__["rule_id"] = rule_id
        __props__.__dict__["service"] = service
        __props__.__dict__["severity"] = severity
        __props__.__dict__["vdomparam"] = vdomparam
        return Otvp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        Action. Valid values: `pass`, `block`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def application(self) -> pulumi.Output[str]:
        """
        Vulnerable applications.
        """
        return pulumi.get(self, "application")

    @property
    @pulumi.getter
    def date(self) -> pulumi.Output[int]:
        """
        Date.
        """
        return pulumi.get(self, "date")

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
    def group(self) -> pulumi.Output[str]:
        """
        Group.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Vulnerable location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def log(self) -> pulumi.Output[str]:
        """
        Enable/disable logging. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "log")

    @property
    @pulumi.getter(name="logPacket")
    def log_packet(self) -> pulumi.Output[str]:
        """
        Enable/disable packet logging. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "log_packet")

    @property
    @pulumi.getter
    def metadatas(self) -> pulumi.Output[Optional[Sequence['outputs.OtvpMetadata']]]:
        """
        Meta data. The structure of `metadata` block is documented below.
        """
        return pulumi.get(self, "metadatas")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Rule name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def os(self) -> pulumi.Output[str]:
        """
        Vulnerable operation systems.
        """
        return pulumi.get(self, "os")

    @property
    @pulumi.getter
    def rev(self) -> pulumi.Output[int]:
        """
        Revision.
        """
        return pulumi.get(self, "rev")

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> pulumi.Output[int]:
        """
        Rule ID.
        """
        return pulumi.get(self, "rule_id")

    @property
    @pulumi.getter
    def service(self) -> pulumi.Output[str]:
        """
        Vulnerable service.
        """
        return pulumi.get(self, "service")

    @property
    @pulumi.getter
    def severity(self) -> pulumi.Output[str]:
        """
        Severity.
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[str]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

