# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['WispArgs', 'Wisp']

@pulumi.input_type
class WispArgs:
    def __init__(__self__, *,
                 server_ip: pulumi.Input[str],
                 server_port: pulumi.Input[int],
                 comment: Optional[pulumi.Input[str]] = None,
                 max_connections: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 outgoing_ip: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Wisp resource.
        :param pulumi.Input[str] server_ip: WISP server IP address.
        :param pulumi.Input[int] server_port: WISP server port (1 - 65535, default = 15868).
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[int] max_connections: Maximum number of web proxy WISP connections (4 - 4096, default = 64).
        :param pulumi.Input[str] name: Server name.
        :param pulumi.Input[str] outgoing_ip: WISP outgoing IP address.
        :param pulumi.Input[int] timeout: Period of time before WISP requests time out (1 - 15 sec, default = 5).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "server_ip", server_ip)
        pulumi.set(__self__, "server_port", server_port)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if max_connections is not None:
            pulumi.set(__self__, "max_connections", max_connections)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if outgoing_ip is not None:
            pulumi.set(__self__, "outgoing_ip", outgoing_ip)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="serverIp")
    def server_ip(self) -> pulumi.Input[str]:
        """
        WISP server IP address.
        """
        return pulumi.get(self, "server_ip")

    @server_ip.setter
    def server_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_ip", value)

    @property
    @pulumi.getter(name="serverPort")
    def server_port(self) -> pulumi.Input[int]:
        """
        WISP server port (1 - 65535, default = 15868).
        """
        return pulumi.get(self, "server_port")

    @server_port.setter
    def server_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "server_port", value)

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
    @pulumi.getter(name="maxConnections")
    def max_connections(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of web proxy WISP connections (4 - 4096, default = 64).
        """
        return pulumi.get(self, "max_connections")

    @max_connections.setter
    def max_connections(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_connections", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Server name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="outgoingIp")
    def outgoing_ip(self) -> Optional[pulumi.Input[str]]:
        """
        WISP outgoing IP address.
        """
        return pulumi.get(self, "outgoing_ip")

    @outgoing_ip.setter
    def outgoing_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "outgoing_ip", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Period of time before WISP requests time out (1 - 15 sec, default = 5).
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

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
class _WispState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 max_connections: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 outgoing_ip: Optional[pulumi.Input[str]] = None,
                 server_ip: Optional[pulumi.Input[str]] = None,
                 server_port: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Wisp resources.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[int] max_connections: Maximum number of web proxy WISP connections (4 - 4096, default = 64).
        :param pulumi.Input[str] name: Server name.
        :param pulumi.Input[str] outgoing_ip: WISP outgoing IP address.
        :param pulumi.Input[str] server_ip: WISP server IP address.
        :param pulumi.Input[int] server_port: WISP server port (1 - 65535, default = 15868).
        :param pulumi.Input[int] timeout: Period of time before WISP requests time out (1 - 15 sec, default = 5).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if max_connections is not None:
            pulumi.set(__self__, "max_connections", max_connections)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if outgoing_ip is not None:
            pulumi.set(__self__, "outgoing_ip", outgoing_ip)
        if server_ip is not None:
            pulumi.set(__self__, "server_ip", server_ip)
        if server_port is not None:
            pulumi.set(__self__, "server_port", server_port)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
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
    @pulumi.getter(name="maxConnections")
    def max_connections(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of web proxy WISP connections (4 - 4096, default = 64).
        """
        return pulumi.get(self, "max_connections")

    @max_connections.setter
    def max_connections(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_connections", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Server name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="outgoingIp")
    def outgoing_ip(self) -> Optional[pulumi.Input[str]]:
        """
        WISP outgoing IP address.
        """
        return pulumi.get(self, "outgoing_ip")

    @outgoing_ip.setter
    def outgoing_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "outgoing_ip", value)

    @property
    @pulumi.getter(name="serverIp")
    def server_ip(self) -> Optional[pulumi.Input[str]]:
        """
        WISP server IP address.
        """
        return pulumi.get(self, "server_ip")

    @server_ip.setter
    def server_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_ip", value)

    @property
    @pulumi.getter(name="serverPort")
    def server_port(self) -> Optional[pulumi.Input[int]]:
        """
        WISP server port (1 - 65535, default = 15868).
        """
        return pulumi.get(self, "server_port")

    @server_port.setter
    def server_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "server_port", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Period of time before WISP requests time out (1 - 15 sec, default = 5).
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

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


class Wisp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 max_connections: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 outgoing_ip: Optional[pulumi.Input[str]] = None,
                 server_ip: Optional[pulumi.Input[str]] = None,
                 server_port: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure Wireless Internet service provider (WISP) servers.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.webproxy.Wisp("trname",
            max_connections=64,
            outgoing_ip="0.0.0.0",
            server_ip="1.1.1.1",
            server_port=15868,
            timeout=5)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        WebProxy Wisp can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:webproxy/wisp:Wisp labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:webproxy/wisp:Wisp labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[int] max_connections: Maximum number of web proxy WISP connections (4 - 4096, default = 64).
        :param pulumi.Input[str] name: Server name.
        :param pulumi.Input[str] outgoing_ip: WISP outgoing IP address.
        :param pulumi.Input[str] server_ip: WISP server IP address.
        :param pulumi.Input[int] server_port: WISP server port (1 - 65535, default = 15868).
        :param pulumi.Input[int] timeout: Period of time before WISP requests time out (1 - 15 sec, default = 5).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WispArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure Wireless Internet service provider (WISP) servers.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.webproxy.Wisp("trname",
            max_connections=64,
            outgoing_ip="0.0.0.0",
            server_ip="1.1.1.1",
            server_port=15868,
            timeout=5)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        WebProxy Wisp can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:webproxy/wisp:Wisp labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:webproxy/wisp:Wisp labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WispArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WispArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 max_connections: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 outgoing_ip: Optional[pulumi.Input[str]] = None,
                 server_ip: Optional[pulumi.Input[str]] = None,
                 server_port: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WispArgs.__new__(WispArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["max_connections"] = max_connections
            __props__.__dict__["name"] = name
            __props__.__dict__["outgoing_ip"] = outgoing_ip
            if server_ip is None and not opts.urn:
                raise TypeError("Missing required property 'server_ip'")
            __props__.__dict__["server_ip"] = server_ip
            if server_port is None and not opts.urn:
                raise TypeError("Missing required property 'server_port'")
            __props__.__dict__["server_port"] = server_port
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["vdomparam"] = vdomparam
        super(Wisp, __self__).__init__(
            'fortios:webproxy/wisp:Wisp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            max_connections: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            outgoing_ip: Optional[pulumi.Input[str]] = None,
            server_ip: Optional[pulumi.Input[str]] = None,
            server_port: Optional[pulumi.Input[int]] = None,
            timeout: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Wisp':
        """
        Get an existing Wisp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[int] max_connections: Maximum number of web proxy WISP connections (4 - 4096, default = 64).
        :param pulumi.Input[str] name: Server name.
        :param pulumi.Input[str] outgoing_ip: WISP outgoing IP address.
        :param pulumi.Input[str] server_ip: WISP server IP address.
        :param pulumi.Input[int] server_port: WISP server port (1 - 65535, default = 15868).
        :param pulumi.Input[int] timeout: Period of time before WISP requests time out (1 - 15 sec, default = 5).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WispState.__new__(_WispState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["max_connections"] = max_connections
        __props__.__dict__["name"] = name
        __props__.__dict__["outgoing_ip"] = outgoing_ip
        __props__.__dict__["server_ip"] = server_ip
        __props__.__dict__["server_port"] = server_port
        __props__.__dict__["timeout"] = timeout
        __props__.__dict__["vdomparam"] = vdomparam
        return Wisp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="maxConnections")
    def max_connections(self) -> pulumi.Output[int]:
        """
        Maximum number of web proxy WISP connections (4 - 4096, default = 64).
        """
        return pulumi.get(self, "max_connections")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Server name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outgoingIp")
    def outgoing_ip(self) -> pulumi.Output[str]:
        """
        WISP outgoing IP address.
        """
        return pulumi.get(self, "outgoing_ip")

    @property
    @pulumi.getter(name="serverIp")
    def server_ip(self) -> pulumi.Output[str]:
        """
        WISP server IP address.
        """
        return pulumi.get(self, "server_ip")

    @property
    @pulumi.getter(name="serverPort")
    def server_port(self) -> pulumi.Output[int]:
        """
        WISP server port (1 - 65535, default = 15868).
        """
        return pulumi.get(self, "server_port")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[int]:
        """
        Period of time before WISP requests time out (1 - 15 sec, default = 5).
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

