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
    'GetNetflowResult',
    'AwaitableGetNetflowResult',
    'get_netflow',
    'get_netflow_output',
]

@pulumi.output_type
class GetNetflowResult:
    """
    A collection of values returned by getNetflow.
    """
    def __init__(__self__, active_flow_timeout=None, collector_ip=None, collector_port=None, id=None, inactive_flow_timeout=None, interface=None, interface_select_method=None, source_ip=None, template_tx_counter=None, template_tx_timeout=None, vdomparam=None):
        if active_flow_timeout and not isinstance(active_flow_timeout, int):
            raise TypeError("Expected argument 'active_flow_timeout' to be a int")
        pulumi.set(__self__, "active_flow_timeout", active_flow_timeout)
        if collector_ip and not isinstance(collector_ip, str):
            raise TypeError("Expected argument 'collector_ip' to be a str")
        pulumi.set(__self__, "collector_ip", collector_ip)
        if collector_port and not isinstance(collector_port, int):
            raise TypeError("Expected argument 'collector_port' to be a int")
        pulumi.set(__self__, "collector_port", collector_port)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if inactive_flow_timeout and not isinstance(inactive_flow_timeout, int):
            raise TypeError("Expected argument 'inactive_flow_timeout' to be a int")
        pulumi.set(__self__, "inactive_flow_timeout", inactive_flow_timeout)
        if interface and not isinstance(interface, str):
            raise TypeError("Expected argument 'interface' to be a str")
        pulumi.set(__self__, "interface", interface)
        if interface_select_method and not isinstance(interface_select_method, str):
            raise TypeError("Expected argument 'interface_select_method' to be a str")
        pulumi.set(__self__, "interface_select_method", interface_select_method)
        if source_ip and not isinstance(source_ip, str):
            raise TypeError("Expected argument 'source_ip' to be a str")
        pulumi.set(__self__, "source_ip", source_ip)
        if template_tx_counter and not isinstance(template_tx_counter, int):
            raise TypeError("Expected argument 'template_tx_counter' to be a int")
        pulumi.set(__self__, "template_tx_counter", template_tx_counter)
        if template_tx_timeout and not isinstance(template_tx_timeout, int):
            raise TypeError("Expected argument 'template_tx_timeout' to be a int")
        pulumi.set(__self__, "template_tx_timeout", template_tx_timeout)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="activeFlowTimeout")
    def active_flow_timeout(self) -> int:
        """
        Timeout to report active flows (1 - 60 min, default = 30).
        """
        return pulumi.get(self, "active_flow_timeout")

    @property
    @pulumi.getter(name="collectorIp")
    def collector_ip(self) -> str:
        """
        Collector IP.
        """
        return pulumi.get(self, "collector_ip")

    @property
    @pulumi.getter(name="collectorPort")
    def collector_port(self) -> int:
        """
        NetFlow collector port number.
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
    @pulumi.getter(name="inactiveFlowTimeout")
    def inactive_flow_timeout(self) -> int:
        """
        Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
        """
        return pulumi.get(self, "inactive_flow_timeout")

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
        Source IP address for communication with the NetFlow agent.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="templateTxCounter")
    def template_tx_counter(self) -> int:
        """
        Counter of flowset records before resending a template flowset record.
        """
        return pulumi.get(self, "template_tx_counter")

    @property
    @pulumi.getter(name="templateTxTimeout")
    def template_tx_timeout(self) -> int:
        """
        Timeout for periodic template flowset transmission (1 - 1440 min, default = 30).
        """
        return pulumi.get(self, "template_tx_timeout")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetNetflowResult(GetNetflowResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetflowResult(
            active_flow_timeout=self.active_flow_timeout,
            collector_ip=self.collector_ip,
            collector_port=self.collector_port,
            id=self.id,
            inactive_flow_timeout=self.inactive_flow_timeout,
            interface=self.interface,
            interface_select_method=self.interface_select_method,
            source_ip=self.source_ip,
            template_tx_counter=self.template_tx_counter,
            template_tx_timeout=self.template_tx_timeout,
            vdomparam=self.vdomparam)


def get_netflow(vdomparam: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetflowResult:
    """
    Use this data source to get information on fortios system netflow


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:system/getNetflow:getNetflow', __args__, opts=opts, typ=GetNetflowResult).value

    return AwaitableGetNetflowResult(
        active_flow_timeout=pulumi.get(__ret__, 'active_flow_timeout'),
        collector_ip=pulumi.get(__ret__, 'collector_ip'),
        collector_port=pulumi.get(__ret__, 'collector_port'),
        id=pulumi.get(__ret__, 'id'),
        inactive_flow_timeout=pulumi.get(__ret__, 'inactive_flow_timeout'),
        interface=pulumi.get(__ret__, 'interface'),
        interface_select_method=pulumi.get(__ret__, 'interface_select_method'),
        source_ip=pulumi.get(__ret__, 'source_ip'),
        template_tx_counter=pulumi.get(__ret__, 'template_tx_counter'),
        template_tx_timeout=pulumi.get(__ret__, 'template_tx_timeout'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_netflow)
def get_netflow_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNetflowResult]:
    """
    Use this data source to get information on fortios system netflow


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
