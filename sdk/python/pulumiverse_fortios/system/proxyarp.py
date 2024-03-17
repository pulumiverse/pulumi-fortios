# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ProxyarpArgs', 'Proxyarp']

@pulumi.input_type
class ProxyarpArgs:
    def __init__(__self__, *,
                 fosid: pulumi.Input[int],
                 interface: pulumi.Input[str],
                 ip: pulumi.Input[str],
                 end_ip: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Proxyarp resource.
        :param pulumi.Input[int] fosid: Unique integer ID of the entry.
        :param pulumi.Input[str] interface: Interface acting proxy-ARP.
        :param pulumi.Input[str] ip: IP address or start IP to be proxied.
        :param pulumi.Input[str] end_ip: End IP of IP range to be proxied.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "fosid", fosid)
        pulumi.set(__self__, "interface", interface)
        pulumi.set(__self__, "ip", ip)
        if end_ip is not None:
            pulumi.set(__self__, "end_ip", end_ip)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Input[int]:
        """
        Unique integer ID of the entry.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: pulumi.Input[int]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Input[str]:
        """
        Interface acting proxy-ARP.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: pulumi.Input[str]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[str]:
        """
        IP address or start IP to be proxied.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> Optional[pulumi.Input[str]]:
        """
        End IP of IP range to be proxied.
        """
        return pulumi.get(self, "end_ip")

    @end_ip.setter
    def end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_ip", value)

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
class _ProxyarpState:
    def __init__(__self__, *,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Proxyarp resources.
        :param pulumi.Input[str] end_ip: End IP of IP range to be proxied.
        :param pulumi.Input[int] fosid: Unique integer ID of the entry.
        :param pulumi.Input[str] interface: Interface acting proxy-ARP.
        :param pulumi.Input[str] ip: IP address or start IP to be proxied.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if end_ip is not None:
            pulumi.set(__self__, "end_ip", end_ip)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> Optional[pulumi.Input[str]]:
        """
        End IP of IP range to be proxied.
        """
        return pulumi.get(self, "end_ip")

    @end_ip.setter
    def end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_ip", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Unique integer ID of the entry.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        Interface acting proxy-ARP.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address or start IP to be proxied.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

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


class Proxyarp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure proxy-ARP.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.system.Proxyarp("trname",
            end_ip="1.1.1.3",
            fosid=1,
            interface="port4",
            ip="1.1.1.1")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        System ProxyArp can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/proxyarp:Proxyarp labelname {{fosid}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/proxyarp:Proxyarp labelname {{fosid}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] end_ip: End IP of IP range to be proxied.
        :param pulumi.Input[int] fosid: Unique integer ID of the entry.
        :param pulumi.Input[str] interface: Interface acting proxy-ARP.
        :param pulumi.Input[str] ip: IP address or start IP to be proxied.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProxyarpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure proxy-ARP.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.system.Proxyarp("trname",
            end_ip="1.1.1.3",
            fosid=1,
            interface="port4",
            ip="1.1.1.1")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        System ProxyArp can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/proxyarp:Proxyarp labelname {{fosid}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/proxyarp:Proxyarp labelname {{fosid}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param ProxyarpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProxyarpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProxyarpArgs.__new__(ProxyarpArgs)

            __props__.__dict__["end_ip"] = end_ip
            if fosid is None and not opts.urn:
                raise TypeError("Missing required property 'fosid'")
            __props__.__dict__["fosid"] = fosid
            if interface is None and not opts.urn:
                raise TypeError("Missing required property 'interface'")
            __props__.__dict__["interface"] = interface
            if ip is None and not opts.urn:
                raise TypeError("Missing required property 'ip'")
            __props__.__dict__["ip"] = ip
            __props__.__dict__["vdomparam"] = vdomparam
        super(Proxyarp, __self__).__init__(
            'fortios:system/proxyarp:Proxyarp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            end_ip: Optional[pulumi.Input[str]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            ip: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Proxyarp':
        """
        Get an existing Proxyarp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] end_ip: End IP of IP range to be proxied.
        :param pulumi.Input[int] fosid: Unique integer ID of the entry.
        :param pulumi.Input[str] interface: Interface acting proxy-ARP.
        :param pulumi.Input[str] ip: IP address or start IP to be proxied.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProxyarpState.__new__(_ProxyarpState)

        __props__.__dict__["end_ip"] = end_ip
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["interface"] = interface
        __props__.__dict__["ip"] = ip
        __props__.__dict__["vdomparam"] = vdomparam
        return Proxyarp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> pulumi.Output[str]:
        """
        End IP of IP range to be proxied.
        """
        return pulumi.get(self, "end_ip")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        Unique integer ID of the entry.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[str]:
        """
        Interface acting proxy-ARP.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        """
        IP address or start IP to be proxied.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

