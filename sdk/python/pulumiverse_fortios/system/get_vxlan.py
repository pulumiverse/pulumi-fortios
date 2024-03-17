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
    'GetVxlanResult',
    'AwaitableGetVxlanResult',
    'get_vxlan',
    'get_vxlan_output',
]

@pulumi.output_type
class GetVxlanResult:
    """
    A collection of values returned by getVxlan.
    """
    def __init__(__self__, dstport=None, id=None, interface=None, ip_version=None, multicast_ttl=None, name=None, remote_ip6s=None, remote_ips=None, vdomparam=None, vni=None):
        if dstport and not isinstance(dstport, int):
            raise TypeError("Expected argument 'dstport' to be a int")
        pulumi.set(__self__, "dstport", dstport)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface and not isinstance(interface, str):
            raise TypeError("Expected argument 'interface' to be a str")
        pulumi.set(__self__, "interface", interface)
        if ip_version and not isinstance(ip_version, str):
            raise TypeError("Expected argument 'ip_version' to be a str")
        pulumi.set(__self__, "ip_version", ip_version)
        if multicast_ttl and not isinstance(multicast_ttl, int):
            raise TypeError("Expected argument 'multicast_ttl' to be a int")
        pulumi.set(__self__, "multicast_ttl", multicast_ttl)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if remote_ip6s and not isinstance(remote_ip6s, list):
            raise TypeError("Expected argument 'remote_ip6s' to be a list")
        pulumi.set(__self__, "remote_ip6s", remote_ip6s)
        if remote_ips and not isinstance(remote_ips, list):
            raise TypeError("Expected argument 'remote_ips' to be a list")
        pulumi.set(__self__, "remote_ips", remote_ips)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if vni and not isinstance(vni, int):
            raise TypeError("Expected argument 'vni' to be a int")
        pulumi.set(__self__, "vni", vni)

    @property
    @pulumi.getter
    def dstport(self) -> int:
        """
        VXLAN destination port (1 - 65535, default = 4789).
        """
        return pulumi.get(self, "dstport")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interface(self) -> str:
        """
        Outgoing interface for VXLAN encapsulated traffic.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> str:
        """
        IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast.
        """
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter(name="multicastTtl")
    def multicast_ttl(self) -> int:
        """
        VXLAN multicast TTL (1-255, default = 0).
        """
        return pulumi.get(self, "multicast_ttl")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        VXLAN device or interface name. Must be a unique interface name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="remoteIp6s")
    def remote_ip6s(self) -> Sequence['outputs.GetVxlanRemoteIp6Result']:
        """
        IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remote_ip6` block is documented below.
        """
        return pulumi.get(self, "remote_ip6s")

    @property
    @pulumi.getter(name="remoteIps")
    def remote_ips(self) -> Sequence['outputs.GetVxlanRemoteIpResult']:
        """
        IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remote_ip` block is documented below.
        """
        return pulumi.get(self, "remote_ips")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vni(self) -> int:
        """
        VXLAN network ID.
        """
        return pulumi.get(self, "vni")


class AwaitableGetVxlanResult(GetVxlanResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVxlanResult(
            dstport=self.dstport,
            id=self.id,
            interface=self.interface,
            ip_version=self.ip_version,
            multicast_ttl=self.multicast_ttl,
            name=self.name,
            remote_ip6s=self.remote_ip6s,
            remote_ips=self.remote_ips,
            vdomparam=self.vdomparam,
            vni=self.vni)


def get_vxlan(name: Optional[str] = None,
              vdomparam: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVxlanResult:
    """
    Use this data source to get information on an fortios system vxlan


    :param str name: Specify the name of the desired system vxlan.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:system/getVxlan:getVxlan', __args__, opts=opts, typ=GetVxlanResult).value

    return AwaitableGetVxlanResult(
        dstport=pulumi.get(__ret__, 'dstport'),
        id=pulumi.get(__ret__, 'id'),
        interface=pulumi.get(__ret__, 'interface'),
        ip_version=pulumi.get(__ret__, 'ip_version'),
        multicast_ttl=pulumi.get(__ret__, 'multicast_ttl'),
        name=pulumi.get(__ret__, 'name'),
        remote_ip6s=pulumi.get(__ret__, 'remote_ip6s'),
        remote_ips=pulumi.get(__ret__, 'remote_ips'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'),
        vni=pulumi.get(__ret__, 'vni'))


@_utilities.lift_output_func(get_vxlan)
def get_vxlan_output(name: Optional[pulumi.Input[str]] = None,
                     vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVxlanResult]:
    """
    Use this data source to get information on an fortios system vxlan


    :param str name: Specify the name of the desired system vxlan.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
