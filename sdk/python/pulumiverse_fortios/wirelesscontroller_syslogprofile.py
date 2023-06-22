# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['WirelesscontrollerSyslogprofileArgs', 'WirelesscontrollerSyslogprofile']

@pulumi.input_type
class WirelesscontrollerSyslogprofileArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 log_level: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_addr_type: Optional[pulumi.Input[str]] = None,
                 server_fqdn: Optional[pulumi.Input[str]] = None,
                 server_ip: Optional[pulumi.Input[str]] = None,
                 server_port: Optional[pulumi.Input[int]] = None,
                 server_status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WirelesscontrollerSyslogprofile resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] log_level: Lowest level of log messages that FortiAP units send to this server (default = information) Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debugging`.
        :param pulumi.Input[str] name: WTP system log server profile name.
        :param pulumi.Input[str] server_addr_type: Syslog server address type (default = IP) Valid values: `fqdn`, `ip`.
        :param pulumi.Input[str] server_fqdn: FQDN of syslog server that FortiAP units send log messages to.
        :param pulumi.Input[str] server_ip: IP address of syslog server that FortiAP units send log messages to.
        :param pulumi.Input[int] server_port: Port number of syslog server that FortiAP units send log messages to (default = 514).
        :param pulumi.Input[str] server_status: Enable/disable FortiAP units to send log messages to a syslog server (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if log_level is not None:
            pulumi.set(__self__, "log_level", log_level)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if server_addr_type is not None:
            pulumi.set(__self__, "server_addr_type", server_addr_type)
        if server_fqdn is not None:
            pulumi.set(__self__, "server_fqdn", server_fqdn)
        if server_ip is not None:
            pulumi.set(__self__, "server_ip", server_ip)
        if server_port is not None:
            pulumi.set(__self__, "server_port", server_port)
        if server_status is not None:
            pulumi.set(__self__, "server_status", server_status)
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
    @pulumi.getter(name="logLevel")
    def log_level(self) -> Optional[pulumi.Input[str]]:
        """
        Lowest level of log messages that FortiAP units send to this server (default = information) Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debugging`.
        """
        return pulumi.get(self, "log_level")

    @log_level.setter
    def log_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_level", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        WTP system log server profile name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serverAddrType")
    def server_addr_type(self) -> Optional[pulumi.Input[str]]:
        """
        Syslog server address type (default = IP) Valid values: `fqdn`, `ip`.
        """
        return pulumi.get(self, "server_addr_type")

    @server_addr_type.setter
    def server_addr_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_addr_type", value)

    @property
    @pulumi.getter(name="serverFqdn")
    def server_fqdn(self) -> Optional[pulumi.Input[str]]:
        """
        FQDN of syslog server that FortiAP units send log messages to.
        """
        return pulumi.get(self, "server_fqdn")

    @server_fqdn.setter
    def server_fqdn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_fqdn", value)

    @property
    @pulumi.getter(name="serverIp")
    def server_ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address of syslog server that FortiAP units send log messages to.
        """
        return pulumi.get(self, "server_ip")

    @server_ip.setter
    def server_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_ip", value)

    @property
    @pulumi.getter(name="serverPort")
    def server_port(self) -> Optional[pulumi.Input[int]]:
        """
        Port number of syslog server that FortiAP units send log messages to (default = 514).
        """
        return pulumi.get(self, "server_port")

    @server_port.setter
    def server_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "server_port", value)

    @property
    @pulumi.getter(name="serverStatus")
    def server_status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FortiAP units to send log messages to a syslog server (default = enable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "server_status")

    @server_status.setter
    def server_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_status", value)

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
class _WirelesscontrollerSyslogprofileState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 log_level: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_addr_type: Optional[pulumi.Input[str]] = None,
                 server_fqdn: Optional[pulumi.Input[str]] = None,
                 server_ip: Optional[pulumi.Input[str]] = None,
                 server_port: Optional[pulumi.Input[int]] = None,
                 server_status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WirelesscontrollerSyslogprofile resources.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] log_level: Lowest level of log messages that FortiAP units send to this server (default = information) Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debugging`.
        :param pulumi.Input[str] name: WTP system log server profile name.
        :param pulumi.Input[str] server_addr_type: Syslog server address type (default = IP) Valid values: `fqdn`, `ip`.
        :param pulumi.Input[str] server_fqdn: FQDN of syslog server that FortiAP units send log messages to.
        :param pulumi.Input[str] server_ip: IP address of syslog server that FortiAP units send log messages to.
        :param pulumi.Input[int] server_port: Port number of syslog server that FortiAP units send log messages to (default = 514).
        :param pulumi.Input[str] server_status: Enable/disable FortiAP units to send log messages to a syslog server (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if log_level is not None:
            pulumi.set(__self__, "log_level", log_level)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if server_addr_type is not None:
            pulumi.set(__self__, "server_addr_type", server_addr_type)
        if server_fqdn is not None:
            pulumi.set(__self__, "server_fqdn", server_fqdn)
        if server_ip is not None:
            pulumi.set(__self__, "server_ip", server_ip)
        if server_port is not None:
            pulumi.set(__self__, "server_port", server_port)
        if server_status is not None:
            pulumi.set(__self__, "server_status", server_status)
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
    @pulumi.getter(name="logLevel")
    def log_level(self) -> Optional[pulumi.Input[str]]:
        """
        Lowest level of log messages that FortiAP units send to this server (default = information) Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debugging`.
        """
        return pulumi.get(self, "log_level")

    @log_level.setter
    def log_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_level", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        WTP system log server profile name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serverAddrType")
    def server_addr_type(self) -> Optional[pulumi.Input[str]]:
        """
        Syslog server address type (default = IP) Valid values: `fqdn`, `ip`.
        """
        return pulumi.get(self, "server_addr_type")

    @server_addr_type.setter
    def server_addr_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_addr_type", value)

    @property
    @pulumi.getter(name="serverFqdn")
    def server_fqdn(self) -> Optional[pulumi.Input[str]]:
        """
        FQDN of syslog server that FortiAP units send log messages to.
        """
        return pulumi.get(self, "server_fqdn")

    @server_fqdn.setter
    def server_fqdn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_fqdn", value)

    @property
    @pulumi.getter(name="serverIp")
    def server_ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address of syslog server that FortiAP units send log messages to.
        """
        return pulumi.get(self, "server_ip")

    @server_ip.setter
    def server_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_ip", value)

    @property
    @pulumi.getter(name="serverPort")
    def server_port(self) -> Optional[pulumi.Input[int]]:
        """
        Port number of syslog server that FortiAP units send log messages to (default = 514).
        """
        return pulumi.get(self, "server_port")

    @server_port.setter
    def server_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "server_port", value)

    @property
    @pulumi.getter(name="serverStatus")
    def server_status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FortiAP units to send log messages to a syslog server (default = enable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "server_status")

    @server_status.setter
    def server_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_status", value)

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


class WirelesscontrollerSyslogprofile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 log_level: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_addr_type: Optional[pulumi.Input[str]] = None,
                 server_fqdn: Optional[pulumi.Input[str]] = None,
                 server_ip: Optional[pulumi.Input[str]] = None,
                 server_port: Optional[pulumi.Input[int]] = None,
                 server_status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure Wireless Termination Points (WTP) system log server profile. Applies to FortiOS Version `>= 7.0.2`.

        ## Import

        WirelessController SyslogProfile can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/wirelesscontrollerSyslogprofile:WirelesscontrollerSyslogprofile labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/wirelesscontrollerSyslogprofile:WirelesscontrollerSyslogprofile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] log_level: Lowest level of log messages that FortiAP units send to this server (default = information) Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debugging`.
        :param pulumi.Input[str] name: WTP system log server profile name.
        :param pulumi.Input[str] server_addr_type: Syslog server address type (default = IP) Valid values: `fqdn`, `ip`.
        :param pulumi.Input[str] server_fqdn: FQDN of syslog server that FortiAP units send log messages to.
        :param pulumi.Input[str] server_ip: IP address of syslog server that FortiAP units send log messages to.
        :param pulumi.Input[int] server_port: Port number of syslog server that FortiAP units send log messages to (default = 514).
        :param pulumi.Input[str] server_status: Enable/disable FortiAP units to send log messages to a syslog server (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WirelesscontrollerSyslogprofileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure Wireless Termination Points (WTP) system log server profile. Applies to FortiOS Version `>= 7.0.2`.

        ## Import

        WirelessController SyslogProfile can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/wirelesscontrollerSyslogprofile:WirelesscontrollerSyslogprofile labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/wirelesscontrollerSyslogprofile:WirelesscontrollerSyslogprofile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WirelesscontrollerSyslogprofileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WirelesscontrollerSyslogprofileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 log_level: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_addr_type: Optional[pulumi.Input[str]] = None,
                 server_fqdn: Optional[pulumi.Input[str]] = None,
                 server_ip: Optional[pulumi.Input[str]] = None,
                 server_port: Optional[pulumi.Input[int]] = None,
                 server_status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WirelesscontrollerSyslogprofileArgs.__new__(WirelesscontrollerSyslogprofileArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["log_level"] = log_level
            __props__.__dict__["name"] = name
            __props__.__dict__["server_addr_type"] = server_addr_type
            __props__.__dict__["server_fqdn"] = server_fqdn
            __props__.__dict__["server_ip"] = server_ip
            __props__.__dict__["server_port"] = server_port
            __props__.__dict__["server_status"] = server_status
            __props__.__dict__["vdomparam"] = vdomparam
        super(WirelesscontrollerSyslogprofile, __self__).__init__(
            'fortios:index/wirelesscontrollerSyslogprofile:WirelesscontrollerSyslogprofile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            log_level: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            server_addr_type: Optional[pulumi.Input[str]] = None,
            server_fqdn: Optional[pulumi.Input[str]] = None,
            server_ip: Optional[pulumi.Input[str]] = None,
            server_port: Optional[pulumi.Input[int]] = None,
            server_status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WirelesscontrollerSyslogprofile':
        """
        Get an existing WirelesscontrollerSyslogprofile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] log_level: Lowest level of log messages that FortiAP units send to this server (default = information) Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debugging`.
        :param pulumi.Input[str] name: WTP system log server profile name.
        :param pulumi.Input[str] server_addr_type: Syslog server address type (default = IP) Valid values: `fqdn`, `ip`.
        :param pulumi.Input[str] server_fqdn: FQDN of syslog server that FortiAP units send log messages to.
        :param pulumi.Input[str] server_ip: IP address of syslog server that FortiAP units send log messages to.
        :param pulumi.Input[int] server_port: Port number of syslog server that FortiAP units send log messages to (default = 514).
        :param pulumi.Input[str] server_status: Enable/disable FortiAP units to send log messages to a syslog server (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WirelesscontrollerSyslogprofileState.__new__(_WirelesscontrollerSyslogprofileState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["log_level"] = log_level
        __props__.__dict__["name"] = name
        __props__.__dict__["server_addr_type"] = server_addr_type
        __props__.__dict__["server_fqdn"] = server_fqdn
        __props__.__dict__["server_ip"] = server_ip
        __props__.__dict__["server_port"] = server_port
        __props__.__dict__["server_status"] = server_status
        __props__.__dict__["vdomparam"] = vdomparam
        return WirelesscontrollerSyslogprofile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="logLevel")
    def log_level(self) -> pulumi.Output[str]:
        """
        Lowest level of log messages that FortiAP units send to this server (default = information) Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debugging`.
        """
        return pulumi.get(self, "log_level")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        WTP system log server profile name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serverAddrType")
    def server_addr_type(self) -> pulumi.Output[str]:
        """
        Syslog server address type (default = IP) Valid values: `fqdn`, `ip`.
        """
        return pulumi.get(self, "server_addr_type")

    @property
    @pulumi.getter(name="serverFqdn")
    def server_fqdn(self) -> pulumi.Output[str]:
        """
        FQDN of syslog server that FortiAP units send log messages to.
        """
        return pulumi.get(self, "server_fqdn")

    @property
    @pulumi.getter(name="serverIp")
    def server_ip(self) -> pulumi.Output[str]:
        """
        IP address of syslog server that FortiAP units send log messages to.
        """
        return pulumi.get(self, "server_ip")

    @property
    @pulumi.getter(name="serverPort")
    def server_port(self) -> pulumi.Output[int]:
        """
        Port number of syslog server that FortiAP units send log messages to (default = 514).
        """
        return pulumi.get(self, "server_port")

    @property
    @pulumi.getter(name="serverStatus")
    def server_status(self) -> pulumi.Output[str]:
        """
        Enable/disable FortiAP units to send log messages to a syslog server (default = enable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "server_status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

