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
    'GetSflowResult',
    'AwaitableGetSflowResult',
    'get_sflow',
    'get_sflow_output',
]

@pulumi.output_type
class GetSflowResult:
    """
    A collection of values returned by getSflow.
    """
    def __init__(__self__, collector_ip=None, collector_port=None, id=None, interface=None, interface_select_method=None, source_ip=None, vdomparam=None):
        if collector_ip and not isinstance(collector_ip, str):
            raise TypeError("Expected argument 'collector_ip' to be a str")
        pulumi.set(__self__, "collector_ip", collector_ip)
        if collector_port and not isinstance(collector_port, int):
            raise TypeError("Expected argument 'collector_port' to be a int")
        pulumi.set(__self__, "collector_port", collector_port)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface and not isinstance(interface, str):
            raise TypeError("Expected argument 'interface' to be a str")
        pulumi.set(__self__, "interface", interface)
        if interface_select_method and not isinstance(interface_select_method, str):
            raise TypeError("Expected argument 'interface_select_method' to be a str")
        pulumi.set(__self__, "interface_select_method", interface_select_method)
        if source_ip and not isinstance(source_ip, str):
            raise TypeError("Expected argument 'source_ip' to be a str")
        pulumi.set(__self__, "source_ip", source_ip)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="collectorIp")
    def collector_ip(self) -> str:
        """
        IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
        """
        return pulumi.get(self, "collector_ip")

    @property
    @pulumi.getter(name="collectorPort")
    def collector_port(self) -> int:
        """
        UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
        """
        return pulumi.get(self, "collector_port")

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
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> str:
        """
        Source IP address for sFlow agent.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSflowResult(GetSflowResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSflowResult(
            collector_ip=self.collector_ip,
            collector_port=self.collector_port,
            id=self.id,
            interface=self.interface,
            interface_select_method=self.interface_select_method,
            source_ip=self.source_ip,
            vdomparam=self.vdomparam)


def get_sflow(vdomparam: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSflowResult:
    """
    Use this data source to get information on fortios system sflow


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:system/getSflow:getSflow', __args__, opts=opts, typ=GetSflowResult).value

    return AwaitableGetSflowResult(
        collector_ip=pulumi.get(__ret__, 'collector_ip'),
        collector_port=pulumi.get(__ret__, 'collector_port'),
        id=pulumi.get(__ret__, 'id'),
        interface=pulumi.get(__ret__, 'interface'),
        interface_select_method=pulumi.get(__ret__, 'interface_select_method'),
        source_ip=pulumi.get(__ret__, 'source_ip'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_sflow)
def get_sflow_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSflowResult]:
    """
    Use this data source to get information on fortios system sflow


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
