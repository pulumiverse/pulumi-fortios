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
    'GetIpiptunnelResult',
    'AwaitableGetIpiptunnelResult',
    'get_ipiptunnel',
    'get_ipiptunnel_output',
]

@pulumi.output_type
class GetIpiptunnelResult:
    """
    A collection of values returned by getIpiptunnel.
    """
    def __init__(__self__, auto_asic_offload=None, id=None, interface=None, local_gw=None, name=None, remote_gw=None, use_sdwan=None, vdomparam=None):
        if auto_asic_offload and not isinstance(auto_asic_offload, str):
            raise TypeError("Expected argument 'auto_asic_offload' to be a str")
        pulumi.set(__self__, "auto_asic_offload", auto_asic_offload)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface and not isinstance(interface, str):
            raise TypeError("Expected argument 'interface' to be a str")
        pulumi.set(__self__, "interface", interface)
        if local_gw and not isinstance(local_gw, str):
            raise TypeError("Expected argument 'local_gw' to be a str")
        pulumi.set(__self__, "local_gw", local_gw)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if remote_gw and not isinstance(remote_gw, str):
            raise TypeError("Expected argument 'remote_gw' to be a str")
        pulumi.set(__self__, "remote_gw", remote_gw)
        if use_sdwan and not isinstance(use_sdwan, str):
            raise TypeError("Expected argument 'use_sdwan' to be a str")
        pulumi.set(__self__, "use_sdwan", use_sdwan)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="autoAsicOffload")
    def auto_asic_offload(self) -> str:
        """
        Enable/disable tunnel ASIC offloading.
        """
        return pulumi.get(self, "auto_asic_offload")

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
        Interface name that is associated with the incoming traffic from available options.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="localGw")
    def local_gw(self) -> str:
        """
        IPv4 address for the local gateway.
        """
        return pulumi.get(self, "local_gw")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        IPIP Tunnel name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="remoteGw")
    def remote_gw(self) -> str:
        """
        IPv4 address for the remote gateway.
        """
        return pulumi.get(self, "remote_gw")

    @property
    @pulumi.getter(name="useSdwan")
    def use_sdwan(self) -> str:
        """
        Enable/disable use of SD-WAN to reach remote gateway.
        """
        return pulumi.get(self, "use_sdwan")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetIpiptunnelResult(GetIpiptunnelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpiptunnelResult(
            auto_asic_offload=self.auto_asic_offload,
            id=self.id,
            interface=self.interface,
            local_gw=self.local_gw,
            name=self.name,
            remote_gw=self.remote_gw,
            use_sdwan=self.use_sdwan,
            vdomparam=self.vdomparam)


def get_ipiptunnel(name: Optional[str] = None,
                   vdomparam: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIpiptunnelResult:
    """
    Use this data source to get information on an fortios system ipiptunnel


    :param str name: Specify the name of the desired system ipiptunnel.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:system/getIpiptunnel:getIpiptunnel', __args__, opts=opts, typ=GetIpiptunnelResult).value

    return AwaitableGetIpiptunnelResult(
        auto_asic_offload=pulumi.get(__ret__, 'auto_asic_offload'),
        id=pulumi.get(__ret__, 'id'),
        interface=pulumi.get(__ret__, 'interface'),
        local_gw=pulumi.get(__ret__, 'local_gw'),
        name=pulumi.get(__ret__, 'name'),
        remote_gw=pulumi.get(__ret__, 'remote_gw'),
        use_sdwan=pulumi.get(__ret__, 'use_sdwan'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_ipiptunnel)
def get_ipiptunnel_output(name: Optional[pulumi.Input[str]] = None,
                          vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIpiptunnelResult]:
    """
    Use this data source to get information on an fortios system ipiptunnel


    :param str name: Specify the name of the desired system ipiptunnel.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
