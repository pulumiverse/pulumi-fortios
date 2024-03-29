# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['Pop3Args', 'Pop3']

@pulumi.input_type
class Pop3Args:
    def __init__(__self__, *,
                 server: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 secure: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Pop3 resource.
        :param pulumi.Input[str] server: {<name_str|ip_str>} server domain name or IP.
        :param pulumi.Input[str] name: POP3 server entry name.
        :param pulumi.Input[int] port: POP3 service port number.
        :param pulumi.Input[str] secure: SSL connection. Valid values: `none`, `starttls`, `pop3s`.
        :param pulumi.Input[str] ssl_min_proto_version: Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "server", server)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if secure is not None:
            pulumi.set(__self__, "secure", secure)
        if ssl_min_proto_version is not None:
            pulumi.set(__self__, "ssl_min_proto_version", ssl_min_proto_version)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def server(self) -> pulumi.Input[str]:
        """
        {<name_str|ip_str>} server domain name or IP.
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: pulumi.Input[str]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        POP3 server entry name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        POP3 service port number.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def secure(self) -> Optional[pulumi.Input[str]]:
        """
        SSL connection. Valid values: `none`, `starttls`, `pop3s`.
        """
        return pulumi.get(self, "secure")

    @secure.setter
    def secure(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secure", value)

    @property
    @pulumi.getter(name="sslMinProtoVersion")
    def ssl_min_proto_version(self) -> Optional[pulumi.Input[str]]:
        """
        Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        """
        return pulumi.get(self, "ssl_min_proto_version")

    @ssl_min_proto_version.setter
    def ssl_min_proto_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_min_proto_version", value)

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
class _Pop3State:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 secure: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Pop3 resources.
        :param pulumi.Input[str] name: POP3 server entry name.
        :param pulumi.Input[int] port: POP3 service port number.
        :param pulumi.Input[str] secure: SSL connection. Valid values: `none`, `starttls`, `pop3s`.
        :param pulumi.Input[str] server: {<name_str|ip_str>} server domain name or IP.
        :param pulumi.Input[str] ssl_min_proto_version: Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if secure is not None:
            pulumi.set(__self__, "secure", secure)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if ssl_min_proto_version is not None:
            pulumi.set(__self__, "ssl_min_proto_version", ssl_min_proto_version)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        POP3 server entry name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        POP3 service port number.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def secure(self) -> Optional[pulumi.Input[str]]:
        """
        SSL connection. Valid values: `none`, `starttls`, `pop3s`.
        """
        return pulumi.get(self, "secure")

    @secure.setter
    def secure(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secure", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        """
        {<name_str|ip_str>} server domain name or IP.
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter(name="sslMinProtoVersion")
    def ssl_min_proto_version(self) -> Optional[pulumi.Input[str]]:
        """
        Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        """
        return pulumi.get(self, "ssl_min_proto_version")

    @ssl_min_proto_version.setter
    def ssl_min_proto_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_min_proto_version", value)

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


class Pop3(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 secure: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        POP3 server entry configuration.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.user.Pop3("trname",
            port=0,
            secure="pop3s",
            server="1.1.1.1",
            ssl_min_proto_version="default")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        User Pop3 can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:user/pop3:Pop3 labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:user/pop3:Pop3 labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: POP3 server entry name.
        :param pulumi.Input[int] port: POP3 service port number.
        :param pulumi.Input[str] secure: SSL connection. Valid values: `none`, `starttls`, `pop3s`.
        :param pulumi.Input[str] server: {<name_str|ip_str>} server domain name or IP.
        :param pulumi.Input[str] ssl_min_proto_version: Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Pop3Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        POP3 server entry configuration.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.user.Pop3("trname",
            port=0,
            secure="pop3s",
            server="1.1.1.1",
            ssl_min_proto_version="default")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        User Pop3 can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:user/pop3:Pop3 labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:user/pop3:Pop3 labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param Pop3Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Pop3Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 secure: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = Pop3Args.__new__(Pop3Args)

            __props__.__dict__["name"] = name
            __props__.__dict__["port"] = port
            __props__.__dict__["secure"] = secure
            if server is None and not opts.urn:
                raise TypeError("Missing required property 'server'")
            __props__.__dict__["server"] = server
            __props__.__dict__["ssl_min_proto_version"] = ssl_min_proto_version
            __props__.__dict__["vdomparam"] = vdomparam
        super(Pop3, __self__).__init__(
            'fortios:user/pop3:Pop3',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            secure: Optional[pulumi.Input[str]] = None,
            server: Optional[pulumi.Input[str]] = None,
            ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Pop3':
        """
        Get an existing Pop3 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: POP3 server entry name.
        :param pulumi.Input[int] port: POP3 service port number.
        :param pulumi.Input[str] secure: SSL connection. Valid values: `none`, `starttls`, `pop3s`.
        :param pulumi.Input[str] server: {<name_str|ip_str>} server domain name or IP.
        :param pulumi.Input[str] ssl_min_proto_version: Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _Pop3State.__new__(_Pop3State)

        __props__.__dict__["name"] = name
        __props__.__dict__["port"] = port
        __props__.__dict__["secure"] = secure
        __props__.__dict__["server"] = server
        __props__.__dict__["ssl_min_proto_version"] = ssl_min_proto_version
        __props__.__dict__["vdomparam"] = vdomparam
        return Pop3(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        POP3 server entry name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        POP3 service port number.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def secure(self) -> pulumi.Output[str]:
        """
        SSL connection. Valid values: `none`, `starttls`, `pop3s`.
        """
        return pulumi.get(self, "secure")

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output[str]:
        """
        {<name_str|ip_str>} server domain name or IP.
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter(name="sslMinProtoVersion")
    def ssl_min_proto_version(self) -> pulumi.Output[str]:
        """
        Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        """
        return pulumi.get(self, "ssl_min_proto_version")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

