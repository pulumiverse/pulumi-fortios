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
    'GetProberesponseResult',
    'AwaitableGetProberesponseResult',
    'get_proberesponse',
    'get_proberesponse_output',
]

@pulumi.output_type
class GetProberesponseResult:
    """
    A collection of values returned by getProberesponse.
    """
    def __init__(__self__, http_probe_value=None, id=None, mode=None, password=None, port=None, security_mode=None, timeout=None, ttl_mode=None, vdomparam=None):
        if http_probe_value and not isinstance(http_probe_value, str):
            raise TypeError("Expected argument 'http_probe_value' to be a str")
        pulumi.set(__self__, "http_probe_value", http_probe_value)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mode and not isinstance(mode, str):
            raise TypeError("Expected argument 'mode' to be a str")
        pulumi.set(__self__, "mode", mode)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if security_mode and not isinstance(security_mode, str):
            raise TypeError("Expected argument 'security_mode' to be a str")
        pulumi.set(__self__, "security_mode", security_mode)
        if timeout and not isinstance(timeout, int):
            raise TypeError("Expected argument 'timeout' to be a int")
        pulumi.set(__self__, "timeout", timeout)
        if ttl_mode and not isinstance(ttl_mode, str):
            raise TypeError("Expected argument 'ttl_mode' to be a str")
        pulumi.set(__self__, "ttl_mode", ttl_mode)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="httpProbeValue")
    def http_probe_value(self) -> str:
        """
        Value to respond to the monitoring server.
        """
        return pulumi.get(self, "http_probe_value")

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
        SLA response mode.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        Twamp respondor password in authentication mode
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        Port number to response.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="securityMode")
    def security_mode(self) -> str:
        """
        Twamp respondor security mode.
        """
        return pulumi.get(self, "security_mode")

    @property
    @pulumi.getter
    def timeout(self) -> int:
        """
        An inactivity timer for a twamp test session.
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter(name="ttlMode")
    def ttl_mode(self) -> str:
        """
        Mode for TWAMP packet TTL modification.
        """
        return pulumi.get(self, "ttl_mode")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetProberesponseResult(GetProberesponseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProberesponseResult(
            http_probe_value=self.http_probe_value,
            id=self.id,
            mode=self.mode,
            password=self.password,
            port=self.port,
            security_mode=self.security_mode,
            timeout=self.timeout,
            ttl_mode=self.ttl_mode,
            vdomparam=self.vdomparam)


def get_proberesponse(vdomparam: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProberesponseResult:
    """
    Use this data source to get information on fortios system proberesponse


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:system/getProberesponse:getProberesponse', __args__, opts=opts, typ=GetProberesponseResult).value

    return AwaitableGetProberesponseResult(
        http_probe_value=pulumi.get(__ret__, 'http_probe_value'),
        id=pulumi.get(__ret__, 'id'),
        mode=pulumi.get(__ret__, 'mode'),
        password=pulumi.get(__ret__, 'password'),
        port=pulumi.get(__ret__, 'port'),
        security_mode=pulumi.get(__ret__, 'security_mode'),
        timeout=pulumi.get(__ret__, 'timeout'),
        ttl_mode=pulumi.get(__ret__, 'ttl_mode'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_proberesponse)
def get_proberesponse_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProberesponseResult]:
    """
    Use this data source to get information on fortios system proberesponse


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
