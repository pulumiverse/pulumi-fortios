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

__all__ = ['VdomexceptionArgs', 'Vdomexception']

@pulumi.input_type
class VdomexceptionArgs:
    def __init__(__self__, *,
                 object: pulumi.Input[str],
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 oid: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input['VdomexceptionVdomArgs']]]] = None):
        """
        The set of arguments for constructing a Vdomexception resource.
        :param pulumi.Input[str] object: Name of the configuration object that can be configured independently for all VDOMs.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[int] fosid: Index <1-4096>.
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[int] oid: Object ID.
        :param pulumi.Input[str] scope: Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration. Valid values: `all`, `inclusive`, `exclusive`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[Sequence[pulumi.Input['VdomexceptionVdomArgs']]] vdoms: Names of the VDOMs. The structure of `vdom` block is documented below.
        """
        pulumi.set(__self__, "object", object)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if oid is not None:
            pulumi.set(__self__, "oid", oid)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vdoms is not None:
            pulumi.set(__self__, "vdoms", vdoms)

    @property
    @pulumi.getter
    def object(self) -> pulumi.Input[str]:
        """
        Name of the configuration object that can be configured independently for all VDOMs.
        """
        return pulumi.get(self, "object")

    @object.setter
    def object(self, value: pulumi.Input[str]):
        pulumi.set(self, "object", value)

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
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Index <1-4096>.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

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
    @pulumi.getter
    def oid(self) -> Optional[pulumi.Input[int]]:
        """
        Object ID.
        """
        return pulumi.get(self, "oid")

    @oid.setter
    def oid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "oid", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration. Valid values: `all`, `inclusive`, `exclusive`.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

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
    def vdoms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VdomexceptionVdomArgs']]]]:
        """
        Names of the VDOMs. The structure of `vdom` block is documented below.
        """
        return pulumi.get(self, "vdoms")

    @vdoms.setter
    def vdoms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VdomexceptionVdomArgs']]]]):
        pulumi.set(self, "vdoms", value)


@pulumi.input_type
class _VdomexceptionState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 object: Optional[pulumi.Input[str]] = None,
                 oid: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input['VdomexceptionVdomArgs']]]] = None):
        """
        Input properties used for looking up and filtering Vdomexception resources.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[int] fosid: Index <1-4096>.
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] object: Name of the configuration object that can be configured independently for all VDOMs.
        :param pulumi.Input[int] oid: Object ID.
        :param pulumi.Input[str] scope: Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration. Valid values: `all`, `inclusive`, `exclusive`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[Sequence[pulumi.Input['VdomexceptionVdomArgs']]] vdoms: Names of the VDOMs. The structure of `vdom` block is documented below.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if object is not None:
            pulumi.set(__self__, "object", object)
        if oid is not None:
            pulumi.set(__self__, "oid", oid)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vdoms is not None:
            pulumi.set(__self__, "vdoms", vdoms)

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
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Index <1-4096>.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

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
    @pulumi.getter
    def object(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the configuration object that can be configured independently for all VDOMs.
        """
        return pulumi.get(self, "object")

    @object.setter
    def object(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object", value)

    @property
    @pulumi.getter
    def oid(self) -> Optional[pulumi.Input[int]]:
        """
        Object ID.
        """
        return pulumi.get(self, "oid")

    @oid.setter
    def oid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "oid", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration. Valid values: `all`, `inclusive`, `exclusive`.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

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
    def vdoms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VdomexceptionVdomArgs']]]]:
        """
        Names of the VDOMs. The structure of `vdom` block is documented below.
        """
        return pulumi.get(self, "vdoms")

    @vdoms.setter
    def vdoms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VdomexceptionVdomArgs']]]]):
        pulumi.set(self, "vdoms", value)


class Vdomexception(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 object: Optional[pulumi.Input[str]] = None,
                 oid: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VdomexceptionVdomArgs']]]]] = None,
                 __props__=None):
        """
        Global configuration objects that can be configured independently for all VDOMs or for the defined VDOM scope.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.system.Vdomexception("trname",
            fosid=1,
            object="log.fortianalyzer.setting",
            oid=7150,
            scope="all")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        System VdomException can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/vdomexception:Vdomexception labelname {{fosid}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/vdomexception:Vdomexception labelname {{fosid}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[int] fosid: Index <1-4096>.
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] object: Name of the configuration object that can be configured independently for all VDOMs.
        :param pulumi.Input[int] oid: Object ID.
        :param pulumi.Input[str] scope: Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration. Valid values: `all`, `inclusive`, `exclusive`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VdomexceptionVdomArgs']]]] vdoms: Names of the VDOMs. The structure of `vdom` block is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VdomexceptionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Global configuration objects that can be configured independently for all VDOMs or for the defined VDOM scope.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.system.Vdomexception("trname",
            fosid=1,
            object="log.fortianalyzer.setting",
            oid=7150,
            scope="all")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        System VdomException can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/vdomexception:Vdomexception labelname {{fosid}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/vdomexception:Vdomexception labelname {{fosid}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param VdomexceptionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VdomexceptionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 object: Optional[pulumi.Input[str]] = None,
                 oid: Optional[pulumi.Input[int]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VdomexceptionVdomArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VdomexceptionArgs.__new__(VdomexceptionArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["get_all_tables"] = get_all_tables
            if object is None and not opts.urn:
                raise TypeError("Missing required property 'object'")
            __props__.__dict__["object"] = object
            __props__.__dict__["oid"] = oid
            __props__.__dict__["scope"] = scope
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["vdoms"] = vdoms
        super(Vdomexception, __self__).__init__(
            'fortios:system/vdomexception:Vdomexception',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            object: Optional[pulumi.Input[str]] = None,
            oid: Optional[pulumi.Input[int]] = None,
            scope: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VdomexceptionVdomArgs']]]]] = None) -> 'Vdomexception':
        """
        Get an existing Vdomexception resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[int] fosid: Index <1-4096>.
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] object: Name of the configuration object that can be configured independently for all VDOMs.
        :param pulumi.Input[int] oid: Object ID.
        :param pulumi.Input[str] scope: Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration. Valid values: `all`, `inclusive`, `exclusive`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VdomexceptionVdomArgs']]]] vdoms: Names of the VDOMs. The structure of `vdom` block is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VdomexceptionState.__new__(_VdomexceptionState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["object"] = object
        __props__.__dict__["oid"] = oid
        __props__.__dict__["scope"] = scope
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["vdoms"] = vdoms
        return Vdomexception(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        Index <1-4096>.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        """
        Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        """
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter
    def object(self) -> pulumi.Output[str]:
        """
        Name of the configuration object that can be configured independently for all VDOMs.
        """
        return pulumi.get(self, "object")

    @property
    @pulumi.getter
    def oid(self) -> pulumi.Output[int]:
        """
        Object ID.
        """
        return pulumi.get(self, "oid")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration. Valid values: `all`, `inclusive`, `exclusive`.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vdoms(self) -> pulumi.Output[Optional[Sequence['outputs.VdomexceptionVdom']]]:
        """
        Names of the VDOMs. The structure of `vdom` block is documented below.
        """
        return pulumi.get(self, "vdoms")

