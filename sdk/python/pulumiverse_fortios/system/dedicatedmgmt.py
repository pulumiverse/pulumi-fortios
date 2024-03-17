# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DedicatedmgmtArgs', 'Dedicatedmgmt']

@pulumi.input_type
class DedicatedmgmtArgs:
    def __init__(__self__, *,
                 default_gateway: Optional[pulumi.Input[str]] = None,
                 dhcp_end_ip: Optional[pulumi.Input[str]] = None,
                 dhcp_netmask: Optional[pulumi.Input[str]] = None,
                 dhcp_server: Optional[pulumi.Input[str]] = None,
                 dhcp_start_ip: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Dedicatedmgmt resource.
        :param pulumi.Input[str] default_gateway: Default gateway for dedicated management interface.
        :param pulumi.Input[str] dhcp_end_ip: DHCP end IP for dedicated management.
        :param pulumi.Input[str] dhcp_netmask: DHCP netmask.
        :param pulumi.Input[str] dhcp_server: Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dhcp_start_ip: DHCP start IP for dedicated management.
        :param pulumi.Input[str] interface: Dedicated management interface.
        :param pulumi.Input[str] status: Enable/disable dedicated management. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if default_gateway is not None:
            pulumi.set(__self__, "default_gateway", default_gateway)
        if dhcp_end_ip is not None:
            pulumi.set(__self__, "dhcp_end_ip", dhcp_end_ip)
        if dhcp_netmask is not None:
            pulumi.set(__self__, "dhcp_netmask", dhcp_netmask)
        if dhcp_server is not None:
            pulumi.set(__self__, "dhcp_server", dhcp_server)
        if dhcp_start_ip is not None:
            pulumi.set(__self__, "dhcp_start_ip", dhcp_start_ip)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="defaultGateway")
    def default_gateway(self) -> Optional[pulumi.Input[str]]:
        """
        Default gateway for dedicated management interface.
        """
        return pulumi.get(self, "default_gateway")

    @default_gateway.setter
    def default_gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_gateway", value)

    @property
    @pulumi.getter(name="dhcpEndIp")
    def dhcp_end_ip(self) -> Optional[pulumi.Input[str]]:
        """
        DHCP end IP for dedicated management.
        """
        return pulumi.get(self, "dhcp_end_ip")

    @dhcp_end_ip.setter
    def dhcp_end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dhcp_end_ip", value)

    @property
    @pulumi.getter(name="dhcpNetmask")
    def dhcp_netmask(self) -> Optional[pulumi.Input[str]]:
        """
        DHCP netmask.
        """
        return pulumi.get(self, "dhcp_netmask")

    @dhcp_netmask.setter
    def dhcp_netmask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dhcp_netmask", value)

    @property
    @pulumi.getter(name="dhcpServer")
    def dhcp_server(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "dhcp_server")

    @dhcp_server.setter
    def dhcp_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dhcp_server", value)

    @property
    @pulumi.getter(name="dhcpStartIp")
    def dhcp_start_ip(self) -> Optional[pulumi.Input[str]]:
        """
        DHCP start IP for dedicated management.
        """
        return pulumi.get(self, "dhcp_start_ip")

    @dhcp_start_ip.setter
    def dhcp_start_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dhcp_start_ip", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        Dedicated management interface.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable dedicated management. Valid values: `enable`, `disable`.
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
class _DedicatedmgmtState:
    def __init__(__self__, *,
                 default_gateway: Optional[pulumi.Input[str]] = None,
                 dhcp_end_ip: Optional[pulumi.Input[str]] = None,
                 dhcp_netmask: Optional[pulumi.Input[str]] = None,
                 dhcp_server: Optional[pulumi.Input[str]] = None,
                 dhcp_start_ip: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Dedicatedmgmt resources.
        :param pulumi.Input[str] default_gateway: Default gateway for dedicated management interface.
        :param pulumi.Input[str] dhcp_end_ip: DHCP end IP for dedicated management.
        :param pulumi.Input[str] dhcp_netmask: DHCP netmask.
        :param pulumi.Input[str] dhcp_server: Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dhcp_start_ip: DHCP start IP for dedicated management.
        :param pulumi.Input[str] interface: Dedicated management interface.
        :param pulumi.Input[str] status: Enable/disable dedicated management. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if default_gateway is not None:
            pulumi.set(__self__, "default_gateway", default_gateway)
        if dhcp_end_ip is not None:
            pulumi.set(__self__, "dhcp_end_ip", dhcp_end_ip)
        if dhcp_netmask is not None:
            pulumi.set(__self__, "dhcp_netmask", dhcp_netmask)
        if dhcp_server is not None:
            pulumi.set(__self__, "dhcp_server", dhcp_server)
        if dhcp_start_ip is not None:
            pulumi.set(__self__, "dhcp_start_ip", dhcp_start_ip)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="defaultGateway")
    def default_gateway(self) -> Optional[pulumi.Input[str]]:
        """
        Default gateway for dedicated management interface.
        """
        return pulumi.get(self, "default_gateway")

    @default_gateway.setter
    def default_gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_gateway", value)

    @property
    @pulumi.getter(name="dhcpEndIp")
    def dhcp_end_ip(self) -> Optional[pulumi.Input[str]]:
        """
        DHCP end IP for dedicated management.
        """
        return pulumi.get(self, "dhcp_end_ip")

    @dhcp_end_ip.setter
    def dhcp_end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dhcp_end_ip", value)

    @property
    @pulumi.getter(name="dhcpNetmask")
    def dhcp_netmask(self) -> Optional[pulumi.Input[str]]:
        """
        DHCP netmask.
        """
        return pulumi.get(self, "dhcp_netmask")

    @dhcp_netmask.setter
    def dhcp_netmask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dhcp_netmask", value)

    @property
    @pulumi.getter(name="dhcpServer")
    def dhcp_server(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "dhcp_server")

    @dhcp_server.setter
    def dhcp_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dhcp_server", value)

    @property
    @pulumi.getter(name="dhcpStartIp")
    def dhcp_start_ip(self) -> Optional[pulumi.Input[str]]:
        """
        DHCP start IP for dedicated management.
        """
        return pulumi.get(self, "dhcp_start_ip")

    @dhcp_start_ip.setter
    def dhcp_start_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dhcp_start_ip", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        Dedicated management interface.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable dedicated management. Valid values: `enable`, `disable`.
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


class Dedicatedmgmt(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_gateway: Optional[pulumi.Input[str]] = None,
                 dhcp_end_ip: Optional[pulumi.Input[str]] = None,
                 dhcp_netmask: Optional[pulumi.Input[str]] = None,
                 dhcp_server: Optional[pulumi.Input[str]] = None,
                 dhcp_start_ip: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure dedicated management.

        ## Import

        System DedicatedMgmt can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/dedicatedmgmt:Dedicatedmgmt labelname SystemDedicatedMgmt
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/dedicatedmgmt:Dedicatedmgmt labelname SystemDedicatedMgmt
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_gateway: Default gateway for dedicated management interface.
        :param pulumi.Input[str] dhcp_end_ip: DHCP end IP for dedicated management.
        :param pulumi.Input[str] dhcp_netmask: DHCP netmask.
        :param pulumi.Input[str] dhcp_server: Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dhcp_start_ip: DHCP start IP for dedicated management.
        :param pulumi.Input[str] interface: Dedicated management interface.
        :param pulumi.Input[str] status: Enable/disable dedicated management. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DedicatedmgmtArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure dedicated management.

        ## Import

        System DedicatedMgmt can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/dedicatedmgmt:Dedicatedmgmt labelname SystemDedicatedMgmt
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/dedicatedmgmt:Dedicatedmgmt labelname SystemDedicatedMgmt
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param DedicatedmgmtArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DedicatedmgmtArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_gateway: Optional[pulumi.Input[str]] = None,
                 dhcp_end_ip: Optional[pulumi.Input[str]] = None,
                 dhcp_netmask: Optional[pulumi.Input[str]] = None,
                 dhcp_server: Optional[pulumi.Input[str]] = None,
                 dhcp_start_ip: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DedicatedmgmtArgs.__new__(DedicatedmgmtArgs)

            __props__.__dict__["default_gateway"] = default_gateway
            __props__.__dict__["dhcp_end_ip"] = dhcp_end_ip
            __props__.__dict__["dhcp_netmask"] = dhcp_netmask
            __props__.__dict__["dhcp_server"] = dhcp_server
            __props__.__dict__["dhcp_start_ip"] = dhcp_start_ip
            __props__.__dict__["interface"] = interface
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(Dedicatedmgmt, __self__).__init__(
            'fortios:system/dedicatedmgmt:Dedicatedmgmt',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_gateway: Optional[pulumi.Input[str]] = None,
            dhcp_end_ip: Optional[pulumi.Input[str]] = None,
            dhcp_netmask: Optional[pulumi.Input[str]] = None,
            dhcp_server: Optional[pulumi.Input[str]] = None,
            dhcp_start_ip: Optional[pulumi.Input[str]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Dedicatedmgmt':
        """
        Get an existing Dedicatedmgmt resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_gateway: Default gateway for dedicated management interface.
        :param pulumi.Input[str] dhcp_end_ip: DHCP end IP for dedicated management.
        :param pulumi.Input[str] dhcp_netmask: DHCP netmask.
        :param pulumi.Input[str] dhcp_server: Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dhcp_start_ip: DHCP start IP for dedicated management.
        :param pulumi.Input[str] interface: Dedicated management interface.
        :param pulumi.Input[str] status: Enable/disable dedicated management. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DedicatedmgmtState.__new__(_DedicatedmgmtState)

        __props__.__dict__["default_gateway"] = default_gateway
        __props__.__dict__["dhcp_end_ip"] = dhcp_end_ip
        __props__.__dict__["dhcp_netmask"] = dhcp_netmask
        __props__.__dict__["dhcp_server"] = dhcp_server
        __props__.__dict__["dhcp_start_ip"] = dhcp_start_ip
        __props__.__dict__["interface"] = interface
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return Dedicatedmgmt(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultGateway")
    def default_gateway(self) -> pulumi.Output[str]:
        """
        Default gateway for dedicated management interface.
        """
        return pulumi.get(self, "default_gateway")

    @property
    @pulumi.getter(name="dhcpEndIp")
    def dhcp_end_ip(self) -> pulumi.Output[str]:
        """
        DHCP end IP for dedicated management.
        """
        return pulumi.get(self, "dhcp_end_ip")

    @property
    @pulumi.getter(name="dhcpNetmask")
    def dhcp_netmask(self) -> pulumi.Output[str]:
        """
        DHCP netmask.
        """
        return pulumi.get(self, "dhcp_netmask")

    @property
    @pulumi.getter(name="dhcpServer")
    def dhcp_server(self) -> pulumi.Output[str]:
        """
        Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "dhcp_server")

    @property
    @pulumi.getter(name="dhcpStartIp")
    def dhcp_start_ip(self) -> pulumi.Output[str]:
        """
        DHCP start IP for dedicated management.
        """
        return pulumi.get(self, "dhcp_start_ip")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[str]:
        """
        Dedicated management interface.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable dedicated management. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

