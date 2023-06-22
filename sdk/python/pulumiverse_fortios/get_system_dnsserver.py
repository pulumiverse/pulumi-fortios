# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSystemDnsserverResult',
    'AwaitableGetSystemDnsserverResult',
    'get_system_dnsserver',
    'get_system_dnsserver_output',
]

@pulumi.output_type
class GetSystemDnsserverResult:
    """
    A collection of values returned by getSystemDnsserver.
    """
    def __init__(__self__, dnsfilter_profile=None, doh=None, id=None, mode=None, name=None, vdomparam=None):
        if dnsfilter_profile and not isinstance(dnsfilter_profile, str):
            raise TypeError("Expected argument 'dnsfilter_profile' to be a str")
        pulumi.set(__self__, "dnsfilter_profile", dnsfilter_profile)
        if doh and not isinstance(doh, str):
            raise TypeError("Expected argument 'doh' to be a str")
        pulumi.set(__self__, "doh", doh)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mode and not isinstance(mode, str):
            raise TypeError("Expected argument 'mode' to be a str")
        pulumi.set(__self__, "mode", mode)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dnsfilterProfile")
    def dnsfilter_profile(self) -> str:
        """
        DNS filter profile.
        """
        return pulumi.get(self, "dnsfilter_profile")

    @property
    @pulumi.getter
    def doh(self) -> str:
        """
        DNS over HTTPS.
        """
        return pulumi.get(self, "doh")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def mode(self) -> str:
        """
        DNS server mode.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        DNS server name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemDnsserverResult(GetSystemDnsserverResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemDnsserverResult(
            dnsfilter_profile=self.dnsfilter_profile,
            doh=self.doh,
            id=self.id,
            mode=self.mode,
            name=self.name,
            vdomparam=self.vdomparam)


def get_system_dnsserver(name: Optional[str] = None,
                         vdomparam: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemDnsserverResult:
    """
    Use this data source to get information on an fortios system dnsserver


    :param str name: Specify the name of the desired system dnsserver.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemDnsserver:getSystemDnsserver', __args__, opts=opts, typ=GetSystemDnsserverResult).value

    return AwaitableGetSystemDnsserverResult(
        dnsfilter_profile=__ret__.dnsfilter_profile,
        doh=__ret__.doh,
        id=__ret__.id,
        mode=__ret__.mode,
        name=__ret__.name,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_system_dnsserver)
def get_system_dnsserver_output(name: Optional[pulumi.Input[str]] = None,
                                vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemDnsserverResult]:
    """
    Use this data source to get information on an fortios system dnsserver


    :param str name: Specify the name of the desired system dnsserver.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
