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

__all__ = ['GeoipoverrideArgs', 'Geoipoverride']

@pulumi.input_type
class GeoipoverrideArgs:
    def __init__(__self__, *,
                 country_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ip6_ranges: Optional[pulumi.Input[Sequence[pulumi.Input['GeoipoverrideIp6RangeArgs']]]] = None,
                 ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input['GeoipoverrideIpRangeArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Geoipoverride resource.
        :param pulumi.Input[str] country_id: Two character Country ID code.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[Sequence[pulumi.Input['GeoipoverrideIp6RangeArgs']]] ip6_ranges: Table of IPv6 ranges assigned to country. The structure of `ip6_range` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['GeoipoverrideIpRangeArgs']]] ip_ranges: Table of IP ranges assigned to country. The structure of `ip_range` block is documented below.
        :param pulumi.Input[str] name: Location name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if country_id is not None:
            pulumi.set(__self__, "country_id", country_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if ip6_ranges is not None:
            pulumi.set(__self__, "ip6_ranges", ip6_ranges)
        if ip_ranges is not None:
            pulumi.set(__self__, "ip_ranges", ip_ranges)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="countryId")
    def country_id(self) -> Optional[pulumi.Input[str]]:
        """
        Two character Country ID code.
        """
        return pulumi.get(self, "country_id")

    @country_id.setter
    def country_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "country_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

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
    @pulumi.getter(name="ip6Ranges")
    def ip6_ranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GeoipoverrideIp6RangeArgs']]]]:
        """
        Table of IPv6 ranges assigned to country. The structure of `ip6_range` block is documented below.
        """
        return pulumi.get(self, "ip6_ranges")

    @ip6_ranges.setter
    def ip6_ranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GeoipoverrideIp6RangeArgs']]]]):
        pulumi.set(self, "ip6_ranges", value)

    @property
    @pulumi.getter(name="ipRanges")
    def ip_ranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GeoipoverrideIpRangeArgs']]]]:
        """
        Table of IP ranges assigned to country. The structure of `ip_range` block is documented below.
        """
        return pulumi.get(self, "ip_ranges")

    @ip_ranges.setter
    def ip_ranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GeoipoverrideIpRangeArgs']]]]):
        pulumi.set(self, "ip_ranges", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Location name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
class _GeoipoverrideState:
    def __init__(__self__, *,
                 country_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ip6_ranges: Optional[pulumi.Input[Sequence[pulumi.Input['GeoipoverrideIp6RangeArgs']]]] = None,
                 ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input['GeoipoverrideIpRangeArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Geoipoverride resources.
        :param pulumi.Input[str] country_id: Two character Country ID code.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[Sequence[pulumi.Input['GeoipoverrideIp6RangeArgs']]] ip6_ranges: Table of IPv6 ranges assigned to country. The structure of `ip6_range` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['GeoipoverrideIpRangeArgs']]] ip_ranges: Table of IP ranges assigned to country. The structure of `ip_range` block is documented below.
        :param pulumi.Input[str] name: Location name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if country_id is not None:
            pulumi.set(__self__, "country_id", country_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if ip6_ranges is not None:
            pulumi.set(__self__, "ip6_ranges", ip6_ranges)
        if ip_ranges is not None:
            pulumi.set(__self__, "ip_ranges", ip_ranges)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="countryId")
    def country_id(self) -> Optional[pulumi.Input[str]]:
        """
        Two character Country ID code.
        """
        return pulumi.get(self, "country_id")

    @country_id.setter
    def country_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "country_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

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
    @pulumi.getter(name="ip6Ranges")
    def ip6_ranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GeoipoverrideIp6RangeArgs']]]]:
        """
        Table of IPv6 ranges assigned to country. The structure of `ip6_range` block is documented below.
        """
        return pulumi.get(self, "ip6_ranges")

    @ip6_ranges.setter
    def ip6_ranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GeoipoverrideIp6RangeArgs']]]]):
        pulumi.set(self, "ip6_ranges", value)

    @property
    @pulumi.getter(name="ipRanges")
    def ip_ranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GeoipoverrideIpRangeArgs']]]]:
        """
        Table of IP ranges assigned to country. The structure of `ip_range` block is documented below.
        """
        return pulumi.get(self, "ip_ranges")

    @ip_ranges.setter
    def ip_ranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GeoipoverrideIpRangeArgs']]]]):
        pulumi.set(self, "ip_ranges", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Location name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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


class Geoipoverride(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 country_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ip6_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GeoipoverrideIp6RangeArgs']]]]] = None,
                 ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GeoipoverrideIpRangeArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.system.Geoipoverride("trname", description="TEST for country")
        ```

        ## Import

        System GeoipOverride can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/geoipoverride:Geoipoverride labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/geoipoverride:Geoipoverride labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] country_id: Two character Country ID code.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GeoipoverrideIp6RangeArgs']]]] ip6_ranges: Table of IPv6 ranges assigned to country. The structure of `ip6_range` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GeoipoverrideIpRangeArgs']]]] ip_ranges: Table of IP ranges assigned to country. The structure of `ip_range` block is documented below.
        :param pulumi.Input[str] name: Location name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[GeoipoverrideArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.system.Geoipoverride("trname", description="TEST for country")
        ```

        ## Import

        System GeoipOverride can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/geoipoverride:Geoipoverride labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/geoipoverride:Geoipoverride labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param GeoipoverrideArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GeoipoverrideArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 country_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ip6_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GeoipoverrideIp6RangeArgs']]]]] = None,
                 ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GeoipoverrideIpRangeArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GeoipoverrideArgs.__new__(GeoipoverrideArgs)

            __props__.__dict__["country_id"] = country_id
            __props__.__dict__["description"] = description
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["ip6_ranges"] = ip6_ranges
            __props__.__dict__["ip_ranges"] = ip_ranges
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(Geoipoverride, __self__).__init__(
            'fortios:system/geoipoverride:Geoipoverride',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            country_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            ip6_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GeoipoverrideIp6RangeArgs']]]]] = None,
            ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GeoipoverrideIpRangeArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Geoipoverride':
        """
        Get an existing Geoipoverride resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] country_id: Two character Country ID code.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GeoipoverrideIp6RangeArgs']]]] ip6_ranges: Table of IPv6 ranges assigned to country. The structure of `ip6_range` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GeoipoverrideIpRangeArgs']]]] ip_ranges: Table of IP ranges assigned to country. The structure of `ip_range` block is documented below.
        :param pulumi.Input[str] name: Location name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GeoipoverrideState.__new__(_GeoipoverrideState)

        __props__.__dict__["country_id"] = country_id
        __props__.__dict__["description"] = description
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["ip6_ranges"] = ip6_ranges
        __props__.__dict__["ip_ranges"] = ip_ranges
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return Geoipoverride(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="countryId")
    def country_id(self) -> pulumi.Output[str]:
        """
        Two character Country ID code.
        """
        return pulumi.get(self, "country_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description.
        """
        return pulumi.get(self, "description")

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
    @pulumi.getter(name="ip6Ranges")
    def ip6_ranges(self) -> pulumi.Output[Optional[Sequence['outputs.GeoipoverrideIp6Range']]]:
        """
        Table of IPv6 ranges assigned to country. The structure of `ip6_range` block is documented below.
        """
        return pulumi.get(self, "ip6_ranges")

    @property
    @pulumi.getter(name="ipRanges")
    def ip_ranges(self) -> pulumi.Output[Optional[Sequence['outputs.GeoipoverrideIpRange']]]:
        """
        Table of IP ranges assigned to country. The structure of `ip_range` block is documented below.
        """
        return pulumi.get(self, "ip_ranges")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Location name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[str]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

