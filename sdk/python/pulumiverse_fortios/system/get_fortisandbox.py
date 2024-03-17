# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetFortisandboxResult',
    'AwaitableGetFortisandboxResult',
    'get_fortisandbox',
    'get_fortisandbox_output',
]

@pulumi.output_type
class GetFortisandboxResult:
    """
    A collection of values returned by getFortisandbox.
    """
    def __init__(__self__, email=None, enc_algorithm=None, forticloud=None, id=None, inline_scan=None, interface=None, interface_select_method=None, server=None, source_ip=None, ssl_min_proto_version=None, status=None, vdomparam=None):
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if enc_algorithm and not isinstance(enc_algorithm, str):
            raise TypeError("Expected argument 'enc_algorithm' to be a str")
        pulumi.set(__self__, "enc_algorithm", enc_algorithm)
        if forticloud and not isinstance(forticloud, str):
            raise TypeError("Expected argument 'forticloud' to be a str")
        pulumi.set(__self__, "forticloud", forticloud)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if inline_scan and not isinstance(inline_scan, str):
            raise TypeError("Expected argument 'inline_scan' to be a str")
        pulumi.set(__self__, "inline_scan", inline_scan)
        if interface and not isinstance(interface, str):
            raise TypeError("Expected argument 'interface' to be a str")
        pulumi.set(__self__, "interface", interface)
        if interface_select_method and not isinstance(interface_select_method, str):
            raise TypeError("Expected argument 'interface_select_method' to be a str")
        pulumi.set(__self__, "interface_select_method", interface_select_method)
        if server and not isinstance(server, str):
            raise TypeError("Expected argument 'server' to be a str")
        pulumi.set(__self__, "server", server)
        if source_ip and not isinstance(source_ip, str):
            raise TypeError("Expected argument 'source_ip' to be a str")
        pulumi.set(__self__, "source_ip", source_ip)
        if ssl_min_proto_version and not isinstance(ssl_min_proto_version, str):
            raise TypeError("Expected argument 'ssl_min_proto_version' to be a str")
        pulumi.set(__self__, "ssl_min_proto_version", ssl_min_proto_version)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        Notifier email address.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="encAlgorithm")
    def enc_algorithm(self) -> str:
        """
        Configure the level of SSL protection for secure communication with FortiSandbox.
        """
        return pulumi.get(self, "enc_algorithm")

    @property
    @pulumi.getter
    def forticloud(self) -> str:
        """
        Enable/disable FortiSandbox Cloud.
        """
        return pulumi.get(self, "forticloud")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="inlineScan")
    def inline_scan(self) -> str:
        """
        Enable/disable FortiSandbox inline scan.
        """
        return pulumi.get(self, "inline_scan")

    @property
    @pulumi.getter
    def interface(self) -> str:
        """
        Specify outgoing interface to reach server.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="interfaceSelectMethod")
    def interface_select_method(self) -> str:
        """
        Specify how to select outgoing interface to reach server.
        """
        return pulumi.get(self, "interface_select_method")

    @property
    @pulumi.getter
    def server(self) -> str:
        """
        IPv4 or IPv6 address of the remote FortiSandbox.
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> str:
        """
        Source IP address for communications to FortiSandbox.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="sslMinProtoVersion")
    def ssl_min_proto_version(self) -> str:
        """
        Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        """
        return pulumi.get(self, "ssl_min_proto_version")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Enable/disable FortiSandbox.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetFortisandboxResult(GetFortisandboxResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFortisandboxResult(
            email=self.email,
            enc_algorithm=self.enc_algorithm,
            forticloud=self.forticloud,
            id=self.id,
            inline_scan=self.inline_scan,
            interface=self.interface,
            interface_select_method=self.interface_select_method,
            server=self.server,
            source_ip=self.source_ip,
            ssl_min_proto_version=self.ssl_min_proto_version,
            status=self.status,
            vdomparam=self.vdomparam)


def get_fortisandbox(vdomparam: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFortisandboxResult:
    """
    Use this data source to get information on fortios system fortisandbox


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:system/getFortisandbox:getFortisandbox', __args__, opts=opts, typ=GetFortisandboxResult).value

    return AwaitableGetFortisandboxResult(
        email=pulumi.get(__ret__, 'email'),
        enc_algorithm=pulumi.get(__ret__, 'enc_algorithm'),
        forticloud=pulumi.get(__ret__, 'forticloud'),
        id=pulumi.get(__ret__, 'id'),
        inline_scan=pulumi.get(__ret__, 'inline_scan'),
        interface=pulumi.get(__ret__, 'interface'),
        interface_select_method=pulumi.get(__ret__, 'interface_select_method'),
        server=pulumi.get(__ret__, 'server'),
        source_ip=pulumi.get(__ret__, 'source_ip'),
        ssl_min_proto_version=pulumi.get(__ret__, 'ssl_min_proto_version'),
        status=pulumi.get(__ret__, 'status'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_fortisandbox)
def get_fortisandbox_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFortisandboxResult]:
    """
    Use this data source to get information on fortios system fortisandbox


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
