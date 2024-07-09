# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['PeripshaperArgs', 'Peripshaper']

@pulumi.input_type
class PeripshaperArgs:
    def __init__(__self__, *,
                 bandwidth_unit: Optional[pulumi.Input[str]] = None,
                 diffserv_forward: Optional[pulumi.Input[str]] = None,
                 diffserv_reverse: Optional[pulumi.Input[str]] = None,
                 diffservcode_forward: Optional[pulumi.Input[str]] = None,
                 diffservcode_rev: Optional[pulumi.Input[str]] = None,
                 max_bandwidth: Optional[pulumi.Input[int]] = None,
                 max_concurrent_session: Optional[pulumi.Input[int]] = None,
                 max_concurrent_tcp_session: Optional[pulumi.Input[int]] = None,
                 max_concurrent_udp_session: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Peripshaper resource.
        :param pulumi.Input[str] bandwidth_unit: Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
        :param pulumi.Input[str] diffserv_forward: Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] diffserv_reverse: Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] diffservcode_forward: Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
        :param pulumi.Input[str] diffservcode_rev: Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
        :param pulumi.Input[int] max_bandwidth: Upper bandwidth limit enforced by this shaper. 0 means no limit. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.15, 7.0.6-7.0.15, >= 7.2.1: 0 - 80000000.
        :param pulumi.Input[int] max_concurrent_session: Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        :param pulumi.Input[int] max_concurrent_tcp_session: Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        :param pulumi.Input[int] max_concurrent_udp_session: Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        :param pulumi.Input[str] name: Traffic shaper name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if bandwidth_unit is not None:
            pulumi.set(__self__, "bandwidth_unit", bandwidth_unit)
        if diffserv_forward is not None:
            pulumi.set(__self__, "diffserv_forward", diffserv_forward)
        if diffserv_reverse is not None:
            pulumi.set(__self__, "diffserv_reverse", diffserv_reverse)
        if diffservcode_forward is not None:
            pulumi.set(__self__, "diffservcode_forward", diffservcode_forward)
        if diffservcode_rev is not None:
            pulumi.set(__self__, "diffservcode_rev", diffservcode_rev)
        if max_bandwidth is not None:
            pulumi.set(__self__, "max_bandwidth", max_bandwidth)
        if max_concurrent_session is not None:
            pulumi.set(__self__, "max_concurrent_session", max_concurrent_session)
        if max_concurrent_tcp_session is not None:
            pulumi.set(__self__, "max_concurrent_tcp_session", max_concurrent_tcp_session)
        if max_concurrent_udp_session is not None:
            pulumi.set(__self__, "max_concurrent_udp_session", max_concurrent_udp_session)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="bandwidthUnit")
    def bandwidth_unit(self) -> Optional[pulumi.Input[str]]:
        """
        Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
        """
        return pulumi.get(self, "bandwidth_unit")

    @bandwidth_unit.setter
    def bandwidth_unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bandwidth_unit", value)

    @property
    @pulumi.getter(name="diffservForward")
    def diffserv_forward(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "diffserv_forward")

    @diffserv_forward.setter
    def diffserv_forward(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "diffserv_forward", value)

    @property
    @pulumi.getter(name="diffservReverse")
    def diffserv_reverse(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "diffserv_reverse")

    @diffserv_reverse.setter
    def diffserv_reverse(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "diffserv_reverse", value)

    @property
    @pulumi.getter(name="diffservcodeForward")
    def diffservcode_forward(self) -> Optional[pulumi.Input[str]]:
        """
        Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
        """
        return pulumi.get(self, "diffservcode_forward")

    @diffservcode_forward.setter
    def diffservcode_forward(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "diffservcode_forward", value)

    @property
    @pulumi.getter(name="diffservcodeRev")
    def diffservcode_rev(self) -> Optional[pulumi.Input[str]]:
        """
        Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
        """
        return pulumi.get(self, "diffservcode_rev")

    @diffservcode_rev.setter
    def diffservcode_rev(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "diffservcode_rev", value)

    @property
    @pulumi.getter(name="maxBandwidth")
    def max_bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Upper bandwidth limit enforced by this shaper. 0 means no limit. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.15, 7.0.6-7.0.15, >= 7.2.1: 0 - 80000000.
        """
        return pulumi.get(self, "max_bandwidth")

    @max_bandwidth.setter
    def max_bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_bandwidth", value)

    @property
    @pulumi.getter(name="maxConcurrentSession")
    def max_concurrent_session(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        """
        return pulumi.get(self, "max_concurrent_session")

    @max_concurrent_session.setter
    def max_concurrent_session(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_concurrent_session", value)

    @property
    @pulumi.getter(name="maxConcurrentTcpSession")
    def max_concurrent_tcp_session(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        """
        return pulumi.get(self, "max_concurrent_tcp_session")

    @max_concurrent_tcp_session.setter
    def max_concurrent_tcp_session(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_concurrent_tcp_session", value)

    @property
    @pulumi.getter(name="maxConcurrentUdpSession")
    def max_concurrent_udp_session(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        """
        return pulumi.get(self, "max_concurrent_udp_session")

    @max_concurrent_udp_session.setter
    def max_concurrent_udp_session(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_concurrent_udp_session", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Traffic shaper name.
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
class _PeripshaperState:
    def __init__(__self__, *,
                 bandwidth_unit: Optional[pulumi.Input[str]] = None,
                 diffserv_forward: Optional[pulumi.Input[str]] = None,
                 diffserv_reverse: Optional[pulumi.Input[str]] = None,
                 diffservcode_forward: Optional[pulumi.Input[str]] = None,
                 diffservcode_rev: Optional[pulumi.Input[str]] = None,
                 max_bandwidth: Optional[pulumi.Input[int]] = None,
                 max_concurrent_session: Optional[pulumi.Input[int]] = None,
                 max_concurrent_tcp_session: Optional[pulumi.Input[int]] = None,
                 max_concurrent_udp_session: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Peripshaper resources.
        :param pulumi.Input[str] bandwidth_unit: Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
        :param pulumi.Input[str] diffserv_forward: Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] diffserv_reverse: Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] diffservcode_forward: Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
        :param pulumi.Input[str] diffservcode_rev: Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
        :param pulumi.Input[int] max_bandwidth: Upper bandwidth limit enforced by this shaper. 0 means no limit. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.15, 7.0.6-7.0.15, >= 7.2.1: 0 - 80000000.
        :param pulumi.Input[int] max_concurrent_session: Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        :param pulumi.Input[int] max_concurrent_tcp_session: Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        :param pulumi.Input[int] max_concurrent_udp_session: Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        :param pulumi.Input[str] name: Traffic shaper name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if bandwidth_unit is not None:
            pulumi.set(__self__, "bandwidth_unit", bandwidth_unit)
        if diffserv_forward is not None:
            pulumi.set(__self__, "diffserv_forward", diffserv_forward)
        if diffserv_reverse is not None:
            pulumi.set(__self__, "diffserv_reverse", diffserv_reverse)
        if diffservcode_forward is not None:
            pulumi.set(__self__, "diffservcode_forward", diffservcode_forward)
        if diffservcode_rev is not None:
            pulumi.set(__self__, "diffservcode_rev", diffservcode_rev)
        if max_bandwidth is not None:
            pulumi.set(__self__, "max_bandwidth", max_bandwidth)
        if max_concurrent_session is not None:
            pulumi.set(__self__, "max_concurrent_session", max_concurrent_session)
        if max_concurrent_tcp_session is not None:
            pulumi.set(__self__, "max_concurrent_tcp_session", max_concurrent_tcp_session)
        if max_concurrent_udp_session is not None:
            pulumi.set(__self__, "max_concurrent_udp_session", max_concurrent_udp_session)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="bandwidthUnit")
    def bandwidth_unit(self) -> Optional[pulumi.Input[str]]:
        """
        Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
        """
        return pulumi.get(self, "bandwidth_unit")

    @bandwidth_unit.setter
    def bandwidth_unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bandwidth_unit", value)

    @property
    @pulumi.getter(name="diffservForward")
    def diffserv_forward(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "diffserv_forward")

    @diffserv_forward.setter
    def diffserv_forward(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "diffserv_forward", value)

    @property
    @pulumi.getter(name="diffservReverse")
    def diffserv_reverse(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "diffserv_reverse")

    @diffserv_reverse.setter
    def diffserv_reverse(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "diffserv_reverse", value)

    @property
    @pulumi.getter(name="diffservcodeForward")
    def diffservcode_forward(self) -> Optional[pulumi.Input[str]]:
        """
        Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
        """
        return pulumi.get(self, "diffservcode_forward")

    @diffservcode_forward.setter
    def diffservcode_forward(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "diffservcode_forward", value)

    @property
    @pulumi.getter(name="diffservcodeRev")
    def diffservcode_rev(self) -> Optional[pulumi.Input[str]]:
        """
        Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
        """
        return pulumi.get(self, "diffservcode_rev")

    @diffservcode_rev.setter
    def diffservcode_rev(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "diffservcode_rev", value)

    @property
    @pulumi.getter(name="maxBandwidth")
    def max_bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Upper bandwidth limit enforced by this shaper. 0 means no limit. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.15, 7.0.6-7.0.15, >= 7.2.1: 0 - 80000000.
        """
        return pulumi.get(self, "max_bandwidth")

    @max_bandwidth.setter
    def max_bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_bandwidth", value)

    @property
    @pulumi.getter(name="maxConcurrentSession")
    def max_concurrent_session(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        """
        return pulumi.get(self, "max_concurrent_session")

    @max_concurrent_session.setter
    def max_concurrent_session(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_concurrent_session", value)

    @property
    @pulumi.getter(name="maxConcurrentTcpSession")
    def max_concurrent_tcp_session(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        """
        return pulumi.get(self, "max_concurrent_tcp_session")

    @max_concurrent_tcp_session.setter
    def max_concurrent_tcp_session(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_concurrent_tcp_session", value)

    @property
    @pulumi.getter(name="maxConcurrentUdpSession")
    def max_concurrent_udp_session(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        """
        return pulumi.get(self, "max_concurrent_udp_session")

    @max_concurrent_udp_session.setter
    def max_concurrent_udp_session(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_concurrent_udp_session", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Traffic shaper name.
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


class Peripshaper(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth_unit: Optional[pulumi.Input[str]] = None,
                 diffserv_forward: Optional[pulumi.Input[str]] = None,
                 diffserv_reverse: Optional[pulumi.Input[str]] = None,
                 diffservcode_forward: Optional[pulumi.Input[str]] = None,
                 diffservcode_rev: Optional[pulumi.Input[str]] = None,
                 max_bandwidth: Optional[pulumi.Input[int]] = None,
                 max_concurrent_session: Optional[pulumi.Input[int]] = None,
                 max_concurrent_tcp_session: Optional[pulumi.Input[int]] = None,
                 max_concurrent_udp_session: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure per-IP traffic shaper.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.firewall.shaper.Peripshaper("trname",
            bandwidth_unit="kbps",
            diffserv_forward="disable",
            diffserv_reverse="disable",
            diffservcode_forward="000000",
            diffservcode_rev="000000",
            max_bandwidth=1024,
            max_concurrent_session=33)
        ```

        ## Import

        FirewallShaper PerIpShaper can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:firewall/shaper/peripshaper:Peripshaper labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:firewall/shaper/peripshaper:Peripshaper labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bandwidth_unit: Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
        :param pulumi.Input[str] diffserv_forward: Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] diffserv_reverse: Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] diffservcode_forward: Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
        :param pulumi.Input[str] diffservcode_rev: Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
        :param pulumi.Input[int] max_bandwidth: Upper bandwidth limit enforced by this shaper. 0 means no limit. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.15, 7.0.6-7.0.15, >= 7.2.1: 0 - 80000000.
        :param pulumi.Input[int] max_concurrent_session: Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        :param pulumi.Input[int] max_concurrent_tcp_session: Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        :param pulumi.Input[int] max_concurrent_udp_session: Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        :param pulumi.Input[str] name: Traffic shaper name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PeripshaperArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure per-IP traffic shaper.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.firewall.shaper.Peripshaper("trname",
            bandwidth_unit="kbps",
            diffserv_forward="disable",
            diffserv_reverse="disable",
            diffservcode_forward="000000",
            diffservcode_rev="000000",
            max_bandwidth=1024,
            max_concurrent_session=33)
        ```

        ## Import

        FirewallShaper PerIpShaper can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:firewall/shaper/peripshaper:Peripshaper labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:firewall/shaper/peripshaper:Peripshaper labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param PeripshaperArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PeripshaperArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth_unit: Optional[pulumi.Input[str]] = None,
                 diffserv_forward: Optional[pulumi.Input[str]] = None,
                 diffserv_reverse: Optional[pulumi.Input[str]] = None,
                 diffservcode_forward: Optional[pulumi.Input[str]] = None,
                 diffservcode_rev: Optional[pulumi.Input[str]] = None,
                 max_bandwidth: Optional[pulumi.Input[int]] = None,
                 max_concurrent_session: Optional[pulumi.Input[int]] = None,
                 max_concurrent_tcp_session: Optional[pulumi.Input[int]] = None,
                 max_concurrent_udp_session: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PeripshaperArgs.__new__(PeripshaperArgs)

            __props__.__dict__["bandwidth_unit"] = bandwidth_unit
            __props__.__dict__["diffserv_forward"] = diffserv_forward
            __props__.__dict__["diffserv_reverse"] = diffserv_reverse
            __props__.__dict__["diffservcode_forward"] = diffservcode_forward
            __props__.__dict__["diffservcode_rev"] = diffservcode_rev
            __props__.__dict__["max_bandwidth"] = max_bandwidth
            __props__.__dict__["max_concurrent_session"] = max_concurrent_session
            __props__.__dict__["max_concurrent_tcp_session"] = max_concurrent_tcp_session
            __props__.__dict__["max_concurrent_udp_session"] = max_concurrent_udp_session
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(Peripshaper, __self__).__init__(
            'fortios:firewall/shaper/peripshaper:Peripshaper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bandwidth_unit: Optional[pulumi.Input[str]] = None,
            diffserv_forward: Optional[pulumi.Input[str]] = None,
            diffserv_reverse: Optional[pulumi.Input[str]] = None,
            diffservcode_forward: Optional[pulumi.Input[str]] = None,
            diffservcode_rev: Optional[pulumi.Input[str]] = None,
            max_bandwidth: Optional[pulumi.Input[int]] = None,
            max_concurrent_session: Optional[pulumi.Input[int]] = None,
            max_concurrent_tcp_session: Optional[pulumi.Input[int]] = None,
            max_concurrent_udp_session: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Peripshaper':
        """
        Get an existing Peripshaper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bandwidth_unit: Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
        :param pulumi.Input[str] diffserv_forward: Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] diffserv_reverse: Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] diffservcode_forward: Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
        :param pulumi.Input[str] diffservcode_rev: Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
        :param pulumi.Input[int] max_bandwidth: Upper bandwidth limit enforced by this shaper. 0 means no limit. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.15, 7.0.6-7.0.15, >= 7.2.1: 0 - 80000000.
        :param pulumi.Input[int] max_concurrent_session: Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        :param pulumi.Input[int] max_concurrent_tcp_session: Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        :param pulumi.Input[int] max_concurrent_udp_session: Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        :param pulumi.Input[str] name: Traffic shaper name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PeripshaperState.__new__(_PeripshaperState)

        __props__.__dict__["bandwidth_unit"] = bandwidth_unit
        __props__.__dict__["diffserv_forward"] = diffserv_forward
        __props__.__dict__["diffserv_reverse"] = diffserv_reverse
        __props__.__dict__["diffservcode_forward"] = diffservcode_forward
        __props__.__dict__["diffservcode_rev"] = diffservcode_rev
        __props__.__dict__["max_bandwidth"] = max_bandwidth
        __props__.__dict__["max_concurrent_session"] = max_concurrent_session
        __props__.__dict__["max_concurrent_tcp_session"] = max_concurrent_tcp_session
        __props__.__dict__["max_concurrent_udp_session"] = max_concurrent_udp_session
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return Peripshaper(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bandwidthUnit")
    def bandwidth_unit(self) -> pulumi.Output[str]:
        """
        Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
        """
        return pulumi.get(self, "bandwidth_unit")

    @property
    @pulumi.getter(name="diffservForward")
    def diffserv_forward(self) -> pulumi.Output[str]:
        """
        Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "diffserv_forward")

    @property
    @pulumi.getter(name="diffservReverse")
    def diffserv_reverse(self) -> pulumi.Output[str]:
        """
        Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "diffserv_reverse")

    @property
    @pulumi.getter(name="diffservcodeForward")
    def diffservcode_forward(self) -> pulumi.Output[str]:
        """
        Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
        """
        return pulumi.get(self, "diffservcode_forward")

    @property
    @pulumi.getter(name="diffservcodeRev")
    def diffservcode_rev(self) -> pulumi.Output[str]:
        """
        Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
        """
        return pulumi.get(self, "diffservcode_rev")

    @property
    @pulumi.getter(name="maxBandwidth")
    def max_bandwidth(self) -> pulumi.Output[int]:
        """
        Upper bandwidth limit enforced by this shaper. 0 means no limit. Units depend on the bandwidth-unit setting. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000. On FortiOS versions 6.4.10-6.4.15, 7.0.6-7.0.15, >= 7.2.1: 0 - 80000000.
        """
        return pulumi.get(self, "max_bandwidth")

    @property
    @pulumi.getter(name="maxConcurrentSession")
    def max_concurrent_session(self) -> pulumi.Output[int]:
        """
        Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        """
        return pulumi.get(self, "max_concurrent_session")

    @property
    @pulumi.getter(name="maxConcurrentTcpSession")
    def max_concurrent_tcp_session(self) -> pulumi.Output[int]:
        """
        Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        """
        return pulumi.get(self, "max_concurrent_tcp_session")

    @property
    @pulumi.getter(name="maxConcurrentUdpSession")
    def max_concurrent_udp_session(self) -> pulumi.Output[int]:
        """
        Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
        """
        return pulumi.get(self, "max_concurrent_udp_session")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Traffic shaper name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[str]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

