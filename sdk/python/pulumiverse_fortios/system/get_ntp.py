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

__all__ = [
    'GetNtpResult',
    'AwaitableGetNtpResult',
    'get_ntp',
    'get_ntp_output',
]

@pulumi.output_type
class GetNtpResult:
    """
    A collection of values returned by getNtp.
    """
    def __init__(__self__, authentication=None, id=None, interfaces=None, key=None, key_id=None, key_type=None, ntpservers=None, ntpsync=None, server_mode=None, source_ip=None, source_ip6=None, syncinterval=None, type=None, vdomparam=None):
        if authentication and not isinstance(authentication, str):
            raise TypeError("Expected argument 'authentication' to be a str")
        pulumi.set(__self__, "authentication", authentication)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interfaces and not isinstance(interfaces, list):
            raise TypeError("Expected argument 'interfaces' to be a list")
        pulumi.set(__self__, "interfaces", interfaces)
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        pulumi.set(__self__, "key", key)
        if key_id and not isinstance(key_id, int):
            raise TypeError("Expected argument 'key_id' to be a int")
        pulumi.set(__self__, "key_id", key_id)
        if key_type and not isinstance(key_type, str):
            raise TypeError("Expected argument 'key_type' to be a str")
        pulumi.set(__self__, "key_type", key_type)
        if ntpservers and not isinstance(ntpservers, list):
            raise TypeError("Expected argument 'ntpservers' to be a list")
        pulumi.set(__self__, "ntpservers", ntpservers)
        if ntpsync and not isinstance(ntpsync, str):
            raise TypeError("Expected argument 'ntpsync' to be a str")
        pulumi.set(__self__, "ntpsync", ntpsync)
        if server_mode and not isinstance(server_mode, str):
            raise TypeError("Expected argument 'server_mode' to be a str")
        pulumi.set(__self__, "server_mode", server_mode)
        if source_ip and not isinstance(source_ip, str):
            raise TypeError("Expected argument 'source_ip' to be a str")
        pulumi.set(__self__, "source_ip", source_ip)
        if source_ip6 and not isinstance(source_ip6, str):
            raise TypeError("Expected argument 'source_ip6' to be a str")
        pulumi.set(__self__, "source_ip6", source_ip6)
        if syncinterval and not isinstance(syncinterval, int):
            raise TypeError("Expected argument 'syncinterval' to be a int")
        pulumi.set(__self__, "syncinterval", syncinterval)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def authentication(self) -> str:
        """
        Enable/disable MD5/SHA1 authentication.
        """
        return pulumi.get(self, "authentication")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interfaces(self) -> Sequence['outputs.GetNtpInterfaceResult']:
        """
        Specify outgoing interface to reach server.
        """
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Key for MD5/SHA1 authentication.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> int:
        """
        Key ID for authentication.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> str:
        """
        Select NTP authentication type.
        """
        return pulumi.get(self, "key_type")

    @property
    @pulumi.getter
    def ntpservers(self) -> Sequence['outputs.GetNtpNtpserverResult']:
        """
        Configure the FortiGate to connect to any available third-party NTP server. The structure of `ntpserver` block is documented below.
        """
        return pulumi.get(self, "ntpservers")

    @property
    @pulumi.getter
    def ntpsync(self) -> str:
        """
        Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.
        """
        return pulumi.get(self, "ntpsync")

    @property
    @pulumi.getter(name="serverMode")
    def server_mode(self) -> str:
        """
        Enable/disable FortiGate NTP Server Mode. Your FortiGate becomes an NTP server for other devices on your network. The FortiGate relays NTP requests to its configured NTP server.
        """
        return pulumi.get(self, "server_mode")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> str:
        """
        Source IP address for communication to the NTP server.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="sourceIp6")
    def source_ip6(self) -> str:
        """
        Source IPv6 address for communication to the NTP server.
        """
        return pulumi.get(self, "source_ip6")

    @property
    @pulumi.getter
    def syncinterval(self) -> int:
        """
        NTP synchronization interval (1 - 1440 min).
        """
        return pulumi.get(self, "syncinterval")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Use the FortiGuard NTP server or any other available NTP Server.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetNtpResult(GetNtpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNtpResult(
            authentication=self.authentication,
            id=self.id,
            interfaces=self.interfaces,
            key=self.key,
            key_id=self.key_id,
            key_type=self.key_type,
            ntpservers=self.ntpservers,
            ntpsync=self.ntpsync,
            server_mode=self.server_mode,
            source_ip=self.source_ip,
            source_ip6=self.source_ip6,
            syncinterval=self.syncinterval,
            type=self.type,
            vdomparam=self.vdomparam)


def get_ntp(vdomparam: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNtpResult:
    """
    Use this data source to get information on fortios system ntp


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:system/getNtp:getNtp', __args__, opts=opts, typ=GetNtpResult).value

    return AwaitableGetNtpResult(
        authentication=pulumi.get(__ret__, 'authentication'),
        id=pulumi.get(__ret__, 'id'),
        interfaces=pulumi.get(__ret__, 'interfaces'),
        key=pulumi.get(__ret__, 'key'),
        key_id=pulumi.get(__ret__, 'key_id'),
        key_type=pulumi.get(__ret__, 'key_type'),
        ntpservers=pulumi.get(__ret__, 'ntpservers'),
        ntpsync=pulumi.get(__ret__, 'ntpsync'),
        server_mode=pulumi.get(__ret__, 'server_mode'),
        source_ip=pulumi.get(__ret__, 'source_ip'),
        source_ip6=pulumi.get(__ret__, 'source_ip6'),
        syncinterval=pulumi.get(__ret__, 'syncinterval'),
        type=pulumi.get(__ret__, 'type'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_ntp)
def get_ntp_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNtpResult]:
    """
    Use this data source to get information on fortios system ntp


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
