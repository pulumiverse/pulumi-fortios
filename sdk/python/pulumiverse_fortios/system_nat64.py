# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SystemNat64Args', 'SystemNat64']

@pulumi.input_type
class SystemNat64Args:
    def __init__(__self__, *,
                 nat64_prefix: pulumi.Input[str],
                 always_synthesize_aaaa_record: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 generate_ipv6_fragment_header: Optional[pulumi.Input[str]] = None,
                 nat46_force_ipv4_packet_forwarding: Optional[pulumi.Input[str]] = None,
                 secondary_prefix_status: Optional[pulumi.Input[str]] = None,
                 secondary_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input['SystemNat64SecondaryPrefixArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemNat64 resource.
        :param pulumi.Input[str] nat64_prefix: NAT64 prefix must be ::/96 (default = 64:ff9b::/96).
        :param pulumi.Input[str] always_synthesize_aaaa_record: Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] generate_ipv6_fragment_header: Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] nat46_force_ipv4_packet_forwarding: Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] secondary_prefix_status: Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['SystemNat64SecondaryPrefixArgs']]] secondary_prefixes: Secondary NAT64 prefix. The structure of `secondary_prefix` block is documented below.
        :param pulumi.Input[str] status: Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "nat64_prefix", nat64_prefix)
        if always_synthesize_aaaa_record is not None:
            pulumi.set(__self__, "always_synthesize_aaaa_record", always_synthesize_aaaa_record)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if generate_ipv6_fragment_header is not None:
            pulumi.set(__self__, "generate_ipv6_fragment_header", generate_ipv6_fragment_header)
        if nat46_force_ipv4_packet_forwarding is not None:
            pulumi.set(__self__, "nat46_force_ipv4_packet_forwarding", nat46_force_ipv4_packet_forwarding)
        if secondary_prefix_status is not None:
            pulumi.set(__self__, "secondary_prefix_status", secondary_prefix_status)
        if secondary_prefixes is not None:
            pulumi.set(__self__, "secondary_prefixes", secondary_prefixes)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="nat64Prefix")
    def nat64_prefix(self) -> pulumi.Input[str]:
        """
        NAT64 prefix must be ::/96 (default = 64:ff9b::/96).
        """
        return pulumi.get(self, "nat64_prefix")

    @nat64_prefix.setter
    def nat64_prefix(self, value: pulumi.Input[str]):
        pulumi.set(self, "nat64_prefix", value)

    @property
    @pulumi.getter(name="alwaysSynthesizeAaaaRecord")
    def always_synthesize_aaaa_record(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "always_synthesize_aaaa_record")

    @always_synthesize_aaaa_record.setter
    def always_synthesize_aaaa_record(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "always_synthesize_aaaa_record", value)

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
    @pulumi.getter(name="generateIpv6FragmentHeader")
    def generate_ipv6_fragment_header(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "generate_ipv6_fragment_header")

    @generate_ipv6_fragment_header.setter
    def generate_ipv6_fragment_header(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "generate_ipv6_fragment_header", value)

    @property
    @pulumi.getter(name="nat46ForceIpv4PacketForwarding")
    def nat46_force_ipv4_packet_forwarding(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "nat46_force_ipv4_packet_forwarding")

    @nat46_force_ipv4_packet_forwarding.setter
    def nat46_force_ipv4_packet_forwarding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat46_force_ipv4_packet_forwarding", value)

    @property
    @pulumi.getter(name="secondaryPrefixStatus")
    def secondary_prefix_status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "secondary_prefix_status")

    @secondary_prefix_status.setter
    def secondary_prefix_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_prefix_status", value)

    @property
    @pulumi.getter(name="secondaryPrefixes")
    def secondary_prefixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemNat64SecondaryPrefixArgs']]]]:
        """
        Secondary NAT64 prefix. The structure of `secondary_prefix` block is documented below.
        """
        return pulumi.get(self, "secondary_prefixes")

    @secondary_prefixes.setter
    def secondary_prefixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemNat64SecondaryPrefixArgs']]]]):
        pulumi.set(self, "secondary_prefixes", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

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
class _SystemNat64State:
    def __init__(__self__, *,
                 always_synthesize_aaaa_record: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 generate_ipv6_fragment_header: Optional[pulumi.Input[str]] = None,
                 nat46_force_ipv4_packet_forwarding: Optional[pulumi.Input[str]] = None,
                 nat64_prefix: Optional[pulumi.Input[str]] = None,
                 secondary_prefix_status: Optional[pulumi.Input[str]] = None,
                 secondary_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input['SystemNat64SecondaryPrefixArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemNat64 resources.
        :param pulumi.Input[str] always_synthesize_aaaa_record: Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] generate_ipv6_fragment_header: Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] nat46_force_ipv4_packet_forwarding: Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] nat64_prefix: NAT64 prefix must be ::/96 (default = 64:ff9b::/96).
        :param pulumi.Input[str] secondary_prefix_status: Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['SystemNat64SecondaryPrefixArgs']]] secondary_prefixes: Secondary NAT64 prefix. The structure of `secondary_prefix` block is documented below.
        :param pulumi.Input[str] status: Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if always_synthesize_aaaa_record is not None:
            pulumi.set(__self__, "always_synthesize_aaaa_record", always_synthesize_aaaa_record)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if generate_ipv6_fragment_header is not None:
            pulumi.set(__self__, "generate_ipv6_fragment_header", generate_ipv6_fragment_header)
        if nat46_force_ipv4_packet_forwarding is not None:
            pulumi.set(__self__, "nat46_force_ipv4_packet_forwarding", nat46_force_ipv4_packet_forwarding)
        if nat64_prefix is not None:
            pulumi.set(__self__, "nat64_prefix", nat64_prefix)
        if secondary_prefix_status is not None:
            pulumi.set(__self__, "secondary_prefix_status", secondary_prefix_status)
        if secondary_prefixes is not None:
            pulumi.set(__self__, "secondary_prefixes", secondary_prefixes)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="alwaysSynthesizeAaaaRecord")
    def always_synthesize_aaaa_record(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "always_synthesize_aaaa_record")

    @always_synthesize_aaaa_record.setter
    def always_synthesize_aaaa_record(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "always_synthesize_aaaa_record", value)

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
    @pulumi.getter(name="generateIpv6FragmentHeader")
    def generate_ipv6_fragment_header(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "generate_ipv6_fragment_header")

    @generate_ipv6_fragment_header.setter
    def generate_ipv6_fragment_header(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "generate_ipv6_fragment_header", value)

    @property
    @pulumi.getter(name="nat46ForceIpv4PacketForwarding")
    def nat46_force_ipv4_packet_forwarding(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "nat46_force_ipv4_packet_forwarding")

    @nat46_force_ipv4_packet_forwarding.setter
    def nat46_force_ipv4_packet_forwarding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat46_force_ipv4_packet_forwarding", value)

    @property
    @pulumi.getter(name="nat64Prefix")
    def nat64_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        NAT64 prefix must be ::/96 (default = 64:ff9b::/96).
        """
        return pulumi.get(self, "nat64_prefix")

    @nat64_prefix.setter
    def nat64_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat64_prefix", value)

    @property
    @pulumi.getter(name="secondaryPrefixStatus")
    def secondary_prefix_status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "secondary_prefix_status")

    @secondary_prefix_status.setter
    def secondary_prefix_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_prefix_status", value)

    @property
    @pulumi.getter(name="secondaryPrefixes")
    def secondary_prefixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemNat64SecondaryPrefixArgs']]]]:
        """
        Secondary NAT64 prefix. The structure of `secondary_prefix` block is documented below.
        """
        return pulumi.get(self, "secondary_prefixes")

    @secondary_prefixes.setter
    def secondary_prefixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemNat64SecondaryPrefixArgs']]]]):
        pulumi.set(self, "secondary_prefixes", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

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


class SystemNat64(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 always_synthesize_aaaa_record: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 generate_ipv6_fragment_header: Optional[pulumi.Input[str]] = None,
                 nat46_force_ipv4_packet_forwarding: Optional[pulumi.Input[str]] = None,
                 nat64_prefix: Optional[pulumi.Input[str]] = None,
                 secondary_prefix_status: Optional[pulumi.Input[str]] = None,
                 secondary_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemNat64SecondaryPrefixArgs']]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure NAT64. Applies to FortiOS Version `<= 7.0.0`.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.SystemNat64("trname",
            always_synthesize_aaaa_record="enable",
            generate_ipv6_fragment_header="disable",
            nat46_force_ipv4_packet_forwarding="disable",
            nat64_prefix="2001:1:2:3::/96",
            secondary_prefix_status="disable",
            status="disable")
        ```

        ## Import

        System Nat64 can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/systemNat64:SystemNat64 labelname SystemNat64
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/systemNat64:SystemNat64 labelname SystemNat64
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] always_synthesize_aaaa_record: Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] generate_ipv6_fragment_header: Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] nat46_force_ipv4_packet_forwarding: Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] nat64_prefix: NAT64 prefix must be ::/96 (default = 64:ff9b::/96).
        :param pulumi.Input[str] secondary_prefix_status: Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemNat64SecondaryPrefixArgs']]]] secondary_prefixes: Secondary NAT64 prefix. The structure of `secondary_prefix` block is documented below.
        :param pulumi.Input[str] status: Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemNat64Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure NAT64. Applies to FortiOS Version `<= 7.0.0`.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.SystemNat64("trname",
            always_synthesize_aaaa_record="enable",
            generate_ipv6_fragment_header="disable",
            nat46_force_ipv4_packet_forwarding="disable",
            nat64_prefix="2001:1:2:3::/96",
            secondary_prefix_status="disable",
            status="disable")
        ```

        ## Import

        System Nat64 can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/systemNat64:SystemNat64 labelname SystemNat64
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/systemNat64:SystemNat64 labelname SystemNat64
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemNat64Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemNat64Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 always_synthesize_aaaa_record: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 generate_ipv6_fragment_header: Optional[pulumi.Input[str]] = None,
                 nat46_force_ipv4_packet_forwarding: Optional[pulumi.Input[str]] = None,
                 nat64_prefix: Optional[pulumi.Input[str]] = None,
                 secondary_prefix_status: Optional[pulumi.Input[str]] = None,
                 secondary_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemNat64SecondaryPrefixArgs']]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemNat64Args.__new__(SystemNat64Args)

            __props__.__dict__["always_synthesize_aaaa_record"] = always_synthesize_aaaa_record
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["generate_ipv6_fragment_header"] = generate_ipv6_fragment_header
            __props__.__dict__["nat46_force_ipv4_packet_forwarding"] = nat46_force_ipv4_packet_forwarding
            if nat64_prefix is None and not opts.urn:
                raise TypeError("Missing required property 'nat64_prefix'")
            __props__.__dict__["nat64_prefix"] = nat64_prefix
            __props__.__dict__["secondary_prefix_status"] = secondary_prefix_status
            __props__.__dict__["secondary_prefixes"] = secondary_prefixes
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemNat64, __self__).__init__(
            'fortios:index/systemNat64:SystemNat64',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            always_synthesize_aaaa_record: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            generate_ipv6_fragment_header: Optional[pulumi.Input[str]] = None,
            nat46_force_ipv4_packet_forwarding: Optional[pulumi.Input[str]] = None,
            nat64_prefix: Optional[pulumi.Input[str]] = None,
            secondary_prefix_status: Optional[pulumi.Input[str]] = None,
            secondary_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemNat64SecondaryPrefixArgs']]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemNat64':
        """
        Get an existing SystemNat64 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] always_synthesize_aaaa_record: Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] generate_ipv6_fragment_header: Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] nat46_force_ipv4_packet_forwarding: Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] nat64_prefix: NAT64 prefix must be ::/96 (default = 64:ff9b::/96).
        :param pulumi.Input[str] secondary_prefix_status: Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemNat64SecondaryPrefixArgs']]]] secondary_prefixes: Secondary NAT64 prefix. The structure of `secondary_prefix` block is documented below.
        :param pulumi.Input[str] status: Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemNat64State.__new__(_SystemNat64State)

        __props__.__dict__["always_synthesize_aaaa_record"] = always_synthesize_aaaa_record
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["generate_ipv6_fragment_header"] = generate_ipv6_fragment_header
        __props__.__dict__["nat46_force_ipv4_packet_forwarding"] = nat46_force_ipv4_packet_forwarding
        __props__.__dict__["nat64_prefix"] = nat64_prefix
        __props__.__dict__["secondary_prefix_status"] = secondary_prefix_status
        __props__.__dict__["secondary_prefixes"] = secondary_prefixes
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemNat64(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alwaysSynthesizeAaaaRecord")
    def always_synthesize_aaaa_record(self) -> pulumi.Output[str]:
        """
        Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "always_synthesize_aaaa_record")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="generateIpv6FragmentHeader")
    def generate_ipv6_fragment_header(self) -> pulumi.Output[str]:
        """
        Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "generate_ipv6_fragment_header")

    @property
    @pulumi.getter(name="nat46ForceIpv4PacketForwarding")
    def nat46_force_ipv4_packet_forwarding(self) -> pulumi.Output[str]:
        """
        Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "nat46_force_ipv4_packet_forwarding")

    @property
    @pulumi.getter(name="nat64Prefix")
    def nat64_prefix(self) -> pulumi.Output[str]:
        """
        NAT64 prefix must be ::/96 (default = 64:ff9b::/96).
        """
        return pulumi.get(self, "nat64_prefix")

    @property
    @pulumi.getter(name="secondaryPrefixStatus")
    def secondary_prefix_status(self) -> pulumi.Output[str]:
        """
        Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "secondary_prefix_status")

    @property
    @pulumi.getter(name="secondaryPrefixes")
    def secondary_prefixes(self) -> pulumi.Output[Optional[Sequence['outputs.SystemNat64SecondaryPrefix']]]:
        """
        Secondary NAT64 prefix. The structure of `secondary_prefix` block is documented below.
        """
        return pulumi.get(self, "secondary_prefixes")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

