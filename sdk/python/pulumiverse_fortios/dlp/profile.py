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

__all__ = ['ProfileArgs', 'Profile']

@pulumi.input_type
class ProfileArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 dlp_log: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 extended_log: Optional[pulumi.Input[str]] = None,
                 feature_set: Optional[pulumi.Input[str]] = None,
                 full_archive_proto: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 nac_quar_log: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 replacemsg_group: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileRuleArgs']]]] = None,
                 summary_proto: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Profile resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dlp_log: Enable/disable DLP logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] extended_log: Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] feature_set: Flow/proxy feature set. Valid values: `flow`, `proxy`.
        :param pulumi.Input[str] full_archive_proto: Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] nac_quar_log: Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: Name of the DLP profile.
        :param pulumi.Input[str] replacemsg_group: Replacement message group used by this DLP profile.
        :param pulumi.Input[Sequence[pulumi.Input['ProfileRuleArgs']]] rules: Set up DLP rules for this profile. The structure of `rule` block is documented below.
        :param pulumi.Input[str] summary_proto: Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dlp_log is not None:
            pulumi.set(__self__, "dlp_log", dlp_log)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if extended_log is not None:
            pulumi.set(__self__, "extended_log", extended_log)
        if feature_set is not None:
            pulumi.set(__self__, "feature_set", feature_set)
        if full_archive_proto is not None:
            pulumi.set(__self__, "full_archive_proto", full_archive_proto)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if nac_quar_log is not None:
            pulumi.set(__self__, "nac_quar_log", nac_quar_log)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if replacemsg_group is not None:
            pulumi.set(__self__, "replacemsg_group", replacemsg_group)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if summary_proto is not None:
            pulumi.set(__self__, "summary_proto", summary_proto)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="dlpLog")
    def dlp_log(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable DLP logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "dlp_log")

    @dlp_log.setter
    def dlp_log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dlp_log", value)

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
    @pulumi.getter(name="extendedLog")
    def extended_log(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "extended_log")

    @extended_log.setter
    def extended_log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extended_log", value)

    @property
    @pulumi.getter(name="featureSet")
    def feature_set(self) -> Optional[pulumi.Input[str]]:
        """
        Flow/proxy feature set. Valid values: `flow`, `proxy`.
        """
        return pulumi.get(self, "feature_set")

    @feature_set.setter
    def feature_set(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "feature_set", value)

    @property
    @pulumi.getter(name="fullArchiveProto")
    def full_archive_proto(self) -> Optional[pulumi.Input[str]]:
        """
        Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        """
        return pulumi.get(self, "full_archive_proto")

    @full_archive_proto.setter
    def full_archive_proto(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "full_archive_proto", value)

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
    @pulumi.getter(name="nacQuarLog")
    def nac_quar_log(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "nac_quar_log")

    @nac_quar_log.setter
    def nac_quar_log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nac_quar_log", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the DLP profile.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="replacemsgGroup")
    def replacemsg_group(self) -> Optional[pulumi.Input[str]]:
        """
        Replacement message group used by this DLP profile.
        """
        return pulumi.get(self, "replacemsg_group")

    @replacemsg_group.setter
    def replacemsg_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replacemsg_group", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProfileRuleArgs']]]]:
        """
        Set up DLP rules for this profile. The structure of `rule` block is documented below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter(name="summaryProto")
    def summary_proto(self) -> Optional[pulumi.Input[str]]:
        """
        Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        """
        return pulumi.get(self, "summary_proto")

    @summary_proto.setter
    def summary_proto(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "summary_proto", value)

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
class _ProfileState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 dlp_log: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 extended_log: Optional[pulumi.Input[str]] = None,
                 feature_set: Optional[pulumi.Input[str]] = None,
                 full_archive_proto: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 nac_quar_log: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 replacemsg_group: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileRuleArgs']]]] = None,
                 summary_proto: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Profile resources.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dlp_log: Enable/disable DLP logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] extended_log: Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] feature_set: Flow/proxy feature set. Valid values: `flow`, `proxy`.
        :param pulumi.Input[str] full_archive_proto: Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] nac_quar_log: Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: Name of the DLP profile.
        :param pulumi.Input[str] replacemsg_group: Replacement message group used by this DLP profile.
        :param pulumi.Input[Sequence[pulumi.Input['ProfileRuleArgs']]] rules: Set up DLP rules for this profile. The structure of `rule` block is documented below.
        :param pulumi.Input[str] summary_proto: Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dlp_log is not None:
            pulumi.set(__self__, "dlp_log", dlp_log)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if extended_log is not None:
            pulumi.set(__self__, "extended_log", extended_log)
        if feature_set is not None:
            pulumi.set(__self__, "feature_set", feature_set)
        if full_archive_proto is not None:
            pulumi.set(__self__, "full_archive_proto", full_archive_proto)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if nac_quar_log is not None:
            pulumi.set(__self__, "nac_quar_log", nac_quar_log)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if replacemsg_group is not None:
            pulumi.set(__self__, "replacemsg_group", replacemsg_group)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if summary_proto is not None:
            pulumi.set(__self__, "summary_proto", summary_proto)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="dlpLog")
    def dlp_log(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable DLP logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "dlp_log")

    @dlp_log.setter
    def dlp_log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dlp_log", value)

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
    @pulumi.getter(name="extendedLog")
    def extended_log(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "extended_log")

    @extended_log.setter
    def extended_log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extended_log", value)

    @property
    @pulumi.getter(name="featureSet")
    def feature_set(self) -> Optional[pulumi.Input[str]]:
        """
        Flow/proxy feature set. Valid values: `flow`, `proxy`.
        """
        return pulumi.get(self, "feature_set")

    @feature_set.setter
    def feature_set(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "feature_set", value)

    @property
    @pulumi.getter(name="fullArchiveProto")
    def full_archive_proto(self) -> Optional[pulumi.Input[str]]:
        """
        Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        """
        return pulumi.get(self, "full_archive_proto")

    @full_archive_proto.setter
    def full_archive_proto(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "full_archive_proto", value)

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
    @pulumi.getter(name="nacQuarLog")
    def nac_quar_log(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "nac_quar_log")

    @nac_quar_log.setter
    def nac_quar_log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nac_quar_log", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the DLP profile.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="replacemsgGroup")
    def replacemsg_group(self) -> Optional[pulumi.Input[str]]:
        """
        Replacement message group used by this DLP profile.
        """
        return pulumi.get(self, "replacemsg_group")

    @replacemsg_group.setter
    def replacemsg_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replacemsg_group", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProfileRuleArgs']]]]:
        """
        Set up DLP rules for this profile. The structure of `rule` block is documented below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter(name="summaryProto")
    def summary_proto(self) -> Optional[pulumi.Input[str]]:
        """
        Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        """
        return pulumi.get(self, "summary_proto")

    @summary_proto.setter
    def summary_proto(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "summary_proto", value)

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


class Profile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dlp_log: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 extended_log: Optional[pulumi.Input[str]] = None,
                 feature_set: Optional[pulumi.Input[str]] = None,
                 full_archive_proto: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 nac_quar_log: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 replacemsg_group: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProfileRuleArgs']]]]] = None,
                 summary_proto: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure DLP profiles. Applies to FortiOS Version `>= 7.2.0`.

        ## Import

        Dlp Profile can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:dlp/profile:Profile labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:dlp/profile:Profile labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dlp_log: Enable/disable DLP logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] extended_log: Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] feature_set: Flow/proxy feature set. Valid values: `flow`, `proxy`.
        :param pulumi.Input[str] full_archive_proto: Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] nac_quar_log: Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: Name of the DLP profile.
        :param pulumi.Input[str] replacemsg_group: Replacement message group used by this DLP profile.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProfileRuleArgs']]]] rules: Set up DLP rules for this profile. The structure of `rule` block is documented below.
        :param pulumi.Input[str] summary_proto: Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProfileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure DLP profiles. Applies to FortiOS Version `>= 7.2.0`.

        ## Import

        Dlp Profile can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:dlp/profile:Profile labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:dlp/profile:Profile labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param ProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dlp_log: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 extended_log: Optional[pulumi.Input[str]] = None,
                 feature_set: Optional[pulumi.Input[str]] = None,
                 full_archive_proto: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 nac_quar_log: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 replacemsg_group: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProfileRuleArgs']]]]] = None,
                 summary_proto: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProfileArgs.__new__(ProfileArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["dlp_log"] = dlp_log
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["extended_log"] = extended_log
            __props__.__dict__["feature_set"] = feature_set
            __props__.__dict__["full_archive_proto"] = full_archive_proto
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["nac_quar_log"] = nac_quar_log
            __props__.__dict__["name"] = name
            __props__.__dict__["replacemsg_group"] = replacemsg_group
            __props__.__dict__["rules"] = rules
            __props__.__dict__["summary_proto"] = summary_proto
            __props__.__dict__["vdomparam"] = vdomparam
        super(Profile, __self__).__init__(
            'fortios:dlp/profile:Profile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            dlp_log: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            extended_log: Optional[pulumi.Input[str]] = None,
            feature_set: Optional[pulumi.Input[str]] = None,
            full_archive_proto: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            nac_quar_log: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            replacemsg_group: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProfileRuleArgs']]]]] = None,
            summary_proto: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Profile':
        """
        Get an existing Profile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dlp_log: Enable/disable DLP logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] extended_log: Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] feature_set: Flow/proxy feature set. Valid values: `flow`, `proxy`.
        :param pulumi.Input[str] full_archive_proto: Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] nac_quar_log: Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: Name of the DLP profile.
        :param pulumi.Input[str] replacemsg_group: Replacement message group used by this DLP profile.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProfileRuleArgs']]]] rules: Set up DLP rules for this profile. The structure of `rule` block is documented below.
        :param pulumi.Input[str] summary_proto: Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProfileState.__new__(_ProfileState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["dlp_log"] = dlp_log
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["extended_log"] = extended_log
        __props__.__dict__["feature_set"] = feature_set
        __props__.__dict__["full_archive_proto"] = full_archive_proto
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["nac_quar_log"] = nac_quar_log
        __props__.__dict__["name"] = name
        __props__.__dict__["replacemsg_group"] = replacemsg_group
        __props__.__dict__["rules"] = rules
        __props__.__dict__["summary_proto"] = summary_proto
        __props__.__dict__["vdomparam"] = vdomparam
        return Profile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="dlpLog")
    def dlp_log(self) -> pulumi.Output[str]:
        """
        Enable/disable DLP logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "dlp_log")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="extendedLog")
    def extended_log(self) -> pulumi.Output[str]:
        """
        Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "extended_log")

    @property
    @pulumi.getter(name="featureSet")
    def feature_set(self) -> pulumi.Output[str]:
        """
        Flow/proxy feature set. Valid values: `flow`, `proxy`.
        """
        return pulumi.get(self, "feature_set")

    @property
    @pulumi.getter(name="fullArchiveProto")
    def full_archive_proto(self) -> pulumi.Output[str]:
        """
        Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        """
        return pulumi.get(self, "full_archive_proto")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        """
        Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        """
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter(name="nacQuarLog")
    def nac_quar_log(self) -> pulumi.Output[str]:
        """
        Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "nac_quar_log")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the DLP profile.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="replacemsgGroup")
    def replacemsg_group(self) -> pulumi.Output[str]:
        """
        Replacement message group used by this DLP profile.
        """
        return pulumi.get(self, "replacemsg_group")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Optional[Sequence['outputs.ProfileRule']]]:
        """
        Set up DLP rules for this profile. The structure of `rule` block is documented below.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="summaryProto")
    def summary_proto(self) -> pulumi.Output[str]:
        """
        Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
        """
        return pulumi.get(self, "summary_proto")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

